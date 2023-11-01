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

// SecurityPolicyLifecycleStateEnum Enum with underlying type: string
type SecurityPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityPolicyLifecycleStateEnum
const (
	SecurityPolicyLifecycleStateCreating SecurityPolicyLifecycleStateEnum = "CREATING"
	SecurityPolicyLifecycleStateUpdating SecurityPolicyLifecycleStateEnum = "UPDATING"
	SecurityPolicyLifecycleStateActive   SecurityPolicyLifecycleStateEnum = "ACTIVE"
	SecurityPolicyLifecycleStateFailed   SecurityPolicyLifecycleStateEnum = "FAILED"
	SecurityPolicyLifecycleStateDeleting SecurityPolicyLifecycleStateEnum = "DELETING"
	SecurityPolicyLifecycleStateDeleted  SecurityPolicyLifecycleStateEnum = "DELETED"
)

var mappingSecurityPolicyLifecycleStateEnum = map[string]SecurityPolicyLifecycleStateEnum{
	"CREATING": SecurityPolicyLifecycleStateCreating,
	"UPDATING": SecurityPolicyLifecycleStateUpdating,
	"ACTIVE":   SecurityPolicyLifecycleStateActive,
	"FAILED":   SecurityPolicyLifecycleStateFailed,
	"DELETING": SecurityPolicyLifecycleStateDeleting,
	"DELETED":  SecurityPolicyLifecycleStateDeleted,
}

var mappingSecurityPolicyLifecycleStateEnumLowerCase = map[string]SecurityPolicyLifecycleStateEnum{
	"creating": SecurityPolicyLifecycleStateCreating,
	"updating": SecurityPolicyLifecycleStateUpdating,
	"active":   SecurityPolicyLifecycleStateActive,
	"failed":   SecurityPolicyLifecycleStateFailed,
	"deleting": SecurityPolicyLifecycleStateDeleting,
	"deleted":  SecurityPolicyLifecycleStateDeleted,
}

// GetSecurityPolicyLifecycleStateEnumValues Enumerates the set of values for SecurityPolicyLifecycleStateEnum
func GetSecurityPolicyLifecycleStateEnumValues() []SecurityPolicyLifecycleStateEnum {
	values := make([]SecurityPolicyLifecycleStateEnum, 0)
	for _, v := range mappingSecurityPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityPolicyLifecycleStateEnum
func GetSecurityPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSecurityPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyLifecycleStateEnum(val string) (SecurityPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
