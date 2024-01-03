// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// QueryOperationTypeEnum Enum with underlying type: string
type QueryOperationTypeEnum string

// Set of constants representing the allowable values for QueryOperationTypeEnum
const (
	QueryOperationTypeExecuteQueryJob QueryOperationTypeEnum = "EXECUTE_QUERY_JOB"
	QueryOperationTypeExecutePurgeJob QueryOperationTypeEnum = "EXECUTE_PURGE_JOB"
)

var mappingQueryOperationTypeEnum = map[string]QueryOperationTypeEnum{
	"EXECUTE_QUERY_JOB": QueryOperationTypeExecuteQueryJob,
	"EXECUTE_PURGE_JOB": QueryOperationTypeExecutePurgeJob,
}

var mappingQueryOperationTypeEnumLowerCase = map[string]QueryOperationTypeEnum{
	"execute_query_job": QueryOperationTypeExecuteQueryJob,
	"execute_purge_job": QueryOperationTypeExecutePurgeJob,
}

// GetQueryOperationTypeEnumValues Enumerates the set of values for QueryOperationTypeEnum
func GetQueryOperationTypeEnumValues() []QueryOperationTypeEnum {
	values := make([]QueryOperationTypeEnum, 0)
	for _, v := range mappingQueryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryOperationTypeEnumStringValues Enumerates the set of values in String for QueryOperationTypeEnum
func GetQueryOperationTypeEnumStringValues() []string {
	return []string{
		"EXECUTE_QUERY_JOB",
		"EXECUTE_PURGE_JOB",
	}
}

// GetMappingQueryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryOperationTypeEnum(val string) (QueryOperationTypeEnum, bool) {
	enum, ok := mappingQueryOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
