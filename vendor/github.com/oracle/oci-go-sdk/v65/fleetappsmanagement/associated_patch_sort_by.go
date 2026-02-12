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

// AssociatedPatchSortByEnum Enum with underlying type: string
type AssociatedPatchSortByEnum string

// Set of constants representing the allowable values for AssociatedPatchSortByEnum
const (
	AssociatedPatchSortByTimeCreated AssociatedPatchSortByEnum = "timeCreated"
	AssociatedPatchSortByPatchName   AssociatedPatchSortByEnum = "patchName"
)

var mappingAssociatedPatchSortByEnum = map[string]AssociatedPatchSortByEnum{
	"timeCreated": AssociatedPatchSortByTimeCreated,
	"patchName":   AssociatedPatchSortByPatchName,
}

var mappingAssociatedPatchSortByEnumLowerCase = map[string]AssociatedPatchSortByEnum{
	"timecreated": AssociatedPatchSortByTimeCreated,
	"patchname":   AssociatedPatchSortByPatchName,
}

// GetAssociatedPatchSortByEnumValues Enumerates the set of values for AssociatedPatchSortByEnum
func GetAssociatedPatchSortByEnumValues() []AssociatedPatchSortByEnum {
	values := make([]AssociatedPatchSortByEnum, 0)
	for _, v := range mappingAssociatedPatchSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedPatchSortByEnumStringValues Enumerates the set of values in String for AssociatedPatchSortByEnum
func GetAssociatedPatchSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"patchName",
	}
}

// GetMappingAssociatedPatchSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedPatchSortByEnum(val string) (AssociatedPatchSortByEnum, bool) {
	enum, ok := mappingAssociatedPatchSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
