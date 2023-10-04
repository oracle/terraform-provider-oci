// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// MaintenanceNotificationTypeEnum Enum with underlying type: string
type MaintenanceNotificationTypeEnum string

// Set of constants representing the allowable values for MaintenanceNotificationTypeEnum
const (
	MaintenanceNotificationTypeScheduledMaintenanceWithReleaseNotes    MaintenanceNotificationTypeEnum = "SCHEDULED_MAINTENANCE_WITH_RELEASE_NOTES"
	MaintenanceNotificationTypeScheduledMaintenanceWithoutReleaseNotes MaintenanceNotificationTypeEnum = "SCHEDULED_MAINTENANCE_WITHOUT_RELEASE_NOTES"
	MaintenanceNotificationTypeRescheduledMaintenance                  MaintenanceNotificationTypeEnum = "RESCHEDULED_MAINTENANCE"
	MaintenanceNotificationTypeScheduledMaintenanceReminder            MaintenanceNotificationTypeEnum = "SCHEDULED_MAINTENANCE_REMINDER"
	MaintenanceNotificationTypeScheduledMaintenanceCompleted           MaintenanceNotificationTypeEnum = "SCHEDULED_MAINTENANCE_COMPLETED"
)

var mappingMaintenanceNotificationTypeEnum = map[string]MaintenanceNotificationTypeEnum{
	"SCHEDULED_MAINTENANCE_WITH_RELEASE_NOTES":    MaintenanceNotificationTypeScheduledMaintenanceWithReleaseNotes,
	"SCHEDULED_MAINTENANCE_WITHOUT_RELEASE_NOTES": MaintenanceNotificationTypeScheduledMaintenanceWithoutReleaseNotes,
	"RESCHEDULED_MAINTENANCE":                     MaintenanceNotificationTypeRescheduledMaintenance,
	"SCHEDULED_MAINTENANCE_REMINDER":              MaintenanceNotificationTypeScheduledMaintenanceReminder,
	"SCHEDULED_MAINTENANCE_COMPLETED":             MaintenanceNotificationTypeScheduledMaintenanceCompleted,
}

var mappingMaintenanceNotificationTypeEnumLowerCase = map[string]MaintenanceNotificationTypeEnum{
	"scheduled_maintenance_with_release_notes":    MaintenanceNotificationTypeScheduledMaintenanceWithReleaseNotes,
	"scheduled_maintenance_without_release_notes": MaintenanceNotificationTypeScheduledMaintenanceWithoutReleaseNotes,
	"rescheduled_maintenance":                     MaintenanceNotificationTypeRescheduledMaintenance,
	"scheduled_maintenance_reminder":              MaintenanceNotificationTypeScheduledMaintenanceReminder,
	"scheduled_maintenance_completed":             MaintenanceNotificationTypeScheduledMaintenanceCompleted,
}

// GetMaintenanceNotificationTypeEnumValues Enumerates the set of values for MaintenanceNotificationTypeEnum
func GetMaintenanceNotificationTypeEnumValues() []MaintenanceNotificationTypeEnum {
	values := make([]MaintenanceNotificationTypeEnum, 0)
	for _, v := range mappingMaintenanceNotificationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceNotificationTypeEnumStringValues Enumerates the set of values in String for MaintenanceNotificationTypeEnum
func GetMaintenanceNotificationTypeEnumStringValues() []string {
	return []string{
		"SCHEDULED_MAINTENANCE_WITH_RELEASE_NOTES",
		"SCHEDULED_MAINTENANCE_WITHOUT_RELEASE_NOTES",
		"RESCHEDULED_MAINTENANCE",
		"SCHEDULED_MAINTENANCE_REMINDER",
		"SCHEDULED_MAINTENANCE_COMPLETED",
	}
}

// GetMappingMaintenanceNotificationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceNotificationTypeEnum(val string) (MaintenanceNotificationTypeEnum, bool) {
	enum, ok := mappingMaintenanceNotificationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
