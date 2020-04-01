package oci

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/fatih/color"

	"github.com/hashicorp/terraform/backend/local"
	"github.com/mitchellh/cli"

	"github.com/hashicorp/hcl/hcl/fmtcmd"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

const (
	exportUserAgentFormatter        = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform-oci-exporter/%s)"
	defaultTmpStateFile             = "terraform.tfstate.tmp"
	varsFile                        = "vars.tf"
	missingRequiredAttributeWarning = `Warning: There are one or more 'Required' attributes for which a value could not be discovered.
This may be expected behavior from the service, which may prevent discovery of certain sensitive attributes or secrets.
Run 'terraform plan' against the generated configuration files to get more information about the missing values.`
)

var referenceMap map[string]string
var vars map[string]string
var resourceNameCount map[string]int
var resourcesMap map[string]*schema.Resource
var datasourcesMap map[string]*schema.Resource
var compartmentScopeServices []string
var isMissingRequiredAttributes bool

func init() {
	resourceNameCount = map[string]int{}
	vars = map[string]string{}
	referenceMap = map[string]string{}

	compartmentScopeServices = make([]string, len(compartmentResourceGraphs))
	idx := 0
	for mode := range compartmentResourceGraphs {
		compartmentScopeServices[idx] = mode
		idx++
	}

	isMissingRequiredAttributes = false
}

func printResourceGraphResources(resourceGraphs map[string]TerraformResourceGraph, scope string) error {
	for graphName, resourceGraph := range resourceGraphs {
		// Need a set here because the same resource type may have multiple associations in the same graph
		// This avoids adding duplicates of those resource types
		resourceSet := map[string]bool{}
		for _, association := range resourceGraph {
			for _, hint := range association {
				if _, isResource := resourcesMap[hint.resourceClass]; isResource {
					resourceSet[hint.resourceClass] = true
				}
			}
		}

		if len(resourceSet) > 0 {
			supportedResources := make([]string, len(resourceSet))
			idx := 0
			for resourceClass := range resourceSet {
				supportedResources[idx] = resourceClass
				idx++
			}

			sort.Strings(supportedResources)
			log.Printf("%s (%s-scope resources)", graphName, scope)
			log.Printf("===========")
			for _, resourceClass := range supportedResources {
				log.Printf("- %s", resourceClass)
			}
			log.Println("")
		}
	}
	return nil
}

func RunListExportableResourcesCommand() error {
	resourcesMap = ResourcesMap()
	datasourcesMap = DataSourcesMap()

	log.Println("List of Discoverable Oracle Cloud Infrastructure Resources")

	if err := printResourceGraphResources(tenancyResourceGraphs, "tenancy"); err != nil {
		return err
	}

	if err := printResourceGraphResources(compartmentResourceGraphs, "compartment"); err != nil {
		return err
	}

	return nil
}

type ExportCommandArgs struct {
	CompartmentId   *string
	CompartmentName *string
	IDs             []string
	Services        []string
	OutputDir       *string
	GenerateState   bool
}

func RunExportCommand(args *ExportCommandArgs) error {
	resourcesMap = ResourcesMap()
	datasourcesMap = DataSourcesMap()

	if err := args.validate(); err != nil {
		return err
	}

	clients := &OracleClients{}
	userAgentString := fmt.Sprintf(exportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, Version)
	httpClient := buildHttpClient()

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClient, err := buildConfigureClientFn(oci_common.DefaultConfigProvider(), httpClient)
	if err != nil {
		return err
	}

	configureClientWithUserAgent := func(client *oci_common.BaseClient) error {
		if err := configureClient(client); err != nil {
			return err
		}
		client.UserAgent = userAgentString
		return nil
	}

	err = createSDKClients(clients, oci_common.DefaultConfigProvider(), configureClientWithUserAgent)
	if err != nil {
		return err
	}

	if args.CompartmentName != nil && *args.CompartmentName != "" {
		var err error
		args.CompartmentId, err = resolveCompartmentId(clients, args.CompartmentName)
		if err != nil {
			return err
		}
	}

	return runExportCommand(clients, args)
}

