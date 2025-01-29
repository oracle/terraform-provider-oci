// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DatabaseHostedInEnum Enum with underlying type: string
type DatabaseHostedInEnum string

// Set of constants representing the allowable values for DatabaseHostedInEnum
const (
	DatabaseHostedInCloud    DatabaseHostedInEnum = "CLOUD"
	DatabaseHostedInExternal DatabaseHostedInEnum = "EXTERNAL"
)

var mappingDatabaseHostedInEnum = map[string]DatabaseHostedInEnum{
	"CLOUD":    DatabaseHostedInCloud,
	"EXTERNAL": DatabaseHostedInExternal,
}

var mappingDatabaseHostedInEnumLowerCase = map[string]DatabaseHostedInEnum{
	"cloud":    DatabaseHostedInCloud,
	"external": DatabaseHostedInExternal,
}

// GetDatabaseHostedInEnumValues Enumerates the set of values for DatabaseHostedInEnum
func GetDatabaseHostedInEnumValues() []DatabaseHostedInEnum {
	values := make([]DatabaseHostedInEnum, 0)
	for _, v := range mappingDatabaseHostedInEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseHostedInEnumStringValues Enumerates the set of values in String for DatabaseHostedInEnum
func GetDatabaseHostedInEnumStringValues() []string {
	return []string{
		"CLOUD",
		"EXTERNAL",
	}
}

// GetMappingDatabaseHostedInEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseHostedInEnum(val string) (DatabaseHostedInEnum, bool) {
	enum, ok := mappingDatabaseHostedInEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
