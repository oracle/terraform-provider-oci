package commonexport

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func init() {
	RegisterAvailabilityResourceGraph()
}

func (tr *TerraformResource) GetTerraformReference() string {
	return fmt.Sprintf("%s.%s", tr.TerraformClass, tr.TerraformName)
}

func (resource *OCIResource) HasFreeformTag(tagKey string) bool {
	if freeformTags, exists := resource.SourceAttributes["freeform_tags"]; exists {
		if freeformTagMap, ok := freeformTags.(map[string]interface{}); ok {
			if _, hasFreeFormTag := freeformTagMap[tagKey]; hasFreeFormTag {
				return true
			}
		}
	}

	return false
}

func (resource *OCIResource) HasDefinedTag(tagKey string, tagValue string) bool {
	if definedTags, exists := resource.SourceAttributes["defined_tags"]; exists {
		if definedTagMap, ok := definedTags.(map[string]interface{}); ok {
			if definedTagValue, hasDefinedTag := definedTagMap[tagKey]; hasDefinedTag {
				return definedTagValue == tagValue
			}
		}
	}

	return false
}

func (ociRes *OCIResource) GetHCLString(builder *strings.Builder, interpolationMap map[string]string) error {
	// Remove any potential cyclical references from the interpolation map
	selfReference := ociRes.GetTerraformReference()
	resourceInterpolationMap := map[string]string{}
	for value, interpolation := range interpolationMap {
		if !strings.Contains(interpolation, selfReference) {
			resourceInterpolationMap[value] = interpolation
		}
	}

	if ociRes.GetHclStringFn != nil {
		return ociRes.GetHclStringFn(builder, ociRes, resourceInterpolationMap)
	}
	return GetHclStringFromGenericMap(builder, ociRes, resourceInterpolationMap)
}

func (tr *TerraformResource) GetHclReferenceIdString() string {
	if tr.TerraformReferenceIdString != "" {
		return TfHclVersionvar.GetSingleExpHclString(tr.TerraformReferenceIdString)
	}
	return TfHclVersionvar.GetDoubleExpHclString(tr.GetTerraformReference(), "id")
}

