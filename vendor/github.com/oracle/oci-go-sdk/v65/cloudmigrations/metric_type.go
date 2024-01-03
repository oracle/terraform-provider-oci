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

// MetricTypeEnum Enum with underlying type: string
type MetricTypeEnum string

// Set of constants representing the allowable values for MetricTypeEnum
const (
	MetricTypeAuto       MetricTypeEnum = "AUTO"
	MetricTypeHistorical MetricTypeEnum = "HISTORICAL"
	MetricTypeRuntime    MetricTypeEnum = "RUNTIME"
)

var mappingMetricTypeEnum = map[string]MetricTypeEnum{
	"AUTO":       MetricTypeAuto,
	"HISTORICAL": MetricTypeHistorical,
	"RUNTIME":    MetricTypeRuntime,
}

var mappingMetricTypeEnumLowerCase = map[string]MetricTypeEnum{
	"auto":       MetricTypeAuto,
	"historical": MetricTypeHistorical,
	"runtime":    MetricTypeRuntime,
}

// GetMetricTypeEnumValues Enumerates the set of values for MetricTypeEnum
func GetMetricTypeEnumValues() []MetricTypeEnum {
	values := make([]MetricTypeEnum, 0)
	for _, v := range mappingMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricTypeEnumStringValues Enumerates the set of values in String for MetricTypeEnum
func GetMetricTypeEnumStringValues() []string {
	return []string{
		"AUTO",
		"HISTORICAL",
		"RUNTIME",
	}
}

// GetMappingMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricTypeEnum(val string) (MetricTypeEnum, bool) {
	enum, ok := mappingMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
