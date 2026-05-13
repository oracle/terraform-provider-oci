// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsSqlReportLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsSqlReportLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsSqlReportLifecycleStateEnum
const (
	DatabaseToolsSqlReportLifecycleStateActive  DatabaseToolsSqlReportLifecycleStateEnum = "ACTIVE"
	DatabaseToolsSqlReportLifecycleStateDeleted DatabaseToolsSqlReportLifecycleStateEnum = "DELETED"
)

var mappingDatabaseToolsSqlReportLifecycleStateEnum = map[string]DatabaseToolsSqlReportLifecycleStateEnum{
	"ACTIVE":  DatabaseToolsSqlReportLifecycleStateActive,
	"DELETED": DatabaseToolsSqlReportLifecycleStateDeleted,
}

var mappingDatabaseToolsSqlReportLifecycleStateEnumLowerCase = map[string]DatabaseToolsSqlReportLifecycleStateEnum{
	"active":  DatabaseToolsSqlReportLifecycleStateActive,
	"deleted": DatabaseToolsSqlReportLifecycleStateDeleted,
}

// GetDatabaseToolsSqlReportLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsSqlReportLifecycleStateEnum
func GetDatabaseToolsSqlReportLifecycleStateEnumValues() []DatabaseToolsSqlReportLifecycleStateEnum {
	values := make([]DatabaseToolsSqlReportLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsSqlReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsSqlReportLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsSqlReportLifecycleStateEnum
func GetDatabaseToolsSqlReportLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingDatabaseToolsSqlReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsSqlReportLifecycleStateEnum(val string) (DatabaseToolsSqlReportLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsSqlReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
