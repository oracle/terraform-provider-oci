// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
	SortByDisplayName SortByEnum = "displayName"
	SortByTimeCreated SortByEnum = "timeCreated"
)

var mappingSortByEnum = map[string]SortByEnum{
	"displayName": SortByDisplayName,
	"timeCreated": SortByTimeCreated,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
	values := make([]SortByEnum, 0)
	for _, v := range mappingSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByEnumStringValues Enumerates the set of values in String for SortByEnum
func GetSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByEnum(val string) (SortByEnum, bool) {
	mappingSortByEnumIgnoreCase := make(map[string]SortByEnum)
	for k, v := range mappingSortByEnum {
		mappingSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
