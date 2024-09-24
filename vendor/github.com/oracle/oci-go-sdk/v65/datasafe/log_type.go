// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// LogTypeEnum Enum with underlying type: string
type LogTypeEnum string

// Set of constants representing the allowable values for LogTypeEnum
const (
	LogTypeViolationLog LogTypeEnum = "VIOLATION_LOG"
)

var mappingLogTypeEnum = map[string]LogTypeEnum{
	"VIOLATION_LOG": LogTypeViolationLog,
}

var mappingLogTypeEnumLowerCase = map[string]LogTypeEnum{
	"violation_log": LogTypeViolationLog,
}

// GetLogTypeEnumValues Enumerates the set of values for LogTypeEnum
func GetLogTypeEnumValues() []LogTypeEnum {
	values := make([]LogTypeEnum, 0)
	for _, v := range mappingLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogTypeEnumStringValues Enumerates the set of values in String for LogTypeEnum
func GetLogTypeEnumStringValues() []string {
	return []string{
		"VIOLATION_LOG",
	}
}

// GetMappingLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogTypeEnum(val string) (LogTypeEnum, bool) {
	enum, ok := mappingLogTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
