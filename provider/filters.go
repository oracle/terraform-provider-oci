package provider

import (
	"reflect"
	"regexp"

	"fmt"

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
		if regex, regexOk := fSet["regex"]; regexOk && regex != nil {
			if strVal, strValOk := regex.(string); strValOk {
				isReg = strVal == "1" || strVal == "true"
			} else if boolVal, boolValOk := regex.(bool); boolValOk {
				isReg = boolVal
			}
		}

		check := func(filterVal string, propertyVal string) bool {
			if isReg {
				re, err := regexp.Compile(filterVal)
				if err != nil {
					panic(fmt.Errorf(`Invalid regular expression "%s" for "%s" filter`, filterVal, keyword))
				}
				return re.MatchString(propertyVal)
			}

			return filterVal == propertyVal
		}

		orComparator := func(item map[string]interface{}) bool {
			actualValue, valueExists := item[keyword]
			if !valueExists {
				return false
			}

			// We use reflection to determine whether the underlying type of the filtering attribute is a string or
			// array of strings. Mainly used because the property could be an SDK enum with underlying string type.
			// TODO: We should store SDK enum values in state as strings prior to calling ApplyFilters, to avoid using reflection
			rValue := reflect.ValueOf(actualValue)
			rType := rValue.Type()

			isStringArray := (rType.Kind() == reflect.Slice || rType.Kind() == reflect.Array) && rType.Elem().Kind() == reflect.String
			isString := rType.Kind() == reflect.String
			if !isStringArray && !isString {
				// property is neither a string nor array of strings, so it can be filtered out
				return false
			}

			for _, filterValue := range fSet["values"].([]interface{}) {
				if isStringArray {
					arrLen := rValue.Len()
					for i := 0; i < arrLen; i++ {
						if check(filterValue.(string), rValue.Index(i).String()) {
							return true
						}
					}
				} else if check(filterValue.(string), rValue.String()) {
					return true
				}
			}
			return false
		}

		items = filter(items, orComparator)
	}

	return items
}

func filter(items []map[string]interface{}, comparator func(map[string]interface{}) bool) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	for _, item := range items {
		if comparator(item) {
			res = append(res, item)
		}
	}
	return res
}
