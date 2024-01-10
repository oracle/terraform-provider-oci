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

// DrPlanStepExecutionStatusEnum Enum with underlying type: string
type DrPlanStepExecutionStatusEnum string

// Set of constants representing the allowable values for DrPlanStepExecutionStatusEnum
const (
	DrPlanStepExecutionStatusQueued          DrPlanStepExecutionStatusEnum = "QUEUED"
	DrPlanStepExecutionStatusDisabled        DrPlanStepExecutionStatusEnum = "DISABLED"
	DrPlanStepExecutionStatusInProgress      DrPlanStepExecutionStatusEnum = "IN_PROGRESS"
	DrPlanStepExecutionStatusSucceeded       DrPlanStepExecutionStatusEnum = "SUCCEEDED"
	DrPlanStepExecutionStatusFailed          DrPlanStepExecutionStatusEnum = "FAILED"
	DrPlanStepExecutionStatusFailedIgnored   DrPlanStepExecutionStatusEnum = "FAILED_IGNORED"
	DrPlanStepExecutionStatusTimedOut        DrPlanStepExecutionStatusEnum = "TIMED_OUT"
	DrPlanStepExecutionStatusTimedOutIgnored DrPlanStepExecutionStatusEnum = "TIMED_OUT_IGNORED"
	DrPlanStepExecutionStatusPaused          DrPlanStepExecutionStatusEnum = "PAUSED"
	DrPlanStepExecutionStatusCanceled        DrPlanStepExecutionStatusEnum = "CANCELED"
)

var mappingDrPlanStepExecutionStatusEnum = map[string]DrPlanStepExecutionStatusEnum{
	"QUEUED":            DrPlanStepExecutionStatusQueued,
	"DISABLED":          DrPlanStepExecutionStatusDisabled,
	"IN_PROGRESS":       DrPlanStepExecutionStatusInProgress,
	"SUCCEEDED":         DrPlanStepExecutionStatusSucceeded,
	"FAILED":            DrPlanStepExecutionStatusFailed,
	"FAILED_IGNORED":    DrPlanStepExecutionStatusFailedIgnored,
	"TIMED_OUT":         DrPlanStepExecutionStatusTimedOut,
	"TIMED_OUT_IGNORED": DrPlanStepExecutionStatusTimedOutIgnored,
	"PAUSED":            DrPlanStepExecutionStatusPaused,
	"CANCELED":          DrPlanStepExecutionStatusCanceled,
}

var mappingDrPlanStepExecutionStatusEnumLowerCase = map[string]DrPlanStepExecutionStatusEnum{
	"queued":            DrPlanStepExecutionStatusQueued,
	"disabled":          DrPlanStepExecutionStatusDisabled,
	"in_progress":       DrPlanStepExecutionStatusInProgress,
	"succeeded":         DrPlanStepExecutionStatusSucceeded,
	"failed":            DrPlanStepExecutionStatusFailed,
	"failed_ignored":    DrPlanStepExecutionStatusFailedIgnored,
	"timed_out":         DrPlanStepExecutionStatusTimedOut,
	"timed_out_ignored": DrPlanStepExecutionStatusTimedOutIgnored,
	"paused":            DrPlanStepExecutionStatusPaused,
	"canceled":          DrPlanStepExecutionStatusCanceled,
}

// GetDrPlanStepExecutionStatusEnumValues Enumerates the set of values for DrPlanStepExecutionStatusEnum
func GetDrPlanStepExecutionStatusEnumValues() []DrPlanStepExecutionStatusEnum {
	values := make([]DrPlanStepExecutionStatusEnum, 0)
	for _, v := range mappingDrPlanStepExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanStepExecutionStatusEnumStringValues Enumerates the set of values in String for DrPlanStepExecutionStatusEnum
func GetDrPlanStepExecutionStatusEnumStringValues() []string {
	return []string{
		"QUEUED",
		"DISABLED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"FAILED_IGNORED",
		"TIMED_OUT",
		"TIMED_OUT_IGNORED",
		"PAUSED",
		"CANCELED",
	}
}

// GetMappingDrPlanStepExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanStepExecutionStatusEnum(val string) (DrPlanStepExecutionStatusEnum, bool) {
	enum, ok := mappingDrPlanStepExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
