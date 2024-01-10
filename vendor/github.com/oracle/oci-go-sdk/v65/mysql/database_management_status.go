// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// DatabaseManagementStatusEnum Enum with underlying type: string
type DatabaseManagementStatusEnum string

// Set of constants representing the allowable values for DatabaseManagementStatusEnum
const (
	DatabaseManagementStatusEnabled  DatabaseManagementStatusEnum = "ENABLED"
	DatabaseManagementStatusDisabled DatabaseManagementStatusEnum = "DISABLED"
)

var mappingDatabaseManagementStatusEnum = map[string]DatabaseManagementStatusEnum{
	"ENABLED":  DatabaseManagementStatusEnabled,
	"DISABLED": DatabaseManagementStatusDisabled,
}

var mappingDatabaseManagementStatusEnumLowerCase = map[string]DatabaseManagementStatusEnum{
	"enabled":  DatabaseManagementStatusEnabled,
	"disabled": DatabaseManagementStatusDisabled,
}

// GetDatabaseManagementStatusEnumValues Enumerates the set of values for DatabaseManagementStatusEnum
func GetDatabaseManagementStatusEnumValues() []DatabaseManagementStatusEnum {
	values := make([]DatabaseManagementStatusEnum, 0)
	for _, v := range mappingDatabaseManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseManagementStatusEnumStringValues Enumerates the set of values in String for DatabaseManagementStatusEnum
func GetDatabaseManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseManagementStatusEnum(val string) (DatabaseManagementStatusEnum, bool) {
	enum, ok := mappingDatabaseManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
