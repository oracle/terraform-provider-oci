// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeExecuteSql                  OperationTypeEnum = "EXECUTE_SQL"
	OperationTypeExecuteSqlScheduledDeletion OperationTypeEnum = "EXECUTE_SQL_SCHEDULED_DELETION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"EXECUTE_SQL":                    OperationTypeExecuteSql,
	"EXECUTE_SQL_SCHEDULED_DELETION": OperationTypeExecuteSqlScheduledDeletion,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"execute_sql":                    OperationTypeExecuteSql,
	"execute_sql_scheduled_deletion": OperationTypeExecuteSqlScheduledDeletion,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"EXECUTE_SQL",
		"EXECUTE_SQL_SCHEDULED_DELETION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
