// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	"syscall"

	"github.com/hashicorp/terraform/backend/local"

	"github.com/mitchellh/cli"

	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

const (
	exportUserAgentFormatter        = "Oracle-GoSDK/%s (go/%s; %s/%s; terraform-oci-exporter/%s)"
	defaultTmpStateFile             = "terraform.tfstate.tmp"
	varsFile                        = "vars.tf"
	providerFile                    = "provider.tf"
	missingRequiredAttributeWarning = `Warning: There are one or more 'Required' attributes for which a value could not be discovered.
This may be expected behavior from the service, which may prevent discovery of certain sensitive attributes or secrets.
Placeholder values have been added for such attributes with a comment "Required attribute not found in discovery, placeholder value set to avoid plan failure".
These missing attributes are also added to the lifecycle ignore_changes.`
	placeholderValueForMissingAttribute = `<placeholder for missing required attribute>`
	EnvLogFile                          = "TF_LOG_PATH"
)

var referenceMap map[string]string
var vars map[string]string
var resourceNameCount map[string]int
var resourcesMap map[string]*schema.Resource
var datasourcesMap map[string]*schema.Resource
var compartmentScopeServices []string
var isMissingRequiredAttributes bool
var exportConfigProvider oci_common.ConfigurationProvider
var tfHclVersion TfHclVersion

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
	TFVersion       *TfHclVersion
}

func RunExportCommand(args *ExportCommandArgs) (error, Status) {
	resourcesMap = ResourcesMap()
	datasourcesMap = DataSourcesMap()

	if err := args.validate(); err != nil {
		return err, StatusFail
	}

	tfHclVersion = *args.TFVersion

	r := &schema.Resource{
		Schema: schemaMap(),
	}
	d := r.Data(nil)

	err := readEnvironmentVars(d)
	if err != nil {
		return err, StatusFail
	}

	clients, err := getExportConfig(d)
	if err != nil {
		return err, StatusFail
	}

	if args.CompartmentName != nil && *args.CompartmentName != "" {
		var err error
		args.CompartmentId, err = resolveCompartmentId(clients.(*OracleClients), args.CompartmentName)
		if err != nil {
			return err, StatusFail
		}
	}

	args.finalizeServices()

	tenancyOcid, exists := clients.(*OracleClients).configuration["tenancy_ocid"]
	if !exists {
		return fmt.Errorf("[ERROR] could not get a tenancy OCID during initialization"), StatusFail
	}

	ctx := createResourceDiscoveryContext(clients.(*OracleClients), args, tenancyOcid)

	if err := runExportCommand(ctx); err != nil {
		return err, StatusFail
	}
	if len(ctx.errorList) > 0 {
		// If the errors were from discovery of resources return partial success status
		ctx.printErrors()
		return nil, StatusPartialSuccess
	}
	return nil, StatusSuccess
}

func convertStringSliceToSet(slice []string, omitEmptyStrings bool) map[string]bool {
	result := map[string]bool{}
	for _, item := range slice {
		if omitEmptyStrings && item == "" {
			continue
		}
		result[item] = false
	}
	return result
}

func (args *ExportCommandArgs) finalizeServices() {
	if len(args.Services) == 0 {
		args.Services = compartmentScopeServices
	}

	// Dedupes possible repeating services from command line and sorts them
	finalServices := []string{}
	serviceSet := convertStringSliceToSet(args.Services, true)
	for service := range serviceSet {
		finalServices = append(finalServices, service)
	}
	args.Services = finalServices
	sort.Strings(args.Services)
}

// Validate export command arguments and returns nil if there are no issues
func (args *ExportCommandArgs) validate() error {
	if args.OutputDir == nil || *args.OutputDir == "" {
		return fmt.Errorf("[ERROR] no output directory specified")
	}

	path, err := os.Stat(*args.OutputDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("[ERROR] output_path does not exist: %s", err)
	}

	if !path.IsDir() {
		return fmt.Errorf("[ERROR] output_path %s should be a directory", *args.OutputDir)
	}

	return nil
}

