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

// SqlTuningTaskStatusTypesEnum Enum with underlying type: string
type SqlTuningTaskStatusTypesEnum string

// Set of constants representing the allowable values for SqlTuningTaskStatusTypesEnum
const (
	SqlTuningTaskStatusTypesCompleted   SqlTuningTaskStatusTypesEnum = "COMPLETED"
	SqlTuningTaskStatusTypesInitial     SqlTuningTaskStatusTypesEnum = "INITIAL"
	SqlTuningTaskStatusTypesExecuting   SqlTuningTaskStatusTypesEnum = "EXECUTING"
	SqlTuningTaskStatusTypesInterrupted SqlTuningTaskStatusTypesEnum = "INTERRUPTED"
	SqlTuningTaskStatusTypesError       SqlTuningTaskStatusTypesEnum = "ERROR"
)

var mappingSqlTuningTaskStatusTypesEnum = map[string]SqlTuningTaskStatusTypesEnum{
	"COMPLETED":   SqlTuningTaskStatusTypesCompleted,
	"INITIAL":     SqlTuningTaskStatusTypesInitial,
	"EXECUTING":   SqlTuningTaskStatusTypesExecuting,
	"INTERRUPTED": SqlTuningTaskStatusTypesInterrupted,
	"ERROR":       SqlTuningTaskStatusTypesError,
}

var mappingSqlTuningTaskStatusTypesEnumLowerCase = map[string]SqlTuningTaskStatusTypesEnum{
	"completed":   SqlTuningTaskStatusTypesCompleted,
	"initial":     SqlTuningTaskStatusTypesInitial,
	"executing":   SqlTuningTaskStatusTypesExecuting,
	"interrupted": SqlTuningTaskStatusTypesInterrupted,
	"error":       SqlTuningTaskStatusTypesError,
}

// GetSqlTuningTaskStatusTypesEnumValues Enumerates the set of values for SqlTuningTaskStatusTypesEnum
func GetSqlTuningTaskStatusTypesEnumValues() []SqlTuningTaskStatusTypesEnum {
	values := make([]SqlTuningTaskStatusTypesEnum, 0)
	for _, v := range mappingSqlTuningTaskStatusTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningTaskStatusTypesEnumStringValues Enumerates the set of values in String for SqlTuningTaskStatusTypesEnum
func GetSqlTuningTaskStatusTypesEnumStringValues() []string {
	return []string{
		"COMPLETED",
		"INITIAL",
		"EXECUTING",
		"INTERRUPTED",
		"ERROR",
	}
}

// GetMappingSqlTuningTaskStatusTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningTaskStatusTypesEnum(val string) (SqlTuningTaskStatusTypesEnum, bool) {
	enum, ok := mappingSqlTuningTaskStatusTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
