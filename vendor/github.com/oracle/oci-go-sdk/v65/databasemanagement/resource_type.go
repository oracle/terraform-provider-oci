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

// ResourceTypeEnum Enum with underlying type: string
type ResourceTypeEnum string

// Set of constants representing the allowable values for ResourceTypeEnum
const (
	ResourceTypeOracleDb ResourceTypeEnum = "ORACLE_DB"
)

var mappingResourceTypeEnum = map[string]ResourceTypeEnum{
	"ORACLE_DB": ResourceTypeOracleDb,
}

var mappingResourceTypeEnumLowerCase = map[string]ResourceTypeEnum{
	"oracle_db": ResourceTypeOracleDb,
}

// GetResourceTypeEnumValues Enumerates the set of values for ResourceTypeEnum
func GetResourceTypeEnumValues() []ResourceTypeEnum {
	values := make([]ResourceTypeEnum, 0)
	for _, v := range mappingResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceTypeEnumStringValues Enumerates the set of values in String for ResourceTypeEnum
func GetResourceTypeEnumStringValues() []string {
	return []string{
		"ORACLE_DB",
	}
}

// GetMappingResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceTypeEnum(val string) (ResourceTypeEnum, bool) {
	enum, ok := mappingResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
