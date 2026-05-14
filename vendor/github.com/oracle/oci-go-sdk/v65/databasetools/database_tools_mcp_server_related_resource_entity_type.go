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

// DatabaseToolsMcpServerRelatedResourceEntityTypeEnum Enum with underlying type: string
type DatabaseToolsMcpServerRelatedResourceEntityTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpServerRelatedResourceEntityTypeEnum
const (
	DatabaseToolsMcpServerRelatedResourceEntityTypeAutonomousdatabase DatabaseToolsMcpServerRelatedResourceEntityTypeEnum = "AUTONOMOUSDATABASE"
	DatabaseToolsMcpServerRelatedResourceEntityTypeDatabase           DatabaseToolsMcpServerRelatedResourceEntityTypeEnum = "DATABASE"
	DatabaseToolsMcpServerRelatedResourceEntityTypePluggabledatabase  DatabaseToolsMcpServerRelatedResourceEntityTypeEnum = "PLUGGABLEDATABASE"
	DatabaseToolsMcpServerRelatedResourceEntityTypeMysqldbsystem      DatabaseToolsMcpServerRelatedResourceEntityTypeEnum = "MYSQLDBSYSTEM"
)

var mappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnum = map[string]DatabaseToolsMcpServerRelatedResourceEntityTypeEnum{
	"AUTONOMOUSDATABASE": DatabaseToolsMcpServerRelatedResourceEntityTypeAutonomousdatabase,
	"DATABASE":           DatabaseToolsMcpServerRelatedResourceEntityTypeDatabase,
	"PLUGGABLEDATABASE":  DatabaseToolsMcpServerRelatedResourceEntityTypePluggabledatabase,
	"MYSQLDBSYSTEM":      DatabaseToolsMcpServerRelatedResourceEntityTypeMysqldbsystem,
}

var mappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnumLowerCase = map[string]DatabaseToolsMcpServerRelatedResourceEntityTypeEnum{
	"autonomousdatabase": DatabaseToolsMcpServerRelatedResourceEntityTypeAutonomousdatabase,
	"database":           DatabaseToolsMcpServerRelatedResourceEntityTypeDatabase,
	"pluggabledatabase":  DatabaseToolsMcpServerRelatedResourceEntityTypePluggabledatabase,
	"mysqldbsystem":      DatabaseToolsMcpServerRelatedResourceEntityTypeMysqldbsystem,
}

// GetDatabaseToolsMcpServerRelatedResourceEntityTypeEnumValues Enumerates the set of values for DatabaseToolsMcpServerRelatedResourceEntityTypeEnum
func GetDatabaseToolsMcpServerRelatedResourceEntityTypeEnumValues() []DatabaseToolsMcpServerRelatedResourceEntityTypeEnum {
	values := make([]DatabaseToolsMcpServerRelatedResourceEntityTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpServerRelatedResourceEntityTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpServerRelatedResourceEntityTypeEnum
func GetDatabaseToolsMcpServerRelatedResourceEntityTypeEnumStringValues() []string {
	return []string{
		"AUTONOMOUSDATABASE",
		"DATABASE",
		"PLUGGABLEDATABASE",
		"MYSQLDBSYSTEM",
	}
}

// GetMappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnum(val string) (DatabaseToolsMcpServerRelatedResourceEntityTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpServerRelatedResourceEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
