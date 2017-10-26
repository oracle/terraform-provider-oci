package main

import (
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

		check := func(filterVal, propertyVal string) bool {
			if isReg {
				re, err := regexp.Compile(filterVal)
				if err != nil {
					panic(fmt.Errorf(`Invalid regular expression "%s" for "%s" filter`, filterVal, keyword))
				}
				return re.MatchString(propertyVal)
			} else {
				return filterVal == propertyVal
			}
		}

		orComparator := func(item map[string]interface{}) bool {
			for _, val := range fSet["values"].([]interface{}) {

				// if the property contains an array of strings
				strArr, strArrOk := item[keyword].([]string)
				if strArrOk {
					for _, subStr := range strArr {
						if check(val.(string), subStr) {
							return true
						}
					}
				}

				// if the property is a literal string
				str, strOk := item[keyword].(string)
				if strOk && check(val.(string), str) {
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
