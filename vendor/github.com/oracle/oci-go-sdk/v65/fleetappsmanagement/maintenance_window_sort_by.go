// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// MaintenanceWindowSortByEnum Enum with underlying type: string
type MaintenanceWindowSortByEnum string

// Set of constants representing the allowable values for MaintenanceWindowSortByEnum
const (
	MaintenanceWindowSortByTimeCreated MaintenanceWindowSortByEnum = "timeCreated"
	MaintenanceWindowSortByDisplayName MaintenanceWindowSortByEnum = "displayName"
)

var mappingMaintenanceWindowSortByEnum = map[string]MaintenanceWindowSortByEnum{
	"timeCreated": MaintenanceWindowSortByTimeCreated,
	"displayName": MaintenanceWindowSortByDisplayName,
}

var mappingMaintenanceWindowSortByEnumLowerCase = map[string]MaintenanceWindowSortByEnum{
	"timecreated": MaintenanceWindowSortByTimeCreated,
	"displayname": MaintenanceWindowSortByDisplayName,
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
		"timeCreated",
		"displayName",
	}
}

// GetMappingMaintenanceWindowSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowSortByEnum(val string) (MaintenanceWindowSortByEnum, bool) {
	enum, ok := mappingMaintenanceWindowSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
