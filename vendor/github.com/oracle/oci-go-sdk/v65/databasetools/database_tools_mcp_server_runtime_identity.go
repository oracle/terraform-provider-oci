// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsMcpServerRuntimeIdentityEnum Enum with underlying type: string
type DatabaseToolsMcpServerRuntimeIdentityEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpServerRuntimeIdentityEnum
const (
	DatabaseToolsMcpServerRuntimeIdentityAuthenticatedPrincipal DatabaseToolsMcpServerRuntimeIdentityEnum = "AUTHENTICATED_PRINCIPAL"
	DatabaseToolsMcpServerRuntimeIdentityResourcePrincipal      DatabaseToolsMcpServerRuntimeIdentityEnum = "RESOURCE_PRINCIPAL"
)

var mappingDatabaseToolsMcpServerRuntimeIdentityEnum = map[string]DatabaseToolsMcpServerRuntimeIdentityEnum{
	"AUTHENTICATED_PRINCIPAL": DatabaseToolsMcpServerRuntimeIdentityAuthenticatedPrincipal,
	"RESOURCE_PRINCIPAL":      DatabaseToolsMcpServerRuntimeIdentityResourcePrincipal,
}

var mappingDatabaseToolsMcpServerRuntimeIdentityEnumLowerCase = map[string]DatabaseToolsMcpServerRuntimeIdentityEnum{
	"authenticated_principal": DatabaseToolsMcpServerRuntimeIdentityAuthenticatedPrincipal,
	"resource_principal":      DatabaseToolsMcpServerRuntimeIdentityResourcePrincipal,
}

// GetDatabaseToolsMcpServerRuntimeIdentityEnumValues Enumerates the set of values for DatabaseToolsMcpServerRuntimeIdentityEnum
func GetDatabaseToolsMcpServerRuntimeIdentityEnumValues() []DatabaseToolsMcpServerRuntimeIdentityEnum {
	values := make([]DatabaseToolsMcpServerRuntimeIdentityEnum, 0)
	for _, v := range mappingDatabaseToolsMcpServerRuntimeIdentityEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpServerRuntimeIdentityEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpServerRuntimeIdentityEnum
func GetDatabaseToolsMcpServerRuntimeIdentityEnumStringValues() []string {
	return []string{
		"AUTHENTICATED_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingDatabaseToolsMcpServerRuntimeIdentityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpServerRuntimeIdentityEnum(val string) (DatabaseToolsMcpServerRuntimeIdentityEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpServerRuntimeIdentityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
