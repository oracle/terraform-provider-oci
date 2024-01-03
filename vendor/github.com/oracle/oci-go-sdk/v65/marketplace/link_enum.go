// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// LinkEnumEnum Enum with underlying type: string
type LinkEnumEnum string

// Set of constants representing the allowable values for LinkEnumEnum
const (
	LinkEnumSelf      LinkEnumEnum = "SELF"
	LinkEnumCanonical LinkEnumEnum = "CANONICAL"
	LinkEnumNext      LinkEnumEnum = "NEXT"
	LinkEnumTemplate  LinkEnumEnum = "TEMPLATE"
	LinkEnumPrev      LinkEnumEnum = "PREV"
)

var mappingLinkEnumEnum = map[string]LinkEnumEnum{
	"SELF":      LinkEnumSelf,
	"CANONICAL": LinkEnumCanonical,
	"NEXT":      LinkEnumNext,
	"TEMPLATE":  LinkEnumTemplate,
	"PREV":      LinkEnumPrev,
}

var mappingLinkEnumEnumLowerCase = map[string]LinkEnumEnum{
	"self":      LinkEnumSelf,
	"canonical": LinkEnumCanonical,
	"next":      LinkEnumNext,
	"template":  LinkEnumTemplate,
	"prev":      LinkEnumPrev,
}

// GetLinkEnumEnumValues Enumerates the set of values for LinkEnumEnum
func GetLinkEnumEnumValues() []LinkEnumEnum {
	values := make([]LinkEnumEnum, 0)
	for _, v := range mappingLinkEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetLinkEnumEnumStringValues Enumerates the set of values in String for LinkEnumEnum
func GetLinkEnumEnumStringValues() []string {
	return []string{
		"SELF",
		"CANONICAL",
		"NEXT",
		"TEMPLATE",
		"PREV",
	}
}

// GetMappingLinkEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLinkEnumEnum(val string) (LinkEnumEnum, bool) {
	enum, ok := mappingLinkEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
