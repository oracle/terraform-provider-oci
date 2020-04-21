// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MaintenanceWindow The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window.
type MaintenanceWindow struct {

	// The maintenance window scheduling preference.
	Preference MaintenanceWindowPreferenceEnum `mandatory:"true" json:"preference"`

	// Months during the year when maintenance should be performed.
	Months []Month `mandatory:"false" json:"months"`

	// Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week.
	// For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days.
	// Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed.
	WeeksOfMonth []int `mandatory:"false" json:"weeksOfMonth"`

	// Days during the week when maintenance should be performed.
	DaysOfWeek []DayOfWeek `mandatory:"false" json:"daysOfWeek"`

	// The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
	// - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	HoursOfDay []int `mandatory:"false" json:"hoursOfDay"`

	// Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4.
	LeadTimeInWeeks *int `mandatory:"false" json:"leadTimeInWeeks"`
}

func (m MaintenanceWindow) String() string {
	return common.PointerString(m)
}

// MaintenanceWindowPreferenceEnum Enum with underlying type: string
type MaintenanceWindowPreferenceEnum string

// Set of constants representing the allowable values for MaintenanceWindowPreferenceEnum
const (
	MaintenanceWindowPreferenceNoPreference     MaintenanceWindowPreferenceEnum = "NO_PREFERENCE"
	MaintenanceWindowPreferenceCustomPreference MaintenanceWindowPreferenceEnum = "CUSTOM_PREFERENCE"
)

var mappingMaintenanceWindowPreference = map[string]MaintenanceWindowPreferenceEnum{
	"NO_PREFERENCE":     MaintenanceWindowPreferenceNoPreference,
	"CUSTOM_PREFERENCE": MaintenanceWindowPreferenceCustomPreference,
}

// GetMaintenanceWindowPreferenceEnumValues Enumerates the set of values for MaintenanceWindowPreferenceEnum
func GetMaintenanceWindowPreferenceEnumValues() []MaintenanceWindowPreferenceEnum {
	values := make([]MaintenanceWindowPreferenceEnum, 0)
	for _, v := range mappingMaintenanceWindowPreference {
		values = append(values, v)
	}
	return values
}
