// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"
	"reflect"
	"regexp"
	"strconv"

	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceFiltersSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},

				"values": {
					Type:     schema.TypeList,
					Required: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},

				"regex": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
			},
		},
	}
}

var PrimitiveDataTypes = map[schema.ValueType]bool{
	schema.TypeString: true,
	schema.TypeBool:   true,
	schema.TypeFloat:  true,
	schema.TypeInt:    true,
}

// Process an entity's properties (string or array of strings) by N filter sets of
// keyword:values, where each filter set ANDs and each keyword:values set ORs
func ApplyFilters(filters *schema.Set, items []map[string]interface{}, resourceSchema map[string]*schema.Schema) []map[string]interface{} {
	if filters == nil || filters.Len() == 0 {
		return items
	}

	for _, f := range filters.List() {
		fSet := f.(map[string]interface{})
		keyword := fSet["name"].(string)
		var pathElements []string
		var err error
		if pathElements, err = getFieldPathElements(resourceSchema, keyword); err != nil {
			log.Printf(err.Error())
			pathElements = []string{keyword}
		}

		isReg := false
		if regex, regexOk := fSet["regex"]; regexOk {
			isReg = regex.(bool)
		}

		// create a string equality check strategy based on this filters "regex" flag
		stringsEqual := func(propertyVal string, filterVal string) bool {
			if isReg {
				re, err := regexp.Compile(filterVal)
				if err != nil {
					// todo: when all SetData() fns are refactored to return a possible error, these log statements should
					// be converted to errors for return propagation
					log.Printf(`[WARN] Invalid regular expression "%s" for "%s" filter\n`, filterVal, keyword)
					return false
				}
				return re.MatchString(propertyVal)
			}

			return filterVal == propertyVal
		}

		// build a collection of items from matches against the set of filters
		res := make([]map[string]interface{}, 0)
		for _, item := range items {
			targetVal, targetValOk := getValueFromPath(item, pathElements)
			if targetValOk && orComparator(targetVal, fSet["values"].([]interface{}), stringsEqual) {
				res = append(res, item)
			}
		}
		items = res
	}

	return items
}

func getValueFromPath(item map[string]interface{}, path []string) (targetVal interface{}, targetValOk bool) {
	workingMap := item
	tempWorkingMap := item
	var conversionOk bool
	for _, pathElement := range path[:len(path)-1] {
		// Defensive check for non existent values
		if workingMap[pathElement] == nil {
			return nil, false
		}
		// Check if it is map
		if tempWorkingMap, conversionOk = checkAndConvertMap(workingMap[pathElement]); !conversionOk {
			// if not map then it has to be a nested structure which is modeled as list with exactly one element of type map[string]interface{}
			if tempWorkingMap, conversionOk = checkAndConvertNestedStructure(workingMap[pathElement]); !conversionOk {
				return nil, false
			}
		}
		workingMap = tempWorkingMap
	}

	targetVal, targetValOk = workingMap[path[len(path)-1]]
	return
}

func checkAndConvertMap(element interface{}) (map[string]interface{}, bool) {
	if tempWorkingMap, isOk := element.(map[string]interface{}); isOk {
		return tempWorkingMap, true
	}

	if stringToStrinMap, isOk := element.(map[string]string); isOk {
		return convertToObjectMap(stringToStrinMap), true
	}

	return nil, false
}

func convertToObjectMap(stringTostring map[string]string) map[string]interface{} {
	convertedMap := make(map[string]interface{}, len(stringTostring))
	for key, value := range stringTostring {
		convertedMap[key] = value
	}

	return convertedMap
}

func checkAndConvertNestedStructure(element interface{}) (map[string]interface{}, bool) {
	if convertedList, convertedListOk := element.([]interface{}); convertedListOk && len(convertedList) == 1 {
		workingMap, isOk := convertedList[0].(map[string]interface{})
		return workingMap, isOk
	}

	return nil, false
}

