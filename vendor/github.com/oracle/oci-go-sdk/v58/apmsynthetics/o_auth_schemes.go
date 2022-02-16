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

// OAuthSchemesEnum Enum with underlying type: string
type OAuthSchemesEnum string

// Set of constants representing the allowable values for OAuthSchemesEnum
const (
	OAuthSchemesNone  OAuthSchemesEnum = "NONE"
	OAuthSchemesBasic OAuthSchemesEnum = "BASIC"
)

var mappingOAuthSchemesEnum = map[string]OAuthSchemesEnum{
	"NONE":  OAuthSchemesNone,
	"BASIC": OAuthSchemesBasic,
}

// GetOAuthSchemesEnumValues Enumerates the set of values for OAuthSchemesEnum
func GetOAuthSchemesEnumValues() []OAuthSchemesEnum {
	values := make([]OAuthSchemesEnum, 0)
	for _, v := range mappingOAuthSchemesEnum {
		values = append(values, v)
	}
	return values
}

// GetOAuthSchemesEnumStringValues Enumerates the set of values in String for OAuthSchemesEnum
func GetOAuthSchemesEnumStringValues() []string {
	return []string{
		"NONE",
		"BASIC",
	}
}

// GetMappingOAuthSchemesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOAuthSchemesEnum(val string) (OAuthSchemesEnum, bool) {
	mappingOAuthSchemesEnumIgnoreCase := make(map[string]OAuthSchemesEnum)
	for k, v := range mappingOAuthSchemesEnum {
		mappingOAuthSchemesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingOAuthSchemesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
