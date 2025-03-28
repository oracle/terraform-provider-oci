// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// ManagementOptionEnum Enum with underlying type: string
type ManagementOptionEnum string

// Set of constants representing the allowable values for ManagementOptionEnum
const (
	ManagementOptionBasic    ManagementOptionEnum = "BASIC"
	ManagementOptionAdvanced ManagementOptionEnum = "ADVANCED"
)

var mappingManagementOptionEnum = map[string]ManagementOptionEnum{
	"BASIC":    ManagementOptionBasic,
	"ADVANCED": ManagementOptionAdvanced,
}

var mappingManagementOptionEnumLowerCase = map[string]ManagementOptionEnum{
	"basic":    ManagementOptionBasic,
	"advanced": ManagementOptionAdvanced,
}

// GetManagementOptionEnumValues Enumerates the set of values for ManagementOptionEnum
func GetManagementOptionEnumValues() []ManagementOptionEnum {
	values := make([]ManagementOptionEnum, 0)
	for _, v := range mappingManagementOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementOptionEnumStringValues Enumerates the set of values in String for ManagementOptionEnum
func GetManagementOptionEnumStringValues() []string {
	return []string{
		"BASIC",
		"ADVANCED",
	}
}

// GetMappingManagementOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementOptionEnum(val string) (ManagementOptionEnum, bool) {
	enum, ok := mappingManagementOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