//Converts the filter name which is delimited by '.' into a list of XPath elements
//Read the filter name from left most token and look into schema map to interpret rest of the filter name string
// e.g. for core_instance: freeform_tags.com.oracle.department -> ["freeform_tags", "com.oracle.department"], nil
// e.g. for core_instance: source_details.source_type -> ["source_details", "source_type"], nil
// e.g. for core_instance: source_details.source_type.xyz -> nil, error
func getFieldPathElements(resourceSchema map[string]*schema.Schema, filterName string) ([]string, error) {

	if resourceSchema == nil {
		log.Printf(`[WARN] schema is nil for filter name %s \n`, filterName)
		return nil, fmt.Errorf("schema is nil for filter name %s", filterName)
	}

	tokenizedFields := strings.Split(filterName, ".")

	//validate tokens
	if len(tokenizedFields) == 0 {
		log.Printf(`[WARN] Invalid filter name "%s"  \n`, filterName)
		return nil, fmt.Errorf("invalid filter name %s", filterName)
	}

	if resourceSchema[tokenizedFields[0]] == nil {
		log.Printf(`[WARN] Schema is nil for token %s for filter name "%s"\n`, tokenizedFields[0], filterName)
		return nil, fmt.Errorf("schema is nil for token %s for filter name %s", tokenizedFields[0], filterName)
	}

	var pathElements []string
	currentSchema := resourceSchema
	for index, tokenizedField := range tokenizedFields {
		if fieldSchema, ok := currentSchema[tokenizedField]; ok && isValidSchemaType(fieldSchema) {
			// add current path element to pathElements
			pathElements = append(pathElements, tokenizedField)
			//check if nested
			convertedElementSchema, conversionOk := fieldSchema.Elem.(*schema.Resource)
			if !conversionOk { // No nested structure
				if len(tokenizedFields) > index+1 { // have more tokens to handle
					// if we have more tokens the schema type has to be map else error condition
					if fieldSchema.Type != schema.TypeMap {
						return nil, fmt.Errorf("invalid filter name format found %s", filterName)

					}
					pathElement := strings.Join(tokenizedFields[index+1:], ".")
					pathElements = append(pathElements, pathElement)
				}
				break
			} else {
				// get next schema and handle next token
				currentSchema = convertedElementSchema.Schema
			}
		} else {
			return nil, fmt.Errorf("invalid schema found for filter name %s", filterName)
		}
	}

	if len(pathElements) == 0 {
		return nil, fmt.Errorf("path elements were not initialized properly")
	}

	return pathElements, nil
}

func isValidSchemaType(fieldSchema *schema.Schema) bool {
	if fieldSchema.Type == schema.TypeList || fieldSchema.Type == schema.TypeSet {
		if elemSchema, conversionOk := fieldSchema.Elem.(*schema.Schema); conversionOk && elemSchema.Type == schema.TypeString {
			return true
		} else if fieldSchema.MaxItems == 1 && fieldSchema.MinItems == 1 { //nested structures
			return true
		}
		return false
	}

	return true
}

type StringCheck func(propertyVal string, filterVal string) bool

// orComparator returns true for any filter that matches the target property
func orComparator(target interface{}, filters []interface{}, stringsEqual StringCheck) bool {
	// Use reflection to determine whether the underlying type of the filtering attribute is a string or
	// array of strings. Mainly used because the property could be an SDK enum with underlying string type.
	val := reflect.ValueOf(target)
	valType := val.Type()

	for _, fVal := range filters {
		switch valType.Kind() {
		case reflect.Bool:
			fBool, err := strconv.ParseBool(fVal.(string))
			if err != nil {
				log.Println("[WARN] Filtering against Type Bool field with un-parsable string boolean form")
				return false
			}
			if val.Bool() == fBool {
				return true
			}
		case reflect.Int:
			// the target field is of type int, but the filter values list element type is string, users can supply string
			// or int like `values = [300, "3600"]` but terraform will converts to string, so use ParseInt
			fInt, err := strconv.ParseInt(fVal.(string), 10, 64)
			if err != nil {
				log.Println("[WARN] Filtering against Type Int field with non-int filter value")
				return false
			}
			if val.Int() == fInt {
				return true
			}
		case reflect.Float64:
			// same comment as above for Ints
			fFloat, err := strconv.ParseFloat(fVal.(string), 64)
			if err != nil {
				log.Println("[WARN] Filtering against Type Float field with non-float filter value")
				return false
			}
			if val.Float() == fFloat {
				return true
			}
		case reflect.String:
			if stringsEqual(val.String(), fVal.(string)) {
				return true
			}
		case reflect.Slice, reflect.Array:
			if valType.Elem().Kind() == reflect.String {
				arrLen := val.Len()
				for i := 0; i < arrLen; i++ {
					if stringsEqual(val.Index(i).String(), fVal.(string)) {
						return true
					}
				}
			}
		}
	}
	return false
}
