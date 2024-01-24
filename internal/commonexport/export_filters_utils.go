// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

import (
	"fmt"
	"reflect"
	"strings"
)

func WalkAndGet(path string, data interface{}) []interface{} {
	_, values := walkAndGet(path, data)
	return values
}
func WalkAndCheckField(path string, data interface{}) bool {
	hasField, _ := walkAndGet(path, data)
	return hasField
}

func walkAndGet(path string, data interface{}) (bool, []interface{}) {
	val := reflect.ValueOf(data)

	if data == nil {
		if path == "" {
			return true, []interface{}{}
		}
		return false, []interface{}{}
	}

	if isArray(val.Interface()) {
		var arrayValues []interface{}
		for i := 0; i < val.Len(); i++ {
			foundField, fieldValue := walkAndGet(path, val.Index(i).Interface())
			if foundField {
				arrayValues = append(arrayValues, fieldValue...)
			}
		}
		return len(arrayValues) > 0, arrayValues
	}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			pathFirstElement := strings.SplitN(path, ".", 2)
			println(e.String())
			if e.String() == pathFirstElement[0] {
				var pathReminder = ""
				if len(pathFirstElement) > 1 {
					pathReminder = pathFirstElement[1]
				}
				hasField, value := walkAndGet(pathReminder, v.Interface())
				if !hasField {
					hasField, value = walkAndGet(path, v.Interface())
				}
				return hasField, value
			} else if e.String() == path {
				return walkAndGet("", v.Interface())
			}
		}
	}

	if val.Kind() == reflect.String && path == "" {
		return true, []interface{}{val.Interface()}
	}

	return false, []interface{}{}
}

func isArray(val interface{}) bool { // Go reflect lib can't sometimes detect given value is array
	switch val.(type) {
	case []interface{}:
		return true
	case []string:
		return true
	default:
		return false
	}
}

func isStringArray(val interface{}) bool { // to support locally established arrays
	switch val.(type) {
	case []string:
		return true
	default:
		return false
	}
}

// create deep copies of the filters supplied
func GetFiltersDeepCopy(filters []ResourceFilter) ([]ResourceFilter, error) {
	filtersCopy := make([]ResourceFilter, 0)
	if filters == nil {
		return filtersCopy, nil
	}

	for _, filter := range filters {
		filterCopy, err := getFilterDeepCopy(filter)
		if err != nil {
			return nil, err
		}
		filtersCopy = append(filtersCopy, filterCopy)
	}
	return filtersCopy, nil
}

// create deep copy of the filter supplied
func getFilterDeepCopy(filter ResourceFilter) (ResourceFilter, error) {
	switch filter.(type) {
	case *ResourceTypeFilter:
		resourceTypeFilter, ok := filter.(*ResourceTypeFilter)
		if !ok {
			return nil, fmt.Errorf("unable to convert filter %+v to resource type filter", filter)
		}

		resourceTypeMapCopy := make(map[string]bool)
		for k, v := range resourceTypeFilter.ResourceType {
			resourceTypeMapCopy[k] = v
		}
		return &ResourceTypeFilter{
			resourceTypeMapCopy,
			resourceTypeFilter.ResourceTypeOperator,
		}, nil
	case *FieldValueFilter:
		fieldValueFilter, ok := filter.(*FieldValueFilter)
		if !ok {
			return nil, fmt.Errorf("unable to convert filter %+v to field value type filter", filter)
		}

		fieldValuesCopy := make([]string, len(fieldValueFilter.Values))
		copy(fieldValuesCopy, fieldValueFilter.Values)

		return &FieldValueFilter{
			fieldValueFilter.FieldPath,
			fieldValueFilter.ResourceTypeOperator,
			fieldValuesCopy,
		}, nil
	case *ResourceFieldValueFilter:
		resourceFieldValueFilter, ok := filter.(*ResourceFieldValueFilter)
		if !ok {
			return nil, fmt.Errorf("unable to convert filter %+v to resource field value type filter", filter)
		}

		values := make([]string, len(resourceFieldValueFilter.Values))
		copy(values, resourceFieldValueFilter.Values)

		return &ResourceFieldValueFilter{
			resourceFieldValueFilter.ResourceType,
			resourceFieldValueFilter.ResourceTypeOperator,
			resourceFieldValueFilter.FieldPath,
			values,
		}, nil
	}
	return nil, fmt.Errorf("unable to convert filter %+v", filter)
}
