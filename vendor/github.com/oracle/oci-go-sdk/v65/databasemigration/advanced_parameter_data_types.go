// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// AdvancedParameterDataTypesEnum Enum with underlying type: string
type AdvancedParameterDataTypesEnum string

// Set of constants representing the allowable values for AdvancedParameterDataTypesEnum
const (
	AdvancedParameterDataTypesString  AdvancedParameterDataTypesEnum = "STRING"
	AdvancedParameterDataTypesInteger AdvancedParameterDataTypesEnum = "INTEGER"
	AdvancedParameterDataTypesFloat   AdvancedParameterDataTypesEnum = "FLOAT"
	AdvancedParameterDataTypesBoolean AdvancedParameterDataTypesEnum = "BOOLEAN"
)

var mappingAdvancedParameterDataTypesEnum = map[string]AdvancedParameterDataTypesEnum{
	"STRING":  AdvancedParameterDataTypesString,
	"INTEGER": AdvancedParameterDataTypesInteger,
	"FLOAT":   AdvancedParameterDataTypesFloat,
	"BOOLEAN": AdvancedParameterDataTypesBoolean,
}

var mappingAdvancedParameterDataTypesEnumLowerCase = map[string]AdvancedParameterDataTypesEnum{
	"string":  AdvancedParameterDataTypesString,
	"integer": AdvancedParameterDataTypesInteger,
	"float":   AdvancedParameterDataTypesFloat,
	"boolean": AdvancedParameterDataTypesBoolean,
}

// GetAdvancedParameterDataTypesEnumValues Enumerates the set of values for AdvancedParameterDataTypesEnum
func GetAdvancedParameterDataTypesEnumValues() []AdvancedParameterDataTypesEnum {
	values := make([]AdvancedParameterDataTypesEnum, 0)
	for _, v := range mappingAdvancedParameterDataTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetAdvancedParameterDataTypesEnumStringValues Enumerates the set of values in String for AdvancedParameterDataTypesEnum
func GetAdvancedParameterDataTypesEnumStringValues() []string {
	return []string{
		"STRING",
		"INTEGER",
		"FLOAT",
		"BOOLEAN",
	}
}

// GetMappingAdvancedParameterDataTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdvancedParameterDataTypesEnum(val string) (AdvancedParameterDataTypesEnum, bool) {
	enum, ok := mappingAdvancedParameterDataTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
