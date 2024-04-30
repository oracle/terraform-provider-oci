// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// ParameterScopeEnum Enum with underlying type: string
type ParameterScopeEnum string

// Set of constants representing the allowable values for ParameterScopeEnum
const (
	ParameterScopeMemory ParameterScopeEnum = "MEMORY"
	ParameterScopeSpfile ParameterScopeEnum = "SPFILE"
	ParameterScopeBoth   ParameterScopeEnum = "BOTH"
)

var mappingParameterScopeEnum = map[string]ParameterScopeEnum{
	"MEMORY": ParameterScopeMemory,
	"SPFILE": ParameterScopeSpfile,
	"BOTH":   ParameterScopeBoth,
}

var mappingParameterScopeEnumLowerCase = map[string]ParameterScopeEnum{
	"memory": ParameterScopeMemory,
	"spfile": ParameterScopeSpfile,
	"both":   ParameterScopeBoth,
}

// GetParameterScopeEnumValues Enumerates the set of values for ParameterScopeEnum
func GetParameterScopeEnumValues() []ParameterScopeEnum {
	values := make([]ParameterScopeEnum, 0)
	for _, v := range mappingParameterScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetParameterScopeEnumStringValues Enumerates the set of values in String for ParameterScopeEnum
func GetParameterScopeEnumStringValues() []string {
	return []string{
		"MEMORY",
		"SPFILE",
		"BOTH",
	}
}

// GetMappingParameterScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingParameterScopeEnum(val string) (ParameterScopeEnum, bool) {
	enum, ok := mappingParameterScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
