// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// IdentityTypeEnum Enum with underlying type: string
type IdentityTypeEnum string

// Set of constants representing the allowable values for IdentityTypeEnum
const (
	IdentityTypeOracleDatabaseResourcePrincipal IdentityTypeEnum = "ORACLE_DATABASE_RESOURCE_PRINCIPAL"
)

var mappingIdentityTypeEnum = map[string]IdentityTypeEnum{
	"ORACLE_DATABASE_RESOURCE_PRINCIPAL": IdentityTypeOracleDatabaseResourcePrincipal,
}

var mappingIdentityTypeEnumLowerCase = map[string]IdentityTypeEnum{
	"oracle_database_resource_principal": IdentityTypeOracleDatabaseResourcePrincipal,
}

// GetIdentityTypeEnumValues Enumerates the set of values for IdentityTypeEnum
func GetIdentityTypeEnumValues() []IdentityTypeEnum {
	values := make([]IdentityTypeEnum, 0)
	for _, v := range mappingIdentityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIdentityTypeEnumStringValues Enumerates the set of values in String for IdentityTypeEnum
func GetIdentityTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DATABASE_RESOURCE_PRINCIPAL",
	}
}

// GetMappingIdentityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIdentityTypeEnum(val string) (IdentityTypeEnum, bool) {
	enum, ok := mappingIdentityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