func getExportConfig(d *schema.ResourceData) (interface{}, error) {
	clients := &OracleClients{
		sdkClientMap:  make(map[string]interface{}, len(oracleClientRegistrations.registeredClients)),
		configuration: make(map[string]string),
	}

	userAgentString := fmt.Sprintf(exportUserAgentFormatter, oci_common.Version(), runtime.Version(), runtime.GOOS, runtime.GOARCH, Version)
	httpClient := buildHttpClient()

	sdkConfigProvider, err := getSdkConfigProvider(d, clients)
	if err != nil {
		return nil, err
	}
	exportConfigProvider = sdkConfigProvider

	// Note: In case of Instance Principal auth, the TenancyOCID will return
	// the ocid for the tenancy for the compute instance and not the one for the customer
	// This needs to be updated in future (if we have customer request) in order to enable
	// tenancy level resources to be discovered with Instance Principal auth
	clients.configuration["tenancy_ocid"], err = sdkConfigProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClientLocal, err := buildConfigureClientFn(sdkConfigProvider, httpClient)
	if err != nil {
		return nil, err
	}

	configureClientWithUserAgent := func(client *oci_common.BaseClient) error {
		if err := configureClientLocal(client); err != nil {
			return err
		}
		client.UserAgent = userAgentString
		return nil
	}
	// beware: global variable `configureClient` set here--used elsewhere outside this execution path
	configureClient = configureClientWithUserAgent
	err = createSDKClients(clients, sdkConfigProvider, configureClientWithUserAgent)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func runExportCommand(ctx *resourceDiscoveryContext) error {
	log.Printf("Running export command\n")

	steps, err := getDiscoverResourceSteps(ctx)
	if err != nil {
		return err
	}

	logOutput := os.Stderr
	if logPath := os.Getenv(EnvLogFile); logPath != "" {
		var err error
		logOutput, err = os.OpenFile(logPath, syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, 0666)
		if err == nil {
			// go-plugin/client users go-hclog/log with os.Stderr as DefaultOutput
			os.Stderr = logOutput
			log.SetOutput(logOutput)
		}
	}

	// Discover and build a model of all targeted resources
	for _, step := range steps {
		err := step.discover()
		if err != nil {
			return err
		}
	}

	// Cull any references from the ref map that contain omitted resources
	// This is to avoid omitted resources from being referenced in generated configs
	for _, step := range steps {
		for _, omittedResource := range step.getOmittedResources() {
			for key, reference := range referenceMap {
				if strings.Contains(reference, omittedResource.getTerraformReference()) {
					delete(referenceMap, key)
				}
			}
		}
	}

	defer ctx.printSummary()

	// Generate HCL configs from all discovered resources
	for _, step := range steps {
		if err := step.writeConfiguration(); err != nil {
			return err
		}
	}

	if isMissingRequiredAttributes {
		ctx.summaryStatements = append(ctx.summaryStatements, "")
		ctx.summaryStatements = append(ctx.summaryStatements, missingRequiredAttributeWarning)
		ctx.summaryStatements = append(ctx.summaryStatements, "\nMissing required attributes:\n")
		for key, value := range ctx.missingAttributesPerResource {
			ctx.summaryStatements = append(ctx.summaryStatements, fmt.Sprintf("%s: %s", key, strings.Join(value, ",")))
		}

	}
	region, err := exportConfigProvider.Region()
	if err != nil {
		return err
	}
	vars["region"] = fmt.Sprintf("\"%s\"", region)

	if err := generateProviderFile(ctx.OutputDir); err != nil {
		return err
	}

	if err := generateVarsFile(vars, ctx.OutputDir); err != nil {
		return err
	}

	if err := ctx.postValidate(); err != nil {
		return err
	}

	if ctx.GenerateState {
		// Run init and import commands
		meta := command.Meta{
			Ui: &cli.BasicUi{
				Reader:      os.Stdin,
				Writer:      os.Stdout,
				ErrorWriter: logOutput,
			},
			RunningInAutomation: true,
		}

		initCmd := command.InitCommand{Meta: meta}
		var initArgs []string
		if pluginDir := getEnvSettingWithBlankDefault("provider_bin_path"); pluginDir != "" {
			log.Printf("[INFO] plugin dir: '%s'", pluginDir)
			initArgs = append(initArgs, fmt.Sprintf("-plugin-dir=%v", pluginDir))
		}
		initArgs = append(initArgs, *ctx.OutputDir)
		if errCode := initCmd.Run(initArgs); errCode != 0 {
			return nil
		}

		stateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), local.DefaultStateFilename)
		tmpStateOutputFile := fmt.Sprintf("%s%s%s", *ctx.OutputDir, string(os.PathSeparator), defaultTmpStateFile)
		if err := os.RemoveAll(tmpStateOutputFile); err != nil {
			log.Printf("[WARN] unable to delete existing tmp state file %s", tmpStateOutputFile)
			return err
		}

		for _, resource := range ctx.discoveredResources {
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
				fmt.Sprintf("-config=%s", *ctx.OutputDir),
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

	return nil
}

func getDiscoverResourceSteps(ctx *resourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	return getDiscoverResourceWithGraphSteps(ctx)
}

func getDiscoverResourceWithGraphSteps(ctx *resourceDiscoveryContext) ([]resourceDiscoveryStep, error) {
	result := []resourceDiscoveryStep{}

	tenancyResource := &OCIResource{
		compartmentId: ctx.tenancyOcid,
		TerraformResource: TerraformResource{
			id:             ctx.tenancyOcid,
			terraformClass: "oci_identity_tenancy",
			terraformName:  "export",
		},
	}

	for _, mode := range ctx.Services {
		if resourceGraph, exists := tenancyResourceGraphs[mode]; exists {
			result = append(result, &resourceDiscoveryWithGraph{
				root:                      tenancyResource,
				resourceGraph:             resourceGraph,
				resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
			})

			vars["tenancy_ocid"] = fmt.Sprintf("\"%s\"", ctx.tenancyOcid)
			referenceMap[ctx.tenancyOcid] = tfHclVersion.getVarHclString("tenancy_ocid")
		}
	}

	if ctx.CompartmentId == nil || *ctx.CompartmentId == "" {
		*ctx.CompartmentId = ctx.tenancyOcid
	}
	compartmentResource := &OCIResource{
		compartmentId: *ctx.CompartmentId,
		TerraformResource: TerraformResource{
			id:             *ctx.CompartmentId,
			terraformClass: "oci_identity_compartment",
			terraformName:  "export",
		},
	}

	for _, mode := range ctx.Services {
		if resourceGraph, exists := compartmentResourceGraphs[mode]; exists {
			result = append(result, &resourceDiscoveryWithGraph{
				root:                      compartmentResource,
				resourceGraph:             resourceGraph,
				resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{name: mode, ctx: ctx},
			})

			vars["compartment_ocid"] = fmt.Sprintf("\"%s\"", *ctx.CompartmentId)
			referenceMap[*ctx.CompartmentId] = tfHclVersion.getVarHclString("compartment_ocid")
		}
	}

	return result, nil
}

func findResources(ctx *resourceDiscoveryContext, root *OCIResource, resourceGraph TerraformResourceGraph) ([]*OCIResource, error) {
	// findResources will never return error, it will add the errors encountered to the errorList and print those after the discovery finishes
	// If find resources needs to fail in some scenario, this func needs to be modified to return error instead of continuing discovery
	// Errors so far are API errors or the errors when service/feature is not available
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
		results, err := findResourceFn(ctx.clients, &childType, root)
		if err != nil {
			// add error to the errorList and continue discovering rest of the resources
			ctx.errorList = append(ctx.errorList, &ResourceDiscoveryError{childType.resourceClass, root.terraformName, err, &resourceGraph})
			continue
		}

		if childType.processDiscoveredResourcesFn != nil {
			results, err = childType.processDiscoveredResourcesFn(ctx.clients, results)
			if err != nil {
				// add error to the errorList and continue discovering rest of the resources
				ctx.errorList = append(ctx.errorList, &ResourceDiscoveryError{childType.resourceClass, root.terraformName, err, &resourceGraph})
				continue
			}
		}
		foundResources = append(foundResources, results...)

		for _, resource := range results {
			//referenceMap[resource.id] = resource.getHclReferenceIdString()
			if ctx.expectedResourceIds != nil && len(ctx.expectedResourceIds) > 0 {
				if _, shouldExport := ctx.expectedResourceIds[resource.id]; shouldExport {
					resource.omitFromExport = false
					ctx.expectedResourceIds[resource.id] = true
				} else {
					resource.omitFromExport = !childType.alwaysExportable
				}
			}

			subResources, err := findResources(ctx, resource, resourceGraph)
			if err != nil {
				continue
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

func generateProviderFile(outputDir *string) error {
	providerTmpFile := fmt.Sprintf("%s%s%s.tmp", *outputDir, string(os.PathSeparator), providerFile)
	providerOutputFile := fmt.Sprintf("%s%s%s", *outputDir, string(os.PathSeparator), providerFile)
	file, err := os.OpenFile(providerTmpFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("provider oci {\n\tregion = %s\n}\n", tfHclVersion.getVarHclString("region")))
	if err != nil {
		_ = file.Close()
		return err
	}

	if fErr := file.Close(); fErr != nil {
		return fErr
	}

	if err := os.Rename(providerTmpFile, providerOutputFile); err != nil {
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
	terraformReferenceIdString string // syntax independent interpolation- `resource_type.resource_name.id`
	terraformTypeInfo          *TerraformResourceHints
	omitFromExport             bool
}

func (tr *TerraformResource) getHclReferenceIdString() string {
	if tr.terraformReferenceIdString != "" {
		return tfHclVersion.getSingleExpHclString(tr.terraformReferenceIdString)
	}
	return tfHclVersion.getDoubleExpHclString(tr.getTerraformReference(), "id")
}

func (tr *TerraformResource) getTerraformReference() string {
	return fmt.Sprintf("%s.%s", tr.terraformClass, tr.terraformName)
}

func getHCLStringFromMap(builder *strings.Builder, sourceAttributes map[string]interface{}, resourceSchema *schema.Resource, interpolationMap map[string]string, ociRes *OCIResource, attributePrefix string) error {
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
			case InterpolationString:
				builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, v.value))
				continue
			case string:
				if varOverride, exists := interpolationMap[fmt.Sprintf("%v", v)]; exists {
					v = varOverride
					builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, v))
				} else {
					builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, escapeTFStrings(v)))
				}
				continue
			case int, bool, float64:
				builder.WriteString(fmt.Sprintf("%s = \"%v\"\n", tfAttribute, v))
				continue
			case []interface{}:
				switch tfSchema.Type {
				case schema.TypeString:
					if tfAttribute == "delivery_policy" {
						builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, parseDeliveryPolicy(v[0].(interface{}))))
						continue
					}
				case schema.TypeList, schema.TypeSet:
					switch elem := tfSchema.Elem.(type) {
					case *schema.Resource:
						for i, item := range v {
							if val := item.(map[string]interface{}); val != nil {
								builder.WriteString(fmt.Sprintf("%s {\n", tfAttribute))
								attributePrefixForRecursiveCall := attributePrefix
								if attributePrefix == "" {
									attributePrefixForRecursiveCall = fmt.Sprintf("%s[%d]", tfAttribute, i)
								} else {
									attributePrefixForRecursiveCall = fmt.Sprintf("%s.%s[%d]", attributePrefix, tfAttribute, i)
								}
								if err := getHCLStringFromMap(builder, val, elem, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
									return err
								}
								builder.WriteString("}\n")
							}
						}
						continue
					case *schema.Schema, schema.ValueType, InterpolationString:
						builder.WriteString(fmt.Sprintf("%s = [\n", tfAttribute))
						for _, item := range v {
							switch trueListVal := item.(type) {
							case InterpolationString:
								builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, trueListVal.value))
							case string:
								if varOverride, exists := interpolationMap[fmt.Sprintf("%v", trueListVal)]; exists {
									trueListVal = varOverride
									builder.WriteString(fmt.Sprintf("%v,\n", trueListVal))
								} else {
									builder.WriteString(fmt.Sprintf("%q,\n", escapeTFStrings(trueListVal)))
								}
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
						attributePrefixForRecursiveCall := attributePrefix
						if attributePrefix == "" {
							attributePrefixForRecursiveCall = tfAttribute
						} else {
							attributePrefixForRecursiveCall = attributePrefix + "." + tfAttribute
						}
						if err := getHCLStringFromMap(builder, v, nestedResource, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
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
						case InterpolationString:
							builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, mapVal.value))
						case string:
							if varOverride, exists := interpolationMap[fmt.Sprintf("%v", mapVal)]; exists {
								mapVal = varOverride
								builder.WriteString(fmt.Sprintf("\"%s\" = %v\n", mapKey, mapVal))
							} else {
								builder.WriteString(fmt.Sprintf("\"%s\" = %q\n", mapKey, escapeTFStrings(mapVal)))
							}
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
			/* Set missing value if specified in resource hints. This is to avoid plan failure for existing infrastructure.
			This is only done for required attributes as the Optional attributes will not cause plan failure
			We can extend this in future to provide this option to customer to add default values for attributes
			and add this logic to Optional attributes too */

			if ociRes.terraformTypeInfo == nil {
				ociRes.terraformTypeInfo = &TerraformResourceHints{}
			}

			if ociRes.terraformTypeInfo.defaultValuesForMissingAttributes == nil {
				ociRes.terraformTypeInfo.defaultValuesForMissingAttributes = make(map[string]string)
			}
			if tfAttributeVal, exists := ociRes.terraformTypeInfo.defaultValuesForMissingAttributes[tfAttribute]; exists {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, tfAttributeVal))
			} else {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, placeholderValueForMissingAttribute))
			}
			builder.WriteString("\t#Required attribute not found in discovery, placeholder value set to avoid plan failure\n")
			isMissingRequiredAttributes = true

			/* Add missing required attribute to ignorableRequiredMissingAttributes to be generated in lifecycle ignore_changes */
			if ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes == nil {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes = make(map[string]bool)
			}
			if attributePrefix == "" {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes[tfAttribute] = true
			} else {
				ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes[attributePrefix+"."+tfAttribute] = true
			}

		} else if tfSchema.Optional {
			log.Printf("[INFO] Optional TF attribute '%s' not found in source\n", tfAttribute)
			builder.WriteString(fmt.Sprintf("#%s = <<Optional value not found in discovery>>\n", tfAttribute))
		}
	}
	return nil
}

