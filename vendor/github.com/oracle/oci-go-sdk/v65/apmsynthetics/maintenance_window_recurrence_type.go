// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"strings"
)

// MaintenanceWindowRecurrenceTypeEnum Enum with underlying type: string
type MaintenanceWindowRecurrenceTypeEnum string

// Set of constants representing the allowable values for MaintenanceWindowRecurrenceTypeEnum
const (
	MaintenanceWindowRecurrenceTypeWeekly  MaintenanceWindowRecurrenceTypeEnum = "WEEKLY"
	MaintenanceWindowRecurrenceTypeMonthly MaintenanceWindowRecurrenceTypeEnum = "MONTHLY"
)

var mappingMaintenanceWindowRecurrenceTypeEnum = map[string]MaintenanceWindowRecurrenceTypeEnum{
	"WEEKLY":  MaintenanceWindowRecurrenceTypeWeekly,
	"MONTHLY": MaintenanceWindowRecurrenceTypeMonthly,
}

var mappingMaintenanceWindowRecurrenceTypeEnumLowerCase = map[string]MaintenanceWindowRecurrenceTypeEnum{
	"weekly":  MaintenanceWindowRecurrenceTypeWeekly,
	"monthly": MaintenanceWindowRecurrenceTypeMonthly,
}

// GetMaintenanceWindowRecurrenceTypeEnumValues Enumerates the set of values for MaintenanceWindowRecurrenceTypeEnum
func GetMaintenanceWindowRecurrenceTypeEnumValues() []MaintenanceWindowRecurrenceTypeEnum {
	values := make([]MaintenanceWindowRecurrenceTypeEnum, 0)
	for _, v := range mappingMaintenanceWindowRecurrenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowRecurrenceTypeEnumStringValues Enumerates the set of values in String for MaintenanceWindowRecurrenceTypeEnum
func GetMaintenanceWindowRecurrenceTypeEnumStringValues() []string {
	return []string{
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingMaintenanceWindowRecurrenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowRecurrenceTypeEnum(val string) (MaintenanceWindowRecurrenceTypeEnum, bool) {
	enum, ok := mappingMaintenanceWindowRecurrenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