// Dedupes possible repeating services from command line and sorts them
func (args *ExportCommandArgs) finalizeServices() {
	seenServices := map[string]bool{}
	finalServices := []string{}

	for _, service := range args.Services {
		if _, seen := seenServices[service]; seen {
			continue
		}
		finalServices = append(finalServices, service)
		seenServices[service] = true
	}
	args.Services = finalServices
	sort.Strings(args.Services)
}

// Validate export command arguments and returns nil if there are no issues
func (args *ExportCommandArgs) validate() error {
	path, err := os.Stat(*args.OutputDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("[ERROR] output_path does not exist: %s", err)
	}

	if !path.IsDir() {
		return fmt.Errorf("[ERROR] output_path %s should be a directory", *args.OutputDir)
	}

	return nil
}

func runExportCommand(clients *OracleClients, args *ExportCommandArgs) error {
	if args.OutputDir == nil || *args.OutputDir == "" {
		return fmt.Errorf("[ERROR] no output directory specified")
	}

	stateOutputFile := fmt.Sprintf("%s%s%s", *args.OutputDir, string(os.PathSeparator), local.DefaultStateFilename)
	tmpStateOutputFile := fmt.Sprintf("%s%s%s", *args.OutputDir, string(os.PathSeparator), defaultTmpStateFile)

	log.Printf("Running export command\n")
	if len(args.Services) == 0 {
		args.Services = compartmentScopeServices
	}

	args.finalizeServices()
	generateConfigSteps, err := buildGenerateConfigSteps(args.CompartmentId, args.Services, oci_common.DefaultConfigProvider())
	if err != nil {
		return err
	}

	// Discover and build a model of all targeted resources
	matchResourceIds := map[string]bool{}
	for _, id := range args.IDs {
		if id != "" {
			matchResourceIds[id] = false
		}
	}

	for _, step := range generateConfigSteps {
		// Discover all resources in the compartment
		ociResources, err := findResources(clients, step.root, step.resourceGraph, matchResourceIds)
		if err != nil {
			return err
		}

		// Filter out omitted resources from export
		step.discoveredResources = []*OCIResource{}
		step.omittedResources = []*OCIResource{}
		for _, resource := range ociResources {
			if !resource.omitFromExport {
				referenceMap[resource.id] = resource.getHclReferenceIdString()
				step.discoveredResources = append(step.discoveredResources, resource)
			} else {
				step.omittedResources = append(step.omittedResources, resource)
			}
		}
	}

	// Cull any references from the ref map that contain omitted resources
	// This is to avoid omitted resources from being referenced in generated configs
	for _, step := range generateConfigSteps {

		for _, omittedResource := range step.omittedResources {
			for key, reference := range referenceMap {
				if strings.Contains(reference, omittedResource.getTerraformReference()) {
					delete(referenceMap, key)
				}
			}
		}
	}

	// Generate HCL configs from all discovered resources
	allDiscoveredResources := []*OCIResource{}
	summaryStatements := []string{}
	defer func() {
		for _, statement := range summaryStatements {
			color.Green(statement)
		}
	}()

	for _, step := range generateConfigSteps {
		configOutputFile := fmt.Sprintf("%s%s%s.tf", *args.OutputDir, string(os.PathSeparator), step.stepName)
		tmpConfigOutputFile := fmt.Sprintf("%s%s%s.tf.tmp", *args.OutputDir, string(os.PathSeparator), step.stepName)

		file, err := os.OpenFile(tmpConfigOutputFile, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}

		// Build the HCL config
		// Note that we still build a TF file even if no resources were discovered for this TF file.
		// A user may run this command multiple times and may see stale resources if we don't overwrite the file with
		// an empty one.
		builder := &strings.Builder{}
		builder.WriteString("## This configuration was generated by terraform-provider-oci\n\n")

		for _, resource := range step.discoveredResources {
			log.Printf("[INFO] ===> Generating resource '%s'", resource.getTerraformReference())
			if err := resource.getHCLString(builder, referenceMap); err != nil {
				_ = file.Close()
				return err
			}
			allDiscoveredResources = append(allDiscoveredResources, resource)
		}

		_, err = file.WriteString(builder.String())
		if err != nil {
			_ = file.Close()
			return err
		}

		if fErr := file.Close(); fErr != nil {
			return fErr
		}

		if err := os.Rename(tmpConfigOutputFile, configOutputFile); err != nil {
			return err
		}

		// Format the HCL config
		if err := fmtcmd.Run([]string{configOutputFile}, []string{}, nil, nil, fmtcmd.Options{Write: true}); err != nil {
			return err
		}

		summaryStatements = append(summaryStatements, fmt.Sprintf("Found %d '%s' resources. Generated under '%s'", len(step.discoveredResources), step.stepName, configOutputFile))
	}

	if isMissingRequiredAttributes {
		summaryStatements = append(summaryStatements, "")
		summaryStatements = append(summaryStatements, missingRequiredAttributeWarning)
	}

	if err := generateVarsFile(vars, args.OutputDir); err != nil {
		return err
	}

	if args.GenerateState {
		// Run init and import commands
		meta := command.Meta{
			Ui: &cli.BasicUi{
				Reader:      os.Stdin,
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
			},
			RunningInAutomation: true,
		}

		initCmd := command.InitCommand{Meta: meta}
		var initArgs []string
		if pluginDir := getEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
			log.Printf("[INFO] plugin dir: '%s'", pluginDir)
			initArgs = append(initArgs, fmt.Sprintf("-plugin-dir=%v", pluginDir))
		}
		initArgs = append(initArgs, *args.OutputDir)
		if errCode := initCmd.Run(initArgs); errCode != 0 {
			return nil
		}

		if err := os.RemoveAll(tmpStateOutputFile); err != nil {
			log.Printf("[WARN] unable to delete existing tmp state file %s", tmpStateOutputFile)
			return err
		}

		for _, resource := range allDiscoveredResources {
			log.Printf("[INFO] ===> Importing resource '%s'", resource.getTerraformReference())

			resourceDefinition, exists := resourcesMap[resource.terraformClass]
			if !exists {
				log.Printf("[INFO] skip importing '%s' since it is not a Terraform OCI resource", resource.getTerraformReference())
				continue
			}

			if resourceDefinition.Importer == nil {
				log.Printf("[WARN] unable to import '%s' because import is not supported for '%s'", resource.getTerraformReference(), resource.terraformClass)
				continue
			}

			importCmd := command.ImportCommand{Meta: meta}
			importId := resource.importId
			if len(importId) == 0 {
				importId = resource.id
			}

			importArgs := []string{
				fmt.Sprintf("-config=%s", *args.OutputDir),
				fmt.Sprintf("-state=%s", tmpStateOutputFile),
				resource.getTerraformReference(),
				importId,
			}
			if errCode := importCmd.Run(importArgs); errCode != 0 {
				return fmt.Errorf("[ERROR] terraform import command failed for resource '%s' at id '%s'", resource.getTerraformReference(), importId)
			}
		}

		if _, err := os.Stat(tmpStateOutputFile); !os.IsNotExist(err) {
			if err := os.Rename(tmpStateOutputFile, stateOutputFile); err != nil {
				return err
			}
		}
	}

	if len(matchResourceIds) > 0 {
		missingResourceIds := []string{}
		for resourceId, found := range matchResourceIds {
			if !found {
				missingResourceIds = append(missingResourceIds, resourceId)
			}
		}

		if len(missingResourceIds) > 0 {
			summaryStatements = append(summaryStatements, "")
			summaryStatements = append(summaryStatements, "Warning: The following resource IDs were not found.")
			for _, resourceId := range missingResourceIds {
				summaryStatements = append(summaryStatements, fmt.Sprintf("- %s", resourceId))
			}
			return fmt.Errorf("[ERROR] one or more expected resource ids were not found")
		}
	}

	summaryStatements = append(summaryStatements, "\n=== COMPLETED ===")

	return nil
}