func GetHCLStringFromMap(builder *strings.Builder, sourceAttributes map[string]interface{}, resourceSchema *schema.Resource, interpolationMap map[string]string, ociRes *OCIResource, attributePrefix string) error {
	sortedKeys := make([]string, len(resourceSchema.Schema))
	cnt := 0
	for k := range resourceSchema.Schema {
		sortedKeys[cnt] = k
		cnt++
	}
	sort.Strings(sortedKeys)

	for _, tfAttribute := range sortedKeys {
		tfSchema := resourceSchema.Schema[tfAttribute]
		if tfSchema.Deprecated != "" || (!tfSchema.Required && !tfSchema.Optional) {
			continue
		}

		if attributeVal, exists := sourceAttributes[tfAttribute]; exists {
			utils.Debugf("Writing attribute %s and value %s", tfAttribute, attributeVal)
			switch v := attributeVal.(type) {
			case InterpolationString:
				if ok := FailedResourceReferenceSet[v.ResourceReference]; ok {
					builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, v.Value))
				} else {
					builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, v.Interpolation))
				}
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
						builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, ParseDeliveryPolicy(v[0].(interface{}))))
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
								if err := GetHCLStringFromMap(builder, val, elem, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
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
								if ok := FailedResourceReferenceSet[trueListVal.ResourceReference]; ok {
									builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, trueListVal.Value))
								} else {
									builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, trueListVal.Interpolation))
								}
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
						if err := GetHCLStringFromMap(builder, v, nestedResource, interpolationMap, ociRes, attributePrefixForRecursiveCall); err != nil {
							return err
						}
						builder.WriteString("}\n")
						continue
					}
					return fmt.Errorf("[ERROR] sourceAttribute '%s', tfAttribute '%s': Nested resource type mismatch", tfAttribute, tfAttribute)
				case schema.TypeMap:
					builder.WriteString(fmt.Sprintf("%s = {\n", tfAttribute))

					keys := utils.GetSortedKeys(v)
					for _, mapKey := range keys {
						switch mapVal := v[mapKey].(type) {
						case InterpolationString:
							if ok := FailedResourceReferenceSet[mapVal.ResourceReference]; ok {
								builder.WriteString(fmt.Sprintf("%s = %q\n", tfAttribute, mapVal.Value))
							} else {
								builder.WriteString(fmt.Sprintf("%s = %v\n", tfAttribute, mapVal.Interpolation))
							}
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
				utils.Logf("[INFO] TF attribute '%s' is nil in source\n", tfAttribute)
				if !tfSchema.Required {
					continue
				}
			default:
				utils.Logf("[WARN] TF attribute '%s' is unknown type in source\n", tfAttribute)
			}
		}

		if tfSchema.Required {
			utils.Logf("[WARN] Required TF attribute '%s' not found in source\n", tfAttribute)
			/* Set missing value if specified in resource hints. This is to avoid plan failure for existing infrastructure.
			This is only done for required attributes as the Optional attributes will not cause plan failure
			We can extend this in future to provide this option to customer to add default values for attributes
			and add this logic to Optional attributes too */

			if ociRes.TerraformTypeInfo == nil {
				ociRes.TerraformTypeInfo = &TerraformResourceHints{}
			}

			if ociRes.TerraformTypeInfo.DefaultValuesForMissingAttributes == nil {
				ociRes.TerraformTypeInfo.DefaultValuesForMissingAttributes = make(map[string]interface{})
			}
			if tfAttributeVal, exists := ociRes.TerraformTypeInfo.DefaultValuesForMissingAttributes[tfAttribute]; exists {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, tfAttributeVal))
			} else {
				builder.WriteString(fmt.Sprintf("%s = %q", tfAttribute, globalvar.PlaceholderValueForMissingAttribute))
			}
			builder.WriteString("\t#Required attribute not found in discovery, placeholder value set to avoid plan failure\n")
			IsMissingRequiredAttributes = true

			/* Add missing required attribute to ignorableRequiredMissingAttributes to be generated in lifecycle ignore_changes */
			if ociRes.TerraformTypeInfo.IgnorableRequiredMissingAttributes == nil {
				ociRes.TerraformTypeInfo.IgnorableRequiredMissingAttributes = make(map[string]bool)
			}
			if attributePrefix == "" {
				ociRes.TerraformTypeInfo.IgnorableRequiredMissingAttributes[tfAttribute] = true
			} else {
				ociRes.TerraformTypeInfo.IgnorableRequiredMissingAttributes[attributePrefix+"."+tfAttribute] = true
			}

		} else if tfSchema.Optional {
			utils.Logf("[INFO] Optional TF attribute '%s' not found in source\n", tfAttribute)
			builder.WriteString(fmt.Sprintf("#%s = <<Optional value not found in discovery>>\n", tfAttribute))
		}
	}
	return nil
}

func (args *ExportCommandArgs) FinalizeServices(ctx *ResourceDiscoveryContext) {
	if len(args.Services) == 0 {
		args.Services = CompartmentScopeServices

		/*
			If compartmentId provided is not provided or is a root compartment then discover tenancy scope resources too
		*/
		if args.CompartmentId != nil && (*args.CompartmentId == "" || *args.CompartmentId == ctx.TenancyOcid) {
			args.Services = append(args.Services, TenancyScopeServices...)
		}
	}

	// Dedupes possible repeating services from command line and sorts them
	finalServices := []string{}
	serviceSet := ConvertStringSliceToSet(args.Services, true)
	excludeServicesSet := ConvertStringSliceToSet(args.ExcludeServices, true)
	for service := range serviceSet {
		if _, exists := excludeServicesSet[service]; !exists {
			finalServices = append(finalServices, service)
		}
	}
	args.Services = finalServices
	sort.Strings(args.Services)
}

func ConvertStringSliceToSet(slice []string, omitEmptyStrings bool) map[string]bool {
	result := map[string]bool{}
	for _, item := range slice {
		if omitEmptyStrings && item == "" {
			continue
		}
		result[item] = false
	}
	return result
}

func escapeTFStrings(val string) string {
	val = strings.ReplaceAll(val, "%{", "%%{")
	val = strings.ReplaceAll(val, "${", "$${")
	return val
}
func ParseDeliveryPolicy(policy interface{}) string {
	backoffRetryPolicy := policy.(map[string]interface{})["backoff_retry_policy"].([]interface{})
	maxRetryDuration := backoffRetryPolicy[0].(map[string]interface{})["max_retry_duration"]
	policyType := backoffRetryPolicy[0].(map[string]interface{})["policy_type"]
	return fmt.Sprintf("{\"backoffRetryPolicy\":{\"maxRetryDuration\":%v,\"policyType\":\"%v\"}}", maxRetryDuration, policyType)
}

