// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// MetricTimeWindowEnum Enum with underlying type: string
type MetricTimeWindowEnum string

// Set of constants representing the allowable values for MetricTimeWindowEnum
const (
	MetricTimeWindow1d  MetricTimeWindowEnum = "1d"
	MetricTimeWindow7d  MetricTimeWindowEnum = "7d"
	MetricTimeWindow30d MetricTimeWindowEnum = "30d"
)

var mappingMetricTimeWindowEnum = map[string]MetricTimeWindowEnum{
	"1d":  MetricTimeWindow1d,
	"7d":  MetricTimeWindow7d,
	"30d": MetricTimeWindow30d,
}

var mappingMetricTimeWindowEnumLowerCase = map[string]MetricTimeWindowEnum{
	"1d":  MetricTimeWindow1d,
	"7d":  MetricTimeWindow7d,
	"30d": MetricTimeWindow30d,
}

// GetMetricTimeWindowEnumValues Enumerates the set of values for MetricTimeWindowEnum
func GetMetricTimeWindowEnumValues() []MetricTimeWindowEnum {
	values := make([]MetricTimeWindowEnum, 0)
	for _, v := range mappingMetricTimeWindowEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricTimeWindowEnumStringValues Enumerates the set of values in String for MetricTimeWindowEnum
func GetMetricTimeWindowEnumStringValues() []string {
	return []string{
		"1d",
		"7d",
		"30d",
	}
}

// GetMappingMetricTimeWindowEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricTimeWindowEnum(val string) (MetricTimeWindowEnum, bool) {
	enum, ok := mappingMetricTimeWindowEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
