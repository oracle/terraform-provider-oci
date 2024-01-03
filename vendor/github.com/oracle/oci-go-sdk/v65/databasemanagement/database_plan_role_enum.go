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

// DatabasePlanRoleEnumEnum Enum with underlying type: string
type DatabasePlanRoleEnumEnum string

// Set of constants representing the allowable values for DatabasePlanRoleEnumEnum
const (
	DatabasePlanRoleEnumPrimary DatabasePlanRoleEnumEnum = "PRIMARY"
	DatabasePlanRoleEnumStandby DatabasePlanRoleEnumEnum = "STANDBY"
	DatabasePlanRoleEnumNone    DatabasePlanRoleEnumEnum = "NONE"
)

var mappingDatabasePlanRoleEnumEnum = map[string]DatabasePlanRoleEnumEnum{
	"PRIMARY": DatabasePlanRoleEnumPrimary,
	"STANDBY": DatabasePlanRoleEnumStandby,
	"NONE":    DatabasePlanRoleEnumNone,
}

var mappingDatabasePlanRoleEnumEnumLowerCase = map[string]DatabasePlanRoleEnumEnum{
	"primary": DatabasePlanRoleEnumPrimary,
	"standby": DatabasePlanRoleEnumStandby,
	"none":    DatabasePlanRoleEnumNone,
}

// GetDatabasePlanRoleEnumEnumValues Enumerates the set of values for DatabasePlanRoleEnumEnum
func GetDatabasePlanRoleEnumEnumValues() []DatabasePlanRoleEnumEnum {
	values := make([]DatabasePlanRoleEnumEnum, 0)
	for _, v := range mappingDatabasePlanRoleEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasePlanRoleEnumEnumStringValues Enumerates the set of values in String for DatabasePlanRoleEnumEnum
func GetDatabasePlanRoleEnumEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"NONE",
	}
}

// GetMappingDatabasePlanRoleEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasePlanRoleEnumEnum(val string) (DatabasePlanRoleEnumEnum, bool) {
	enum, ok := mappingDatabasePlanRoleEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