func buildGenerateConfigSteps(compartmentId *string, services []string, configProvider oci_common.ConfigurationProvider) ([]*GenerateConfigStep, error) {
	result := []*GenerateConfigStep{}

	tenancyId, err := configProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	tenancyResource := &OCIResource{
		compartmentId: tenancyId,
		TerraformResource: TerraformResource{
			id:             tenancyId,
			terraformClass: "oci_identity_tenancy",
			terraformName:  "export",
		},
	}

	for _, mode := range services {
		if resourceGraph, exists := tenancyResourceGraphs[mode]; exists {
			result = append(result, &GenerateConfigStep{
				root:          tenancyResource,
				resourceGraph: resourceGraph,
				stepName:      mode,
			})

			vars["tenancy_ocid"] = ""
			referenceMap[tenancyId] = "${var.tenancy_ocid}"
		}
	}

	if compartmentId == nil || *compartmentId == "" {
		*compartmentId = tenancyId
	}
	compartmentResource := &OCIResource{
		compartmentId: *compartmentId,
		TerraformResource: TerraformResource{
			id:             *compartmentId,
			terraformClass: "oci_identity_compartment",
			terraformName:  "export",
		},
	}

	for _, mode := range services {
		if resourceGraph, exists := compartmentResourceGraphs[mode]; exists {
			result = append(result, &GenerateConfigStep{
				root:          compartmentResource,
				resourceGraph: resourceGraph,
				stepName:      mode,
			})

			vars["compartment_ocid"] = ""
			referenceMap[*compartmentId] = "${var.compartment_ocid}"
		}
	}

	return result, nil
}

