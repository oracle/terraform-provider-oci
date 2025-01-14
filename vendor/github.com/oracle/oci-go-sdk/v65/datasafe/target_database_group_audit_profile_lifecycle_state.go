// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TargetDatabaseGroupAuditProfileLifecycleStateEnum Enum with underlying type: string
type TargetDatabaseGroupAuditProfileLifecycleStateEnum string

// Set of constants representing the allowable values for TargetDatabaseGroupAuditProfileLifecycleStateEnum
const (
	TargetDatabaseGroupAuditProfileLifecycleStateCreating       TargetDatabaseGroupAuditProfileLifecycleStateEnum = "CREATING"
	TargetDatabaseGroupAuditProfileLifecycleStateUpdating       TargetDatabaseGroupAuditProfileLifecycleStateEnum = "UPDATING"
	TargetDatabaseGroupAuditProfileLifecycleStateActive         TargetDatabaseGroupAuditProfileLifecycleStateEnum = "ACTIVE"
	TargetDatabaseGroupAuditProfileLifecycleStateDeleting       TargetDatabaseGroupAuditProfileLifecycleStateEnum = "DELETING"
	TargetDatabaseGroupAuditProfileLifecycleStateFailed         TargetDatabaseGroupAuditProfileLifecycleStateEnum = "FAILED"
	TargetDatabaseGroupAuditProfileLifecycleStateNeedsAttention TargetDatabaseGroupAuditProfileLifecycleStateEnum = "NEEDS_ATTENTION"
	TargetDatabaseGroupAuditProfileLifecycleStateDeleted        TargetDatabaseGroupAuditProfileLifecycleStateEnum = "DELETED"
)

var mappingTargetDatabaseGroupAuditProfileLifecycleStateEnum = map[string]TargetDatabaseGroupAuditProfileLifecycleStateEnum{
	"CREATING":        TargetDatabaseGroupAuditProfileLifecycleStateCreating,
	"UPDATING":        TargetDatabaseGroupAuditProfileLifecycleStateUpdating,
	"ACTIVE":          TargetDatabaseGroupAuditProfileLifecycleStateActive,
	"DELETING":        TargetDatabaseGroupAuditProfileLifecycleStateDeleting,
	"FAILED":          TargetDatabaseGroupAuditProfileLifecycleStateFailed,
	"NEEDS_ATTENTION": TargetDatabaseGroupAuditProfileLifecycleStateNeedsAttention,
	"DELETED":         TargetDatabaseGroupAuditProfileLifecycleStateDeleted,
}

var mappingTargetDatabaseGroupAuditProfileLifecycleStateEnumLowerCase = map[string]TargetDatabaseGroupAuditProfileLifecycleStateEnum{
	"creating":        TargetDatabaseGroupAuditProfileLifecycleStateCreating,
	"updating":        TargetDatabaseGroupAuditProfileLifecycleStateUpdating,
	"active":          TargetDatabaseGroupAuditProfileLifecycleStateActive,
	"deleting":        TargetDatabaseGroupAuditProfileLifecycleStateDeleting,
	"failed":          TargetDatabaseGroupAuditProfileLifecycleStateFailed,
	"needs_attention": TargetDatabaseGroupAuditProfileLifecycleStateNeedsAttention,
	"deleted":         TargetDatabaseGroupAuditProfileLifecycleStateDeleted,
}

// GetTargetDatabaseGroupAuditProfileLifecycleStateEnumValues Enumerates the set of values for TargetDatabaseGroupAuditProfileLifecycleStateEnum
func GetTargetDatabaseGroupAuditProfileLifecycleStateEnumValues() []TargetDatabaseGroupAuditProfileLifecycleStateEnum {
	values := make([]TargetDatabaseGroupAuditProfileLifecycleStateEnum, 0)
	for _, v := range mappingTargetDatabaseGroupAuditProfileLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDatabaseGroupAuditProfileLifecycleStateEnumStringValues Enumerates the set of values in String for TargetDatabaseGroupAuditProfileLifecycleStateEnum
func GetTargetDatabaseGroupAuditProfileLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingTargetDatabaseGroupAuditProfileLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDatabaseGroupAuditProfileLifecycleStateEnum(val string) (TargetDatabaseGroupAuditProfileLifecycleStateEnum, bool) {
	enum, ok := mappingTargetDatabaseGroupAuditProfileLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
