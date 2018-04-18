// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"
	"reflect"
	"regexp"
	"strconv"

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

// Process an entity's properties (string or array of strings) by N filter sets of
// keyword:values, where each filter set ANDs and each keyword:values set ORs
func ApplyFilters(filters *schema.Set, items []map[string]interface{}) []map[string]interface{} {
	if filters == nil || filters.Len() == 0 {
		return items
	}

	for _, f := range filters.List() {
		fSet := f.(map[string]interface{})
		keyword := fSet["name"].(string)

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
			targetVal, targetValOk := item[keyword]
			if targetValOk && orComparator(targetVal, fSet["values"].([]interface{}), stringsEqual) {
				res = append(res, item)
			}
		}
		items = res
	}

	return items
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
