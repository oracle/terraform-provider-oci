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

var mappingSqlTuningTaskStatusTypes = map[string]SqlTuningTaskStatusTypesEnum{
	"COMPLETED":   SqlTuningTaskStatusTypesCompleted,
	"INITIAL":     SqlTuningTaskStatusTypesInitial,
	"EXECUTING":   SqlTuningTaskStatusTypesExecuting,
	"INTERRUPTED": SqlTuningTaskStatusTypesInterrupted,
	"ERROR":       SqlTuningTaskStatusTypesError,
}

// GetSqlTuningTaskStatusTypesEnumValues Enumerates the set of values for SqlTuningTaskStatusTypesEnum
func GetSqlTuningTaskStatusTypesEnumValues() []SqlTuningTaskStatusTypesEnum {
	values := make([]SqlTuningTaskStatusTypesEnum, 0)
	for _, v := range mappingSqlTuningTaskStatusTypes {
		values = append(values, v)
	}
	return values
}
