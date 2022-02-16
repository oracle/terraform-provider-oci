// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	mappingDatabaseSubTypeEnumIgnoreCase := make(map[string]DatabaseSubTypeEnum)
	for k, v := range mappingDatabaseSubTypeEnum {
		mappingDatabaseSubTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseSubTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
