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

// SecurityPolicyReportLifecycleStateEnum Enum with underlying type: string
type SecurityPolicyReportLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityPolicyReportLifecycleStateEnum
const (
	SecurityPolicyReportLifecycleStateCreating       SecurityPolicyReportLifecycleStateEnum = "CREATING"
	SecurityPolicyReportLifecycleStateSucceeded      SecurityPolicyReportLifecycleStateEnum = "SUCCEEDED"
	SecurityPolicyReportLifecycleStateUpdating       SecurityPolicyReportLifecycleStateEnum = "UPDATING"
	SecurityPolicyReportLifecycleStateDeleting       SecurityPolicyReportLifecycleStateEnum = "DELETING"
	SecurityPolicyReportLifecycleStateDeleted        SecurityPolicyReportLifecycleStateEnum = "DELETED"
	SecurityPolicyReportLifecycleStateFailed         SecurityPolicyReportLifecycleStateEnum = "FAILED"
	SecurityPolicyReportLifecycleStateNeedsAttention SecurityPolicyReportLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingSecurityPolicyReportLifecycleStateEnum = map[string]SecurityPolicyReportLifecycleStateEnum{
	"CREATING":        SecurityPolicyReportLifecycleStateCreating,
	"SUCCEEDED":       SecurityPolicyReportLifecycleStateSucceeded,
	"UPDATING":        SecurityPolicyReportLifecycleStateUpdating,
	"DELETING":        SecurityPolicyReportLifecycleStateDeleting,
	"DELETED":         SecurityPolicyReportLifecycleStateDeleted,
	"FAILED":          SecurityPolicyReportLifecycleStateFailed,
	"NEEDS_ATTENTION": SecurityPolicyReportLifecycleStateNeedsAttention,
}

var mappingSecurityPolicyReportLifecycleStateEnumLowerCase = map[string]SecurityPolicyReportLifecycleStateEnum{
	"creating":        SecurityPolicyReportLifecycleStateCreating,
	"succeeded":       SecurityPolicyReportLifecycleStateSucceeded,
	"updating":        SecurityPolicyReportLifecycleStateUpdating,
	"deleting":        SecurityPolicyReportLifecycleStateDeleting,
	"deleted":         SecurityPolicyReportLifecycleStateDeleted,
	"failed":          SecurityPolicyReportLifecycleStateFailed,
	"needs_attention": SecurityPolicyReportLifecycleStateNeedsAttention,
}

// GetSecurityPolicyReportLifecycleStateEnumValues Enumerates the set of values for SecurityPolicyReportLifecycleStateEnum
func GetSecurityPolicyReportLifecycleStateEnumValues() []SecurityPolicyReportLifecycleStateEnum {
	values := make([]SecurityPolicyReportLifecycleStateEnum, 0)
	for _, v := range mappingSecurityPolicyReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityPolicyReportLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityPolicyReportLifecycleStateEnum
func GetSecurityPolicyReportLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SUCCEEDED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingSecurityPolicyReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityPolicyReportLifecycleStateEnum(val string) (SecurityPolicyReportLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityPolicyReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
