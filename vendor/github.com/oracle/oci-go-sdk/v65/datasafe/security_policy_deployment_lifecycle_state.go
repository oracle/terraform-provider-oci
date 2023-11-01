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

// SecurityPolicyDeploymentLifecycleStateEnum Enum with underlying type: string
type SecurityPolicyDeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityPolicyDeploymentLifecycleStateEnum
const (
	SecurityPolicyDeploymentLifecycleStateCreating       SecurityPolicyDeploymentLifecycleStateEnum = "CREATING"
	SecurityPolicyDeploymentLifecycleStateUpdating       SecurityPolicyDeploymentLifecycleStateEnum = "UPDATING"
	SecurityPolicyDeploymentLifecycleStateDeployed       SecurityPolicyDeploymentLifecycleStateEnum = "DEPLOYED"
	SecurityPolicyDeploymentLifecycleStateNeedsAttention SecurityPolicyDeploymentLifecycleStateEnum = "NEEDS_ATTENTION"
	SecurityPolicyDeploymentLifecycleStateFailed         SecurityPolicyDeploymentLifecycleStateEnum = "FAILED"
	SecurityPolicyDeploymentLifecycleStateDeleting       SecurityPolicyDeploymentLifecycleStateEnum = "DELETING"
	SecurityPolicyDeploymentLifecycleStateDeleted        SecurityPolicyDeploymentLifecycleStateEnum = "DELETED"
)

var mappingSecurityPolicyDeploymentLifecycleStateEnum = map[string]SecurityPolicyDeploymentLifecycleStateEnum{
	"CREATING":        SecurityPolicyDeploymentLifecycleStateCreating,
	"UPDATING":        SecurityPolicyDeploymentLifecycleStateUpdating,
	"DEPLOYED":        SecurityPolicyDeploymentLifecycleStateDeployed,
	"NEEDS_ATTENTION": SecurityPolicyDeploymentLifecycleStateNeedsAttention,
	"FAILED":          SecurityPolicyDeploymentLifecycleStateFailed,
	"DELETING":        SecurityPolicyDeploymentLifecycleStateDeleting,
	"DELETED":         SecurityPolicyDeploymentLifecycleStateDeleted,
}

var mappingSecurityPolicyDeploymentLifecycleStateEnumLowerCase = map[string]SecurityPolicyDeploymentLifecycleStateEnum{
	"creating":        SecurityPolicyDeploymentLifecycleStateCreating,
	"updating":        SecurityPolicyDeploymentLifecycleStateUpdating,
	"deployed":        SecurityPolicyDeploymentLifecycleStateDeployed,
	"needs_attention": SecurityPolicyDeploymentLifecycleStateNeedsAttention,
	"failed":          SecurityPolicyDeploymentLifecycleStateFailed,
	"deleting":        SecurityPolicyDeploymentLifecycleStateDeleting,
	"deleted":         SecurityPolicyDeploymentLifecycleStateDeleted,
}

// GetSecurityPolicyDeploymentLifecycleStateEnumValues Enumerates the set of values for SecurityPolicyDeploymentLifecycleStateEnum
func GetSecurityPolicyDeploymentLifecycleStateEnumValues() []SecurityPolicyDeploymentLifecycleStateEnum {
	values := make([]SecurityPolicyDeploymentLifecycleStateEnum, 0)
	for _, v := range mappingSecurityPolicyDeploymentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyDeploymentLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityPolicyDeploymentLifecycleStateEnum
func GetSecurityPolicyDeploymentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"DEPLOYED",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSecurityPolicyDeploymentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyDeploymentLifecycleStateEnum(val string) (SecurityPolicyDeploymentLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityPolicyDeploymentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
