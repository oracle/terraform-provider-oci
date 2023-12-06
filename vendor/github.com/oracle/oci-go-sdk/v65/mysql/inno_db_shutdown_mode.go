// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// InnoDbShutdownModeEnum Enum with underlying type: string
type InnoDbShutdownModeEnum string

// Set of constants representing the allowable values for InnoDbShutdownModeEnum
const (
	InnoDbShutdownModeImmediate InnoDbShutdownModeEnum = "IMMEDIATE"
	InnoDbShutdownModeFast      InnoDbShutdownModeEnum = "FAST"
	InnoDbShutdownModeSlow      InnoDbShutdownModeEnum = "SLOW"
)

var mappingInnoDbShutdownModeEnum = map[string]InnoDbShutdownModeEnum{
	"IMMEDIATE": InnoDbShutdownModeImmediate,
	"FAST":      InnoDbShutdownModeFast,
	"SLOW":      InnoDbShutdownModeSlow,
}

var mappingInnoDbShutdownModeEnumLowerCase = map[string]InnoDbShutdownModeEnum{
	"immediate": InnoDbShutdownModeImmediate,
	"fast":      InnoDbShutdownModeFast,
	"slow":      InnoDbShutdownModeSlow,
}

// GetInnoDbShutdownModeEnumValues Enumerates the set of values for InnoDbShutdownModeEnum
func GetInnoDbShutdownModeEnumValues() []InnoDbShutdownModeEnum {
	values := make([]InnoDbShutdownModeEnum, 0)
	for _, v := range mappingInnoDbShutdownModeEnum {
		values = append(values, v)
	}
	return values
}

// GetInnoDbShutdownModeEnumStringValues Enumerates the set of values in String for InnoDbShutdownModeEnum
func GetInnoDbShutdownModeEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"FAST",
		"SLOW",
	}
}

// GetMappingInnoDbShutdownModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInnoDbShutdownModeEnum(val string) (InnoDbShutdownModeEnum, bool) {
	enum, ok := mappingInnoDbShutdownModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
