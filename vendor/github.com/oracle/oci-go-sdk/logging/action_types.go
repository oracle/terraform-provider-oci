// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

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

var mappingActionTypes = map[string]ActionTypesEnum{
	"CREATED":     ActionTypesCreated,
	"UPDATED":     ActionTypesUpdated,
	"DELETED":     ActionTypesDeleted,
	"IN_PROGRESS": ActionTypesInProgress,
	"RELATED":     ActionTypesRelated,
}

// GetActionTypesEnumValues Enumerates the set of values for ActionTypesEnum
func GetActionTypesEnumValues() []ActionTypesEnum {
	values := make([]ActionTypesEnum, 0)
	for _, v := range mappingActionTypes {
		values = append(values, v)
	}
	return values
}
