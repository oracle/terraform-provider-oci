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

// MetricExtensionLifeCycleStatesEnum Enum with underlying type: string
type MetricExtensionLifeCycleStatesEnum string

// Set of constants representing the allowable values for MetricExtensionLifeCycleStatesEnum
const (
	MetricExtensionLifeCycleStatesActive  MetricExtensionLifeCycleStatesEnum = "ACTIVE"
	MetricExtensionLifeCycleStatesDeleted MetricExtensionLifeCycleStatesEnum = "DELETED"
)

var mappingMetricExtensionLifeCycleStatesEnum = map[string]MetricExtensionLifeCycleStatesEnum{
	"ACTIVE":  MetricExtensionLifeCycleStatesActive,
	"DELETED": MetricExtensionLifeCycleStatesDeleted,
}

var mappingMetricExtensionLifeCycleStatesEnumLowerCase = map[string]MetricExtensionLifeCycleStatesEnum{
	"active":  MetricExtensionLifeCycleStatesActive,
	"deleted": MetricExtensionLifeCycleStatesDeleted,
}

// GetMetricExtensionLifeCycleStatesEnumValues Enumerates the set of values for MetricExtensionLifeCycleStatesEnum
func GetMetricExtensionLifeCycleStatesEnumValues() []MetricExtensionLifeCycleStatesEnum {
	values := make([]MetricExtensionLifeCycleStatesEnum, 0)
	for _, v := range mappingMetricExtensionLifeCycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricExtensionLifeCycleStatesEnumStringValues Enumerates the set of values in String for MetricExtensionLifeCycleStatesEnum
func GetMetricExtensionLifeCycleStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingMetricExtensionLifeCycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricExtensionLifeCycleStatesEnum(val string) (MetricExtensionLifeCycleStatesEnum, bool) {
	enum, ok := mappingMetricExtensionLifeCycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
