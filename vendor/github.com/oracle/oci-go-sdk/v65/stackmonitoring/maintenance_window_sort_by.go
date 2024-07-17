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

// MaintenanceWindowSortByEnum Enum with underlying type: string
type MaintenanceWindowSortByEnum string

// Set of constants representing the allowable values for MaintenanceWindowSortByEnum
const (
	MaintenanceWindowSortByName        MaintenanceWindowSortByEnum = "NAME"
	MaintenanceWindowSortByStartTime   MaintenanceWindowSortByEnum = "START_TIME"
	MaintenanceWindowSortByEndTime     MaintenanceWindowSortByEnum = "END_TIME"
	MaintenanceWindowSortByTimeCreated MaintenanceWindowSortByEnum = "TIME_CREATED"
	MaintenanceWindowSortByTimeUpdated MaintenanceWindowSortByEnum = "TIME_UPDATED"
)

var mappingMaintenanceWindowSortByEnum = map[string]MaintenanceWindowSortByEnum{
	"NAME":         MaintenanceWindowSortByName,
	"START_TIME":   MaintenanceWindowSortByStartTime,
	"END_TIME":     MaintenanceWindowSortByEndTime,
	"TIME_CREATED": MaintenanceWindowSortByTimeCreated,
	"TIME_UPDATED": MaintenanceWindowSortByTimeUpdated,
}

var mappingMaintenanceWindowSortByEnumLowerCase = map[string]MaintenanceWindowSortByEnum{
	"name":         MaintenanceWindowSortByName,
	"start_time":   MaintenanceWindowSortByStartTime,
	"end_time":     MaintenanceWindowSortByEndTime,
	"time_created": MaintenanceWindowSortByTimeCreated,
	"time_updated": MaintenanceWindowSortByTimeUpdated,
}

// GetMaintenanceWindowSortByEnumValues Enumerates the set of values for MaintenanceWindowSortByEnum
func GetMaintenanceWindowSortByEnumValues() []MaintenanceWindowSortByEnum {
	values := make([]MaintenanceWindowSortByEnum, 0)
	for _, v := range mappingMaintenanceWindowSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowSortByEnumStringValues Enumerates the set of values in String for MaintenanceWindowSortByEnum
func GetMaintenanceWindowSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"START_TIME",
		"END_TIME",
		"TIME_CREATED",
		"TIME_UPDATED",
	}
}

// GetMappingMaintenanceWindowSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowSortByEnum(val string) (MaintenanceWindowSortByEnum, bool) {
	enum, ok := mappingMaintenanceWindowSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
