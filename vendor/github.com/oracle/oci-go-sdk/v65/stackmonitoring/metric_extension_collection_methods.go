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

// MetricExtensionCollectionMethodsEnum Enum with underlying type: string
type MetricExtensionCollectionMethodsEnum string

// Set of constants representing the allowable values for MetricExtensionCollectionMethodsEnum
const (
	MetricExtensionCollectionMethodsOsCommand MetricExtensionCollectionMethodsEnum = "OS_COMMAND"
	MetricExtensionCollectionMethodsSql       MetricExtensionCollectionMethodsEnum = "SQL"
	MetricExtensionCollectionMethodsJmx       MetricExtensionCollectionMethodsEnum = "JMX"
)

var mappingMetricExtensionCollectionMethodsEnum = map[string]MetricExtensionCollectionMethodsEnum{
	"OS_COMMAND": MetricExtensionCollectionMethodsOsCommand,
	"SQL":        MetricExtensionCollectionMethodsSql,
	"JMX":        MetricExtensionCollectionMethodsJmx,
}

var mappingMetricExtensionCollectionMethodsEnumLowerCase = map[string]MetricExtensionCollectionMethodsEnum{
	"os_command": MetricExtensionCollectionMethodsOsCommand,
	"sql":        MetricExtensionCollectionMethodsSql,
	"jmx":        MetricExtensionCollectionMethodsJmx,
}

// GetMetricExtensionCollectionMethodsEnumValues Enumerates the set of values for MetricExtensionCollectionMethodsEnum
func GetMetricExtensionCollectionMethodsEnumValues() []MetricExtensionCollectionMethodsEnum {
	values := make([]MetricExtensionCollectionMethodsEnum, 0)
	for _, v := range mappingMetricExtensionCollectionMethodsEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricExtensionCollectionMethodsEnumStringValues Enumerates the set of values in String for MetricExtensionCollectionMethodsEnum
func GetMetricExtensionCollectionMethodsEnumStringValues() []string {
	return []string{
		"OS_COMMAND",
		"SQL",
		"JMX",
	}
}

// GetMappingMetricExtensionCollectionMethodsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricExtensionCollectionMethodsEnum(val string) (MetricExtensionCollectionMethodsEnum, bool) {
	enum, ok := mappingMetricExtensionCollectionMethodsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
