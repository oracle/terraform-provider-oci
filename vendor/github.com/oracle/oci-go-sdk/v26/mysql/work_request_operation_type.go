// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

// WorkRequestOperationTypeEnum Enum with underlying type: string
type WorkRequestOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestOperationTypeEnum
const (
	WorkRequestOperationTypeCreateDbsystem  WorkRequestOperationTypeEnum = "CREATE_DBSYSTEM"
	WorkRequestOperationTypeUpdateDbsystem  WorkRequestOperationTypeEnum = "UPDATE_DBSYSTEM"
	WorkRequestOperationTypeDeleteDbsystem  WorkRequestOperationTypeEnum = "DELETE_DBSYSTEM"
	WorkRequestOperationTypeStartDbsystem   WorkRequestOperationTypeEnum = "START_DBSYSTEM"
	WorkRequestOperationTypeStopDbsystem    WorkRequestOperationTypeEnum = "STOP_DBSYSTEM"
	WorkRequestOperationTypeRestartDbsystem WorkRequestOperationTypeEnum = "RESTART_DBSYSTEM"
)

var mappingWorkRequestOperationType = map[string]WorkRequestOperationTypeEnum{
	"CREATE_DBSYSTEM":  WorkRequestOperationTypeCreateDbsystem,
	"UPDATE_DBSYSTEM":  WorkRequestOperationTypeUpdateDbsystem,
	"DELETE_DBSYSTEM":  WorkRequestOperationTypeDeleteDbsystem,
	"START_DBSYSTEM":   WorkRequestOperationTypeStartDbsystem,
	"STOP_DBSYSTEM":    WorkRequestOperationTypeStopDbsystem,
	"RESTART_DBSYSTEM": WorkRequestOperationTypeRestartDbsystem,
}

// GetWorkRequestOperationTypeEnumValues Enumerates the set of values for WorkRequestOperationTypeEnum
func GetWorkRequestOperationTypeEnumValues() []WorkRequestOperationTypeEnum {
	values := make([]WorkRequestOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestOperationType {
		values = append(values, v)
	}
	return values
}
