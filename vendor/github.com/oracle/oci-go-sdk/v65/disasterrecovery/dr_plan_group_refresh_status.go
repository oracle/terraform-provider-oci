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

// DrPlanGroupRefreshStatusEnum Enum with underlying type: string
type DrPlanGroupRefreshStatusEnum string

// Set of constants representing the allowable values for DrPlanGroupRefreshStatusEnum
const (
	DrPlanGroupRefreshStatusGroupAdded    DrPlanGroupRefreshStatusEnum = "GROUP_ADDED"
	DrPlanGroupRefreshStatusGroupDeleted  DrPlanGroupRefreshStatusEnum = "GROUP_DELETED"
	DrPlanGroupRefreshStatusGroupModified DrPlanGroupRefreshStatusEnum = "GROUP_MODIFIED"
)

var mappingDrPlanGroupRefreshStatusEnum = map[string]DrPlanGroupRefreshStatusEnum{
	"GROUP_ADDED":    DrPlanGroupRefreshStatusGroupAdded,
	"GROUP_DELETED":  DrPlanGroupRefreshStatusGroupDeleted,
	"GROUP_MODIFIED": DrPlanGroupRefreshStatusGroupModified,
}

var mappingDrPlanGroupRefreshStatusEnumLowerCase = map[string]DrPlanGroupRefreshStatusEnum{
	"group_added":    DrPlanGroupRefreshStatusGroupAdded,
	"group_deleted":  DrPlanGroupRefreshStatusGroupDeleted,
	"group_modified": DrPlanGroupRefreshStatusGroupModified,
}

// GetDrPlanGroupRefreshStatusEnumValues Enumerates the set of values for DrPlanGroupRefreshStatusEnum
func GetDrPlanGroupRefreshStatusEnumValues() []DrPlanGroupRefreshStatusEnum {
	values := make([]DrPlanGroupRefreshStatusEnum, 0)
	for _, v := range mappingDrPlanGroupRefreshStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrPlanGroupRefreshStatusEnumStringValues Enumerates the set of values in String for DrPlanGroupRefreshStatusEnum
func GetDrPlanGroupRefreshStatusEnumStringValues() []string {
	return []string{
		"GROUP_ADDED",
		"GROUP_DELETED",
		"GROUP_MODIFIED",
	}
}

// GetMappingDrPlanGroupRefreshStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrPlanGroupRefreshStatusEnum(val string) (DrPlanGroupRefreshStatusEnum, bool) {
	enum, ok := mappingDrPlanGroupRefreshStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
