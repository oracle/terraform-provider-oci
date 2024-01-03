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

// LastExecutionStatusEnum Enum with underlying type: string
type LastExecutionStatusEnum string

// Set of constants representing the allowable values for LastExecutionStatusEnum
const (
	LastExecutionStatusFailed    LastExecutionStatusEnum = "FAILED"
	LastExecutionStatusSucceeded LastExecutionStatusEnum = "SUCCEEDED"
)

var mappingLastExecutionStatusEnum = map[string]LastExecutionStatusEnum{
	"FAILED":    LastExecutionStatusFailed,
	"SUCCEEDED": LastExecutionStatusSucceeded,
}

var mappingLastExecutionStatusEnumLowerCase = map[string]LastExecutionStatusEnum{
	"failed":    LastExecutionStatusFailed,
	"succeeded": LastExecutionStatusSucceeded,
}

// GetLastExecutionStatusEnumValues Enumerates the set of values for LastExecutionStatusEnum
func GetLastExecutionStatusEnumValues() []LastExecutionStatusEnum {
	values := make([]LastExecutionStatusEnum, 0)
	for _, v := range mappingLastExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetLastExecutionStatusEnumStringValues Enumerates the set of values in String for LastExecutionStatusEnum
func GetLastExecutionStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingLastExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLastExecutionStatusEnum(val string) (LastExecutionStatusEnum, bool) {
	enum, ok := mappingLastExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
