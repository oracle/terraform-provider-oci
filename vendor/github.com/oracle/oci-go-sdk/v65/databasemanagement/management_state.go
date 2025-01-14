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

// ManagementStateEnum Enum with underlying type: string
type ManagementStateEnum string

// Set of constants representing the allowable values for ManagementStateEnum
const (
	ManagementStateEnabled  ManagementStateEnum = "ENABLED"
	ManagementStateDisabled ManagementStateEnum = "DISABLED"
)

var mappingManagementStateEnum = map[string]ManagementStateEnum{
	"ENABLED":  ManagementStateEnabled,
	"DISABLED": ManagementStateDisabled,
}

var mappingManagementStateEnumLowerCase = map[string]ManagementStateEnum{
	"enabled":  ManagementStateEnabled,
	"disabled": ManagementStateDisabled,
}

// GetManagementStateEnumValues Enumerates the set of values for ManagementStateEnum
func GetManagementStateEnumValues() []ManagementStateEnum {
	values := make([]ManagementStateEnum, 0)
	for _, v := range mappingManagementStateEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementStateEnumStringValues Enumerates the set of values in String for ManagementStateEnum
func GetManagementStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingManagementStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementStateEnum(val string) (ManagementStateEnum, bool) {
	enum, ok := mappingManagementStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
