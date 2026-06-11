// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// ProtectedDatabaseAnalyticsGroupByEnum Enum with underlying type: string
type ProtectedDatabaseAnalyticsGroupByEnum string

// Set of constants representing the allowable values for ProtectedDatabaseAnalyticsGroupByEnum
const (
	ProtectedDatabaseAnalyticsGroupByHealth             ProtectedDatabaseAnalyticsGroupByEnum = "HEALTH"
	ProtectedDatabaseAnalyticsGroupByIsRedoLogsEnabled  ProtectedDatabaseAnalyticsGroupByEnum = "IS_REDO_LOGS_ENABLED"
	ProtectedDatabaseAnalyticsGroupByProtectionPolicyId ProtectedDatabaseAnalyticsGroupByEnum = "PROTECTION_POLICY_ID"
	ProtectedDatabaseAnalyticsGroupByLifecycleState     ProtectedDatabaseAnalyticsGroupByEnum = "LIFECYCLE_STATE"
)

var mappingProtectedDatabaseAnalyticsGroupByEnum = map[string]ProtectedDatabaseAnalyticsGroupByEnum{
	"HEALTH":               ProtectedDatabaseAnalyticsGroupByHealth,
	"IS_REDO_LOGS_ENABLED": ProtectedDatabaseAnalyticsGroupByIsRedoLogsEnabled,
	"PROTECTION_POLICY_ID": ProtectedDatabaseAnalyticsGroupByProtectionPolicyId,
	"LIFECYCLE_STATE":      ProtectedDatabaseAnalyticsGroupByLifecycleState,
}

var mappingProtectedDatabaseAnalyticsGroupByEnumLowerCase = map[string]ProtectedDatabaseAnalyticsGroupByEnum{
	"health":               ProtectedDatabaseAnalyticsGroupByHealth,
	"is_redo_logs_enabled": ProtectedDatabaseAnalyticsGroupByIsRedoLogsEnabled,
	"protection_policy_id": ProtectedDatabaseAnalyticsGroupByProtectionPolicyId,
	"lifecycle_state":      ProtectedDatabaseAnalyticsGroupByLifecycleState,
}

// GetProtectedDatabaseAnalyticsGroupByEnumValues Enumerates the set of values for ProtectedDatabaseAnalyticsGroupByEnum
func GetProtectedDatabaseAnalyticsGroupByEnumValues() []ProtectedDatabaseAnalyticsGroupByEnum {
	values := make([]ProtectedDatabaseAnalyticsGroupByEnum, 0)
	for _, v := range mappingProtectedDatabaseAnalyticsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectedDatabaseAnalyticsGroupByEnumStringValues Enumerates the set of values in String for ProtectedDatabaseAnalyticsGroupByEnum
func GetProtectedDatabaseAnalyticsGroupByEnumStringValues() []string {
	return []string{
		"HEALTH",
		"IS_REDO_LOGS_ENABLED",
		"PROTECTION_POLICY_ID",
		"LIFECYCLE_STATE",
	}
}

// GetMappingProtectedDatabaseAnalyticsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectedDatabaseAnalyticsGroupByEnum(val string) (ProtectedDatabaseAnalyticsGroupByEnum, bool) {
	enum, ok := mappingProtectedDatabaseAnalyticsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
