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

// SqlCollectionLifecycleStateEnum Enum with underlying type: string
type SqlCollectionLifecycleStateEnum string

// Set of constants representing the allowable values for SqlCollectionLifecycleStateEnum
const (
	SqlCollectionLifecycleStateCreating       SqlCollectionLifecycleStateEnum = "CREATING"
	SqlCollectionLifecycleStateUpdating       SqlCollectionLifecycleStateEnum = "UPDATING"
	SqlCollectionLifecycleStateCollecting     SqlCollectionLifecycleStateEnum = "COLLECTING"
	SqlCollectionLifecycleStateCompleted      SqlCollectionLifecycleStateEnum = "COMPLETED"
	SqlCollectionLifecycleStateInactive       SqlCollectionLifecycleStateEnum = "INACTIVE"
	SqlCollectionLifecycleStateFailed         SqlCollectionLifecycleStateEnum = "FAILED"
	SqlCollectionLifecycleStateDeleting       SqlCollectionLifecycleStateEnum = "DELETING"
	SqlCollectionLifecycleStateDeleted        SqlCollectionLifecycleStateEnum = "DELETED"
	SqlCollectionLifecycleStateNeedsAttention SqlCollectionLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingSqlCollectionLifecycleStateEnum = map[string]SqlCollectionLifecycleStateEnum{
	"CREATING":        SqlCollectionLifecycleStateCreating,
	"UPDATING":        SqlCollectionLifecycleStateUpdating,
	"COLLECTING":      SqlCollectionLifecycleStateCollecting,
	"COMPLETED":       SqlCollectionLifecycleStateCompleted,
	"INACTIVE":        SqlCollectionLifecycleStateInactive,
	"FAILED":          SqlCollectionLifecycleStateFailed,
	"DELETING":        SqlCollectionLifecycleStateDeleting,
	"DELETED":         SqlCollectionLifecycleStateDeleted,
	"NEEDS_ATTENTION": SqlCollectionLifecycleStateNeedsAttention,
}

var mappingSqlCollectionLifecycleStateEnumLowerCase = map[string]SqlCollectionLifecycleStateEnum{
	"creating":        SqlCollectionLifecycleStateCreating,
	"updating":        SqlCollectionLifecycleStateUpdating,
	"collecting":      SqlCollectionLifecycleStateCollecting,
	"completed":       SqlCollectionLifecycleStateCompleted,
	"inactive":        SqlCollectionLifecycleStateInactive,
	"failed":          SqlCollectionLifecycleStateFailed,
	"deleting":        SqlCollectionLifecycleStateDeleting,
	"deleted":         SqlCollectionLifecycleStateDeleted,
	"needs_attention": SqlCollectionLifecycleStateNeedsAttention,
}

// GetSqlCollectionLifecycleStateEnumValues Enumerates the set of values for SqlCollectionLifecycleStateEnum
func GetSqlCollectionLifecycleStateEnumValues() []SqlCollectionLifecycleStateEnum {
	values := make([]SqlCollectionLifecycleStateEnum, 0)
	for _, v := range mappingSqlCollectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlCollectionLifecycleStateEnumStringValues Enumerates the set of values in String for SqlCollectionLifecycleStateEnum
func GetSqlCollectionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"COLLECTING",
		"COMPLETED",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingSqlCollectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlCollectionLifecycleStateEnum(val string) (SqlCollectionLifecycleStateEnum, bool) {
	enum, ok := mappingSqlCollectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
