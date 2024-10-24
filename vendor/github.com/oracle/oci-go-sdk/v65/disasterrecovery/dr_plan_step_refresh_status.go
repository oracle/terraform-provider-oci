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

// DrPlanStepRefreshStatusEnum Enum with underlying type: string
type DrPlanStepRefreshStatusEnum string

// Set of constants representing the allowable values for DrPlanStepRefreshStatusEnum
const (
	DrPlanStepRefreshStatusStepAdded   DrPlanStepRefreshStatusEnum = "STEP_ADDED"
	DrPlanStepRefreshStatusStepDeleted DrPlanStepRefreshStatusEnum = "STEP_DELETED"
)

var mappingDrPlanStepRefreshStatusEnum = map[string]DrPlanStepRefreshStatusEnum{
	"STEP_ADDED":   DrPlanStepRefreshStatusStepAdded,
	"STEP_DELETED": DrPlanStepRefreshStatusStepDeleted,
}

var mappingDrPlanStepRefreshStatusEnumLowerCase = map[string]DrPlanStepRefreshStatusEnum{
	"step_added":   DrPlanStepRefreshStatusStepAdded,
	"step_deleted": DrPlanStepRefreshStatusStepDeleted,
}

// GetDrPlanStepRefreshStatusEnumValues Enumerates the set of values for DrPlanStepRefreshStatusEnum
func GetDrPlanStepRefreshStatusEnumValues() []DrPlanStepRefreshStatusEnum {
	values := make([]DrPlanStepRefreshStatusEnum, 0)
	for _, v := range mappingDrPlanStepRefreshStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanStepRefreshStatusEnumStringValues Enumerates the set of values in String for DrPlanStepRefreshStatusEnum
func GetDrPlanStepRefreshStatusEnumStringValues() []string {
	return []string{
		"STEP_ADDED",
		"STEP_DELETED",
	}
}

// GetMappingDrPlanStepRefreshStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanStepRefreshStatusEnum(val string) (DrPlanStepRefreshStatusEnum, bool) {
	enum, ok := mappingDrPlanStepRefreshStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
