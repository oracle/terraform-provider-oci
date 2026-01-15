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

// CatalogItemSortByEnum Enum with underlying type: string
type CatalogItemSortByEnum string

// Set of constants representing the allowable values for CatalogItemSortByEnum
const (
	CatalogItemSortByTimeCreated             CatalogItemSortByEnum = "timeCreated"
	CatalogItemSortByDisplayName             CatalogItemSortByEnum = "displayName"
	CatalogItemSortByTimeBackfillLastChecked CatalogItemSortByEnum = "timeBackfillLastChecked"
)

var mappingCatalogItemSortByEnum = map[string]CatalogItemSortByEnum{
	"timeCreated":             CatalogItemSortByTimeCreated,
	"displayName":             CatalogItemSortByDisplayName,
	"timeBackfillLastChecked": CatalogItemSortByTimeBackfillLastChecked,
}

var mappingCatalogItemSortByEnumLowerCase = map[string]CatalogItemSortByEnum{
	"timecreated":             CatalogItemSortByTimeCreated,
	"displayname":             CatalogItemSortByDisplayName,
	"timebackfilllastchecked": CatalogItemSortByTimeBackfillLastChecked,
}

// GetCatalogItemSortByEnumValues Enumerates the set of values for CatalogItemSortByEnum
func GetCatalogItemSortByEnumValues() []CatalogItemSortByEnum {
	values := make([]CatalogItemSortByEnum, 0)
	for _, v := range mappingCatalogItemSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogItemSortByEnumStringValues Enumerates the set of values in String for CatalogItemSortByEnum
func GetCatalogItemSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"timeBackfillLastChecked",
	}
}

// GetMappingCatalogItemSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogItemSortByEnum(val string) (CatalogItemSortByEnum, bool) {
	enum, ok := mappingCatalogItemSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
