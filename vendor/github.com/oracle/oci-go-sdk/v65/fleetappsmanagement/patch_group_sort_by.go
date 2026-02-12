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

// PatchGroupSortByEnum Enum with underlying type: string
type PatchGroupSortByEnum string

// Set of constants representing the allowable values for PatchGroupSortByEnum
const (
	PatchGroupSortByTimeCreated PatchGroupSortByEnum = "timeCreated"
	PatchGroupSortByDisplayName PatchGroupSortByEnum = "displayName"
)

var mappingPatchGroupSortByEnum = map[string]PatchGroupSortByEnum{
	"timeCreated": PatchGroupSortByTimeCreated,
	"displayName": PatchGroupSortByDisplayName,
}

var mappingPatchGroupSortByEnumLowerCase = map[string]PatchGroupSortByEnum{
	"timecreated": PatchGroupSortByTimeCreated,
	"displayname": PatchGroupSortByDisplayName,
}

// GetPatchGroupSortByEnumValues Enumerates the set of values for PatchGroupSortByEnum
func GetPatchGroupSortByEnumValues() []PatchGroupSortByEnum {
	values := make([]PatchGroupSortByEnum, 0)
	for _, v := range mappingPatchGroupSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchGroupSortByEnumStringValues Enumerates the set of values in String for PatchGroupSortByEnum
func GetPatchGroupSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingPatchGroupSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchGroupSortByEnum(val string) (PatchGroupSortByEnum, bool) {
	enum, ok := mappingPatchGroupSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
