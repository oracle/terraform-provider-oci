// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the <a href="https://docs.oracle.com/en-us/iaas/jms/doc/java-download.html">Java Download</a> feature of Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// TokenSortByEnum Enum with underlying type: string
type TokenSortByEnum string

// Set of constants representing the allowable values for TokenSortByEnum
const (
	TokenSortByTimeCreated TokenSortByEnum = "timeCreated"
	TokenSortByTimeExpires TokenSortByEnum = "timeExpires"
	TokenSortByState       TokenSortByEnum = "state"
	TokenSortByDisplayName TokenSortByEnum = "displayName"
	TokenSortByJavaVersion TokenSortByEnum = "javaVersion"
)

var mappingTokenSortByEnum = map[string]TokenSortByEnum{
	"timeCreated": TokenSortByTimeCreated,
	"timeExpires": TokenSortByTimeExpires,
	"state":       TokenSortByState,
	"displayName": TokenSortByDisplayName,
	"javaVersion": TokenSortByJavaVersion,
}

var mappingTokenSortByEnumLowerCase = map[string]TokenSortByEnum{
	"timecreated": TokenSortByTimeCreated,
	"timeexpires": TokenSortByTimeExpires,
	"state":       TokenSortByState,
	"displayname": TokenSortByDisplayName,
	"javaversion": TokenSortByJavaVersion,
}

// GetTokenSortByEnumValues Enumerates the set of values for TokenSortByEnum
func GetTokenSortByEnumValues() []TokenSortByEnum {
	values := make([]TokenSortByEnum, 0)
	for _, v := range mappingTokenSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetTokenSortByEnumStringValues Enumerates the set of values in String for TokenSortByEnum
func GetTokenSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeExpires",
		"state",
		"displayName",
		"javaVersion",
	}
}

// GetMappingTokenSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTokenSortByEnum(val string) (TokenSortByEnum, bool) {
	enum, ok := mappingTokenSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
