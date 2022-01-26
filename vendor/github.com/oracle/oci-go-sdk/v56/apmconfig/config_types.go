// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

// ConfigTypesEnum Enum with underlying type: string
type ConfigTypesEnum string

// Set of constants representing the allowable values for ConfigTypesEnum
const (
	ConfigTypesSpanFilter  ConfigTypesEnum = "SPAN_FILTER"
	ConfigTypesMetricGroup ConfigTypesEnum = "METRIC_GROUP"
	ConfigTypesApdex       ConfigTypesEnum = "APDEX"
)

var mappingConfigTypes = map[string]ConfigTypesEnum{
	"SPAN_FILTER":  ConfigTypesSpanFilter,
	"METRIC_GROUP": ConfigTypesMetricGroup,
	"APDEX":        ConfigTypesApdex,
}

// GetConfigTypesEnumValues Enumerates the set of values for ConfigTypesEnum
func GetConfigTypesEnumValues() []ConfigTypesEnum {
	values := make([]ConfigTypesEnum, 0)
	for _, v := range mappingConfigTypes {
		values = append(values, v)
	}
	return values
}
