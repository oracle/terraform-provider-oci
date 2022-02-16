// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusCancelled  WorkRequestStatusEnum = "CANCELLED"
	WorkRequestStatusCancelling WorkRequestStatusEnum = "CANCELLING"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusInprogress WorkRequestStatusEnum = "INPROGRESS"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":   WorkRequestStatusAccepted,
	"CANCELLED":  WorkRequestStatusCancelled,
	"CANCELLING": WorkRequestStatusCancelling,
	"FAILED":     WorkRequestStatusFailed,
	"INPROGRESS": WorkRequestStatusInprogress,
	"SUCCEEDED":  WorkRequestStatusSucceeded,
}

// GetWorkRequestStatusEnumValues Enumerates the set of values for WorkRequestStatusEnum
func GetWorkRequestStatusEnumValues() []WorkRequestStatusEnum {
	values := make([]WorkRequestStatusEnum, 0)
	for _, v := range mappingWorkRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusEnumStringValues Enumerates the set of values in String for WorkRequestStatusEnum
func GetWorkRequestStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CANCELLED",
		"CANCELLING",
		"FAILED",
		"INPROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	mappingWorkRequestStatusEnumIgnoreCase := make(map[string]WorkRequestStatusEnum)
	for k, v := range mappingWorkRequestStatusEnum {
		mappingWorkRequestStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