// Validate export command arguments and returns nil if there are no issues
func (args *ExportCommandArgs) Validate() error {
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

	if args.Parallelism < 1 {
		return fmt.Errorf("[ERROR] invalid value for arument parallelism, specify a value >= 1")
	}

	// validate and extract variables_resource_level
	if args.VarsExportResourceLevel != nil {
		VarsExportForResourceLevel, err = extractVarsExportResourceLevel(args.VarsExportResourceLevel)
		if err != nil {
			utils.Logln(err.Error())
			return err
		}
	}

	// validate and extract variables_global_level
	if args.VarExportGlobalLevel != nil {
		VarsExportForGlobalLevel, err = extractVarsExportGlobalLevel(args.VarExportGlobalLevel)
		if err != nil {
			utils.Logln(err.Error())
			return err
		}
	}
	return nil
}

func (ctx *ResourceDiscoveryContext) GetResourceHint(resourceClass string) (*TerraformResourceHints, error) {
	if hints, exists := ctx.ResourceHintsLookup[resourceClass]; exists {
		return hints, nil
	}

	// If no resource hint could be found, just return a simple hint for now to unblock
	return nil, fmt.Errorf("[ERROR] resource type '%s' is not supported by resource discovery", resourceClass)
}

func (ctx *ResourceDiscoveryContext) AddErrorToList(error *ResourceDiscoveryError) {
	ctx.CtxLock.Lock()
	defer ctx.CtxLock.Unlock()
	ctx.ErrorList.Errors = append(ctx.ErrorList.Errors, error)

}

func (ctx *ResourceDiscoveryContext) PostValidate() {
	// Check that all expected resource IDs were found, if any were given
	var missingResourceIds []string
	for resourceId, found := range ctx.ExpectedResourceIds {
		if !found {
			missingResourceIds = append(missingResourceIds, resourceId)
		}
	}

	if len(missingResourceIds) > 0 {
		ctx.SummaryStatements = append(ctx.SummaryStatements, "")
		ctx.SummaryStatements = append(ctx.SummaryStatements, "Warning: The following resource IDs were not found.")
		for _, resourceId := range missingResourceIds {
			ctx.SummaryStatements = append(ctx.SummaryStatements, fmt.Sprintf("- %s", resourceId))
		}
		rdError := &ResourceDiscoveryError{
			"",
			"",
			fmt.Errorf("[ERROR] one or more expected resource ids were not found"),
			nil}

		ctx.AddErrorToList(rdError)
	}
}

func (ctx *ResourceDiscoveryContext) PrintSummary() {

	ctx.SummaryStatements = append(ctx.SummaryStatements, "=== COMPLETED ===")

	for _, statement := range ctx.SummaryStatements {
		utils.Logln(utils.Green(statement))
	}
	utils.Logln(utils.Green("========= PERFORMANCE SUMMARY New Branch=========="))
	utils.Logln(utils.Green(fmt.Sprintf("Total resources: %v", len(ctx.DiscoveredResources))))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken for discovering all services: %v", ctx.TimeTakenToDiscover)))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken for generating state of all services: %v", ctx.TimeTakenToGenerateState)))
	utils.Logln(utils.Green(fmt.Sprintf("Total time taken by entire export: %v", ctx.TimeTakenForEntireExport)))
}

