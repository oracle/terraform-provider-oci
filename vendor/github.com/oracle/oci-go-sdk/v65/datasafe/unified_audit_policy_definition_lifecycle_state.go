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

// UnifiedAuditPolicyDefinitionLifecycleStateEnum Enum with underlying type: string
type UnifiedAuditPolicyDefinitionLifecycleStateEnum string

// Set of constants representing the allowable values for UnifiedAuditPolicyDefinitionLifecycleStateEnum
const (
	UnifiedAuditPolicyDefinitionLifecycleStateCreating       UnifiedAuditPolicyDefinitionLifecycleStateEnum = "CREATING"
	UnifiedAuditPolicyDefinitionLifecycleStateUpdating       UnifiedAuditPolicyDefinitionLifecycleStateEnum = "UPDATING"
	UnifiedAuditPolicyDefinitionLifecycleStateActive         UnifiedAuditPolicyDefinitionLifecycleStateEnum = "ACTIVE"
	UnifiedAuditPolicyDefinitionLifecycleStateInactive       UnifiedAuditPolicyDefinitionLifecycleStateEnum = "INACTIVE"
	UnifiedAuditPolicyDefinitionLifecycleStateFailed         UnifiedAuditPolicyDefinitionLifecycleStateEnum = "FAILED"
	UnifiedAuditPolicyDefinitionLifecycleStateDeleting       UnifiedAuditPolicyDefinitionLifecycleStateEnum = "DELETING"
	UnifiedAuditPolicyDefinitionLifecycleStateNeedsAttention UnifiedAuditPolicyDefinitionLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingUnifiedAuditPolicyDefinitionLifecycleStateEnum = map[string]UnifiedAuditPolicyDefinitionLifecycleStateEnum{
	"CREATING":        UnifiedAuditPolicyDefinitionLifecycleStateCreating,
	"UPDATING":        UnifiedAuditPolicyDefinitionLifecycleStateUpdating,
	"ACTIVE":          UnifiedAuditPolicyDefinitionLifecycleStateActive,
	"INACTIVE":        UnifiedAuditPolicyDefinitionLifecycleStateInactive,
	"FAILED":          UnifiedAuditPolicyDefinitionLifecycleStateFailed,
	"DELETING":        UnifiedAuditPolicyDefinitionLifecycleStateDeleting,
	"NEEDS_ATTENTION": UnifiedAuditPolicyDefinitionLifecycleStateNeedsAttention,
}

var mappingUnifiedAuditPolicyDefinitionLifecycleStateEnumLowerCase = map[string]UnifiedAuditPolicyDefinitionLifecycleStateEnum{
	"creating":        UnifiedAuditPolicyDefinitionLifecycleStateCreating,
	"updating":        UnifiedAuditPolicyDefinitionLifecycleStateUpdating,
	"active":          UnifiedAuditPolicyDefinitionLifecycleStateActive,
	"inactive":        UnifiedAuditPolicyDefinitionLifecycleStateInactive,
	"failed":          UnifiedAuditPolicyDefinitionLifecycleStateFailed,
	"deleting":        UnifiedAuditPolicyDefinitionLifecycleStateDeleting,
	"needs_attention": UnifiedAuditPolicyDefinitionLifecycleStateNeedsAttention,
}

// GetUnifiedAuditPolicyDefinitionLifecycleStateEnumValues Enumerates the set of values for UnifiedAuditPolicyDefinitionLifecycleStateEnum
func GetUnifiedAuditPolicyDefinitionLifecycleStateEnumValues() []UnifiedAuditPolicyDefinitionLifecycleStateEnum {
	values := make([]UnifiedAuditPolicyDefinitionLifecycleStateEnum, 0)
	for _, v := range mappingUnifiedAuditPolicyDefinitionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAuditPolicyDefinitionLifecycleStateEnumStringValues Enumerates the set of values in String for UnifiedAuditPolicyDefinitionLifecycleStateEnum
func GetUnifiedAuditPolicyDefinitionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingUnifiedAuditPolicyDefinitionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAuditPolicyDefinitionLifecycleStateEnum(val string) (UnifiedAuditPolicyDefinitionLifecycleStateEnum, bool) {
	enum, ok := mappingUnifiedAuditPolicyDefinitionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
