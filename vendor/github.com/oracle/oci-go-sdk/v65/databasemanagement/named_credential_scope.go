// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// NamedCredentialScopeEnum Enum with underlying type: string
type NamedCredentialScopeEnum string

// Set of constants representing the allowable values for NamedCredentialScopeEnum
const (
	NamedCredentialScopeResource NamedCredentialScopeEnum = "RESOURCE"
	NamedCredentialScopeGlobal   NamedCredentialScopeEnum = "GLOBAL"
)

var mappingNamedCredentialScopeEnum = map[string]NamedCredentialScopeEnum{
	"RESOURCE": NamedCredentialScopeResource,
	"GLOBAL":   NamedCredentialScopeGlobal,
}

var mappingNamedCredentialScopeEnumLowerCase = map[string]NamedCredentialScopeEnum{
	"resource": NamedCredentialScopeResource,
	"global":   NamedCredentialScopeGlobal,
}

// GetNamedCredentialScopeEnumValues Enumerates the set of values for NamedCredentialScopeEnum
func GetNamedCredentialScopeEnumValues() []NamedCredentialScopeEnum {
	values := make([]NamedCredentialScopeEnum, 0)
	for _, v := range mappingNamedCredentialScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetNamedCredentialScopeEnumStringValues Enumerates the set of values in String for NamedCredentialScopeEnum
func GetNamedCredentialScopeEnumStringValues() []string {
	return []string{
		"RESOURCE",
		"GLOBAL",
	}
}

// GetMappingNamedCredentialScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNamedCredentialScopeEnum(val string) (NamedCredentialScopeEnum, bool) {
	enum, ok := mappingNamedCredentialScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
