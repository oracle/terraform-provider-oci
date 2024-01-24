// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package commonexport

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

var (
	// regex to match Type filter
	resourceTypeFilterPattern, _ = regexp.Compile("^Type(=|!=)([Aa-zZ_,]+)$")
	// regex to match Type attribute filter
	resourceFieldValueFilterPattern, _ = regexp.Compile("^Type=([Aa-zZ_]+);AttrName=([Aa-zZ0-9-_.]+);Value(=|!=)([Aa-zZ0-9_.:, -/]+)$")
	// regex to match global type filter
	fieldValueFilterPattern, _ = regexp.Compile("^AttrName=([Aa-zZ0-9-_.]+);Value(=|!=)([Aa-zZ0-9_.:, -/]+)$")
)

// Reference: https://pkg.go.dev/flag#hdr-Command_line_flag_syntax
type Filter []ResourceFilter

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (f *Filter) String() string {
	return fmt.Sprint(*f)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (f *Filter) Set(value string) error {
	resourceFilter, err := f.ParseFilter(value)
	if err != nil {
		return err
	}
	*f = append(*f, resourceFilter)
	return nil
}

// --filter="Type=oci_core_vcn"  --filter"=Type=oci_core_subnet
// AND, OR operator
// Type=oci_core_vcn
// Type=oci_core_vcn, oci_core_instance
// Type!=oci_core_vcn
// Type!=oci_core_vcn, oci_core_instance
// Type=oci_core_vcn;AttrName=defined_tags.example-namespace.example-key;Value!=example-value
// Type=oci_core_vcn;AttrName=defined_tags.example-namespace.example-key;Value=example-value
// AttrName=defined_tags.example-namespace.example-key;Value=example-value
// AttrName=defined_tags.example-namespace.example-key;Value!=example-value
// AttrName=defined_tags.example-namespace.example-key;Value!=example-value,example-value1
// Upon parsing of above combinations:
// if multiple Types are passed, they will be treated as separate Type filters
// if one type, one attribute and array values is passed, it will be one filter
// AttrName will not accept a list of names
// Type will accept one value when used with AttrName and values
func (f *Filter) ParseFilter(rawFilter string) (ResourceFilter, error) {
	if resourceTypeFilterPattern.MatchString(rawFilter) {
		// matches resource type filter

		// extract operator
		groups := resourceTypeFilterPattern.FindStringSubmatch(rawFilter)
		filterOperator, err := getFilterOperator(groups[1])
		if err != nil {
			return nil, err
		}

		typeValues := make(map[string]bool)
		for _, typeValue := range strings.Split(groups[2], ",") {
			typeValues[strings.ToLower(strings.TrimSpace(typeValue))] = true
		}

		return &ResourceTypeFilter{
			ResourceType:         typeValues,
			ResourceTypeOperator: filterOperator,
		}, nil
	} else if fieldValueFilterPattern.MatchString(rawFilter) {
		// matches filed value filter type

		// extract operator
		groups := fieldValueFilterPattern.FindStringSubmatch(rawFilter)
		filterOperator, err := getFilterOperator(groups[2])
		if err != nil {
			return nil, err
		}

		fieldValues := strings.Split(groups[3], ",")
		for idx, fieldValue := range fieldValues {
			fieldValues[idx] = strings.TrimSpace(fieldValue)
		}

		if !validateFieldValueFilterValue(groups[1]) {
			return nil, fmt.Errorf("unsupported filter argument supplied %s, please check the filter documentation", rawFilter)
		}

		return &FieldValueFilter{
			FieldPath:            groups[1],
			ResourceTypeOperator: filterOperator,
			Values:               fieldValues,
		}, nil
	}
	return nil, fmt.Errorf("invalid value %s provided. Unable to parse filter. Please refer to filter documentation", rawFilter)
}

// convert a delimited string of values into a slice
func convertArrStringToSlice(valuesStr string, delimiter string) []string {
	fieldValues := strings.Split(valuesStr, delimiter)
	for idx, fieldValue := range fieldValues {
		fieldValues[idx] = strings.TrimSpace(fieldValue)
	}
	return fieldValues
}

// get the FilterOperator from the string representation
func getFilterOperator(input string) (FilterOperator, error) {
	if input == globalvar.EqualToOperator {
		return INCLUDE, nil
	} else if input == globalvar.NotEqualToOperator {
		return EXCLUDE, nil
	}

	// default to include
	return INCLUDE, fmt.Errorf("invalid filter operator passed %v", input)
}

// validate if the field value passed for the filter is supported
func validateFieldValueFilterValue(fieldValue string) bool {
	fieldValueFirstElement := strings.SplitN(fieldValue, ".", 2)
	return len(fieldValueFirstElement) >= 1
	/*if len(fieldValueFirstElement) < 1 {
		return false
	}
	//_, ok := globalvar.Supported_global_filter_resource_attributes[fieldValueFirstElement[0]]
	return true*/
}
