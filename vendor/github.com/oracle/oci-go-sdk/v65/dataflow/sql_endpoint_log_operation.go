// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// SqlEndpointLogOperationEnum Enum with underlying type: string
type SqlEndpointLogOperationEnum string

// Set of constants representing the allowable values for SqlEndpointLogOperationEnum
const (
	SqlEndpointLogOperationEnable  SqlEndpointLogOperationEnum = "ENABLE"
	SqlEndpointLogOperationDisable SqlEndpointLogOperationEnum = "DISABLE"
	SqlEndpointLogOperationDelete  SqlEndpointLogOperationEnum = "DELETE"
)

var mappingSqlEndpointLogOperationEnum = map[string]SqlEndpointLogOperationEnum{
	"ENABLE":  SqlEndpointLogOperationEnable,
	"DISABLE": SqlEndpointLogOperationDisable,
	"DELETE":  SqlEndpointLogOperationDelete,
}

var mappingSqlEndpointLogOperationEnumLowerCase = map[string]SqlEndpointLogOperationEnum{
	"enable":  SqlEndpointLogOperationEnable,
	"disable": SqlEndpointLogOperationDisable,
	"delete":  SqlEndpointLogOperationDelete,
}

// GetSqlEndpointLogOperationEnumValues Enumerates the set of values for SqlEndpointLogOperationEnum
func GetSqlEndpointLogOperationEnumValues() []SqlEndpointLogOperationEnum {
	values := make([]SqlEndpointLogOperationEnum, 0)
	for _, v := range mappingSqlEndpointLogOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlEndpointLogOperationEnumStringValues Enumerates the set of values in String for SqlEndpointLogOperationEnum
func GetSqlEndpointLogOperationEnumStringValues() []string {
	return []string{
		"ENABLE",
		"DISABLE",
		"DELETE",
	}
}

// GetMappingSqlEndpointLogOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlEndpointLogOperationEnum(val string) (SqlEndpointLogOperationEnum, bool) {
	enum, ok := mappingSqlEndpointLogOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
