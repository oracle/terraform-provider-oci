// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AutomaticDrPlanExecutionSubmissionStatusEnum Enum with underlying type: string
type AutomaticDrPlanExecutionSubmissionStatusEnum string

// Set of constants representing the allowable values for AutomaticDrPlanExecutionSubmissionStatusEnum
const (
	AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionStartedSuccessfully     AutomaticDrPlanExecutionSubmissionStatusEnum = "AUTOMATIC_DR_EXECUTION_STARTED_SUCCESSFULLY"
	AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionBlockedValidationFailed AutomaticDrPlanExecutionSubmissionStatusEnum = "AUTOMATIC_DR_EXECUTION_BLOCKED_VALIDATION_FAILED"
)

var mappingAutomaticDrPlanExecutionSubmissionStatusEnum = map[string]AutomaticDrPlanExecutionSubmissionStatusEnum{
	"AUTOMATIC_DR_EXECUTION_STARTED_SUCCESSFULLY":      AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionStartedSuccessfully,
	"AUTOMATIC_DR_EXECUTION_BLOCKED_VALIDATION_FAILED": AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionBlockedValidationFailed,
}

var mappingAutomaticDrPlanExecutionSubmissionStatusEnumLowerCase = map[string]AutomaticDrPlanExecutionSubmissionStatusEnum{
	"automatic_dr_execution_started_successfully":      AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionStartedSuccessfully,
	"automatic_dr_execution_blocked_validation_failed": AutomaticDrPlanExecutionSubmissionStatusAutomaticDrExecutionBlockedValidationFailed,
}

// GetAutomaticDrPlanExecutionSubmissionStatusEnumValues Enumerates the set of values for AutomaticDrPlanExecutionSubmissionStatusEnum
func GetAutomaticDrPlanExecutionSubmissionStatusEnumValues() []AutomaticDrPlanExecutionSubmissionStatusEnum {
	values := make([]AutomaticDrPlanExecutionSubmissionStatusEnum, 0)
	for _, v := range mappingAutomaticDrPlanExecutionSubmissionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticDrPlanExecutionSubmissionStatusEnumStringValues Enumerates the set of values in String for AutomaticDrPlanExecutionSubmissionStatusEnum
func GetAutomaticDrPlanExecutionSubmissionStatusEnumStringValues() []string {
	return []string{
		"AUTOMATIC_DR_EXECUTION_STARTED_SUCCESSFULLY",
		"AUTOMATIC_DR_EXECUTION_BLOCKED_VALIDATION_FAILED",
	}
}

// GetMappingAutomaticDrPlanExecutionSubmissionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticDrPlanExecutionSubmissionStatusEnum(val string) (AutomaticDrPlanExecutionSubmissionStatusEnum, bool) {
	enum, ok := mappingAutomaticDrPlanExecutionSubmissionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
