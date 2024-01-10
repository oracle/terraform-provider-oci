// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// DrPlanGroupExecutionStatusEnum Enum with underlying type: string
type DrPlanGroupExecutionStatusEnum string

// Set of constants representing the allowable values for DrPlanGroupExecutionStatusEnum
const (
	DrPlanGroupExecutionStatusQueued               DrPlanGroupExecutionStatusEnum = "QUEUED"
	DrPlanGroupExecutionStatusDisabled             DrPlanGroupExecutionStatusEnum = "DISABLED"
	DrPlanGroupExecutionStatusInProgress           DrPlanGroupExecutionStatusEnum = "IN_PROGRESS"
	DrPlanGroupExecutionStatusSucceeded            DrPlanGroupExecutionStatusEnum = "SUCCEEDED"
	DrPlanGroupExecutionStatusSucceededWithWarning DrPlanGroupExecutionStatusEnum = "SUCCEEDED_WITH_WARNING"
	DrPlanGroupExecutionStatusFailed               DrPlanGroupExecutionStatusEnum = "FAILED"
	DrPlanGroupExecutionStatusFailedIgnored        DrPlanGroupExecutionStatusEnum = "FAILED_IGNORED"
	DrPlanGroupExecutionStatusTimedOut             DrPlanGroupExecutionStatusEnum = "TIMED_OUT"
	DrPlanGroupExecutionStatusTimedOutIgnored      DrPlanGroupExecutionStatusEnum = "TIMED_OUT_IGNORED"
	DrPlanGroupExecutionStatusPaused               DrPlanGroupExecutionStatusEnum = "PAUSED"
	DrPlanGroupExecutionStatusCanceled             DrPlanGroupExecutionStatusEnum = "CANCELED"
)

var mappingDrPlanGroupExecutionStatusEnum = map[string]DrPlanGroupExecutionStatusEnum{
	"QUEUED":                 DrPlanGroupExecutionStatusQueued,
	"DISABLED":               DrPlanGroupExecutionStatusDisabled,
	"IN_PROGRESS":            DrPlanGroupExecutionStatusInProgress,
	"SUCCEEDED":              DrPlanGroupExecutionStatusSucceeded,
	"SUCCEEDED_WITH_WARNING": DrPlanGroupExecutionStatusSucceededWithWarning,
	"FAILED":                 DrPlanGroupExecutionStatusFailed,
	"FAILED_IGNORED":         DrPlanGroupExecutionStatusFailedIgnored,
	"TIMED_OUT":              DrPlanGroupExecutionStatusTimedOut,
	"TIMED_OUT_IGNORED":      DrPlanGroupExecutionStatusTimedOutIgnored,
	"PAUSED":                 DrPlanGroupExecutionStatusPaused,
	"CANCELED":               DrPlanGroupExecutionStatusCanceled,
}

var mappingDrPlanGroupExecutionStatusEnumLowerCase = map[string]DrPlanGroupExecutionStatusEnum{
	"queued":                 DrPlanGroupExecutionStatusQueued,
	"disabled":               DrPlanGroupExecutionStatusDisabled,
	"in_progress":            DrPlanGroupExecutionStatusInProgress,
	"succeeded":              DrPlanGroupExecutionStatusSucceeded,
	"succeeded_with_warning": DrPlanGroupExecutionStatusSucceededWithWarning,
	"failed":                 DrPlanGroupExecutionStatusFailed,
	"failed_ignored":         DrPlanGroupExecutionStatusFailedIgnored,
	"timed_out":              DrPlanGroupExecutionStatusTimedOut,
	"timed_out_ignored":      DrPlanGroupExecutionStatusTimedOutIgnored,
	"paused":                 DrPlanGroupExecutionStatusPaused,
	"canceled":               DrPlanGroupExecutionStatusCanceled,
}

// GetDrPlanGroupExecutionStatusEnumValues Enumerates the set of values for DrPlanGroupExecutionStatusEnum
func GetDrPlanGroupExecutionStatusEnumValues() []DrPlanGroupExecutionStatusEnum {
	values := make([]DrPlanGroupExecutionStatusEnum, 0)
	for _, v := range mappingDrPlanGroupExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanGroupExecutionStatusEnumStringValues Enumerates the set of values in String for DrPlanGroupExecutionStatusEnum
func GetDrPlanGroupExecutionStatusEnumStringValues() []string {
	return []string{
		"QUEUED",
		"DISABLED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"SUCCEEDED_WITH_WARNING",
		"FAILED",
		"FAILED_IGNORED",
		"TIMED_OUT",
		"TIMED_OUT_IGNORED",
		"PAUSED",
		"CANCELED",
	}
}

// GetMappingDrPlanGroupExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanGroupExecutionStatusEnum(val string) (DrPlanGroupExecutionStatusEnum, bool) {
	enum, ok := mappingDrPlanGroupExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
