// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MaintenanceWindowLifecycleDetailsEnum Enum with underlying type: string
type MaintenanceWindowLifecycleDetailsEnum string

// Set of constants representing the allowable values for MaintenanceWindowLifecycleDetailsEnum
const (
	MaintenanceWindowLifecycleDetailsInProgress MaintenanceWindowLifecycleDetailsEnum = "IN_PROGRESS"
	MaintenanceWindowLifecycleDetailsScheduled  MaintenanceWindowLifecycleDetailsEnum = "SCHEDULED"
	MaintenanceWindowLifecycleDetailsCompleted  MaintenanceWindowLifecycleDetailsEnum = "COMPLETED"
)

var mappingMaintenanceWindowLifecycleDetailsEnum = map[string]MaintenanceWindowLifecycleDetailsEnum{
	"IN_PROGRESS": MaintenanceWindowLifecycleDetailsInProgress,
	"SCHEDULED":   MaintenanceWindowLifecycleDetailsScheduled,
	"COMPLETED":   MaintenanceWindowLifecycleDetailsCompleted,
}

var mappingMaintenanceWindowLifecycleDetailsEnumLowerCase = map[string]MaintenanceWindowLifecycleDetailsEnum{
	"in_progress": MaintenanceWindowLifecycleDetailsInProgress,
	"scheduled":   MaintenanceWindowLifecycleDetailsScheduled,
	"completed":   MaintenanceWindowLifecycleDetailsCompleted,
}

// GetMaintenanceWindowLifecycleDetailsEnumValues Enumerates the set of values for MaintenanceWindowLifecycleDetailsEnum
func GetMaintenanceWindowLifecycleDetailsEnumValues() []MaintenanceWindowLifecycleDetailsEnum {
	values := make([]MaintenanceWindowLifecycleDetailsEnum, 0)
	for _, v := range mappingMaintenanceWindowLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowLifecycleDetailsEnumStringValues Enumerates the set of values in String for MaintenanceWindowLifecycleDetailsEnum
func GetMaintenanceWindowLifecycleDetailsEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SCHEDULED",
		"COMPLETED",
	}
}

// GetMappingMaintenanceWindowLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowLifecycleDetailsEnum(val string) (MaintenanceWindowLifecycleDetailsEnum, bool) {
	enum, ok := mappingMaintenanceWindowLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
