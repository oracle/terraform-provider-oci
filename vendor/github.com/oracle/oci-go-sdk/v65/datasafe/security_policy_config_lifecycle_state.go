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

// SecurityPolicyConfigLifecycleStateEnum Enum with underlying type: string
type SecurityPolicyConfigLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityPolicyConfigLifecycleStateEnum
const (
	SecurityPolicyConfigLifecycleStateCreating       SecurityPolicyConfigLifecycleStateEnum = "CREATING"
	SecurityPolicyConfigLifecycleStateUpdating       SecurityPolicyConfigLifecycleStateEnum = "UPDATING"
	SecurityPolicyConfigLifecycleStateActive         SecurityPolicyConfigLifecycleStateEnum = "ACTIVE"
	SecurityPolicyConfigLifecycleStateFailed         SecurityPolicyConfigLifecycleStateEnum = "FAILED"
	SecurityPolicyConfigLifecycleStateNeedsAttention SecurityPolicyConfigLifecycleStateEnum = "NEEDS_ATTENTION"
	SecurityPolicyConfigLifecycleStateDeleting       SecurityPolicyConfigLifecycleStateEnum = "DELETING"
	SecurityPolicyConfigLifecycleStateDeleted        SecurityPolicyConfigLifecycleStateEnum = "DELETED"
)

var mappingSecurityPolicyConfigLifecycleStateEnum = map[string]SecurityPolicyConfigLifecycleStateEnum{
	"CREATING":        SecurityPolicyConfigLifecycleStateCreating,
	"UPDATING":        SecurityPolicyConfigLifecycleStateUpdating,
	"ACTIVE":          SecurityPolicyConfigLifecycleStateActive,
	"FAILED":          SecurityPolicyConfigLifecycleStateFailed,
	"NEEDS_ATTENTION": SecurityPolicyConfigLifecycleStateNeedsAttention,
	"DELETING":        SecurityPolicyConfigLifecycleStateDeleting,
	"DELETED":         SecurityPolicyConfigLifecycleStateDeleted,
}

var mappingSecurityPolicyConfigLifecycleStateEnumLowerCase = map[string]SecurityPolicyConfigLifecycleStateEnum{
	"creating":        SecurityPolicyConfigLifecycleStateCreating,
	"updating":        SecurityPolicyConfigLifecycleStateUpdating,
	"active":          SecurityPolicyConfigLifecycleStateActive,
	"failed":          SecurityPolicyConfigLifecycleStateFailed,
	"needs_attention": SecurityPolicyConfigLifecycleStateNeedsAttention,
	"deleting":        SecurityPolicyConfigLifecycleStateDeleting,
	"deleted":         SecurityPolicyConfigLifecycleStateDeleted,
}

// GetSecurityPolicyConfigLifecycleStateEnumValues Enumerates the set of values for SecurityPolicyConfigLifecycleStateEnum
func GetSecurityPolicyConfigLifecycleStateEnumValues() []SecurityPolicyConfigLifecycleStateEnum {
	values := make([]SecurityPolicyConfigLifecycleStateEnum, 0)
	for _, v := range mappingSecurityPolicyConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyConfigLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityPolicyConfigLifecycleStateEnum
func GetSecurityPolicyConfigLifecycleStateEnumStringValues() []string {
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

// GetMappingSecurityPolicyConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyConfigLifecycleStateEnum(val string) (SecurityPolicyConfigLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityPolicyConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
