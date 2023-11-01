// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// MySqlDatabaseStatusEnum Enum with underlying type: string
type MySqlDatabaseStatusEnum string

// Set of constants representing the allowable values for MySqlDatabaseStatusEnum
const (
	MySqlDatabaseStatusUp      MySqlDatabaseStatusEnum = "UP"
	MySqlDatabaseStatusDown    MySqlDatabaseStatusEnum = "DOWN"
	MySqlDatabaseStatusUnknown MySqlDatabaseStatusEnum = "UNKNOWN"
)

var mappingMySqlDatabaseStatusEnum = map[string]MySqlDatabaseStatusEnum{
	"UP":      MySqlDatabaseStatusUp,
	"DOWN":    MySqlDatabaseStatusDown,
	"UNKNOWN": MySqlDatabaseStatusUnknown,
}

var mappingMySqlDatabaseStatusEnumLowerCase = map[string]MySqlDatabaseStatusEnum{
	"up":      MySqlDatabaseStatusUp,
	"down":    MySqlDatabaseStatusDown,
	"unknown": MySqlDatabaseStatusUnknown,
}

// GetMySqlDatabaseStatusEnumValues Enumerates the set of values for MySqlDatabaseStatusEnum
func GetMySqlDatabaseStatusEnumValues() []MySqlDatabaseStatusEnum {
	values := make([]MySqlDatabaseStatusEnum, 0)
	for _, v := range mappingMySqlDatabaseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlDatabaseStatusEnumStringValues Enumerates the set of values in String for MySqlDatabaseStatusEnum
func GetMySqlDatabaseStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingMySqlDatabaseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlDatabaseStatusEnum(val string) (MySqlDatabaseStatusEnum, bool) {
	enum, ok := mappingMySqlDatabaseStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
