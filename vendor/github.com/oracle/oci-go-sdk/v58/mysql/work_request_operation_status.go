// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

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

var mappingWorkRequestOperationStatusEnum = map[string]WorkRequestOperationStatusEnum{
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
	for _, v := range mappingWorkRequestOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestOperationStatusEnumStringValues Enumerates the set of values in String for WorkRequestOperationStatusEnum
func GetWorkRequestOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestOperationStatusEnum(val string) (WorkRequestOperationStatusEnum, bool) {
	mappingWorkRequestOperationStatusEnumIgnoreCase := make(map[string]WorkRequestOperationStatusEnum)
	for k, v := range mappingWorkRequestOperationStatusEnum {
		mappingWorkRequestOperationStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestOperationStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
