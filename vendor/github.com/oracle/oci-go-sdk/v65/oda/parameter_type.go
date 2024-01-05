// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// ParameterTypeEnum Enum with underlying type: string
type ParameterTypeEnum string

// Set of constants representing the allowable values for ParameterTypeEnum
const (
	ParameterTypeString  ParameterTypeEnum = "STRING"
	ParameterTypeInteger ParameterTypeEnum = "INTEGER"
	ParameterTypeFloat   ParameterTypeEnum = "FLOAT"
	ParameterTypeBoolean ParameterTypeEnum = "BOOLEAN"
	ParameterTypeSecure  ParameterTypeEnum = "SECURE"
)

var mappingParameterTypeEnum = map[string]ParameterTypeEnum{
	"STRING":  ParameterTypeString,
	"INTEGER": ParameterTypeInteger,
	"FLOAT":   ParameterTypeFloat,
	"BOOLEAN": ParameterTypeBoolean,
	"SECURE":  ParameterTypeSecure,
}

var mappingParameterTypeEnumLowerCase = map[string]ParameterTypeEnum{
	"string":  ParameterTypeString,
	"integer": ParameterTypeInteger,
	"float":   ParameterTypeFloat,
	"boolean": ParameterTypeBoolean,
	"secure":  ParameterTypeSecure,
}

// GetParameterTypeEnumValues Enumerates the set of values for ParameterTypeEnum
func GetParameterTypeEnumValues() []ParameterTypeEnum {
	values := make([]ParameterTypeEnum, 0)
	for _, v := range mappingParameterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterTypeEnumStringValues Enumerates the set of values in String for ParameterTypeEnum
func GetParameterTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"INTEGER",
		"FLOAT",
		"BOOLEAN",
		"SECURE",
	}
}

// GetMappingParameterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterTypeEnum(val string) (ParameterTypeEnum, bool) {
	enum, ok := mappingParameterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
