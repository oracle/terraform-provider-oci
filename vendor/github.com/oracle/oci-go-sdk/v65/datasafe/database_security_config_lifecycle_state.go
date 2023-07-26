// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseSecurityConfigLifecycleStateEnum Enum with underlying type: string
type DatabaseSecurityConfigLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseSecurityConfigLifecycleStateEnum
const (
	DatabaseSecurityConfigLifecycleStateCreating       DatabaseSecurityConfigLifecycleStateEnum = "CREATING"
	DatabaseSecurityConfigLifecycleStateUpdating       DatabaseSecurityConfigLifecycleStateEnum = "UPDATING"
	DatabaseSecurityConfigLifecycleStateActive         DatabaseSecurityConfigLifecycleStateEnum = "ACTIVE"
	DatabaseSecurityConfigLifecycleStateFailed         DatabaseSecurityConfigLifecycleStateEnum = "FAILED"
	DatabaseSecurityConfigLifecycleStateNeedsAttention DatabaseSecurityConfigLifecycleStateEnum = "NEEDS_ATTENTION"
	DatabaseSecurityConfigLifecycleStateDeleting       DatabaseSecurityConfigLifecycleStateEnum = "DELETING"
	DatabaseSecurityConfigLifecycleStateDeleted        DatabaseSecurityConfigLifecycleStateEnum = "DELETED"
)

var mappingDatabaseSecurityConfigLifecycleStateEnum = map[string]DatabaseSecurityConfigLifecycleStateEnum{
	"CREATING":        DatabaseSecurityConfigLifecycleStateCreating,
	"UPDATING":        DatabaseSecurityConfigLifecycleStateUpdating,
	"ACTIVE":          DatabaseSecurityConfigLifecycleStateActive,
	"FAILED":          DatabaseSecurityConfigLifecycleStateFailed,
	"NEEDS_ATTENTION": DatabaseSecurityConfigLifecycleStateNeedsAttention,
	"DELETING":        DatabaseSecurityConfigLifecycleStateDeleting,
	"DELETED":         DatabaseSecurityConfigLifecycleStateDeleted,
}

var mappingDatabaseSecurityConfigLifecycleStateEnumLowerCase = map[string]DatabaseSecurityConfigLifecycleStateEnum{
	"creating":        DatabaseSecurityConfigLifecycleStateCreating,
	"updating":        DatabaseSecurityConfigLifecycleStateUpdating,
	"active":          DatabaseSecurityConfigLifecycleStateActive,
	"failed":          DatabaseSecurityConfigLifecycleStateFailed,
	"needs_attention": DatabaseSecurityConfigLifecycleStateNeedsAttention,
	"deleting":        DatabaseSecurityConfigLifecycleStateDeleting,
	"deleted":         DatabaseSecurityConfigLifecycleStateDeleted,
}

// GetDatabaseSecurityConfigLifecycleStateEnumValues Enumerates the set of values for DatabaseSecurityConfigLifecycleStateEnum
func GetDatabaseSecurityConfigLifecycleStateEnumValues() []DatabaseSecurityConfigLifecycleStateEnum {
	values := make([]DatabaseSecurityConfigLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseSecurityConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSecurityConfigLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseSecurityConfigLifecycleStateEnum
func GetDatabaseSecurityConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingDatabaseSecurityConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSecurityConfigLifecycleStateEnum(val string) (DatabaseSecurityConfigLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseSecurityConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
