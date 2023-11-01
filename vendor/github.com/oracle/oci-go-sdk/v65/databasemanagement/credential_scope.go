// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// CredentialScopeEnum Enum with underlying type: string
type CredentialScopeEnum string

// Set of constants representing the allowable values for CredentialScopeEnum
const (
	CredentialScopeResource CredentialScopeEnum = "RESOURCE"
	CredentialScopeGlobal   CredentialScopeEnum = "GLOBAL"
)

var mappingCredentialScopeEnum = map[string]CredentialScopeEnum{
	"RESOURCE": CredentialScopeResource,
	"GLOBAL":   CredentialScopeGlobal,
}

var mappingCredentialScopeEnumLowerCase = map[string]CredentialScopeEnum{
	"resource": CredentialScopeResource,
	"global":   CredentialScopeGlobal,
}

// GetCredentialScopeEnumValues Enumerates the set of values for CredentialScopeEnum
func GetCredentialScopeEnumValues() []CredentialScopeEnum {
	values := make([]CredentialScopeEnum, 0)
	for _, v := range mappingCredentialScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialScopeEnumStringValues Enumerates the set of values in String for CredentialScopeEnum
func GetCredentialScopeEnumStringValues() []string {
	return []string{
		"RESOURCE",
		"GLOBAL",
	}
}

// GetMappingCredentialScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialScopeEnum(val string) (CredentialScopeEnum, bool) {
	enum, ok := mappingCredentialScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
