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

// WorkRequestStatusEnum Enum with underlying type: string
type WorkRequestStatusEnum string

// Set of constants representing the allowable values for WorkRequestStatusEnum
const (
	WorkRequestStatusAccepted   WorkRequestStatusEnum = "ACCEPTED"
	WorkRequestStatusCanceled   WorkRequestStatusEnum = "CANCELED"
	WorkRequestStatusFailed     WorkRequestStatusEnum = "FAILED"
	WorkRequestStatusInProgress WorkRequestStatusEnum = "IN_PROGRESS"
	WorkRequestStatusSucceeded  WorkRequestStatusEnum = "SUCCEEDED"
)

var mappingWorkRequestStatusEnum = map[string]WorkRequestStatusEnum{
	"ACCEPTED":    WorkRequestStatusAccepted,
	"CANCELED":    WorkRequestStatusCanceled,
	"FAILED":      WorkRequestStatusFailed,
	"IN_PROGRESS": WorkRequestStatusInProgress,
	"SUCCEEDED":   WorkRequestStatusSucceeded,
}

var mappingWorkRequestStatusEnumLowerCase = map[string]WorkRequestStatusEnum{
	"accepted":    WorkRequestStatusAccepted,
	"canceled":    WorkRequestStatusCanceled,
	"failed":      WorkRequestStatusFailed,
	"in_progress": WorkRequestStatusInProgress,
	"succeeded":   WorkRequestStatusSucceeded,
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
		"CANCELED",
		"FAILED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingWorkRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusEnum(val string) (WorkRequestStatusEnum, bool) {
	enum, ok := mappingWorkRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
