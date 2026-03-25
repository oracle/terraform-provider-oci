// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// ResourceGroupSortByEnum Enum with underlying type: string
type ResourceGroupSortByEnum string

// Set of constants representing the allowable values for ResourceGroupSortByEnum
const (
	ResourceGroupSortById           ResourceGroupSortByEnum = "id"
	ResourceGroupSortByResourceName ResourceGroupSortByEnum = "resourceName"
	ResourceGroupSortByDisplayName  ResourceGroupSortByEnum = "displayName"
	ResourceGroupSortByDescription  ResourceGroupSortByEnum = "description"
	ResourceGroupSortByCapacity     ResourceGroupSortByEnum = "capacity"
	ResourceGroupSortByStatus       ResourceGroupSortByEnum = "status"
)

var mappingResourceGroupSortByEnum = map[string]ResourceGroupSortByEnum{
	"id":           ResourceGroupSortById,
	"resourceName": ResourceGroupSortByResourceName,
	"displayName":  ResourceGroupSortByDisplayName,
	"description":  ResourceGroupSortByDescription,
	"capacity":     ResourceGroupSortByCapacity,
	"status":       ResourceGroupSortByStatus,
}

var mappingResourceGroupSortByEnumLowerCase = map[string]ResourceGroupSortByEnum{
	"id":           ResourceGroupSortById,
	"resourcename": ResourceGroupSortByResourceName,
	"displayname":  ResourceGroupSortByDisplayName,
	"description":  ResourceGroupSortByDescription,
	"capacity":     ResourceGroupSortByCapacity,
	"status":       ResourceGroupSortByStatus,
}

// GetResourceGroupSortByEnumValues Enumerates the set of values for ResourceGroupSortByEnum
func GetResourceGroupSortByEnumValues() []ResourceGroupSortByEnum {
	values := make([]ResourceGroupSortByEnum, 0)
	for _, v := range mappingResourceGroupSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceGroupSortByEnumStringValues Enumerates the set of values in String for ResourceGroupSortByEnum
func GetResourceGroupSortByEnumStringValues() []string {
	return []string{
		"id",
		"resourceName",
		"displayName",
		"description",
		"capacity",
		"status",
	}
}

// GetMappingResourceGroupSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceGroupSortByEnum(val string) (ResourceGroupSortByEnum, bool) {
	enum, ok := mappingResourceGroupSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
