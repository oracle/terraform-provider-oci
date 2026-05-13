// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// CredentialRelatedResourceTypeEnum Enum with underlying type: string
type CredentialRelatedResourceTypeEnum string

// Set of constants representing the allowable values for CredentialRelatedResourceTypeEnum
const (
	CredentialRelatedResourceTypeDatabaseToolsIdentity CredentialRelatedResourceTypeEnum = "DATABASE_TOOLS_IDENTITY"
)

var mappingCredentialRelatedResourceTypeEnum = map[string]CredentialRelatedResourceTypeEnum{
	"DATABASE_TOOLS_IDENTITY": CredentialRelatedResourceTypeDatabaseToolsIdentity,
}

var mappingCredentialRelatedResourceTypeEnumLowerCase = map[string]CredentialRelatedResourceTypeEnum{
	"database_tools_identity": CredentialRelatedResourceTypeDatabaseToolsIdentity,
}

// GetCredentialRelatedResourceTypeEnumValues Enumerates the set of values for CredentialRelatedResourceTypeEnum
func GetCredentialRelatedResourceTypeEnumValues() []CredentialRelatedResourceTypeEnum {
	values := make([]CredentialRelatedResourceTypeEnum, 0)
	for _, v := range mappingCredentialRelatedResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialRelatedResourceTypeEnumStringValues Enumerates the set of values in String for CredentialRelatedResourceTypeEnum
func GetCredentialRelatedResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE_TOOLS_IDENTITY",
	}
}

// GetMappingCredentialRelatedResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialRelatedResourceTypeEnum(val string) (CredentialRelatedResourceTypeEnum, bool) {
	enum, ok := mappingCredentialRelatedResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
