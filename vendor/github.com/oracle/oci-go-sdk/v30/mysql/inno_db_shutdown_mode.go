// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

// InnoDbShutdownModeEnum Enum with underlying type: string
type InnoDbShutdownModeEnum string

// Set of constants representing the allowable values for InnoDbShutdownModeEnum
const (
	InnoDbShutdownModeImmediate InnoDbShutdownModeEnum = "IMMEDIATE"
	InnoDbShutdownModeFast      InnoDbShutdownModeEnum = "FAST"
	InnoDbShutdownModeSlow      InnoDbShutdownModeEnum = "SLOW"
)

var mappingInnoDbShutdownMode = map[string]InnoDbShutdownModeEnum{
	"IMMEDIATE": InnoDbShutdownModeImmediate,
	"FAST":      InnoDbShutdownModeFast,
	"SLOW":      InnoDbShutdownModeSlow,
}

// GetInnoDbShutdownModeEnumValues Enumerates the set of values for InnoDbShutdownModeEnum
func GetInnoDbShutdownModeEnumValues() []InnoDbShutdownModeEnum {
	values := make([]InnoDbShutdownModeEnum, 0)
	for _, v := range mappingInnoDbShutdownMode {
		values = append(values, v)
	}
	return values
}
