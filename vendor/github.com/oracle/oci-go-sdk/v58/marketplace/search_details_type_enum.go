// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// SearchDetailsTypeEnumEnum Enum with underlying type: string
type SearchDetailsTypeEnumEnum string

// Set of constants representing the allowable values for SearchDetailsTypeEnumEnum
const (
	SearchDetailsTypeEnumFreeText   SearchDetailsTypeEnumEnum = "FreeText"
	SearchDetailsTypeEnumStructured SearchDetailsTypeEnumEnum = "Structured"
)

var mappingSearchDetailsTypeEnumEnum = map[string]SearchDetailsTypeEnumEnum{
	"FreeText":   SearchDetailsTypeEnumFreeText,
	"Structured": SearchDetailsTypeEnumStructured,
}

// GetSearchDetailsTypeEnumEnumValues Enumerates the set of values for SearchDetailsTypeEnumEnum
func GetSearchDetailsTypeEnumEnumValues() []SearchDetailsTypeEnumEnum {
	values := make([]SearchDetailsTypeEnumEnum, 0)
	for _, v := range mappingSearchDetailsTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchDetailsTypeEnumEnumStringValues Enumerates the set of values in String for SearchDetailsTypeEnumEnum
func GetSearchDetailsTypeEnumEnumStringValues() []string {
	return []string{
		"FreeText",
		"Structured",
	}
}

// GetMappingSearchDetailsTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchDetailsTypeEnumEnum(val string) (SearchDetailsTypeEnumEnum, bool) {
	mappingSearchDetailsTypeEnumEnumIgnoreCase := make(map[string]SearchDetailsTypeEnumEnum)
	for k, v := range mappingSearchDetailsTypeEnumEnum {
		mappingSearchDetailsTypeEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSearchDetailsTypeEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