func (resource *OCIResource) hasFreeformTag(tagKey string) bool {
	if freeformTags, exists := resource.sourceAttributes["freeform_tags"]; exists {
		if freeformTagMap, ok := freeformTags.(map[string]interface{}); ok {
			if _, hasFreeFormTag := freeformTagMap[tagKey]; hasFreeFormTag {
				return true
			}
		}
	}

	return false
}

func (resource *OCIResource) hasDefinedTag(tagKey string, tagValue string) bool {
	if definedTags, exists := resource.sourceAttributes["defined_tags"]; exists {
		if definedTagMap, ok := definedTags.(map[string]interface{}); ok {
			if definedTagValue, hasDefinedTag := definedTagMap[tagKey]; hasDefinedTag {
				return definedTagValue == tagValue
			}
		}
	}

	return false
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
	if err := getHCLStringFromMap(builder, ociRes.sourceAttributes, resourceSchema, interpolationMap, ociRes, ""); err != nil {
		return err
	}

	if ociRes.terraformTypeInfo != nil && len(ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes) > 0 {
		builder.WriteString("\n# Required attributes that were not found in discovery have been added to " +
			"lifecycle ignore_changes")
		builder.WriteString("\n# This is done to avoid terraform plan failure for the existing infrastructure")
		builder.WriteString("\nlifecycle {\n" +
			"ignore_changes = [")

		missingAttributes := make([]string, 0, len(ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes))

		for attribute := range ociRes.terraformTypeInfo.ignorableRequiredMissingAttributes {
			missingAttributes = append(missingAttributes, tfHclVersion.getReference(attribute))
		}
		builder.WriteString(strings.Join(missingAttributes, ","))

		builder.WriteString("]\n" +
			"}\n")
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

	if !tfMeta.DiscoversWithSingularDatasource() {
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
	} else if d.Id() != "" {
		// Result is from a singular datasource that hasn't had its state voided (hence d.Id() is non-empty)
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
	} else {
		log.Printf("[DEBUG] singular data source not able to find resource")
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

// This function should only be used to escape TF-characters in strings
func escapeTFStrings(val string) string {
	val = strings.ReplaceAll(val, "%{", "%%{")
	val = strings.ReplaceAll(val, "${", "$${")
	return val
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

	rootCompartment, err := exportConfigProvider.TenancyOCID()
	if err != nil {
		return nil, err
	}
	req.CompartmentId = &rootCompartment

	recursiveSearch := true
	req.CompartmentIdInSubtree = &recursiveSearch

	for {
		resp, err := clients.identityClient().ListCompartments(context.Background(), req)
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

func readEnvironmentVars(d *schema.ResourceData) error {

	if err := d.Set(authAttrName, getProviderEnvSettingWithDefault(authAttrName, authAPIKeySetting)); err != nil {
		return err
	}
	if err := d.Set(configFileProfileAttrName, getProviderEnvSettingWithDefault(configFileProfileAttrName, "")); err != nil {
		return err
	}
	if region := getProviderEnvSettingWithDefault(regionAttrName, ""); region != "" {
		if err := d.Set(regionAttrName, region); err != nil {
			return err
		}
	}

	if tenancyOcid := getProviderEnvSettingWithDefault(tenancyOcidAttrName, ""); tenancyOcid != "" {
		if err := d.Set(tenancyOcidAttrName, tenancyOcid); err != nil {
			return err
		}
	}

	if userOcid := getProviderEnvSettingWithDefault(userOcidAttrName, ""); userOcid != "" {
		if err := d.Set(userOcidAttrName, userOcid); err != nil {
			return err
		}
	}
	if fingerprint := getProviderEnvSettingWithDefault(fingerprintAttrName, ""); fingerprint != "" {
		if err := d.Set(fingerprintAttrName, fingerprint); err != nil {
			return err
		}
	}
	if privateKey := getProviderEnvSettingWithDefault(privateKeyAttrName, ""); privateKey != "" {
		if err := d.Set(privateKeyAttrName, privateKey); err != nil {
			return err
		}
	}
	if privateKeyPath := getProviderEnvSettingWithDefault(privateKeyPathAttrName, ""); privateKeyPath != "" {
		if err := d.Set(privateKeyPathAttrName, privateKeyPath); err != nil {
			return err
		}
	}
	if privateKeyPassword := getProviderEnvSettingWithDefault(privateKeyPasswordAttrName, ""); privateKeyPassword != "" {
		if err := d.Set(privateKeyPasswordAttrName, privateKeyPassword); err != nil {
			return err
		}
	}
	return nil
}
