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

// AuthenticationTypeEnum Enum with underlying type: string
type AuthenticationTypeEnum string

// Set of constants representing the allowable values for AuthenticationTypeEnum
const (
	AuthenticationTypeToken    AuthenticationTypeEnum = "TOKEN"
	AuthenticationTypePassword AuthenticationTypeEnum = "PASSWORD"
)

var mappingAuthenticationTypeEnum = map[string]AuthenticationTypeEnum{
	"TOKEN":    AuthenticationTypeToken,
	"PASSWORD": AuthenticationTypePassword,
}

var mappingAuthenticationTypeEnumLowerCase = map[string]AuthenticationTypeEnum{
	"token":    AuthenticationTypeToken,
	"password": AuthenticationTypePassword,
}

// GetAuthenticationTypeEnumValues Enumerates the set of values for AuthenticationTypeEnum
func GetAuthenticationTypeEnumValues() []AuthenticationTypeEnum {
	values := make([]AuthenticationTypeEnum, 0)
	for _, v := range mappingAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationTypeEnumStringValues Enumerates the set of values in String for AuthenticationTypeEnum
func GetAuthenticationTypeEnumStringValues() []string {
	return []string{
		"TOKEN",
		"PASSWORD",
	}
}

// GetMappingAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationTypeEnum(val string) (AuthenticationTypeEnum, bool) {
	enum, ok := mappingAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
