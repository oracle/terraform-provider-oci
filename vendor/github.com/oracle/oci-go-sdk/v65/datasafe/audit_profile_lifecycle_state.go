// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AuditProfileLifecycleStateEnum Enum with underlying type: string
type AuditProfileLifecycleStateEnum string

// Set of constants representing the allowable values for AuditProfileLifecycleStateEnum
const (
	AuditProfileLifecycleStateCreating       AuditProfileLifecycleStateEnum = "CREATING"
	AuditProfileLifecycleStateUpdating       AuditProfileLifecycleStateEnum = "UPDATING"
	AuditProfileLifecycleStateActive         AuditProfileLifecycleStateEnum = "ACTIVE"
	AuditProfileLifecycleStateDeleting       AuditProfileLifecycleStateEnum = "DELETING"
	AuditProfileLifecycleStateFailed         AuditProfileLifecycleStateEnum = "FAILED"
	AuditProfileLifecycleStateNeedsAttention AuditProfileLifecycleStateEnum = "NEEDS_ATTENTION"
	AuditProfileLifecycleStateDeleted        AuditProfileLifecycleStateEnum = "DELETED"
)

var mappingAuditProfileLifecycleStateEnum = map[string]AuditProfileLifecycleStateEnum{
	"CREATING":        AuditProfileLifecycleStateCreating,
	"UPDATING":        AuditProfileLifecycleStateUpdating,
	"ACTIVE":          AuditProfileLifecycleStateActive,
	"DELETING":        AuditProfileLifecycleStateDeleting,
	"FAILED":          AuditProfileLifecycleStateFailed,
	"NEEDS_ATTENTION": AuditProfileLifecycleStateNeedsAttention,
	"DELETED":         AuditProfileLifecycleStateDeleted,
}

var mappingAuditProfileLifecycleStateEnumLowerCase = map[string]AuditProfileLifecycleStateEnum{
	"creating":        AuditProfileLifecycleStateCreating,
	"updating":        AuditProfileLifecycleStateUpdating,
	"active":          AuditProfileLifecycleStateActive,
	"deleting":        AuditProfileLifecycleStateDeleting,
	"failed":          AuditProfileLifecycleStateFailed,
	"needs_attention": AuditProfileLifecycleStateNeedsAttention,
	"deleted":         AuditProfileLifecycleStateDeleted,
}

// GetAuditProfileLifecycleStateEnumValues Enumerates the set of values for AuditProfileLifecycleStateEnum
func GetAuditProfileLifecycleStateEnumValues() []AuditProfileLifecycleStateEnum {
	values := make([]AuditProfileLifecycleStateEnum, 0)
	for _, v := range mappingAuditProfileLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditProfileLifecycleStateEnumStringValues Enumerates the set of values in String for AuditProfileLifecycleStateEnum
func GetAuditProfileLifecycleStateEnumStringValues() []string {
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

// GetMappingAuditProfileLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditProfileLifecycleStateEnum(val string) (AuditProfileLifecycleStateEnum, bool) {
	enum, ok := mappingAuditProfileLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
