// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// RequestAuthenticationSchemesForScriptedRestEnum Enum with underlying type: string
type RequestAuthenticationSchemesForScriptedRestEnum string

// Set of constants representing the allowable values for RequestAuthenticationSchemesForScriptedRestEnum
const (
	RequestAuthenticationSchemesForScriptedRestNone              RequestAuthenticationSchemesForScriptedRestEnum = "NONE"
	RequestAuthenticationSchemesForScriptedRestResourcePrincipal RequestAuthenticationSchemesForScriptedRestEnum = "RESOURCE_PRINCIPAL"
)

var mappingRequestAuthenticationSchemesForScriptedRestEnum = map[string]RequestAuthenticationSchemesForScriptedRestEnum{
	"NONE":               RequestAuthenticationSchemesForScriptedRestNone,
	"RESOURCE_PRINCIPAL": RequestAuthenticationSchemesForScriptedRestResourcePrincipal,
}

var mappingRequestAuthenticationSchemesForScriptedRestEnumLowerCase = map[string]RequestAuthenticationSchemesForScriptedRestEnum{
	"none":               RequestAuthenticationSchemesForScriptedRestNone,
	"resource_principal": RequestAuthenticationSchemesForScriptedRestResourcePrincipal,
}

// GetRequestAuthenticationSchemesForScriptedRestEnumValues Enumerates the set of values for RequestAuthenticationSchemesForScriptedRestEnum
func GetRequestAuthenticationSchemesForScriptedRestEnumValues() []RequestAuthenticationSchemesForScriptedRestEnum {
	values := make([]RequestAuthenticationSchemesForScriptedRestEnum, 0)
	for _, v := range mappingRequestAuthenticationSchemesForScriptedRestEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestAuthenticationSchemesForScriptedRestEnumStringValues Enumerates the set of values in String for RequestAuthenticationSchemesForScriptedRestEnum
func GetRequestAuthenticationSchemesForScriptedRestEnumStringValues() []string {
	return []string{
		"NONE",
		"RESOURCE_PRINCIPAL",
	}
}

// GetMappingRequestAuthenticationSchemesForScriptedRestEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestAuthenticationSchemesForScriptedRestEnum(val string) (RequestAuthenticationSchemesForScriptedRestEnum, bool) {
	enum, ok := mappingRequestAuthenticationSchemesForScriptedRestEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
