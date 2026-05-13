// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsMcpServerLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsMcpServerLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpServerLifecycleStateEnum
const (
	DatabaseToolsMcpServerLifecycleStateCreating       DatabaseToolsMcpServerLifecycleStateEnum = "CREATING"
	DatabaseToolsMcpServerLifecycleStateUpdating       DatabaseToolsMcpServerLifecycleStateEnum = "UPDATING"
	DatabaseToolsMcpServerLifecycleStateActive         DatabaseToolsMcpServerLifecycleStateEnum = "ACTIVE"
	DatabaseToolsMcpServerLifecycleStateDeleting       DatabaseToolsMcpServerLifecycleStateEnum = "DELETING"
	DatabaseToolsMcpServerLifecycleStateDeleted        DatabaseToolsMcpServerLifecycleStateEnum = "DELETED"
	DatabaseToolsMcpServerLifecycleStateFailed         DatabaseToolsMcpServerLifecycleStateEnum = "FAILED"
	DatabaseToolsMcpServerLifecycleStateNeedsAttention DatabaseToolsMcpServerLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDatabaseToolsMcpServerLifecycleStateEnum = map[string]DatabaseToolsMcpServerLifecycleStateEnum{
	"CREATING":        DatabaseToolsMcpServerLifecycleStateCreating,
	"UPDATING":        DatabaseToolsMcpServerLifecycleStateUpdating,
	"ACTIVE":          DatabaseToolsMcpServerLifecycleStateActive,
	"DELETING":        DatabaseToolsMcpServerLifecycleStateDeleting,
	"DELETED":         DatabaseToolsMcpServerLifecycleStateDeleted,
	"FAILED":          DatabaseToolsMcpServerLifecycleStateFailed,
	"NEEDS_ATTENTION": DatabaseToolsMcpServerLifecycleStateNeedsAttention,
}

var mappingDatabaseToolsMcpServerLifecycleStateEnumLowerCase = map[string]DatabaseToolsMcpServerLifecycleStateEnum{
	"creating":        DatabaseToolsMcpServerLifecycleStateCreating,
	"updating":        DatabaseToolsMcpServerLifecycleStateUpdating,
	"active":          DatabaseToolsMcpServerLifecycleStateActive,
	"deleting":        DatabaseToolsMcpServerLifecycleStateDeleting,
	"deleted":         DatabaseToolsMcpServerLifecycleStateDeleted,
	"failed":          DatabaseToolsMcpServerLifecycleStateFailed,
	"needs_attention": DatabaseToolsMcpServerLifecycleStateNeedsAttention,
}

// GetDatabaseToolsMcpServerLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsMcpServerLifecycleStateEnum
func GetDatabaseToolsMcpServerLifecycleStateEnumValues() []DatabaseToolsMcpServerLifecycleStateEnum {
	values := make([]DatabaseToolsMcpServerLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsMcpServerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpServerLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpServerLifecycleStateEnum
func GetDatabaseToolsMcpServerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDatabaseToolsMcpServerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpServerLifecycleStateEnum(val string) (DatabaseToolsMcpServerLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpServerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
