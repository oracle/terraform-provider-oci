// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

// WorkRequestOperationStatusEnum Enum with underlying type: string
type WorkRequestOperationStatusEnum string

// Set of constants representing the allowable values for WorkRequestOperationStatusEnum
const (
	WorkRequestOperationStatusAccepted   WorkRequestOperationStatusEnum = "ACCEPTED"
	WorkRequestOperationStatusInProgress WorkRequestOperationStatusEnum = "IN_PROGRESS"
	WorkRequestOperationStatusFailed     WorkRequestOperationStatusEnum = "FAILED"
	WorkRequestOperationStatusSucceeded  WorkRequestOperationStatusEnum = "SUCCEEDED"
	WorkRequestOperationStatusCanceling  WorkRequestOperationStatusEnum = "CANCELING"
	WorkRequestOperationStatusCanceled   WorkRequestOperationStatusEnum = "CANCELED"
)

var mappingWorkRequestOperationStatus = map[string]WorkRequestOperationStatusEnum{
	"ACCEPTED":    WorkRequestOperationStatusAccepted,
	"IN_PROGRESS": WorkRequestOperationStatusInProgress,
	"FAILED":      WorkRequestOperationStatusFailed,
	"SUCCEEDED":   WorkRequestOperationStatusSucceeded,
	"CANCELING":   WorkRequestOperationStatusCanceling,
	"CANCELED":    WorkRequestOperationStatusCanceled,
}

// GetWorkRequestOperationStatusEnumValues Enumerates the set of values for WorkRequestOperationStatusEnum
func GetWorkRequestOperationStatusEnumValues() []WorkRequestOperationStatusEnum {
	values := make([]WorkRequestOperationStatusEnum, 0)
	for _, v := range mappingWorkRequestOperationStatus {
		values = append(values, v)
	}
	return values
}
