// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// LogSetKeyTypesEnum Enum with underlying type: string
type LogSetKeyTypesEnum string

// Set of constants representing the allowable values for LogSetKeyTypesEnum
const (
	LogSetKeyTypesObjectPath LogSetKeyTypesEnum = "OBJECT_PATH"
)

var mappingLogSetKeyTypesEnum = map[string]LogSetKeyTypesEnum{
	"OBJECT_PATH": LogSetKeyTypesObjectPath,
}

var mappingLogSetKeyTypesEnumLowerCase = map[string]LogSetKeyTypesEnum{
	"object_path": LogSetKeyTypesObjectPath,
}

// GetLogSetKeyTypesEnumValues Enumerates the set of values for LogSetKeyTypesEnum
func GetLogSetKeyTypesEnumValues() []LogSetKeyTypesEnum {
	values := make([]LogSetKeyTypesEnum, 0)
	for _, v := range mappingLogSetKeyTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogSetKeyTypesEnumStringValues Enumerates the set of values in String for LogSetKeyTypesEnum
func GetLogSetKeyTypesEnumStringValues() []string {
	return []string{
		"OBJECT_PATH",
	}
}

// GetMappingLogSetKeyTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogSetKeyTypesEnum(val string) (LogSetKeyTypesEnum, bool) {
	enum, ok := mappingLogSetKeyTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
