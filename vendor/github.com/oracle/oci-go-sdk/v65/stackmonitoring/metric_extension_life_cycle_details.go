// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MetricExtensionLifeCycleDetailsEnum Enum with underlying type: string
type MetricExtensionLifeCycleDetailsEnum string

// Set of constants representing the allowable values for MetricExtensionLifeCycleDetailsEnum
const (
	MetricExtensionLifeCycleDetailsDraft     MetricExtensionLifeCycleDetailsEnum = "DRAFT"
	MetricExtensionLifeCycleDetailsPublished MetricExtensionLifeCycleDetailsEnum = "PUBLISHED"
)

var mappingMetricExtensionLifeCycleDetailsEnum = map[string]MetricExtensionLifeCycleDetailsEnum{
	"DRAFT":     MetricExtensionLifeCycleDetailsDraft,
	"PUBLISHED": MetricExtensionLifeCycleDetailsPublished,
}

var mappingMetricExtensionLifeCycleDetailsEnumLowerCase = map[string]MetricExtensionLifeCycleDetailsEnum{
	"draft":     MetricExtensionLifeCycleDetailsDraft,
	"published": MetricExtensionLifeCycleDetailsPublished,
}

// GetMetricExtensionLifeCycleDetailsEnumValues Enumerates the set of values for MetricExtensionLifeCycleDetailsEnum
func GetMetricExtensionLifeCycleDetailsEnumValues() []MetricExtensionLifeCycleDetailsEnum {
	values := make([]MetricExtensionLifeCycleDetailsEnum, 0)
	for _, v := range mappingMetricExtensionLifeCycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricExtensionLifeCycleDetailsEnumStringValues Enumerates the set of values in String for MetricExtensionLifeCycleDetailsEnum
func GetMetricExtensionLifeCycleDetailsEnumStringValues() []string {
	return []string{
		"DRAFT",
		"PUBLISHED",
	}
}

// GetMappingMetricExtensionLifeCycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricExtensionLifeCycleDetailsEnum(val string) (MetricExtensionLifeCycleDetailsEnum, bool) {
	enum, ok := mappingMetricExtensionLifeCycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
