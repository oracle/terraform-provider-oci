// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MetricExtensionSortByEnum Enum with underlying type: string
type MetricExtensionSortByEnum string

// Set of constants representing the allowable values for MetricExtensionSortByEnum
const (
	MetricExtensionSortByName        MetricExtensionSortByEnum = "NAME"
	MetricExtensionSortByTimeCreated MetricExtensionSortByEnum = "TIME_CREATED"
)

var mappingMetricExtensionSortByEnum = map[string]MetricExtensionSortByEnum{
	"NAME":         MetricExtensionSortByName,
	"TIME_CREATED": MetricExtensionSortByTimeCreated,
}

var mappingMetricExtensionSortByEnumLowerCase = map[string]MetricExtensionSortByEnum{
	"name":         MetricExtensionSortByName,
	"time_created": MetricExtensionSortByTimeCreated,
}

// GetMetricExtensionSortByEnumValues Enumerates the set of values for MetricExtensionSortByEnum
func GetMetricExtensionSortByEnumValues() []MetricExtensionSortByEnum {
	values := make([]MetricExtensionSortByEnum, 0)
	for _, v := range mappingMetricExtensionSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricExtensionSortByEnumStringValues Enumerates the set of values in String for MetricExtensionSortByEnum
func GetMetricExtensionSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingMetricExtensionSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricExtensionSortByEnum(val string) (MetricExtensionSortByEnum, bool) {
	enum, ok := mappingMetricExtensionSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
