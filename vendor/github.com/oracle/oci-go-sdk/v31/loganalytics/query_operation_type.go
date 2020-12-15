// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// QueryOperationTypeEnum Enum with underlying type: string
type QueryOperationTypeEnum string

// Set of constants representing the allowable values for QueryOperationTypeEnum
const (
	QueryOperationTypeExecuteQueryJob QueryOperationTypeEnum = "EXECUTE_QUERY_JOB"
	QueryOperationTypeExecutePurgeJob QueryOperationTypeEnum = "EXECUTE_PURGE_JOB"
)

var mappingQueryOperationType = map[string]QueryOperationTypeEnum{
	"EXECUTE_QUERY_JOB": QueryOperationTypeExecuteQueryJob,
	"EXECUTE_PURGE_JOB": QueryOperationTypeExecutePurgeJob,
}

// GetQueryOperationTypeEnumValues Enumerates the set of values for QueryOperationTypeEnum
func GetQueryOperationTypeEnumValues() []QueryOperationTypeEnum {
	values := make([]QueryOperationTypeEnum, 0)
	for _, v := range mappingQueryOperationType {
		values = append(values, v)
	}
	return values
}