func (ctx *ResourceDiscoveryContext) PrintErrors() ([]string, []string) {
	utils.Logln(utils.Yellow("\n\n[WARN] Resource discovery finished with errors listed below:\n"))
	var notDiscoveredParentResources []string
	var notDiscoveredChildResources []string
	for _, resourceDiscoveryError := range ctx.ErrorList.Errors {
		if resourceDiscoveryError.ResourceType == "" || ctx.TargetSpecificResources {
			utils.Logln(utils.Yellow(resourceDiscoveryError.Error.Error()))

		} else if resourceDiscoveryError.ParentResource == "export" {
			utils.Logln(utils.Yellow(fmt.Sprintf("Error discovering `%s` resources: %s", resourceDiscoveryError.ResourceType, resourceDiscoveryError.Error.Error())))
			partiallyResourcesDiscovered := "Error discovering " + resourceDiscoveryError.ResourceType + " resources: " + resourceDiscoveryError.Error.Error()
			notDiscoveredParentResources = append(notDiscoveredParentResources, partiallyResourcesDiscovered)

		} else {
			utils.Logln(utils.Yellow(fmt.Sprintf("Error discovering `%s` resources for %s: %s", resourceDiscoveryError.ResourceType, resourceDiscoveryError.ParentResource, resourceDiscoveryError.Error.Error())))
		}
		/* log child resources if exist and were not discovered because of error in parent resource discovery*/
		if resourceDiscoveryError.ResourceGraph != nil && !ctx.TargetSpecificResources {
			var notFoundChildren []string
			getNotFoundChildren(resourceDiscoveryError.ResourceType, resourceDiscoveryError.ResourceGraph, &notFoundChildren)
			if len(notFoundChildren) > 0 {
				utils.Logln(utils.Yellow(fmt.Sprintf("\tFollowing child resources were also not discovered due to parent error: %v", strings.Join(notFoundChildren, ", "))))
				notFoundChildResources := "\tFollowing child resources were also not discovered due to parent error: " + strings.Join(notFoundChildren, ", ")
				notDiscoveredChildResources = append(notDiscoveredChildResources, notFoundChildResources)
			}
		}
	}

	return notDiscoveredParentResources, notDiscoveredChildResources
}

func (h *TerraformResourceHints) DiscoversWithSingularDatasource() bool {
	return h.DatasourceItemsAttr == ""
}

