// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// DbManagementFeatureEnum Enum with underlying type: string
type DbManagementFeatureEnum string

// Set of constants representing the allowable values for DbManagementFeatureEnum
const (
	DbManagementFeatureDiagnosticsAndPerformance DbManagementFeatureEnum = "DIAGNOSTICS_AND_PERFORMANCE"
	DbManagementFeatureDbLifecycleManagement     DbManagementFeatureEnum = "DB_LIFECYCLE_MANAGEMENT"
	DbManagementFeatureSqlwatch                  DbManagementFeatureEnum = "SQLWATCH"
)

var mappingDbManagementFeatureEnum = map[string]DbManagementFeatureEnum{
	"DIAGNOSTICS_AND_PERFORMANCE": DbManagementFeatureDiagnosticsAndPerformance,
	"DB_LIFECYCLE_MANAGEMENT":     DbManagementFeatureDbLifecycleManagement,
	"SQLWATCH":                    DbManagementFeatureSqlwatch,
}

var mappingDbManagementFeatureEnumLowerCase = map[string]DbManagementFeatureEnum{
	"diagnostics_and_performance": DbManagementFeatureDiagnosticsAndPerformance,
	"db_lifecycle_management":     DbManagementFeatureDbLifecycleManagement,
	"sqlwatch":                    DbManagementFeatureSqlwatch,
}

// GetDbManagementFeatureEnumValues Enumerates the set of values for DbManagementFeatureEnum
func GetDbManagementFeatureEnumValues() []DbManagementFeatureEnum {
	values := make([]DbManagementFeatureEnum, 0)
	for _, v := range mappingDbManagementFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetDbManagementFeatureEnumStringValues Enumerates the set of values in String for DbManagementFeatureEnum
func GetDbManagementFeatureEnumStringValues() []string {
	return []string{
		"DIAGNOSTICS_AND_PERFORMANCE",
		"DB_LIFECYCLE_MANAGEMENT",
		"SQLWATCH",
	}
}

// GetMappingDbManagementFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbManagementFeatureEnum(val string) (DbManagementFeatureEnum, bool) {
	enum, ok := mappingDbManagementFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
