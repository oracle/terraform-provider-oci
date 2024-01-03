// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// DatabaseConnectionTypesEnum Enum with underlying type: string
type DatabaseConnectionTypesEnum string

// Set of constants representing the allowable values for DatabaseConnectionTypesEnum
const (
	DatabaseConnectionTypesManual         DatabaseConnectionTypesEnum = "MANUAL"
	DatabaseConnectionTypesAutonomous     DatabaseConnectionTypesEnum = "AUTONOMOUS"
	DatabaseConnectionTypesUserManagedOci DatabaseConnectionTypesEnum = "USER_MANAGED_OCI"
)

var mappingDatabaseConnectionTypesEnum = map[string]DatabaseConnectionTypesEnum{
	"MANUAL":           DatabaseConnectionTypesManual,
	"AUTONOMOUS":       DatabaseConnectionTypesAutonomous,
	"USER_MANAGED_OCI": DatabaseConnectionTypesUserManagedOci,
}

var mappingDatabaseConnectionTypesEnumLowerCase = map[string]DatabaseConnectionTypesEnum{
	"manual":           DatabaseConnectionTypesManual,
	"autonomous":       DatabaseConnectionTypesAutonomous,
	"user_managed_oci": DatabaseConnectionTypesUserManagedOci,
}

// GetDatabaseConnectionTypesEnumValues Enumerates the set of values for DatabaseConnectionTypesEnum
func GetDatabaseConnectionTypesEnumValues() []DatabaseConnectionTypesEnum {
	values := make([]DatabaseConnectionTypesEnum, 0)
	for _, v := range mappingDatabaseConnectionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionTypesEnumStringValues Enumerates the set of values in String for DatabaseConnectionTypesEnum
func GetDatabaseConnectionTypesEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTONOMOUS",
		"USER_MANAGED_OCI",
	}
}

// GetMappingDatabaseConnectionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionTypesEnum(val string) (DatabaseConnectionTypesEnum, bool) {
	enum, ok := mappingDatabaseConnectionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