func FindResourcesGeneric(ctx *ResourceDiscoveryContext, tfMeta *TerraformResourceAssociation, parent *OCIResource, resourceGraph *TerraformResourceGraph) ([]*OCIResource, error) {
	results := []*OCIResource{}
	clients := ctx.Clients

	utils.Logf("[INFO] discovering resources with data source '%s'\n", tfMeta.DatasourceClass)
	utils.Debugf("[DEBUG] discovering resources with data source '%s'\n", tfMeta.DatasourceClass)
	datasource := DatasourcesMap[tfMeta.DatasourceClass]
	d := datasource.TestResourceData()
	d.Set("compartment_id", parent.CompartmentId)

	for queryAttributeName, queryValue := range tfMeta.DatasourceQueryParams {
		utils.Logf("[INFO] adding datasource query attribute '%s' from parent attribute '%s'\n", queryAttributeName, queryValue)
		utils.Debugf("[DEBUG] adding datasource query attribute '%s' from parent attribute '%s'\n", queryAttributeName, queryValue)

		if queryValue == "" || queryValue == "id" {
			d.Set(queryAttributeName, parent.Id)
		} else if strings.HasPrefix(queryValue, "'") && strings.HasSuffix(queryValue, "'") { // Anything encapsulated in ' ' means to use the literal value
			d.Set(queryAttributeName, queryValue[1:len(queryValue)-1])
		} else if val, ok := parent.SourceAttributes[queryValue]; ok {
			d.Set(queryAttributeName, val)
		} else {
			utils.Logf("[WARN] no attribute '%s' found in parent '%s', returning no results for this resource\n", queryValue, parent.GetTerraformReference())
			return results, nil
		}
	}
	utils.Debugf("[DEBUG] Initiating GET Datasource Call for %s compartment %s", tfMeta.DatasourceClass, parent.CompartmentId)
	if err := datasource.Read(d, clients); err != nil {
		utils.Debugf("[DEBUG] GET Datasource Call Failure for %s compartment %s\nError: %s", tfMeta.DatasourceClass, parent.CompartmentId, err)
		return results, err
	}
	utils.Debugf("[DEBUG] GET Datasource Call Success for %s compartment %s", tfMeta.DatasourceClass, parent.CompartmentId)
	if !tfMeta.DiscoversWithSingularDatasource() {
		// Results are from a plural datasource
		itemSchema := datasource.Schema[tfMeta.DatasourceItemsAttr]
		elemResource, ok := itemSchema.Elem.(*schema.Resource)
		if !ok {
			return results, fmt.Errorf("[ERROR] element schema is not of a resource")
		}
		datasourceItemsAttribute := tfMeta.DatasourceItemsAttr

		if tfMeta.IsDatasourceCollection {
			collectionItemSchema := elemResource.Schema["items"]

			elemResource, ok = collectionItemSchema.Elem.(*schema.Resource)
			if !ok {
				return results, fmt.Errorf("[ERROR] collection element schema is not of a resource")
			}
			datasourceItemsAttribute = tfMeta.DatasourceItemsAttr + ".0.items"
		}

		foundItems, _ := d.GetOkExists(datasourceItemsAttribute)
		for idx, item := range foundItems.([]interface{}) {
			if itemMap, ok := item.(map[string]interface{}); ok {
				if state, exists := itemMap["state"].(string); exists && len(tfMeta.DiscoverableLifecycleStates) > 0 {
					discoverable := false

					if utils.GetEnvSettingWithBlankDefault(globalvar.DiscoverAllStatesEnv) == "1" {
						utils.Logf("Skipping Lifecycle State Check as TF_DISCOVER_ALL_STATES is set")
						discoverable = true
					} else {
						for _, val := range tfMeta.DiscoverableLifecycleStates {
							if strings.EqualFold(state, val) {
								discoverable = true
								break
							}
						}
					}

					if !discoverable {
						continue
					}
				}
			}
			var resource *OCIResource
			var err error
			if tfMeta.RequireResourceRefresh {
				resourceSchema := ResourcesMap[tfMeta.ResourceClass]
				r := resourceSchema.TestResourceData()

				// Use resource to fill in all attributes (likely because the datasource doesn't return complete info)
				if tfMeta.GetIdFn != nil {
					tmpResource, err := generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", datasourceItemsAttribute, idx), tfMeta, parent)
					if err != nil {
						rdError := &ResourceDiscoveryError{
							ResourceType:   tfMeta.ResourceClass,
							ParentResource: parent.TerraformName,
							Error:          fmt.Errorf("[ERROR] error generating temporary resource from resource data returned in list datasource read: %v ", err),
							ResourceGraph:  resourceGraph}
						ctx.AddErrorToList(rdError)
						continue
					}

					itemId, err := tfMeta.GetIdFn(tmpResource)
					if err != nil {
						rdError := &ResourceDiscoveryError{
							ResourceType:   tfMeta.ResourceClass,
							ParentResource: parent.TerraformName,
							Error:          fmt.Errorf("[ERROR] failed to get a composite ID for the resource: %v ", err),
							ResourceGraph:  resourceGraph}
						ctx.AddErrorToList(rdError)
						continue
					}
					r.SetId(itemId)
				} else if idSchema, exists := elemResource.Schema["id"]; exists && idSchema.Type == schema.TypeString {
					itemId := item.(map[string]interface{})["id"]
					r.SetId(itemId.(string))
				} else {
					rdError := &ResourceDiscoveryError{
						ResourceType:   tfMeta.ResourceClass,
						ParentResource: parent.TerraformName,
						Error: fmt.Errorf("[ERROR] elements in datasource '%s' are missing an 'id' field and is unable to generate an id",
							tfMeta.DatasourceClass),
						ResourceGraph: resourceGraph}
					ctx.AddErrorToList(rdError)
					continue
				}

				if err = resourceSchema.Read(r, clients); err != nil {
					rdError := &ResourceDiscoveryError{
						ResourceType:   tfMeta.ResourceClass,
						ParentResource: parent.TerraformName,
						Error:          fmt.Errorf("[ERROR] error refreshing resource using resource read: %v ", err),
						ResourceGraph:  resourceGraph}
					ctx.AddErrorToList(rdError)
					continue
				}
				// If state was voided because of error in Read (r.Id() is empty)
				if r.Id() == "" {
					rdError := &ResourceDiscoveryError{
						ResourceType:   tfMeta.ResourceClass,
						ParentResource: parent.TerraformName,
						Error:          fmt.Errorf("[ERROR] error refreshing resource using resource read, state voided"),
						ResourceGraph:  resourceGraph}
					ctx.AddErrorToList(rdError)
					continue
				}
				resource, err = generateOciResourceFromResourceData(r, r, resourceSchema.Schema, "", tfMeta, parent)
				if err != nil {
					rdError := &ResourceDiscoveryError{
						ResourceType:   tfMeta.ResourceClass,
						ParentResource: parent.TerraformName,
						Error:          fmt.Errorf("[ERROR] error generating resource from resource data returned in resource read: %v ", err),
						ResourceGraph:  resourceGraph}
					ctx.AddErrorToList(rdError)
					continue
				}
			} else {
				resource, err = generateOciResourceFromResourceData(d, item, elemResource.Schema, fmt.Sprintf("%s.%v", datasourceItemsAttribute, idx), tfMeta, parent)
				if err != nil {
					rdError := &ResourceDiscoveryError{
						ResourceType:   tfMeta.ResourceClass,
						ParentResource: parent.TerraformName,
						Error:          fmt.Errorf("[ERROR] error generating resource from resource data returned in list datasource read: %v ", err),
						ResourceGraph:  resourceGraph}
					ctx.AddErrorToList(rdError)
					continue
				}
			}

			if resource.TerraformName, err = GenerateTerraformNameFromResource(resource.SourceAttributes, elemResource.Schema); err != nil {
				resource.TerraformName = fmt.Sprintf("%s_%s", parent.TerraformName, tfMeta.ResourceAbbreviation)
				resource.TerraformName = CheckDuplicateResourceName(resource.TerraformName)
			}

			results = append(results, resource)
		}
	} else if d.Id() != "" {
		// Result is from a singular datasource that hasn't had its state voided (hence d.Id() is non-empty)
		resource, err := generateOciResourceFromResourceData(d, d, datasource.Schema, "", tfMeta, parent)
		if err != nil {
			return results, err
		}

		if resource.TerraformName, err = GenerateTerraformNameFromResource(resource.SourceAttributes, datasource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.TerraformName, tfMeta.ResourceAbbreviation)
			resource.TerraformName = CheckDuplicateResourceName(resource.TerraformName)
		}

		discoverable := true
		if state, ok := resource.SourceAttributes["state"]; ok && len(tfMeta.DiscoverableLifecycleStates) > 0 {
			discoverable = false

			if utils.GetEnvSettingWithBlankDefault(globalvar.DiscoverAllStatesEnv) == "1" {
				utils.Logf("Skipping Lifecycle State Check as TF_DISCOVER_ALL_STATES is set")
				discoverable = true
			} else {
				for _, val := range tfMeta.DiscoverableLifecycleStates {
					if strings.EqualFold(state.(string), val) {
						discoverable = true
						break
					}
				}
			}
		}

		if discoverable {
			results = append(results, resource)
		}
	} else {
		utils.Debugf("[DEBUG] singular data source not able to find resource")
	}

	return results, nil
}