func findResources(clients *OracleClients, root *OCIResource, resourceGraph TerraformResourceGraph, exportableResourceIds map[string]bool) ([]*OCIResource, error) {
	foundResources := []*OCIResource{}

	childResourceTypes, exists := resourceGraph[root.terraformClass]
	if !exists {
		return foundResources, nil
	}

	log.Printf("[INFO] resource discovery: visiting %s\n", root.getTerraformReference())

	for _, childType := range childResourceTypes {
		findResourceFn := findResourcesGeneric
		if childType.findResourcesOverrideFn != nil {
			findResourceFn = childType.findResourcesOverrideFn
		}
		results, err := findResourceFn(clients, &childType, root)
		if err != nil {
			return foundResources, err
		}

		if childType.processDiscoveredResourcesFn != nil {
			results, err = childType.processDiscoveredResourcesFn(clients, results)
			if err != nil {
				return foundResources, err
			}
		}
		foundResources = append(foundResources, results...)

		for _, resource := range results {
			//referenceMap[resource.id] = resource.getHclReferenceIdString()
			if exportableResourceIds != nil && len(exportableResourceIds) > 0 {
				if _, shouldExport := exportableResourceIds[resource.id]; shouldExport {
					resource.omitFromExport = false
					exportableResourceIds[resource.id] = true
				} else {
					resource.omitFromExport = !childType.alwaysExportable
				}
			}

			subResources, err := findResources(clients, resource, resourceGraph, exportableResourceIds)
			if err != nil {
				return foundResources, err
			}
			foundResources = append(foundResources, subResources...)
		}
	}

	return foundResources, nil
}

func generateVarsFile(vars map[string]string, outputDir *string) error {
	varsTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), varsFile)
	varsOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), varsFile)
	file, err := os.OpenFile(varsTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	for variable, defaultVal := range vars {
		if defaultVal != "" {
			_, _ = file.WriteString(fmt.Sprintf("variable %s { default = %s }\n", variable, defaultVal))
		} else {
			_, _ = file.WriteString(fmt.Sprintf("variable %s {}\n", variable))
		}
	}

	if err := file.Close(); err != nil {
		return err
	}

	if err := os.Rename(varsTmpFile, varsOutputFile); err != nil {
		return err
	}

	return nil
}

type OCIResource struct {
	TerraformResource
	compartmentId    string
	rawResource      interface{}
	sourceAttributes map[string]interface{}
	getHclStringFn   func(*strings.Builder, *OCIResource, map[string]string) error
	parent           *OCIResource
}

type TerraformResource struct {
	id                         string
	importId                   string
	terraformClass             string
	terraformName              string
	terraformReferenceIdString string
	terraformTypeInfo          *TerraformResourceHints
	omitFromExport             bool
}

func (tr *TerraformResource) getHclReferenceIdString() string {
	if tr.terraformReferenceIdString != "" {
		return fmt.Sprintf("${%s}", tr.terraformReferenceIdString)
	}
	return fmt.Sprintf("${%s.id}", tr.getTerraformReference())
}

func (tr *TerraformResource) getTerraformReference() string {
	return fmt.Sprintf("%s.%s", tr.terraformClass, tr.terraformName)
}

