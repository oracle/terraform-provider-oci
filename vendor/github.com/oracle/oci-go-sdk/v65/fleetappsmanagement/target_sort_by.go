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

// TargetSortByEnum Enum with underlying type: string
type TargetSortByEnum string

// Set of constants representing the allowable values for TargetSortByEnum
const (
	TargetSortByDisplayName         TargetSortByEnum = "displayName"
	TargetSortByProduct             TargetSortByEnum = "product"
	TargetSortByResourceDisplayName TargetSortByEnum = "resourceDisplayName"
)

var mappingTargetSortByEnum = map[string]TargetSortByEnum{
	"displayName":         TargetSortByDisplayName,
	"product":             TargetSortByProduct,
	"resourceDisplayName": TargetSortByResourceDisplayName,
}

var mappingTargetSortByEnumLowerCase = map[string]TargetSortByEnum{
	"displayname":         TargetSortByDisplayName,
	"product":             TargetSortByProduct,
	"resourcedisplayname": TargetSortByResourceDisplayName,
}

// GetTargetSortByEnumValues Enumerates the set of values for TargetSortByEnum
func GetTargetSortByEnumValues() []TargetSortByEnum {
	values := make([]TargetSortByEnum, 0)
	for _, v := range mappingTargetSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetSortByEnumStringValues Enumerates the set of values in String for TargetSortByEnum
func GetTargetSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"product",
		"resourceDisplayName",
	}
}

// GetMappingTargetSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetSortByEnum(val string) (TargetSortByEnum, bool) {
	enum, ok := mappingTargetSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
