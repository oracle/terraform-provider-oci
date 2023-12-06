package commonexport

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func exportAttributeAsVariable(sourceAttributes map[string]interface{}, resourceType string, resourceName string, interpolationMap map[string]string) error {

	// handle user input both flags
	if len(VarsExportForResourceLevel) > 0 && len(VarsExportForGlobalLevel) > 0 {
		return exportAttributeForResourceAndGlobalLevel(sourceAttributes, resourceType, resourceName, VarsExportForResourceLevel, VarsExportForGlobalLevel, interpolationMap)
	}

	// export attribute per resource level
	if len(VarsExportForResourceLevel) > 0 {
		return exportAttributeForResourceLevel(sourceAttributes, resourceType, resourceName, VarsExportForResourceLevel, interpolationMap)
	}

	// export attribute as global level
	if len(VarsExportForGlobalLevel) > 0 {
		return exportAttributeForGlobalLevel(sourceAttributes, resourceName, VarsExportForGlobalLevel, interpolationMap)
	}

	// if non of the flags provided, export variable global from default list
	return exportAttributeFromDefaultList(globalvar.DefaultListAttributeExportAsVariable, sourceAttributes, resourceName, interpolationMap)
}

/* Functions for handling variables_resource_level */
// Return map of resource type and attribute
func extractVarsExportResourceLevel(exportVars []string) (map[string][]string, error) {
	result := map[string][]string{}
	//
	for _, item := range exportVars {
		if !strings.Contains(item, globalvar.DotDelimiter) {
			return nil, fmt.Errorf("[ERROR] variables_resource_level is in wrong format of resourceType.attribute: %v", item)
		}
		resourceTypeAndAttribute := strings.Split(item, globalvar.DotDelimiter)
		if len(resourceTypeAndAttribute) > 2 {
			// TODO: handle nested attributes later
			return nil, fmt.Errorf("[ERROR] variables_resource_level only support top level attribute following format resourceType.attribute: %v", item)
		}
		//isRdSupport, err := resourcediscovery.isResourceSupportImport(resourceTypeAndAttribute[0])
		//if !isRdSupport || err != nil {
		//	return nil, fmt.Errorf("[ERROR] this resource is incorrect or not supported by Resource Discovery: %v", item)
		//}
		// Assuming variables_resource_level for top level attribute with format resourceType.attribute
		result[resourceTypeAndAttribute[0]] = append(result[resourceTypeAndAttribute[0]], resourceTypeAndAttribute[1])
	}
	return result, nil
}

func exportAttributeForResourceLevel(sourceAttributes map[string]interface{}, resourceType string, resourceName string, varsExportResourceLevel map[string][]string, interpolationMap map[string]string) error {
	if attributeList, exist := varsExportResourceLevel[resourceType]; exist {
		for _, attribute := range attributeList {
			utils.Debugf("[DEBUG] Exporting attribute %s of resource %s, sourceAttributes[attribute]: %v", attribute, resourceName, sourceAttributes[attribute])
			if _, exist = sourceAttributes[attribute]; exist {
				attributeVal := sourceAttributes[attribute].(string)
				variableName := utils.GetVarNameFromAttributeOfResources(attribute, resourceType, resourceName)
				utils.Debugf("[DEBUG] Exporting attribute %s of resource %s with value %s and variableName: %s", attribute, resourceName, attributeVal, variableName)
				Vars[variableName] = fmt.Sprintf("\"%s\"", attributeVal)
				interpolationMap[attributeVal] = TfHclVersionvar.GetVarHclString(variableName)
			}
		}
	}
	return nil
}

/* Functions for handling variables_global_level */
// Return list of attributes will be encapsulate as global variable
func extractVarsExportGlobalLevel(attributes []string) ([]string, error) {
	var result []string
	for _, attr := range attributes {
		// TODO: Handle nested attribute
		if strings.Contains(attr, globalvar.DotDelimiter) {
			return nil, fmt.Errorf("[ERROR] variables_global_level only support top level attribute: %v", attr)
		}
		result = append(result, attr)
	}
	return result, nil
}

func exportAttributeForGlobalLevel(sourceAttributes map[string]interface{}, resourceName string, varsExportGlobalLevel []string, interpolationMap map[string]string) error {
	for _, tfAttribute := range varsExportGlobalLevel {
		if attributeVal, exist := sourceAttributes[tfAttribute]; exist {
			utils.Debugf("[DEBUG] Exporting attribute %s of resource %s", tfAttribute, resourceName)
			variableName := getVarNameFromAttributeAndValue(tfAttribute, attributeVal.(string))
			Vars[variableName] = fmt.Sprintf("\"%s\"", attributeVal)
			interpolationMap[attributeVal.(string)] = TfHclVersionvar.GetVarHclString(variableName)
		}
	}
	return nil
}

func getVarNameFromAttributeAndValue(attribute string, value string) string {
	var isStringContainSpecialChar = regexp.MustCompile("\\W+")
	// check if value contain special character
	if isStringContainSpecialChar.MatchString(value) {
		value = isStringContainSpecialChar.ReplaceAllString(value, "-")
	}

	// following format attribute--val
	return fmt.Sprintf(globalvar.VariableGlobalLevelFormat, attribute, value)
}

/* Functions for handling special cases*/

// Handle both flags provided
func exportAttributeForResourceAndGlobalLevel(sourceAttributes map[string]interface{}, resourceType string, resourceName string, varsExportResourceLevel map[string][]string, varsExportGlobalLevel []string, interpolationMap map[string]string) error {
	// export attribute per resource level is higher priority than global level
	if _, exist := varsExportResourceLevel[resourceType]; exist {
		return exportAttributeForResourceLevel(sourceAttributes, resourceType, resourceName, VarsExportForResourceLevel, interpolationMap)
	}
	return exportAttributeForGlobalLevel(sourceAttributes, resourceName, VarsExportForGlobalLevel, interpolationMap)
}

// Handle no flags provided
func exportAttributeFromDefaultList(defaultList []string, sourceAttributes map[string]interface{}, resourceName string, interpolationMap map[string]string) error {
	return exportAttributeForGlobalLevel(sourceAttributes, resourceName, defaultList, interpolationMap)
}
