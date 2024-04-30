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

// DatabaseSubTypeEnum Enum with underlying type: string
type DatabaseSubTypeEnum string

// Set of constants representing the allowable values for DatabaseSubTypeEnum
const (
	DatabaseSubTypeCdb    DatabaseSubTypeEnum = "CDB"
	DatabaseSubTypePdb    DatabaseSubTypeEnum = "PDB"
	DatabaseSubTypeNonCdb DatabaseSubTypeEnum = "NON_CDB"
	DatabaseSubTypeAcd    DatabaseSubTypeEnum = "ACD"
	DatabaseSubTypeAdb    DatabaseSubTypeEnum = "ADB"
)

var mappingDatabaseSubTypeEnum = map[string]DatabaseSubTypeEnum{
	"CDB":     DatabaseSubTypeCdb,
	"PDB":     DatabaseSubTypePdb,
	"NON_CDB": DatabaseSubTypeNonCdb,
	"ACD":     DatabaseSubTypeAcd,
	"ADB":     DatabaseSubTypeAdb,
}

var mappingDatabaseSubTypeEnumLowerCase = map[string]DatabaseSubTypeEnum{
	"cdb":     DatabaseSubTypeCdb,
	"pdb":     DatabaseSubTypePdb,
	"non_cdb": DatabaseSubTypeNonCdb,
	"acd":     DatabaseSubTypeAcd,
	"adb":     DatabaseSubTypeAdb,
}

// GetDatabaseSubTypeEnumValues Enumerates the set of values for DatabaseSubTypeEnum
func GetDatabaseSubTypeEnumValues() []DatabaseSubTypeEnum {
	values := make([]DatabaseSubTypeEnum, 0)
	for _, v := range mappingDatabaseSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseSubTypeEnumStringValues Enumerates the set of values in String for DatabaseSubTypeEnum
func GetDatabaseSubTypeEnumStringValues() []string {
	return []string{
		"CDB",
		"PDB",
		"NON_CDB",
		"ACD",
		"ADB",
	}
}

// GetMappingDatabaseSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseSubTypeEnum(val string) (DatabaseSubTypeEnum, bool) {
	enum, ok := mappingDatabaseSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