func getHCLStringFromMap(builder *strings.Builder, sourceAttributes map[string]interface{}, resourceSchema *schema.Resource, interpolationMap map[string]string) error {
	sortedKeys := make([]string, len(resourceSchema.Schema))
	cnt := 0
	for k := range resourceSchema.Schema {
		sortedKeys[cnt] = k
		cnt++
	}
	sort.Strings(sortedKeys)

	for _, tfAttribute := range sortedKeys {
		tfSchema := resourceSchema.Schema[tfAttribute]
		if tfSchema.Deprecated != "" || tfSchema.Removed != "" || (!tfSchema.Required && !tfSchema.Optional) {
			continue
		}

		if attributeVal, exists := sourceAttributes[tfAttribute]; exists {
			switch v := attributeVal.(type) {
			case string:
				if varOverride, exists := interpolationMap[fmt.Sprintf("%v", v)]; exists {
					v = varOverride
				}
				builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, v))
				continue
			case int, bool, float64:
				builder.WriteString(fmt.Sprintf("%s = \"%v\"\n", tfAttribute, v))
				continue
			case []interface{}:
				switch tfSchema.Type {
				case schema.TypeList, schema.TypeSet:
					switch elem := tfSchema.Elem.(type) {
					case *schema.Resource:
						for _, item := range v {
							if val := item.(map[string]interface{}); val != nil {
								builder.WriteString(fmt.Sprintf("%s {\n", tfAttribute))
								if err := getHCLStringFromMap(builder, val, elem, interpolationMap); err != nil {
									return err
								}
								builder.WriteString("}\n")
							}
						}
						continue
					case *schema.Schema, schema.ValueType:
						builder.WriteString(fmt.Sprintf("%s = [\n", tfAttribute))
						for _, item := range v {
							switch trueListVal := item.(type) {
							case string:
								if varOverride, exists := interpolationMap[fmt.Sprintf("%v", trueListVal)]; exists {
									trueListVal = varOverride
								}
								builder.WriteString(fmt.Sprintf("%q,\n", trueListVal))
							case int, bool, float64:
								builder.WriteString(fmt.Sprintf("\"%v\",\n", trueListVal))
							default:
								return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': List element type mismatch", tfAttribute, tfAttribute)
							}
						}
						builder.WriteString("]\n")
						continue
					}

					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': List element is neither schema.Resource or schema.Schema", tfAttribute, tfAttribute)
				}
			case map[string]interface{}:
				switch tfSchema.Type {
				case schema.TypeList:
					if nestedResource := tfSchema.Elem.(*schema.Resource); nestedResource != nil {
						builder.WriteString(fmt.Sprintf("%s {\n", tfAttribute))
						if err := getHCLStringFromMap(builder, v, nestedResource, interpolationMap); err != nil {
							return err
						}
						builder.WriteString("}\n")
						continue
					}
					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': Nested resource type mismatch", tfAttribute, tfAttribute)
				case schema.TypeMap:
					builder.WriteString(fmt.Sprintf("%s = {\n", tfAttribute))

					keys := getSortedKeys(v)
					for _, mapKey := range keys {
						switch mapVal := v[mapKey].(type) {
						case string:
							if varOverride, exists := interpolationMap[fmt.Sprintf("%v", mapVal)]; exists {
								mapVal = varOverride
							}
							builder.WriteString(fmt.Sprintf("\"%s\" = %q\n", mapKey, mapVal))
						case int, bool, float64:
							builder.WriteString(fmt.Sprintf("\"%s\" = \"%v\"\n", mapKey, mapVal))
						default:
							builder.WriteString(fmt.Sprintf("#%s = <<Placeholder due to complex map value>>\n", mapKey))
						}
					}
					builder.WriteString("}\n")
					continue
				default:
					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': Source attribute is nested object but TF attribute is not", tfAttribute, tfAttribute)
				}
			case nil:
				log.Printf("[INFO] TF attribute '%s' is nil in source\n", tfAttribute)
				if !tfSchema.Required {
					continue
				}
			default:
				log.Printf("[WARN] TF attribute '%s' is unknown type in source\n", tfAttribute)
			}
		}

		if tfSchema.Required {
			log.Printf("[WARN] Required TF attribute '%s' not found in source\n", tfAttribute)
			builder.WriteString(fmt.Sprintf("#%s = <<Required attribute not found in discovery>>\n", tfAttribute))
			isMissingRequiredAttributes = true
		} else if tfSchema.Optional {
			log.Printf("[INFO] Optional TF attribute '%s' not found in source\n", tfAttribute)
			builder.WriteString(fmt.Sprintf("#%s = <<Optional value not found in discovery>>\n", tfAttribute))
		}
	}
	return nil
}

