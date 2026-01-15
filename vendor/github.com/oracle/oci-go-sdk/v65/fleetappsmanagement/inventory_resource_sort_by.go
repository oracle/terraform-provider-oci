// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// InventoryResourceSortByEnum Enum with underlying type: string
type InventoryResourceSortByEnum string

// Set of constants representing the allowable values for InventoryResourceSortByEnum
const (
	InventoryResourceSortByTimeCreated InventoryResourceSortByEnum = "timeCreated"
	InventoryResourceSortByDisplayName InventoryResourceSortByEnum = "displayName"
)

var mappingInventoryResourceSortByEnum = map[string]InventoryResourceSortByEnum{
	"timeCreated": InventoryResourceSortByTimeCreated,
	"displayName": InventoryResourceSortByDisplayName,
}

var mappingInventoryResourceSortByEnumLowerCase = map[string]InventoryResourceSortByEnum{
	"timecreated": InventoryResourceSortByTimeCreated,
	"displayname": InventoryResourceSortByDisplayName,
}

// GetInventoryResourceSortByEnumValues Enumerates the set of values for InventoryResourceSortByEnum
func GetInventoryResourceSortByEnumValues() []InventoryResourceSortByEnum {
	values := make([]InventoryResourceSortByEnum, 0)
	for _, v := range mappingInventoryResourceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetInventoryResourceSortByEnumStringValues Enumerates the set of values in String for InventoryResourceSortByEnum
func GetInventoryResourceSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingInventoryResourceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInventoryResourceSortByEnum(val string) (InventoryResourceSortByEnum, bool) {
	enum, ok := mappingInventoryResourceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
