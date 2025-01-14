// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsIdentityLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsIdentityLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsIdentityLifecycleStateEnum
const (
	DatabaseToolsIdentityLifecycleStateCreating       DatabaseToolsIdentityLifecycleStateEnum = "CREATING"
	DatabaseToolsIdentityLifecycleStateUpdating       DatabaseToolsIdentityLifecycleStateEnum = "UPDATING"
	DatabaseToolsIdentityLifecycleStateActive         DatabaseToolsIdentityLifecycleStateEnum = "ACTIVE"
	DatabaseToolsIdentityLifecycleStateDeleting       DatabaseToolsIdentityLifecycleStateEnum = "DELETING"
	DatabaseToolsIdentityLifecycleStateDeleted        DatabaseToolsIdentityLifecycleStateEnum = "DELETED"
	DatabaseToolsIdentityLifecycleStateFailed         DatabaseToolsIdentityLifecycleStateEnum = "FAILED"
	DatabaseToolsIdentityLifecycleStateNeedsAttention DatabaseToolsIdentityLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDatabaseToolsIdentityLifecycleStateEnum = map[string]DatabaseToolsIdentityLifecycleStateEnum{
	"CREATING":        DatabaseToolsIdentityLifecycleStateCreating,
	"UPDATING":        DatabaseToolsIdentityLifecycleStateUpdating,
	"ACTIVE":          DatabaseToolsIdentityLifecycleStateActive,
	"DELETING":        DatabaseToolsIdentityLifecycleStateDeleting,
	"DELETED":         DatabaseToolsIdentityLifecycleStateDeleted,
	"FAILED":          DatabaseToolsIdentityLifecycleStateFailed,
	"NEEDS_ATTENTION": DatabaseToolsIdentityLifecycleStateNeedsAttention,
}

var mappingDatabaseToolsIdentityLifecycleStateEnumLowerCase = map[string]DatabaseToolsIdentityLifecycleStateEnum{
	"creating":        DatabaseToolsIdentityLifecycleStateCreating,
	"updating":        DatabaseToolsIdentityLifecycleStateUpdating,
	"active":          DatabaseToolsIdentityLifecycleStateActive,
	"deleting":        DatabaseToolsIdentityLifecycleStateDeleting,
	"deleted":         DatabaseToolsIdentityLifecycleStateDeleted,
	"failed":          DatabaseToolsIdentityLifecycleStateFailed,
	"needs_attention": DatabaseToolsIdentityLifecycleStateNeedsAttention,
}

// GetDatabaseToolsIdentityLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsIdentityLifecycleStateEnum
func GetDatabaseToolsIdentityLifecycleStateEnumValues() []DatabaseToolsIdentityLifecycleStateEnum {
	values := make([]DatabaseToolsIdentityLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsIdentityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsIdentityLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsIdentityLifecycleStateEnum
func GetDatabaseToolsIdentityLifecycleStateEnumStringValues() []string {
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

// GetMappingDatabaseToolsIdentityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsIdentityLifecycleStateEnum(val string) (DatabaseToolsIdentityLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsIdentityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
