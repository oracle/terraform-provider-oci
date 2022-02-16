// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

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
	ActionTypesFailed     ActionTypesEnum = "FAILED"
)

var mappingActionTypesEnum = map[string]ActionTypesEnum{
	"CREATED":     ActionTypesCreated,
	"UPDATED":     ActionTypesUpdated,
	"DELETED":     ActionTypesDeleted,
	"IN_PROGRESS": ActionTypesInProgress,
	"FAILED":      ActionTypesFailed,
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
		"FAILED",
	}
}

// GetMappingActionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypesEnum(val string) (ActionTypesEnum, bool) {
	mappingActionTypesEnumIgnoreCase := make(map[string]ActionTypesEnum)
	for k, v := range mappingActionTypesEnum {
		mappingActionTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingActionTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
