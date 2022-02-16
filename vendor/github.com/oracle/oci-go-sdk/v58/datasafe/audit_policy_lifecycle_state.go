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

// AuditPolicyLifecycleStateEnum Enum with underlying type: string
type AuditPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for AuditPolicyLifecycleStateEnum
const (
	AuditPolicyLifecycleStateCreating       AuditPolicyLifecycleStateEnum = "CREATING"
	AuditPolicyLifecycleStateUpdating       AuditPolicyLifecycleStateEnum = "UPDATING"
	AuditPolicyLifecycleStateActive         AuditPolicyLifecycleStateEnum = "ACTIVE"
	AuditPolicyLifecycleStateFailed         AuditPolicyLifecycleStateEnum = "FAILED"
	AuditPolicyLifecycleStateNeedsAttention AuditPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	AuditPolicyLifecycleStateDeleting       AuditPolicyLifecycleStateEnum = "DELETING"
	AuditPolicyLifecycleStateDeleted        AuditPolicyLifecycleStateEnum = "DELETED"
)

var mappingAuditPolicyLifecycleStateEnum = map[string]AuditPolicyLifecycleStateEnum{
	"CREATING":        AuditPolicyLifecycleStateCreating,
	"UPDATING":        AuditPolicyLifecycleStateUpdating,
	"ACTIVE":          AuditPolicyLifecycleStateActive,
	"FAILED":          AuditPolicyLifecycleStateFailed,
	"NEEDS_ATTENTION": AuditPolicyLifecycleStateNeedsAttention,
	"DELETING":        AuditPolicyLifecycleStateDeleting,
	"DELETED":         AuditPolicyLifecycleStateDeleted,
}

// GetAuditPolicyLifecycleStateEnumValues Enumerates the set of values for AuditPolicyLifecycleStateEnum
func GetAuditPolicyLifecycleStateEnumValues() []AuditPolicyLifecycleStateEnum {
	values := make([]AuditPolicyLifecycleStateEnum, 0)
	for _, v := range mappingAuditPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for AuditPolicyLifecycleStateEnum
func GetAuditPolicyLifecycleStateEnumStringValues() []string {
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

// GetMappingAuditPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditPolicyLifecycleStateEnum(val string) (AuditPolicyLifecycleStateEnum, bool) {
	mappingAuditPolicyLifecycleStateEnumIgnoreCase := make(map[string]AuditPolicyLifecycleStateEnum)
	for k, v := range mappingAuditPolicyLifecycleStateEnum {
		mappingAuditPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAuditPolicyLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
