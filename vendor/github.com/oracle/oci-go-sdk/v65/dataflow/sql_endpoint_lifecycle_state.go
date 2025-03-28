// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// SqlEndpointLifecycleStateEnum Enum with underlying type: string
type SqlEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for SqlEndpointLifecycleStateEnum
const (
	SqlEndpointLifecycleStateCreating       SqlEndpointLifecycleStateEnum = "CREATING"
	SqlEndpointLifecycleStateActive         SqlEndpointLifecycleStateEnum = "ACTIVE"
	SqlEndpointLifecycleStateDeleting       SqlEndpointLifecycleStateEnum = "DELETING"
	SqlEndpointLifecycleStateDeleted        SqlEndpointLifecycleStateEnum = "DELETED"
	SqlEndpointLifecycleStateFailed         SqlEndpointLifecycleStateEnum = "FAILED"
	SqlEndpointLifecycleStateUpdating       SqlEndpointLifecycleStateEnum = "UPDATING"
	SqlEndpointLifecycleStateNeedsAttention SqlEndpointLifecycleStateEnum = "NEEDS_ATTENTION"
	SqlEndpointLifecycleStateInactive       SqlEndpointLifecycleStateEnum = "INACTIVE"
)

var mappingSqlEndpointLifecycleStateEnum = map[string]SqlEndpointLifecycleStateEnum{
	"CREATING":        SqlEndpointLifecycleStateCreating,
	"ACTIVE":          SqlEndpointLifecycleStateActive,
	"DELETING":        SqlEndpointLifecycleStateDeleting,
	"DELETED":         SqlEndpointLifecycleStateDeleted,
	"FAILED":          SqlEndpointLifecycleStateFailed,
	"UPDATING":        SqlEndpointLifecycleStateUpdating,
	"NEEDS_ATTENTION": SqlEndpointLifecycleStateNeedsAttention,
	"INACTIVE":        SqlEndpointLifecycleStateInactive,
}

var mappingSqlEndpointLifecycleStateEnumLowerCase = map[string]SqlEndpointLifecycleStateEnum{
	"creating":        SqlEndpointLifecycleStateCreating,
	"active":          SqlEndpointLifecycleStateActive,
	"deleting":        SqlEndpointLifecycleStateDeleting,
	"deleted":         SqlEndpointLifecycleStateDeleted,
	"failed":          SqlEndpointLifecycleStateFailed,
	"updating":        SqlEndpointLifecycleStateUpdating,
	"needs_attention": SqlEndpointLifecycleStateNeedsAttention,
	"inactive":        SqlEndpointLifecycleStateInactive,
}

// GetSqlEndpointLifecycleStateEnumValues Enumerates the set of values for SqlEndpointLifecycleStateEnum
func GetSqlEndpointLifecycleStateEnumValues() []SqlEndpointLifecycleStateEnum {
	values := make([]SqlEndpointLifecycleStateEnum, 0)
	for _, v := range mappingSqlEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for SqlEndpointLifecycleStateEnum
func GetSqlEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
		"NEEDS_ATTENTION",
		"INACTIVE",
	}
}

// GetMappingSqlEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlEndpointLifecycleStateEnum(val string) (SqlEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingSqlEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
