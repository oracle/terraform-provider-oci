// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// EntityLifecycleStatesEnum Enum with underlying type: string
type EntityLifecycleStatesEnum string

// Set of constants representing the allowable values for EntityLifecycleStatesEnum
const (
	EntityLifecycleStatesActive  EntityLifecycleStatesEnum = "ACTIVE"
	EntityLifecycleStatesDeleted EntityLifecycleStatesEnum = "DELETED"
)

var mappingEntityLifecycleStatesEnum = map[string]EntityLifecycleStatesEnum{
	"ACTIVE":  EntityLifecycleStatesActive,
	"DELETED": EntityLifecycleStatesDeleted,
}

// GetEntityLifecycleStatesEnumValues Enumerates the set of values for EntityLifecycleStatesEnum
func GetEntityLifecycleStatesEnumValues() []EntityLifecycleStatesEnum {
	values := make([]EntityLifecycleStatesEnum, 0)
	for _, v := range mappingEntityLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityLifecycleStatesEnumStringValues Enumerates the set of values in String for EntityLifecycleStatesEnum
func GetEntityLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingEntityLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityLifecycleStatesEnum(val string) (EntityLifecycleStatesEnum, bool) {
	mappingEntityLifecycleStatesEnumIgnoreCase := make(map[string]EntityLifecycleStatesEnum)
	for k, v := range mappingEntityLifecycleStatesEnum {
		mappingEntityLifecycleStatesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingEntityLifecycleStatesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
