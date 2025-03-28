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

// SqlFirewallAllowedSqlLifecycleStateEnum Enum with underlying type: string
type SqlFirewallAllowedSqlLifecycleStateEnum string

// Set of constants representing the allowable values for SqlFirewallAllowedSqlLifecycleStateEnum
const (
	SqlFirewallAllowedSqlLifecycleStateActive  SqlFirewallAllowedSqlLifecycleStateEnum = "ACTIVE"
	SqlFirewallAllowedSqlLifecycleStateDeleted SqlFirewallAllowedSqlLifecycleStateEnum = "DELETED"
	SqlFirewallAllowedSqlLifecycleStateFailed  SqlFirewallAllowedSqlLifecycleStateEnum = "FAILED"
)

var mappingSqlFirewallAllowedSqlLifecycleStateEnum = map[string]SqlFirewallAllowedSqlLifecycleStateEnum{
	"ACTIVE":  SqlFirewallAllowedSqlLifecycleStateActive,
	"DELETED": SqlFirewallAllowedSqlLifecycleStateDeleted,
	"FAILED":  SqlFirewallAllowedSqlLifecycleStateFailed,
}

var mappingSqlFirewallAllowedSqlLifecycleStateEnumLowerCase = map[string]SqlFirewallAllowedSqlLifecycleStateEnum{
	"active":  SqlFirewallAllowedSqlLifecycleStateActive,
	"deleted": SqlFirewallAllowedSqlLifecycleStateDeleted,
	"failed":  SqlFirewallAllowedSqlLifecycleStateFailed,
}

// GetSqlFirewallAllowedSqlLifecycleStateEnumValues Enumerates the set of values for SqlFirewallAllowedSqlLifecycleStateEnum
func GetSqlFirewallAllowedSqlLifecycleStateEnumValues() []SqlFirewallAllowedSqlLifecycleStateEnum {
	values := make([]SqlFirewallAllowedSqlLifecycleStateEnum, 0)
	for _, v := range mappingSqlFirewallAllowedSqlLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlFirewallAllowedSqlLifecycleStateEnumStringValues Enumerates the set of values in String for SqlFirewallAllowedSqlLifecycleStateEnum
func GetSqlFirewallAllowedSqlLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSqlFirewallAllowedSqlLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlFirewallAllowedSqlLifecycleStateEnum(val string) (SqlFirewallAllowedSqlLifecycleStateEnum, bool) {
	enum, ok := mappingSqlFirewallAllowedSqlLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
