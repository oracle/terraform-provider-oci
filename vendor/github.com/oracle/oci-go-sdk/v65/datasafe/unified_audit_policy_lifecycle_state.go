// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UnifiedAuditPolicyLifecycleStateEnum Enum with underlying type: string
type UnifiedAuditPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyLifecycleStateEnum
const (
	UnifiedAuditPolicyLifecycleStateCreating       UnifiedAuditPolicyLifecycleStateEnum = "CREATING"
	UnifiedAuditPolicyLifecycleStateUpdating       UnifiedAuditPolicyLifecycleStateEnum = "UPDATING"
	UnifiedAuditPolicyLifecycleStateActive         UnifiedAuditPolicyLifecycleStateEnum = "ACTIVE"
	UnifiedAuditPolicyLifecycleStateInactive       UnifiedAuditPolicyLifecycleStateEnum = "INACTIVE"
	UnifiedAuditPolicyLifecycleStateFailed         UnifiedAuditPolicyLifecycleStateEnum = "FAILED"
	UnifiedAuditPolicyLifecycleStateDeleting       UnifiedAuditPolicyLifecycleStateEnum = "DELETING"
	UnifiedAuditPolicyLifecycleStateNeedsAttention UnifiedAuditPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	UnifiedAuditPolicyLifecycleStateDeleted        UnifiedAuditPolicyLifecycleStateEnum = "DELETED"
)

var mappingUnifiedAuditPolicyLifecycleStateEnum = map[string]UnifiedAuditPolicyLifecycleStateEnum{
	"CREATING":        UnifiedAuditPolicyLifecycleStateCreating,
	"UPDATING":        UnifiedAuditPolicyLifecycleStateUpdating,
	"ACTIVE":          UnifiedAuditPolicyLifecycleStateActive,
	"INACTIVE":        UnifiedAuditPolicyLifecycleStateInactive,
	"FAILED":          UnifiedAuditPolicyLifecycleStateFailed,
	"DELETING":        UnifiedAuditPolicyLifecycleStateDeleting,
	"NEEDS_ATTENTION": UnifiedAuditPolicyLifecycleStateNeedsAttention,
	"DELETED":         UnifiedAuditPolicyLifecycleStateDeleted,
}

var mappingUnifiedAuditPolicyLifecycleStateEnumLowerCase = map[string]UnifiedAuditPolicyLifecycleStateEnum{
	"creating":        UnifiedAuditPolicyLifecycleStateCreating,
	"updating":        UnifiedAuditPolicyLifecycleStateUpdating,
	"active":          UnifiedAuditPolicyLifecycleStateActive,
	"inactive":        UnifiedAuditPolicyLifecycleStateInactive,
	"failed":          UnifiedAuditPolicyLifecycleStateFailed,
	"deleting":        UnifiedAuditPolicyLifecycleStateDeleting,
	"needs_attention": UnifiedAuditPolicyLifecycleStateNeedsAttention,
	"deleted":         UnifiedAuditPolicyLifecycleStateDeleted,
}

// GetUnifiedAuditPolicyLifecycleStateEnumValues Enumerates the set of values for UnifiedAuditPolicyLifecycleStateEnum
func GetUnifiedAuditPolicyLifecycleStateEnumValues() []UnifiedAuditPolicyLifecycleStateEnum {
	values := make([]UnifiedAuditPolicyLifecycleStateEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyLifecycleStateEnum
func GetUnifiedAuditPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingUnifiedAuditPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyLifecycleStateEnum(val string) (UnifiedAuditPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