func (ociRes *OCIResource) getHCLString(builder *strings.Builder, interpolationMap map[string]string) error {
	// Remove any potential cyclical references from the interpolation map
	selfReference := ociRes.getTerraformReference()
	resourceInterpolationMap := map[string]string{}
	for value, interpolation := range interpolationMap {
		if !strings.Contains(interpolation, selfReference) {
			resourceInterpolationMap[value] = interpolation
		}
	}

	if ociRes.getHclStringFn != nil {
		return ociRes.getHclStringFn(builder, ociRes, resourceInterpolationMap)
	}
	return getHclStringFromGenericMap(builder, ociRes, resourceInterpolationMap)
}

func getHclStringFromGenericMap(builder *strings.Builder, ociRes *OCIResource, interpolationMap map[string]string) error {
	resourceSchema := resourcesMap[ociRes.terraformClass]

	builder.WriteString(fmt.Sprintf("resource %s %s {\n", ociRes.terraformClass, ociRes.terraformName))
	if err := getHCLStringFromMap(builder, ociRes.sourceAttributes, resourceSchema, interpolationMap); err != nil {
		return err
	}
	builder.WriteString("}\n")

	return nil
}

// This function attempts to convert resource data items to a map representation that omits attributes where no value was set.
func convertDatasourceItemToMap(d *schema.ResourceData, itemPrefix string, itemSchema map[string]*schema.Schema) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	for attributeKey, attributeSchema := range itemSchema {
		var key string
		if itemPrefix != "" {
			key = fmt.Sprintf("%s.%s", itemPrefix, attributeKey)
		} else {
			key = attributeKey
		}

		switch attributeSchema.Type {
		case schema.TypeBool, schema.TypeInt, schema.TypeFloat, schema.TypeString:
			if val, exists := d.GetOkExists(key); exists {
				result[attributeKey] = val
			}
		case schema.TypeList:
			switch v := attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					result[attributeKey] = val
				}
			case *schema.Resource:
				if val, exists := d.GetOkExists(key); exists {
					list := val.([]interface{})
					resourceList := make([]interface{}, len(list))
					for idx := range list {
						resourceList[idx], _ = convertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, idx), v.Schema)
					}
					result[attributeKey] = resourceList
				}
			}
		case schema.TypeMap:
			switch attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					result[attributeKey] = val
				}
			default:
				return result, fmt.Errorf("[ERROR] found a non-primitive element in schema for TypeMap attribute '%s'", attributeKey)
			}
		case schema.TypeSet:
			switch v := attributeSchema.Elem.(type) {
			case *schema.Schema, schema.ValueType:
				if val, exists := d.GetOkExists(key); exists {
					setVal := val.(*schema.Set)
					result[attributeKey] = setVal.List()
				}
			case *schema.Resource:
				if val, exists := d.GetOkExists(key); exists {
					setVal := val.(*schema.Set)
					list := setVal.List()
					resourceList := make([]interface{}, len(list))
					for idx, item := range list {
						itemHashCode := setVal.F(item)
						resourceList[idx], _ = convertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, itemHashCode), v.Schema)
					}
					result[attributeKey] = resourceList
				}
			}
		}
	}

	return result, nil
}

