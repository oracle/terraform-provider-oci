// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RunbookSortByEnum Enum with underlying type: string
type RunbookSortByEnum string

// Set of constants representing the allowable values for RunbookSortByEnum
const (
	RunbookSortByTimeCreated RunbookSortByEnum = "timeCreated"
	RunbookSortByDisplayName RunbookSortByEnum = "displayName"
)

var mappingRunbookSortByEnum = map[string]RunbookSortByEnum{
	"timeCreated": RunbookSortByTimeCreated,
	"displayName": RunbookSortByDisplayName,
}

var mappingRunbookSortByEnumLowerCase = map[string]RunbookSortByEnum{
	"timecreated": RunbookSortByTimeCreated,
	"displayname": RunbookSortByDisplayName,
}

// GetRunbookSortByEnumValues Enumerates the set of values for RunbookSortByEnum
func GetRunbookSortByEnumValues() []RunbookSortByEnum {
	values := make([]RunbookSortByEnum, 0)
	for _, v := range mappingRunbookSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookSortByEnumStringValues Enumerates the set of values in String for RunbookSortByEnum
func GetRunbookSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingRunbookSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookSortByEnum(val string) (RunbookSortByEnum, bool) {
	enum, ok := mappingRunbookSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
