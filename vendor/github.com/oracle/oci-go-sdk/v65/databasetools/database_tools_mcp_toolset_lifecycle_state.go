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

// DatabaseToolsMcpToolsetLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsMcpToolsetLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpToolsetLifecycleStateEnum
const (
	DatabaseToolsMcpToolsetLifecycleStateCreating DatabaseToolsMcpToolsetLifecycleStateEnum = "CREATING"
	DatabaseToolsMcpToolsetLifecycleStateUpdating DatabaseToolsMcpToolsetLifecycleStateEnum = "UPDATING"
	DatabaseToolsMcpToolsetLifecycleStateActive   DatabaseToolsMcpToolsetLifecycleStateEnum = "ACTIVE"
	DatabaseToolsMcpToolsetLifecycleStateDeleting DatabaseToolsMcpToolsetLifecycleStateEnum = "DELETING"
	DatabaseToolsMcpToolsetLifecycleStateDeleted  DatabaseToolsMcpToolsetLifecycleStateEnum = "DELETED"
	DatabaseToolsMcpToolsetLifecycleStateFailed   DatabaseToolsMcpToolsetLifecycleStateEnum = "FAILED"
)

var mappingDatabaseToolsMcpToolsetLifecycleStateEnum = map[string]DatabaseToolsMcpToolsetLifecycleStateEnum{
	"CREATING": DatabaseToolsMcpToolsetLifecycleStateCreating,
	"UPDATING": DatabaseToolsMcpToolsetLifecycleStateUpdating,
	"ACTIVE":   DatabaseToolsMcpToolsetLifecycleStateActive,
	"DELETING": DatabaseToolsMcpToolsetLifecycleStateDeleting,
	"DELETED":  DatabaseToolsMcpToolsetLifecycleStateDeleted,
	"FAILED":   DatabaseToolsMcpToolsetLifecycleStateFailed,
}

var mappingDatabaseToolsMcpToolsetLifecycleStateEnumLowerCase = map[string]DatabaseToolsMcpToolsetLifecycleStateEnum{
	"creating": DatabaseToolsMcpToolsetLifecycleStateCreating,
	"updating": DatabaseToolsMcpToolsetLifecycleStateUpdating,
	"active":   DatabaseToolsMcpToolsetLifecycleStateActive,
	"deleting": DatabaseToolsMcpToolsetLifecycleStateDeleting,
	"deleted":  DatabaseToolsMcpToolsetLifecycleStateDeleted,
	"failed":   DatabaseToolsMcpToolsetLifecycleStateFailed,
}

// GetDatabaseToolsMcpToolsetLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsMcpToolsetLifecycleStateEnum
func GetDatabaseToolsMcpToolsetLifecycleStateEnumValues() []DatabaseToolsMcpToolsetLifecycleStateEnum {
	values := make([]DatabaseToolsMcpToolsetLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsMcpToolsetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpToolsetLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpToolsetLifecycleStateEnum
func GetDatabaseToolsMcpToolsetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDatabaseToolsMcpToolsetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpToolsetLifecycleStateEnum(val string) (DatabaseToolsMcpToolsetLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpToolsetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
