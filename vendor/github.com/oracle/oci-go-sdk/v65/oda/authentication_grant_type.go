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

// AuthenticationGrantTypeEnum Enum with underlying type: string
type AuthenticationGrantTypeEnum string

// Set of constants representing the allowable values for AuthenticationGrantTypeEnum
const (
	AuthenticationGrantTypeClientCredentials AuthenticationGrantTypeEnum = "CLIENT_CREDENTIALS"
	AuthenticationGrantTypeAuthorizationCode AuthenticationGrantTypeEnum = "AUTHORIZATION_CODE"
)

var mappingAuthenticationGrantTypeEnum = map[string]AuthenticationGrantTypeEnum{
	"CLIENT_CREDENTIALS": AuthenticationGrantTypeClientCredentials,
	"AUTHORIZATION_CODE": AuthenticationGrantTypeAuthorizationCode,
}

var mappingAuthenticationGrantTypeEnumLowerCase = map[string]AuthenticationGrantTypeEnum{
	"client_credentials": AuthenticationGrantTypeClientCredentials,
	"authorization_code": AuthenticationGrantTypeAuthorizationCode,
}

// GetAuthenticationGrantTypeEnumValues Enumerates the set of values for AuthenticationGrantTypeEnum
func GetAuthenticationGrantTypeEnumValues() []AuthenticationGrantTypeEnum {
	values := make([]AuthenticationGrantTypeEnum, 0)
	for _, v := range mappingAuthenticationGrantTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthenticationGrantTypeEnumStringValues Enumerates the set of values in String for AuthenticationGrantTypeEnum
func GetAuthenticationGrantTypeEnumStringValues() []string {
	return []string{
		"CLIENT_CREDENTIALS",
		"AUTHORIZATION_CODE",
	}
}

// GetMappingAuthenticationGrantTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthenticationGrantTypeEnum(val string) (AuthenticationGrantTypeEnum, bool) {
	enum, ok := mappingAuthenticationGrantTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
