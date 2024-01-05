// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

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
