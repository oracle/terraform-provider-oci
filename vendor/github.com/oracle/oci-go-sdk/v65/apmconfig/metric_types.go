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

// MetricTypesEnum Enum with underlying type: string
type MetricTypesEnum string

// Set of constants representing the allowable values for MetricTypesEnum
const (
	MetricTypesCounter MetricTypesEnum = "COUNTER"
	MetricTypesGauge   MetricTypesEnum = "GAUGE"
)

var mappingMetricTypesEnum = map[string]MetricTypesEnum{
	"COUNTER": MetricTypesCounter,
	"GAUGE":   MetricTypesGauge,
}

var mappingMetricTypesEnumLowerCase = map[string]MetricTypesEnum{
	"counter": MetricTypesCounter,
	"gauge":   MetricTypesGauge,
}

// GetMetricTypesEnumValues Enumerates the set of values for MetricTypesEnum
func GetMetricTypesEnumValues() []MetricTypesEnum {
	values := make([]MetricTypesEnum, 0)
	for _, v := range mappingMetricTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricTypesEnumStringValues Enumerates the set of values in String for MetricTypesEnum
func GetMetricTypesEnumStringValues() []string {
	return []string{
		"COUNTER",
		"GAUGE",
	}
}

// GetMappingMetricTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricTypesEnum(val string) (MetricTypesEnum, bool) {
	enum, ok := mappingMetricTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