// This function attempts to convert resource data items to a map representation that omits attributes where no value was set.
func ConvertDatasourceItemToMap(d *schema.ResourceData, itemPrefix string, itemSchema map[string]*schema.Schema) (map[string]interface{}, error) {
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
						resourceList[idx], _ = ConvertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, idx), v.Schema)
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
						resourceList[idx], _ = ConvertDatasourceItemToMap(d, fmt.Sprintf("%s.%v", key, itemHashCode), v.Schema)
					}
					result[attributeKey] = resourceList
				}
			}
		}
	}

	return result, nil
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
	resourceMap, err := ConvertDatasourceItemToMap(d, itemPrefix, resourceSchema)
	if err != nil {
		return nil, err
	}

	resource := &OCIResource{
		CompartmentId:    parent.CompartmentId,
		SourceAttributes: resourceMap,
		RawResource:      rawResource,
		TerraformResource: TerraformResource{
			TerraformClass:    tfMeta.ResourceClass,
			TerraformTypeInfo: tfMeta.TerraformResourceHints,
		},
		GetHclStringFn: GetHclStringFromGenericMap,
		Parent:         parent,
	}

	if tfMeta.GetIdFn != nil {
		if customId, err := tfMeta.GetIdFn(resource); err == nil {
			resource.Id = customId
		}
	} else if resourceId, resourceIdExists := resourceMap["id"]; resourceIdExists {
		resource.Id = resourceId.(string)
	}

	if resource.Id == "" {
		resource.Id = d.Id()
	}

	if tfMeta.GetHCLStringOverrideFn != nil {
		resource.GetHclStringFn = tfMeta.GetHCLStringOverrideFn
	}

	return resource, nil
}