func findResourcesGeneric(clients *OracleClients, tfMeta *TerraformResourceAssociation, parent *OCIResource) ([]*OCIResource, error) {
	results := []*OCIResource{}

	log.Printf("[INFO] discovering resources with data source '%s'\n", tfMeta.datasourceClass)
	datasource := datasourcesMap[tfMeta.datasourceClass]
	d := datasource.TestResourceData()
	d.Set("compartment_id", parent.compartmentId)

	for queryAttributeName, queryValue := range tfMeta.datasourceQueryParams {
		log.Printf("[INFO] adding datasource query attribute '%s' from parent attribute '%s'\n", queryAttributeName, queryValue)
		if queryValue == "" || queryValue == "id" {
			d.Set(queryAttributeName, parent.id)
		} else if strings.HasPrefix(queryValue, "'") && strings.HasSuffix(queryValue, "'") { // Anything encapsulated in ' ' means to use the literal value
			d.Set(queryAttributeName, queryValue[1:len(queryValue)-1])
		} else if val, ok := parent.sourceAttributes[queryValue]; ok {
			d.Set(queryAttributeName, val)
		} else {
			log.Printf("[WARN] no attribute '%s' found in parent '%s', returning no results for this resource\n", queryValue, parent.getTerraformReference())
			return results, nil
		}
	}

	if err := datasource.Read(d, clients); err != nil {
		return results, err
	}

	if tfMeta.datasourceItemsAttr != "" {
		// Results are from a plural datasource
		itemSchema := datasource.Schema[tfMeta.datasourceItemsAttr]
		elemResource, ok := itemSchema.Elem.(*schema.Resource)
		if !ok {
			return results, fmt.Errorf("[ERROR] element schema is not of a resource")
		}

		foundItems, _ := d.GetOkExists(tfMeta.datasourceItemsAttr)
		for idx, item := range foundItems.([]interface{}) {
			var resource *OCIResource
			var err error
			if tfMeta.requireResourceRefresh {
				resourceSchema := resourcesMap[tfMeta.resourceClass]
				r := resourceSchema.TestResourceData()

				// Use resource to fill in all attributes (likely because the datasource doesn't return complete info)
				if tfMeta.getIdFn != nil {
					tmpResource, err := generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", tfMeta.datasourceItemsAttr, idx), tfMeta, parent)
					if err != nil {
						return results, err
					}

					itemId, err := tfMeta.getIdFn(tmpResource)
					if err != nil {
						return results, err
					}
					r.SetId(itemId)
				} else if idSchema, exists := elemResource.Schema["id"]; exists && idSchema.Type == schema.TypeString {
					itemId := item.(map[string]interface{})["id"]
					r.SetId(itemId.(string))
				} else {
					return results, fmt.Errorf("[ERROR] elements in datasource '%s' are missing an 'id' field and is unable to generate an id", tfMeta.datasourceClass)
				}

				if err = resourceSchema.Read(r, clients); err != nil {
					return results, err
				}

				resource, err = generateOciResourceFromResourceData(r, r, resourceSchema.Schema, "", tfMeta, parent)
				if err != nil {
					return results, err
				}
			} else {
				resource, err = generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", tfMeta.datasourceItemsAttr, idx), tfMeta, parent)
				if err != nil {
					return results, err
				}
			}

			if state, ok := resource.sourceAttributes["state"]; ok && len(tfMeta.discoverableLifecycleStates) > 0 {
				discoverable := false
				for _, val := range tfMeta.discoverableLifecycleStates {
					if strings.EqualFold(state.(string), val) {
						discoverable = true
						break
					}
				}

				if !discoverable {
					continue
				}
			}

			if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, elemResource.Schema); err != nil {
				resource.terraformName = fmt.Sprintf("%s_%s_%d", parent.terraformName, tfMeta.resourceAbbreviation, idx+1)
			}

			results = append(results, resource)
		}
	} else {
		// Result is from a singular datasource
		resource, err := generateOciResourceFromResourceData(d, d, datasource.Schema, "", tfMeta, parent)
		if err != nil {
			return results, err
		}

		if resource.terraformName, err = generateTerraformNameFromResource(resource.sourceAttributes, datasource.Schema); err != nil {
			resource.terraformName = fmt.Sprintf("%s_%s", parent.terraformName, tfMeta.resourceAbbreviation)
		}

		discoverable := true
		if state, ok := resource.sourceAttributes["state"]; ok && len(tfMeta.discoverableLifecycleStates) > 0 {
			discoverable = false
			for _, val := range tfMeta.discoverableLifecycleStates {
				if strings.EqualFold(state.(string), val) {
					discoverable = true
					break
				}
			}
		}

		if discoverable {
			results = append(results, resource)
		}
	}

	return results, nil
}

func getNormalizedTerraformName(source string) string {
	// Only alphanumeric, underscore, and hyphens are allowed. Strip out anything else.
	reg, err := regexp.Compile(`[^a-zA-Z0-9\-\_]+`)
	if err != nil {
		log.Fatal(err)
	}

	result := reg.ReplaceAllString(source, "-")
	result = fmt.Sprintf("export_%s", result)
	return result
}

