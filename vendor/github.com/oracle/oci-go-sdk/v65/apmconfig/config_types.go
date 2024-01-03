// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"strings"
)

// ConfigTypesEnum Enum with underlying type: string
type ConfigTypesEnum string

// Set of constants representing the allowable values for ConfigTypesEnum
const (
	ConfigTypesSpanFilter  ConfigTypesEnum = "SPAN_FILTER"
	ConfigTypesMetricGroup ConfigTypesEnum = "METRIC_GROUP"
	ConfigTypesApdex       ConfigTypesEnum = "APDEX"
	ConfigTypesOptions     ConfigTypesEnum = "OPTIONS"
)

var mappingConfigTypesEnum = map[string]ConfigTypesEnum{
	"SPAN_FILTER":  ConfigTypesSpanFilter,
	"METRIC_GROUP": ConfigTypesMetricGroup,
	"APDEX":        ConfigTypesApdex,
	"OPTIONS":      ConfigTypesOptions,
}

var mappingConfigTypesEnumLowerCase = map[string]ConfigTypesEnum{
	"span_filter":  ConfigTypesSpanFilter,
	"metric_group": ConfigTypesMetricGroup,
	"apdex":        ConfigTypesApdex,
	"options":      ConfigTypesOptions,
}

// GetConfigTypesEnumValues Enumerates the set of values for ConfigTypesEnum
func GetConfigTypesEnumValues() []ConfigTypesEnum {
	values := make([]ConfigTypesEnum, 0)
	for _, v := range mappingConfigTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigTypesEnumStringValues Enumerates the set of values in String for ConfigTypesEnum
func GetConfigTypesEnumStringValues() []string {
	return []string{
		"SPAN_FILTER",
		"METRIC_GROUP",
		"APDEX",
		"OPTIONS",
	}
}

// GetMappingConfigTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigTypesEnum(val string) (ConfigTypesEnum, bool) {
	enum, ok := mappingConfigTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
