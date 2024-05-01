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

// DatabasePlanTypeEnumEnum Enum with underlying type: string
type DatabasePlanTypeEnumEnum string

// Set of constants representing the allowable values for DatabasePlanTypeEnumEnum
const (
	DatabasePlanTypeEnumDatabase DatabasePlanTypeEnumEnum = "DATABASE"
	DatabasePlanTypeEnumProfile  DatabasePlanTypeEnumEnum = "PROFILE"
	DatabasePlanTypeEnumOther    DatabasePlanTypeEnumEnum = "OTHER"
)

var mappingDatabasePlanTypeEnumEnum = map[string]DatabasePlanTypeEnumEnum{
	"DATABASE": DatabasePlanTypeEnumDatabase,
	"PROFILE":  DatabasePlanTypeEnumProfile,
	"OTHER":    DatabasePlanTypeEnumOther,
}

var mappingDatabasePlanTypeEnumEnumLowerCase = map[string]DatabasePlanTypeEnumEnum{
	"database": DatabasePlanTypeEnumDatabase,
	"profile":  DatabasePlanTypeEnumProfile,
	"other":    DatabasePlanTypeEnumOther,
}

// GetDatabasePlanTypeEnumEnumValues Enumerates the set of values for DatabasePlanTypeEnumEnum
func GetDatabasePlanTypeEnumEnumValues() []DatabasePlanTypeEnumEnum {
	values := make([]DatabasePlanTypeEnumEnum, 0)
	for _, v := range mappingDatabasePlanTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasePlanTypeEnumEnumStringValues Enumerates the set of values in String for DatabasePlanTypeEnumEnum
func GetDatabasePlanTypeEnumEnumStringValues() []string {
	return []string{
		"DATABASE",
		"PROFILE",
		"OTHER",
	}
}

// GetMappingDatabasePlanTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasePlanTypeEnumEnum(val string) (DatabasePlanTypeEnumEnum, bool) {
	enum, ok := mappingDatabasePlanTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
