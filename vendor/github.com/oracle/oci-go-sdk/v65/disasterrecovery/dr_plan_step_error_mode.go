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

// DrPlanStepErrorModeEnum Enum with underlying type: string
type DrPlanStepErrorModeEnum string

// Set of constants representing the allowable values for DrPlanStepErrorModeEnum
const (
	DrPlanStepErrorModeStopOnError     DrPlanStepErrorModeEnum = "STOP_ON_ERROR"
	DrPlanStepErrorModeContinueOnError DrPlanStepErrorModeEnum = "CONTINUE_ON_ERROR"
)

var mappingDrPlanStepErrorModeEnum = map[string]DrPlanStepErrorModeEnum{
	"STOP_ON_ERROR":     DrPlanStepErrorModeStopOnError,
	"CONTINUE_ON_ERROR": DrPlanStepErrorModeContinueOnError,
}

var mappingDrPlanStepErrorModeEnumLowerCase = map[string]DrPlanStepErrorModeEnum{
	"stop_on_error":     DrPlanStepErrorModeStopOnError,
	"continue_on_error": DrPlanStepErrorModeContinueOnError,
}

// GetDrPlanStepErrorModeEnumValues Enumerates the set of values for DrPlanStepErrorModeEnum
func GetDrPlanStepErrorModeEnumValues() []DrPlanStepErrorModeEnum {
	values := make([]DrPlanStepErrorModeEnum, 0)
	for _, v := range mappingDrPlanStepErrorModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanStepErrorModeEnumStringValues Enumerates the set of values in String for DrPlanStepErrorModeEnum
func GetDrPlanStepErrorModeEnumStringValues() []string {
	return []string{
		"STOP_ON_ERROR",
		"CONTINUE_ON_ERROR",
	}
}

// GetMappingDrPlanStepErrorModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanStepErrorModeEnum(val string) (DrPlanStepErrorModeEnum, bool) {
	enum, ok := mappingDrPlanStepErrorModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
