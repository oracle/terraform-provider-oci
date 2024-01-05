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

// AuthenticationIdentityProviderEnum Enum with underlying type: string
type AuthenticationIdentityProviderEnum string

// Set of constants representing the allowable values for AuthenticationIdentityProviderEnum
const (
	AuthenticationIdentityProviderGeneric   AuthenticationIdentityProviderEnum = "GENERIC"
	AuthenticationIdentityProviderOam       AuthenticationIdentityProviderEnum = "OAM"
	AuthenticationIdentityProviderGoogle    AuthenticationIdentityProviderEnum = "GOOGLE"
	AuthenticationIdentityProviderMicrosoft AuthenticationIdentityProviderEnum = "MICROSOFT"
)

var mappingAuthenticationIdentityProviderEnum = map[string]AuthenticationIdentityProviderEnum{
	"GENERIC":   AuthenticationIdentityProviderGeneric,
	"OAM":       AuthenticationIdentityProviderOam,
	"GOOGLE":    AuthenticationIdentityProviderGoogle,
	"MICROSOFT": AuthenticationIdentityProviderMicrosoft,
}

var mappingAuthenticationIdentityProviderEnumLowerCase = map[string]AuthenticationIdentityProviderEnum{
	"generic":   AuthenticationIdentityProviderGeneric,
	"oam":       AuthenticationIdentityProviderOam,
	"google":    AuthenticationIdentityProviderGoogle,
	"microsoft": AuthenticationIdentityProviderMicrosoft,
}

// GetAuthenticationIdentityProviderEnumValues Enumerates the set of values for AuthenticationIdentityProviderEnum
func GetAuthenticationIdentityProviderEnumValues() []AuthenticationIdentityProviderEnum {
	values := make([]AuthenticationIdentityProviderEnum, 0)
	for _, v := range mappingAuthenticationIdentityProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationIdentityProviderEnumStringValues Enumerates the set of values in String for AuthenticationIdentityProviderEnum
func GetAuthenticationIdentityProviderEnumStringValues() []string {
	return []string{
		"GENERIC",
		"OAM",
		"GOOGLE",
		"MICROSOFT",
	}
}

// GetMappingAuthenticationIdentityProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationIdentityProviderEnum(val string) (AuthenticationIdentityProviderEnum, bool) {
	enum, ok := mappingAuthenticationIdentityProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
