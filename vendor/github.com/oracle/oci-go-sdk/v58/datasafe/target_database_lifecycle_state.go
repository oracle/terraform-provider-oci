// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// TargetDatabaseLifecycleStateEnum Enum with underlying type: string
type TargetDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for TargetDatabaseLifecycleStateEnum
const (
	TargetDatabaseLifecycleStateCreating       TargetDatabaseLifecycleStateEnum = "CREATING"
	TargetDatabaseLifecycleStateUpdating       TargetDatabaseLifecycleStateEnum = "UPDATING"
	TargetDatabaseLifecycleStateActive         TargetDatabaseLifecycleStateEnum = "ACTIVE"
	TargetDatabaseLifecycleStateInactive       TargetDatabaseLifecycleStateEnum = "INACTIVE"
	TargetDatabaseLifecycleStateDeleting       TargetDatabaseLifecycleStateEnum = "DELETING"
	TargetDatabaseLifecycleStateDeleted        TargetDatabaseLifecycleStateEnum = "DELETED"
	TargetDatabaseLifecycleStateNeedsAttention TargetDatabaseLifecycleStateEnum = "NEEDS_ATTENTION"
	TargetDatabaseLifecycleStateFailed         TargetDatabaseLifecycleStateEnum = "FAILED"
)

var mappingTargetDatabaseLifecycleStateEnum = map[string]TargetDatabaseLifecycleStateEnum{
	"CREATING":        TargetDatabaseLifecycleStateCreating,
	"UPDATING":        TargetDatabaseLifecycleStateUpdating,
	"ACTIVE":          TargetDatabaseLifecycleStateActive,
	"INACTIVE":        TargetDatabaseLifecycleStateInactive,
	"DELETING":        TargetDatabaseLifecycleStateDeleting,
	"DELETED":         TargetDatabaseLifecycleStateDeleted,
	"NEEDS_ATTENTION": TargetDatabaseLifecycleStateNeedsAttention,
	"FAILED":          TargetDatabaseLifecycleStateFailed,
}

// GetTargetDatabaseLifecycleStateEnumValues Enumerates the set of values for TargetDatabaseLifecycleStateEnum
func GetTargetDatabaseLifecycleStateEnumValues() []TargetDatabaseLifecycleStateEnum {
	values := make([]TargetDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingTargetDatabaseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDatabaseLifecycleStateEnumStringValues Enumerates the set of values in String for TargetDatabaseLifecycleStateEnum
func GetTargetDatabaseLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingTargetDatabaseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDatabaseLifecycleStateEnum(val string) (TargetDatabaseLifecycleStateEnum, bool) {
	mappingTargetDatabaseLifecycleStateEnumIgnoreCase := make(map[string]TargetDatabaseLifecycleStateEnum)
	for k, v := range mappingTargetDatabaseLifecycleStateEnum {
		mappingTargetDatabaseLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTargetDatabaseLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
