// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Control Plane API
//
// Use the Application Performance Monitoring Control Plane API to perform operations such as creating, updating,
// deleting and listing APM domains and monitoring the progress of these operations using the work request APIs. For more information, see Application Performance Monitoring (https://docs.cloud.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmcontrolplane

import (
	"strings"
)

// ActionTypesEnum Enum with underlying type: string
type ActionTypesEnum string

// Set of constants representing the allowable values for ActionTypesEnum
const (
	ActionTypesCreated    ActionTypesEnum = "CREATED"
	ActionTypesUpdated    ActionTypesEnum = "UPDATED"
	ActionTypesDeleted    ActionTypesEnum = "DELETED"
	ActionTypesInProgress ActionTypesEnum = "IN_PROGRESS"
	ActionTypesRelated    ActionTypesEnum = "RELATED"
)

var mappingActionTypesEnum = map[string]ActionTypesEnum{
	"CREATED":     ActionTypesCreated,
	"UPDATED":     ActionTypesUpdated,
	"DELETED":     ActionTypesDeleted,
	"IN_PROGRESS": ActionTypesInProgress,
	"RELATED":     ActionTypesRelated,
}

var mappingActionTypesEnumLowerCase = map[string]ActionTypesEnum{
	"created":     ActionTypesCreated,
	"updated":     ActionTypesUpdated,
	"deleted":     ActionTypesDeleted,
	"in_progress": ActionTypesInProgress,
	"related":     ActionTypesRelated,
}

// GetActionTypesEnumValues Enumerates the set of values for ActionTypesEnum
func GetActionTypesEnumValues() []ActionTypesEnum {
	values := make([]ActionTypesEnum, 0)
	for _, v := range mappingActionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypesEnumStringValues Enumerates the set of values in String for ActionTypesEnum
func GetActionTypesEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"RELATED",
	}
}

// GetMappingActionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypesEnum(val string) (ActionTypesEnum, bool) {
	enum, ok := mappingActionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
