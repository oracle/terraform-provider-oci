// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ApiFormatTypeEnum Enum with underlying type: string
type ApiFormatTypeEnum string

// Set of constants representing the allowable values for ApiFormatTypeEnum
const (
	ApiFormatTypeCohere  ApiFormatTypeEnum = "COHERE"
	ApiFormatTypeGeneric ApiFormatTypeEnum = "GENERIC"
)

var mappingApiFormatTypeEnum = map[string]ApiFormatTypeEnum{
	"COHERE":  ApiFormatTypeCohere,
	"GENERIC": ApiFormatTypeGeneric,
}

var mappingApiFormatTypeEnumLowerCase = map[string]ApiFormatTypeEnum{
	"cohere":  ApiFormatTypeCohere,
	"generic": ApiFormatTypeGeneric,
}

// GetApiFormatTypeEnumValues Enumerates the set of values for ApiFormatTypeEnum
func GetApiFormatTypeEnumValues() []ApiFormatTypeEnum {
	values := make([]ApiFormatTypeEnum, 0)
	for _, v := range mappingApiFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApiFormatTypeEnumStringValues Enumerates the set of values in String for ApiFormatTypeEnum
func GetApiFormatTypeEnumStringValues() []string {
	return []string{
		"COHERE",
		"GENERIC",
	}
}

// GetMappingApiFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApiFormatTypeEnum(val string) (ApiFormatTypeEnum, bool) {
	enum, ok := mappingApiFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
