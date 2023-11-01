// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SqlOutParamTypesEnum Enum with underlying type: string
type SqlOutParamTypesEnum string

// Set of constants representing the allowable values for SqlOutParamTypesEnum
const (
	SqlOutParamTypesSqlCursor SqlOutParamTypesEnum = "SQL_CURSOR"
	SqlOutParamTypesArray     SqlOutParamTypesEnum = "ARRAY"
)

var mappingSqlOutParamTypesEnum = map[string]SqlOutParamTypesEnum{
	"SQL_CURSOR": SqlOutParamTypesSqlCursor,
	"ARRAY":      SqlOutParamTypesArray,
}

var mappingSqlOutParamTypesEnumLowerCase = map[string]SqlOutParamTypesEnum{
	"sql_cursor": SqlOutParamTypesSqlCursor,
	"array":      SqlOutParamTypesArray,
}

// GetSqlOutParamTypesEnumValues Enumerates the set of values for SqlOutParamTypesEnum
func GetSqlOutParamTypesEnumValues() []SqlOutParamTypesEnum {
	values := make([]SqlOutParamTypesEnum, 0)
	for _, v := range mappingSqlOutParamTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlOutParamTypesEnumStringValues Enumerates the set of values in String for SqlOutParamTypesEnum
func GetSqlOutParamTypesEnumStringValues() []string {
	return []string{
		"SQL_CURSOR",
		"ARRAY",
	}
}

// GetMappingSqlOutParamTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlOutParamTypesEnum(val string) (SqlOutParamTypesEnum, bool) {
	enum, ok := mappingSqlOutParamTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
