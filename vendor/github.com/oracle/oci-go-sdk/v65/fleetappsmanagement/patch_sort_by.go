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

// PatchSortByEnum Enum with underlying type: string
type PatchSortByEnum string

// Set of constants representing the allowable values for PatchSortByEnum
const (
	PatchSortByTimeCreated PatchSortByEnum = "timeCreated"
	PatchSortByName        PatchSortByEnum = "name"
)

var mappingPatchSortByEnum = map[string]PatchSortByEnum{
	"timeCreated": PatchSortByTimeCreated,
	"name":        PatchSortByName,
}

var mappingPatchSortByEnumLowerCase = map[string]PatchSortByEnum{
	"timecreated": PatchSortByTimeCreated,
	"name":        PatchSortByName,
}

// GetPatchSortByEnumValues Enumerates the set of values for PatchSortByEnum
func GetPatchSortByEnumValues() []PatchSortByEnum {
	values := make([]PatchSortByEnum, 0)
	for _, v := range mappingPatchSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSortByEnumStringValues Enumerates the set of values in String for PatchSortByEnum
func GetPatchSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingPatchSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSortByEnum(val string) (PatchSortByEnum, bool) {
	enum, ok := mappingPatchSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
