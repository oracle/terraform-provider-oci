// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"strings"
)

// RequestAuthenticationSchemesEnum Enum with underlying type: string
type RequestAuthenticationSchemesEnum string

// Set of constants representing the allowable values for RequestAuthenticationSchemesEnum
const (
	RequestAuthenticationSchemesOauth  RequestAuthenticationSchemesEnum = "OAUTH"
	RequestAuthenticationSchemesNone   RequestAuthenticationSchemesEnum = "NONE"
	RequestAuthenticationSchemesBasic  RequestAuthenticationSchemesEnum = "BASIC"
	RequestAuthenticationSchemesBearer RequestAuthenticationSchemesEnum = "BEARER"
)

var mappingRequestAuthenticationSchemesEnum = map[string]RequestAuthenticationSchemesEnum{
	"OAUTH":  RequestAuthenticationSchemesOauth,
	"NONE":   RequestAuthenticationSchemesNone,
	"BASIC":  RequestAuthenticationSchemesBasic,
	"BEARER": RequestAuthenticationSchemesBearer,
}

// GetRequestAuthenticationSchemesEnumValues Enumerates the set of values for RequestAuthenticationSchemesEnum
func GetRequestAuthenticationSchemesEnumValues() []RequestAuthenticationSchemesEnum {
	values := make([]RequestAuthenticationSchemesEnum, 0)
	for _, v := range mappingRequestAuthenticationSchemesEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestAuthenticationSchemesEnumStringValues Enumerates the set of values in String for RequestAuthenticationSchemesEnum
func GetRequestAuthenticationSchemesEnumStringValues() []string {
	return []string{
		"OAUTH",
		"NONE",
		"BASIC",
		"BEARER",
	}
}

// GetMappingRequestAuthenticationSchemesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestAuthenticationSchemesEnum(val string) (RequestAuthenticationSchemesEnum, bool) {
	mappingRequestAuthenticationSchemesEnumIgnoreCase := make(map[string]RequestAuthenticationSchemesEnum)
	for k, v := range mappingRequestAuthenticationSchemesEnum {
		mappingRequestAuthenticationSchemesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRequestAuthenticationSchemesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
