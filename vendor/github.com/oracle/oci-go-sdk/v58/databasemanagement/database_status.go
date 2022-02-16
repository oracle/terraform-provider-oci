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

// DatabaseStatusEnum Enum with underlying type: string
type DatabaseStatusEnum string

// Set of constants representing the allowable values for DatabaseStatusEnum
const (
	DatabaseStatusUp      DatabaseStatusEnum = "UP"
	DatabaseStatusDown    DatabaseStatusEnum = "DOWN"
	DatabaseStatusUnknown DatabaseStatusEnum = "UNKNOWN"
)

var mappingDatabaseStatusEnum = map[string]DatabaseStatusEnum{
	"UP":      DatabaseStatusUp,
	"DOWN":    DatabaseStatusDown,
	"UNKNOWN": DatabaseStatusUnknown,
}

// GetDatabaseStatusEnumValues Enumerates the set of values for DatabaseStatusEnum
func GetDatabaseStatusEnumValues() []DatabaseStatusEnum {
	values := make([]DatabaseStatusEnum, 0)
	for _, v := range mappingDatabaseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseStatusEnumStringValues Enumerates the set of values in String for DatabaseStatusEnum
func GetDatabaseStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingDatabaseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseStatusEnum(val string) (DatabaseStatusEnum, bool) {
	mappingDatabaseStatusEnumIgnoreCase := make(map[string]DatabaseStatusEnum)
	for k, v := range mappingDatabaseStatusEnum {
		mappingDatabaseStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatabaseStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
