// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// NamespaceUsageEnum Enum with underlying type: string
type NamespaceUsageEnum string

// Set of constants representing the allowable values for NamespaceUsageEnum
const (
	NamespaceUsageMetricExtension NamespaceUsageEnum = "METRIC_EXTENSION"
)

var mappingNamespaceUsageEnum = map[string]NamespaceUsageEnum{
	"METRIC_EXTENSION": NamespaceUsageMetricExtension,
}

var mappingNamespaceUsageEnumLowerCase = map[string]NamespaceUsageEnum{
	"metric_extension": NamespaceUsageMetricExtension,
}

// GetNamespaceUsageEnumValues Enumerates the set of values for NamespaceUsageEnum
func GetNamespaceUsageEnumValues() []NamespaceUsageEnum {
	values := make([]NamespaceUsageEnum, 0)
	for _, v := range mappingNamespaceUsageEnum {
		values = append(values, v)
	}
	return values
}

// GetNamespaceUsageEnumStringValues Enumerates the set of values in String for NamespaceUsageEnum
func GetNamespaceUsageEnumStringValues() []string {
	return []string{
		"METRIC_EXTENSION",
	}
}

// GetMappingNamespaceUsageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamespaceUsageEnum(val string) (NamespaceUsageEnum, bool) {
	enum, ok := mappingNamespaceUsageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
