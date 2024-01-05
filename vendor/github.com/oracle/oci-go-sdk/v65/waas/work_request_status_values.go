// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"strings"
)

// WorkRequestStatusValuesEnum Enum with underlying type: string
type WorkRequestStatusValuesEnum string

// Set of constants representing the allowable values for WorkRequestStatusValuesEnum
const (
	WorkRequestStatusValuesAccepted   WorkRequestStatusValuesEnum = "ACCEPTED"
	WorkRequestStatusValuesInProgress WorkRequestStatusValuesEnum = "IN_PROGRESS"
	WorkRequestStatusValuesFailed     WorkRequestStatusValuesEnum = "FAILED"
	WorkRequestStatusValuesSucceeded  WorkRequestStatusValuesEnum = "SUCCEEDED"
	WorkRequestStatusValuesCanceling  WorkRequestStatusValuesEnum = "CANCELING"
	WorkRequestStatusValuesCanceled   WorkRequestStatusValuesEnum = "CANCELED"
)

var mappingWorkRequestStatusValuesEnum = map[string]WorkRequestStatusValuesEnum{
	"ACCEPTED":    WorkRequestStatusValuesAccepted,
	"IN_PROGRESS": WorkRequestStatusValuesInProgress,
	"FAILED":      WorkRequestStatusValuesFailed,
	"SUCCEEDED":   WorkRequestStatusValuesSucceeded,
	"CANCELING":   WorkRequestStatusValuesCanceling,
	"CANCELED":    WorkRequestStatusValuesCanceled,
}

var mappingWorkRequestStatusValuesEnumLowerCase = map[string]WorkRequestStatusValuesEnum{
	"accepted":    WorkRequestStatusValuesAccepted,
	"in_progress": WorkRequestStatusValuesInProgress,
	"failed":      WorkRequestStatusValuesFailed,
	"succeeded":   WorkRequestStatusValuesSucceeded,
	"canceling":   WorkRequestStatusValuesCanceling,
	"canceled":    WorkRequestStatusValuesCanceled,
}

// GetWorkRequestStatusValuesEnumValues Enumerates the set of values for WorkRequestStatusValuesEnum
func GetWorkRequestStatusValuesEnumValues() []WorkRequestStatusValuesEnum {
	values := make([]WorkRequestStatusValuesEnum, 0)
	for _, v := range mappingWorkRequestStatusValuesEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestStatusValuesEnumStringValues Enumerates the set of values in String for WorkRequestStatusValuesEnum
func GetWorkRequestStatusValuesEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkRequestStatusValuesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestStatusValuesEnum(val string) (WorkRequestStatusValuesEnum, bool) {
	enum, ok := mappingWorkRequestStatusValuesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
