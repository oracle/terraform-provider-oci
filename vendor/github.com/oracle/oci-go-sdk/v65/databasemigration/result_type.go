// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// ResultTypeEnum Enum with underlying type: string
type ResultTypeEnum string

// Set of constants representing the allowable values for ResultTypeEnum
const (
	ResultTypeSucceeded ResultTypeEnum = "SUCCEEDED"
	ResultTypeFailed    ResultTypeEnum = "FAILED"
	ResultTypeTimedOut  ResultTypeEnum = "TIMED_OUT"
)

var mappingResultTypeEnum = map[string]ResultTypeEnum{
	"SUCCEEDED": ResultTypeSucceeded,
	"FAILED":    ResultTypeFailed,
	"TIMED_OUT": ResultTypeTimedOut,
}

var mappingResultTypeEnumLowerCase = map[string]ResultTypeEnum{
	"succeeded": ResultTypeSucceeded,
	"failed":    ResultTypeFailed,
	"timed_out": ResultTypeTimedOut,
}

// GetResultTypeEnumValues Enumerates the set of values for ResultTypeEnum
func GetResultTypeEnumValues() []ResultTypeEnum {
	values := make([]ResultTypeEnum, 0)
	for _, v := range mappingResultTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResultTypeEnumStringValues Enumerates the set of values in String for ResultTypeEnum
func GetResultTypeEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"TIMED_OUT",
	}
}

// GetMappingResultTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResultTypeEnum(val string) (ResultTypeEnum, bool) {
	enum, ok := mappingResultTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