func convertResourceDataToMap(schemaMap map[string]*schema.Schema, d *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	for key := range schemaMap {
		if val, ok := d.GetOkExists(key); ok {
			result[key] = val
		}
	}

	return result
}

func generateTerraformNameFromResource(resourceAttributes map[string]interface{}, resourceSchema map[string]*schema.Schema) (string, error) {
	possibleNameAttributes := []string{
		"display_name",
		"name",
	}

	for _, nameAttribute := range possibleNameAttributes {
		if nameSchema, hasNameAttr := resourceSchema[nameAttribute]; hasNameAttr && nameSchema.Type == schema.TypeString {
			if value, exists := resourceAttributes[nameAttribute]; exists {
				terraformName := getNormalizedTerraformName(value.(string))
				if count, resourceNameExists := resourceNameCount[terraformName]; resourceNameExists {
					resourceNameCount[terraformName] = count + 1
					terraformName = fmt.Sprintf("%s_%d", terraformName, count)
				} else {
					resourceNameCount[terraformName] = 1
				}
				return terraformName, nil
			}
		}
	}

	return "", fmt.Errorf("unable to find a suitable name from the resource attributes")
}

func generateOciResourceFromResourceData(d *schema.ResourceData, rawResource interface{}, resourceSchema map[string]*schema.Schema, itemPrefix string, tfMeta *TerraformResourceAssociation, parent *OCIResource) (*OCIResource, error) {
	// The following conversion takes a ResourceData and converts it to a map where null values are preserved (and omitted).
	// Note that we don't use the raw map that Terraform gives us, because it will set zero-values even though the datasource didn't set one.
	//
	// TODO: An improvement on this logic would be to load the resource schema's Read function and invoke that to get the full
	// resource representation for attributes that a datasource might not give. The reasons for not doing this yet are:
	// 1) Adding an extra Read invocation could result in unnecessary traffic overhead against services for every resource we've discovered
	// 2) The result of resource Reads may return TypeSets with nested resources, which are hard to check for non-existent sub-attributes
	//
	// For now, assume that the datasource Read is good enough and add custom logic (via the process functions) if you need to add more info
	resourceMap, err := convertDatasourceItemToMap(d, itemPrefix, resourceSchema)
	if err != nil {
		return nil, err
	}

	resource := &OCIResource{
		compartmentId:    parent.compartmentId,
		sourceAttributes: resourceMap,
		rawResource:      rawResource,
		TerraformResource: TerraformResource{
			terraformClass:    tfMeta.resourceClass,
			terraformTypeInfo: tfMeta.TerraformResourceHints,
		},
		getHclStringFn: getHclStringFromGenericMap,
		parent:         parent,
	}

	if tfMeta.getIdFn != nil {
		if customId, err := tfMeta.getIdFn(resource); err == nil {
			resource.id = customId
		}
	} else if resourceId, resourceIdExists := resourceMap["id"]; resourceIdExists {
		resource.id = resourceId.(string)
	}

	if resource.id == "" {
		resource.id = d.Id()
	}

	if tfMeta.getHCLStringOverrideFn != nil {
		resource.getHclStringFn = tfMeta.getHCLStringOverrideFn
	}

	return resource, nil
}

func resolveCompartmentId(clients *OracleClients, compartmentName *string) (*string, error) {
	req := oci_identity.ListCompartmentsRequest{}

	rootCompartment, err := oci_common.DefaultConfigProvider().TenancyOCID()
	if err != nil {
		return nil, err
	}
	req.CompartmentId = &rootCompartment

	recursiveSearch := true
	req.CompartmentIdInSubtree = &recursiveSearch

	for {
		resp, err := clients.identityClient.ListCompartments(context.Background(), req)
		if err != nil {
			return nil, err
		}

		for _, compartment := range resp.Items {
			if compartment.Name != nil && *compartment.Name == *compartmentName {
				log.Printf("[INFO] resolved compartment name '%s' to compartment id '%s'", *compartmentName, *compartment.Id)
				return compartment.Id, nil
			}
		}

		if resp.OpcNextPage == nil {
			break
		}
		req.Page = resp.OpcNextPage
	}

	return nil, fmt.Errorf("[ERROR] Could not find a compartment named '%s' in your tenancy", *compartmentName)
}
