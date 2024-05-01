// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// DbManagementFeatureEnum Enum with underlying type: string
type DbManagementFeatureEnum string

// Set of constants representing the allowable values for DbManagementFeatureEnum
const (
	DbManagementFeatureDiagnosticsAndManagement DbManagementFeatureEnum = "DIAGNOSTICS_AND_MANAGEMENT"
)

var mappingDbManagementFeatureEnum = map[string]DbManagementFeatureEnum{
	"DIAGNOSTICS_AND_MANAGEMENT": DbManagementFeatureDiagnosticsAndManagement,
}

var mappingDbManagementFeatureEnumLowerCase = map[string]DbManagementFeatureEnum{
	"diagnostics_and_management": DbManagementFeatureDiagnosticsAndManagement,
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
		"DIAGNOSTICS_AND_MANAGEMENT",
	}
}

// GetMappingDbManagementFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbManagementFeatureEnum(val string) (DbManagementFeatureEnum, bool) {
	enum, ok := mappingDbManagementFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
