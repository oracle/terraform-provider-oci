// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// SqlQueryTypesEnum Enum with underlying type: string
type SqlQueryTypesEnum string

// Set of constants representing the allowable values for SqlQueryTypesEnum
const (
	SqlQueryTypesStatement SqlQueryTypesEnum = "STATEMENT"
	SqlQueryTypesSqlScript SqlQueryTypesEnum = "SQL_SCRIPT"
)

var mappingSqlQueryTypesEnum = map[string]SqlQueryTypesEnum{
	"STATEMENT":  SqlQueryTypesStatement,
	"SQL_SCRIPT": SqlQueryTypesSqlScript,
}

var mappingSqlQueryTypesEnumLowerCase = map[string]SqlQueryTypesEnum{
	"statement":  SqlQueryTypesStatement,
	"sql_script": SqlQueryTypesSqlScript,
}

// GetSqlQueryTypesEnumValues Enumerates the set of values for SqlQueryTypesEnum
func GetSqlQueryTypesEnumValues() []SqlQueryTypesEnum {
	values := make([]SqlQueryTypesEnum, 0)
	for _, v := range mappingSqlQueryTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlQueryTypesEnumStringValues Enumerates the set of values in String for SqlQueryTypesEnum
func GetSqlQueryTypesEnumStringValues() []string {
	return []string{
		"STATEMENT",
		"SQL_SCRIPT",
	}
}

// GetMappingSqlQueryTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlQueryTypesEnum(val string) (SqlQueryTypesEnum, bool) {
	enum, ok := mappingSqlQueryTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
