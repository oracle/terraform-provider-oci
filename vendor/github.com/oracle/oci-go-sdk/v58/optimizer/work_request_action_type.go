// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// WorkRequestActionTypeEnum Enum with underlying type: string
type WorkRequestActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestActionTypeEnum
const (
	WorkRequestActionTypeCreated    WorkRequestActionTypeEnum = "CREATED"
	WorkRequestActionTypeUpdated    WorkRequestActionTypeEnum = "UPDATED"
	WorkRequestActionTypeDeleted    WorkRequestActionTypeEnum = "DELETED"
	WorkRequestActionTypeInProgress WorkRequestActionTypeEnum = "IN_PROGRESS"
	WorkRequestActionTypeRelated    WorkRequestActionTypeEnum = "RELATED"
)

var mappingWorkRequestActionTypeEnum = map[string]WorkRequestActionTypeEnum{
	"CREATED":     WorkRequestActionTypeCreated,
	"UPDATED":     WorkRequestActionTypeUpdated,
	"DELETED":     WorkRequestActionTypeDeleted,
	"IN_PROGRESS": WorkRequestActionTypeInProgress,
	"RELATED":     WorkRequestActionTypeRelated,
}

// GetWorkRequestActionTypeEnumValues Enumerates the set of values for WorkRequestActionTypeEnum
func GetWorkRequestActionTypeEnumValues() []WorkRequestActionTypeEnum {
	values := make([]WorkRequestActionTypeEnum, 0)
	for _, v := range mappingWorkRequestActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestActionTypeEnumStringValues Enumerates the set of values in String for WorkRequestActionTypeEnum
func GetWorkRequestActionTypeEnumStringValues() []string {
	return []string{
		"CREATED",
		"UPDATED",
		"DELETED",
		"IN_PROGRESS",
		"RELATED",
	}
}

// GetMappingWorkRequestActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestActionTypeEnum(val string) (WorkRequestActionTypeEnum, bool) {
	mappingWorkRequestActionTypeEnumIgnoreCase := make(map[string]WorkRequestActionTypeEnum)
	for k, v := range mappingWorkRequestActionTypeEnum {
		mappingWorkRequestActionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkRequestActionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
