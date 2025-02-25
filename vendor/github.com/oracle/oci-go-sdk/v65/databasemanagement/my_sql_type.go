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

// MySqlTypeEnum Enum with underlying type: string
type MySqlTypeEnum string

// Set of constants representing the allowable values for MySqlTypeEnum
const (
	MySqlTypeExternal MySqlTypeEnum = "EXTERNAL"
	MySqlTypeMds      MySqlTypeEnum = "MDS"
)

var mappingMySqlTypeEnum = map[string]MySqlTypeEnum{
	"EXTERNAL": MySqlTypeExternal,
	"MDS":      MySqlTypeMds,
}

var mappingMySqlTypeEnumLowerCase = map[string]MySqlTypeEnum{
	"external": MySqlTypeExternal,
	"mds":      MySqlTypeMds,
}

// GetMySqlTypeEnumValues Enumerates the set of values for MySqlTypeEnum
func GetMySqlTypeEnumValues() []MySqlTypeEnum {
	values := make([]MySqlTypeEnum, 0)
	for _, v := range mappingMySqlTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlTypeEnumStringValues Enumerates the set of values in String for MySqlTypeEnum
func GetMySqlTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"MDS",
	}
}

// GetMappingMySqlTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlTypeEnum(val string) (MySqlTypeEnum, bool) {
	enum, ok := mappingMySqlTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
