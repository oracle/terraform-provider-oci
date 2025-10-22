// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// MaintenanceScheduleTypeEnum Enum with underlying type: string
type MaintenanceScheduleTypeEnum string

// Set of constants representing the allowable values for MaintenanceScheduleTypeEnum
const (
	MaintenanceScheduleTypeEarly   MaintenanceScheduleTypeEnum = "EARLY"
	MaintenanceScheduleTypeRegular MaintenanceScheduleTypeEnum = "REGULAR"
)

var mappingMaintenanceScheduleTypeEnum = map[string]MaintenanceScheduleTypeEnum{
	"EARLY":   MaintenanceScheduleTypeEarly,
	"REGULAR": MaintenanceScheduleTypeRegular,
}

var mappingMaintenanceScheduleTypeEnumLowerCase = map[string]MaintenanceScheduleTypeEnum{
	"early":   MaintenanceScheduleTypeEarly,
	"regular": MaintenanceScheduleTypeRegular,
}

// GetMaintenanceScheduleTypeEnumValues Enumerates the set of values for MaintenanceScheduleTypeEnum
func GetMaintenanceScheduleTypeEnumValues() []MaintenanceScheduleTypeEnum {
	values := make([]MaintenanceScheduleTypeEnum, 0)
	for _, v := range mappingMaintenanceScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceScheduleTypeEnumStringValues Enumerates the set of values in String for MaintenanceScheduleTypeEnum
func GetMaintenanceScheduleTypeEnumStringValues() []string {
	return []string{
		"EARLY",
		"REGULAR",
	}
}

// GetMappingMaintenanceScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceScheduleTypeEnum(val string) (MaintenanceScheduleTypeEnum, bool) {
	enum, ok := mappingMaintenanceScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
