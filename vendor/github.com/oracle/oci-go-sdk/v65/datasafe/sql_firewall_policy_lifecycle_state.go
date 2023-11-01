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

// SqlFirewallPolicyLifecycleStateEnum Enum with underlying type: string
type SqlFirewallPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for SqlFirewallPolicyLifecycleStateEnum
const (
	SqlFirewallPolicyLifecycleStateCreating       SqlFirewallPolicyLifecycleStateEnum = "CREATING"
	SqlFirewallPolicyLifecycleStateUpdating       SqlFirewallPolicyLifecycleStateEnum = "UPDATING"
	SqlFirewallPolicyLifecycleStateActive         SqlFirewallPolicyLifecycleStateEnum = "ACTIVE"
	SqlFirewallPolicyLifecycleStateInactive       SqlFirewallPolicyLifecycleStateEnum = "INACTIVE"
	SqlFirewallPolicyLifecycleStateFailed         SqlFirewallPolicyLifecycleStateEnum = "FAILED"
	SqlFirewallPolicyLifecycleStateDeleting       SqlFirewallPolicyLifecycleStateEnum = "DELETING"
	SqlFirewallPolicyLifecycleStateDeleted        SqlFirewallPolicyLifecycleStateEnum = "DELETED"
	SqlFirewallPolicyLifecycleStateNeedsAttention SqlFirewallPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingSqlFirewallPolicyLifecycleStateEnum = map[string]SqlFirewallPolicyLifecycleStateEnum{
	"CREATING":        SqlFirewallPolicyLifecycleStateCreating,
	"UPDATING":        SqlFirewallPolicyLifecycleStateUpdating,
	"ACTIVE":          SqlFirewallPolicyLifecycleStateActive,
	"INACTIVE":        SqlFirewallPolicyLifecycleStateInactive,
	"FAILED":          SqlFirewallPolicyLifecycleStateFailed,
	"DELETING":        SqlFirewallPolicyLifecycleStateDeleting,
	"DELETED":         SqlFirewallPolicyLifecycleStateDeleted,
	"NEEDS_ATTENTION": SqlFirewallPolicyLifecycleStateNeedsAttention,
}

var mappingSqlFirewallPolicyLifecycleStateEnumLowerCase = map[string]SqlFirewallPolicyLifecycleStateEnum{
	"creating":        SqlFirewallPolicyLifecycleStateCreating,
	"updating":        SqlFirewallPolicyLifecycleStateUpdating,
	"active":          SqlFirewallPolicyLifecycleStateActive,
	"inactive":        SqlFirewallPolicyLifecycleStateInactive,
	"failed":          SqlFirewallPolicyLifecycleStateFailed,
	"deleting":        SqlFirewallPolicyLifecycleStateDeleting,
	"deleted":         SqlFirewallPolicyLifecycleStateDeleted,
	"needs_attention": SqlFirewallPolicyLifecycleStateNeedsAttention,
}

// GetSqlFirewallPolicyLifecycleStateEnumValues Enumerates the set of values for SqlFirewallPolicyLifecycleStateEnum
func GetSqlFirewallPolicyLifecycleStateEnumValues() []SqlFirewallPolicyLifecycleStateEnum {
	values := make([]SqlFirewallPolicyLifecycleStateEnum, 0)
	for _, v := range mappingSqlFirewallPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for SqlFirewallPolicyLifecycleStateEnum
func GetSqlFirewallPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingSqlFirewallPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallPolicyLifecycleStateEnum(val string) (SqlFirewallPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingSqlFirewallPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