func GetOciResource(d *schema.ResourceData, resourceSchema map[string]*schema.Schema, compartmentId string, resourceHint *TerraformResourceHints, resourceId string) (*OCIResource, error) {
	resourceMap, err := ConvertDatasourceItemToMap(d, "", resourceSchema)
	if err != nil {
		return nil, err
	}

	resource := &OCIResource{
		CompartmentId:    compartmentId,
		SourceAttributes: resourceMap,
		RawResource:      d,
		TerraformResource: TerraformResource{
			TerraformClass:    resourceHint.ResourceClass,
			TerraformTypeInfo: resourceHint,
		},
		GetHclStringFn: GetHclStringFromGenericMap,
	}

	if resourceId != "" {
		resource.Id = resourceId
	}

	if resource.Id == "" {
		resource.Id = d.Id()
	}

	return resource, nil
}

func getNotFoundChildren(parent string, resourceGraph *TerraformResourceGraph, children *[]string) {
	childResources, exists := (*resourceGraph)[parent]
	if exists {
		for _, child := range childResources {
			*children = append(*children, child.ResourceClass)
			// Avoid recursion if a resource can be nested within itself e.g. compartments
			if child.ResourceClass != parent {
				getNotFoundChildren(child.ResourceClass, resourceGraph, children)
			}
		}
	}
}

func RegisterCompartmentGraphs(servicename string, graph TerraformResourceGraph) {
	if oci_common.CheckForEnabledServices(utils.GetSDKServiceName(servicename)) {
		CompartmentResourceGraphs[servicename] = graph
	}
}

func RegisterTenancyGraphs(servicename string, graph TerraformResourceGraph) {
	if oci_common.CheckForEnabledServices(utils.GetSDKServiceName(servicename)) {
		TenancyResourceGraphs[servicename] = graph
	}
}

func RegisterRelatedResourcesGraph(resourceName string, association []TerraformResourceAssociation) {
	ExportRelatedResourcesGraph[resourceName] = association
}

func BuildAvailabilityResourceGraph(resourceName string, association []TerraformResourceAssociation) {

	if len(availabilityDomainResourceGraph[resourceName]) != 0 {
		for _, val := range association {
			availabilityDomainResourceGraph[resourceName] = append(availabilityDomainResourceGraph[resourceName], val)
		}
	} else {
		availabilityDomainResourceGraph[resourceName] = association
	}
}

func RegisterAvailabilityResourceGraph() {
	RegisterCompartmentGraphs("availability_domain", availabilityDomainResourceGraph)
}

func ConvertResourceDataToMap(schemaMap map[string]*schema.Schema, d *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	for key := range schemaMap {
		if val, ok := d.GetOkExists(key); ok {
			result[key] = val
		}
	}

	return result
}

var GenerateTerraformNameFromResource = func(resourceAttributes map[string]interface{}, resourceSchema map[string]*schema.Schema) (string, error) {
	possibleNameAttributes := []string{
		"display_name",
		"name",
	}

	for _, nameAttribute := range possibleNameAttributes {
		if nameSchema, hasNameAttr := resourceSchema[nameAttribute]; hasNameAttr && nameSchema.Type == schema.TypeString {
			if value, exists := resourceAttributes[nameAttribute]; exists {
				terraformName := getNormalizedTerraformName(value.(string))
				terraformName = CheckDuplicateResourceName(terraformName)
				return terraformName, nil
			}
		}
	}

	return "", fmt.Errorf("unable to find a suitable name from the resource attributes")
}

var CheckDuplicateResourceName = func(terraformName string) string {
	ResourceNameCountLock.Lock()
	defer ResourceNameCountLock.Unlock() // Ensure the lock is released even if a panic occurs

	originalName := terraformName
	utils.Logf("[INFO] Checking Duplicate Resource Name for %s", originalName)
	// Check if resource already exists
	for {
		if _, exists := ResourceNameCount[terraformName]; !exists {
			utils.Logf("[INFO] Exiting Duplicate resource name for %s", originalName)
			break
		}
		count := ResourceNameCount[originalName]
		ResourceNameCount[originalName] = count + 1
		terraformName = fmt.Sprintf("%s_%d", originalName, count)
	}

	ResourceNameCount[terraformName] = 1
	utils.Logf("[INFO] Returning Handled Duplicate resource name for %s as %s", originalName, terraformName)
	return terraformName
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

func GetValidUniqueTerraformName(terraformName string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\-\_]+`)
	terraformName = reg.ReplaceAllString(terraformName, "-")
	terraformName = CheckDuplicateResourceName(terraformName)

	return terraformName
}
