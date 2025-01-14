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

// RuntimeIdentityEnum Enum with underlying type: string
type RuntimeIdentityEnum string

// Set of constants representing the allowable values for RuntimeIdentityEnum
const (
	RuntimeIdentityAuthenticatedPrincipal RuntimeIdentityEnum = "AUTHENTICATED_PRINCIPAL"
	RuntimeIdentityResourcePrincipal      RuntimeIdentityEnum = "RESOURCE_PRINCIPAL"
)

var mappingRuntimeIdentityEnum = map[string]RuntimeIdentityEnum{
	"AUTHENTICATED_PRINCIPAL": RuntimeIdentityAuthenticatedPrincipal,
	"RESOURCE_PRINCIPAL":      RuntimeIdentityResourcePrincipal,
}

var mappingRuntimeIdentityEnumLowerCase = map[string]RuntimeIdentityEnum{
	"authenticated_principal": RuntimeIdentityAuthenticatedPrincipal,
	"resource_principal":      RuntimeIdentityResourcePrincipal,
}

// GetRuntimeIdentityEnumValues Enumerates the set of values for RuntimeIdentityEnum
func GetRuntimeIdentityEnumValues() []RuntimeIdentityEnum {
	values := make([]RuntimeIdentityEnum, 0)
	for _, v := range mappingRuntimeIdentityEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeIdentityEnumStringValues Enumerates the set of values in String for RuntimeIdentityEnum
func GetRuntimeIdentityEnumStringValues() []string {
	return []string{
		"AUTHENTICATED_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingRuntimeIdentityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeIdentityEnum(val string) (RuntimeIdentityEnum, bool) {
	enum, ok := mappingRuntimeIdentityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
