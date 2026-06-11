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

// ProtectedDatabaseAnalyticsMetricEnum Enum with underlying type: string
type ProtectedDatabaseAnalyticsMetricEnum string

// Set of constants representing the allowable values for ProtectedDatabaseAnalyticsMetricEnum
const (
	ProtectedDatabaseAnalyticsMetricCount                ProtectedDatabaseAnalyticsMetricEnum = "COUNT"
	ProtectedDatabaseAnalyticsMetricBackupSpaceUsedInGbs ProtectedDatabaseAnalyticsMetricEnum = "BACKUP_SPACE_USED_IN_GBS"
)

var mappingProtectedDatabaseAnalyticsMetricEnum = map[string]ProtectedDatabaseAnalyticsMetricEnum{
	"COUNT":                    ProtectedDatabaseAnalyticsMetricCount,
	"BACKUP_SPACE_USED_IN_GBS": ProtectedDatabaseAnalyticsMetricBackupSpaceUsedInGbs,
}

var mappingProtectedDatabaseAnalyticsMetricEnumLowerCase = map[string]ProtectedDatabaseAnalyticsMetricEnum{
	"count":                    ProtectedDatabaseAnalyticsMetricCount,
	"backup_space_used_in_gbs": ProtectedDatabaseAnalyticsMetricBackupSpaceUsedInGbs,
}

// GetProtectedDatabaseAnalyticsMetricEnumValues Enumerates the set of values for ProtectedDatabaseAnalyticsMetricEnum
func GetProtectedDatabaseAnalyticsMetricEnumValues() []ProtectedDatabaseAnalyticsMetricEnum {
	values := make([]ProtectedDatabaseAnalyticsMetricEnum, 0)
	for _, v := range mappingProtectedDatabaseAnalyticsMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetProtectedDatabaseAnalyticsMetricEnumStringValues Enumerates the set of values in String for ProtectedDatabaseAnalyticsMetricEnum
func GetProtectedDatabaseAnalyticsMetricEnumStringValues() []string {
	return []string{
		"COUNT",
		"BACKUP_SPACE_USED_IN_GBS",
	}
}

// GetMappingProtectedDatabaseAnalyticsMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProtectedDatabaseAnalyticsMetricEnum(val string) (ProtectedDatabaseAnalyticsMetricEnum, bool) {
	enum, ok := mappingProtectedDatabaseAnalyticsMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
