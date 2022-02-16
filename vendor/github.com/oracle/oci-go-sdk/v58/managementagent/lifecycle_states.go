// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// LifecycleStatesEnum Enum with underlying type: string
type LifecycleStatesEnum string

// Set of constants representing the allowable values for LifecycleStatesEnum
const (
	LifecycleStatesCreating   LifecycleStatesEnum = "CREATING"
	LifecycleStatesUpdating   LifecycleStatesEnum = "UPDATING"
	LifecycleStatesActive     LifecycleStatesEnum = "ACTIVE"
	LifecycleStatesInactive   LifecycleStatesEnum = "INACTIVE"
	LifecycleStatesTerminated LifecycleStatesEnum = "TERMINATED"
	LifecycleStatesDeleting   LifecycleStatesEnum = "DELETING"
	LifecycleStatesDeleted    LifecycleStatesEnum = "DELETED"
	LifecycleStatesFailed     LifecycleStatesEnum = "FAILED"
)

var mappingLifecycleStatesEnum = map[string]LifecycleStatesEnum{
	"CREATING":   LifecycleStatesCreating,
	"UPDATING":   LifecycleStatesUpdating,
	"ACTIVE":     LifecycleStatesActive,
	"INACTIVE":   LifecycleStatesInactive,
	"TERMINATED": LifecycleStatesTerminated,
	"DELETING":   LifecycleStatesDeleting,
	"DELETED":    LifecycleStatesDeleted,
	"FAILED":     LifecycleStatesFailed,
}

// GetLifecycleStatesEnumValues Enumerates the set of values for LifecycleStatesEnum
func GetLifecycleStatesEnumValues() []LifecycleStatesEnum {
	values := make([]LifecycleStatesEnum, 0)
	for _, v := range mappingLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleStatesEnumStringValues Enumerates the set of values in String for LifecycleStatesEnum
func GetLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"TERMINATED",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleStatesEnum(val string) (LifecycleStatesEnum, bool) {
	mappingLifecycleStatesEnumIgnoreCase := make(map[string]LifecycleStatesEnum)
	for k, v := range mappingLifecycleStatesEnum {
		mappingLifecycleStatesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLifecycleStatesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
