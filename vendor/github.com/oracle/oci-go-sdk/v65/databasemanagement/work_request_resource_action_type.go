// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionTypeEnum
const (
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted    WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeRelated    WorkRequestResourceActionTypeEnum = "RELATED"
	WorkRequestResourceActionTypeFailed     WorkRequestResourceActionTypeEnum = "FAILED"
	WorkRequestResourceActionTypeAccepted   WorkRequestResourceActionTypeEnum = "ACCEPTED"
	WorkRequestResourceActionTypeEnabled    WorkRequestResourceActionTypeEnum = "ENABLED"
	WorkRequestResourceActionTypeDisabled   WorkRequestResourceActionTypeEnum = "DISABLED"
)

var mappingWorkRequestResourceActionTypeEnum = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
	"DELETED":     WorkRequestResourceActionTypeDeleted,
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
	"RELATED":     WorkRequestResourceActionTypeRelated,
	"FAILED":      WorkRequestResourceActionTypeFailed,
	"ACCEPTED":    WorkRequestResourceActionTypeAccepted,
	"ENABLED":     WorkRequestResourceActionTypeEnabled,
	"DISABLED":    WorkRequestResourceActionTypeDisabled,
}

var mappingWorkRequestResourceActionTypeEnumLowerCase = map[string]WorkRequestResourceActionTypeEnum{
	"created":     WorkRequestResourceActionTypeCreated,
	"updated":     WorkRequestResourceActionTypeUpdated,
	"deleted":     WorkRequestResourceActionTypeDeleted,
	"in_progress": WorkRequestResourceActionTypeInProgress,
	"related":     WorkRequestResourceActionTypeRelated,
	"failed":      WorkRequestResourceActionTypeFailed,
	"accepted":    WorkRequestResourceActionTypeAccepted,
	"enabled":     WorkRequestResourceActionTypeEnabled,
	"disabled":    WorkRequestResourceActionTypeDisabled,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestResourceActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestResourceActionTypeEnum
func GetWorkRequestResourceActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"RELATED",
		"FAILED",
		"ACCEPTED",
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingWorkRequestResourceActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestResourceActionTypeEnum(val string) (WorkRequestResourceActionTypeEnum, bool) {
	enum, ok := mappingWorkRequestResourceActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
