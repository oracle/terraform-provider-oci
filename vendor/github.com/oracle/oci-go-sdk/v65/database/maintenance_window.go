// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceWindow The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window.
type MaintenanceWindow struct {

	// The maintenance window scheduling preference.
	Preference MaintenanceWindowPreferenceEnum `mandatory:"true" json:"preference"`

	// Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	// *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See Oracle-Managed Infrastructure Maintenance Updates (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information.
	PatchingMode MaintenanceWindowPatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	IsCustomActionTimeoutEnabled *bool `mandatory:"false" json:"isCustomActionTimeoutEnabled"`

	// Determines the amount of time the system will wait before the start of each database server patching operation.
	// Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive).
	CustomActionTimeoutInMins *int `mandatory:"false" json:"customActionTimeoutInMins"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceWindowPreferenceEnum(string(m.Preference)); !ok && m.Preference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Preference: %s. Supported values are: %s.", m.Preference, strings.Join(GetMaintenanceWindowPreferenceEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMaintenanceWindowPatchingModeEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetMaintenanceWindowPatchingModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceWindowPreferenceEnum Enum with underlying type: string
type MaintenanceWindowPreferenceEnum string

// Set of constants representing the allowable values for MaintenanceWindowPreferenceEnum
const (
	MaintenanceWindowPreferenceNoPreference     MaintenanceWindowPreferenceEnum = "NO_PREFERENCE"
	MaintenanceWindowPreferenceCustomPreference MaintenanceWindowPreferenceEnum = "CUSTOM_PREFERENCE"
)

var mappingMaintenanceWindowPreferenceEnum = map[string]MaintenanceWindowPreferenceEnum{
	"NO_PREFERENCE":     MaintenanceWindowPreferenceNoPreference,
	"CUSTOM_PREFERENCE": MaintenanceWindowPreferenceCustomPreference,
}

var mappingMaintenanceWindowPreferenceEnumLowerCase = map[string]MaintenanceWindowPreferenceEnum{
	"no_preference":     MaintenanceWindowPreferenceNoPreference,
	"custom_preference": MaintenanceWindowPreferenceCustomPreference,
}

// GetMaintenanceWindowPreferenceEnumValues Enumerates the set of values for MaintenanceWindowPreferenceEnum
func GetMaintenanceWindowPreferenceEnumValues() []MaintenanceWindowPreferenceEnum {
	values := make([]MaintenanceWindowPreferenceEnum, 0)
	for _, v := range mappingMaintenanceWindowPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowPreferenceEnumStringValues Enumerates the set of values in String for MaintenanceWindowPreferenceEnum
func GetMaintenanceWindowPreferenceEnumStringValues() []string {
	return []string{
		"NO_PREFERENCE",
		"CUSTOM_PREFERENCE",
	}
}

// GetMappingMaintenanceWindowPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowPreferenceEnum(val string) (MaintenanceWindowPreferenceEnum, bool) {
	enum, ok := mappingMaintenanceWindowPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceWindowPatchingModeEnum Enum with underlying type: string
type MaintenanceWindowPatchingModeEnum string

// Set of constants representing the allowable values for MaintenanceWindowPatchingModeEnum
const (
	MaintenanceWindowPatchingModeRolling    MaintenanceWindowPatchingModeEnum = "ROLLING"
	MaintenanceWindowPatchingModeNonrolling MaintenanceWindowPatchingModeEnum = "NONROLLING"
)

var mappingMaintenanceWindowPatchingModeEnum = map[string]MaintenanceWindowPatchingModeEnum{
	"ROLLING":    MaintenanceWindowPatchingModeRolling,
	"NONROLLING": MaintenanceWindowPatchingModeNonrolling,
}

var mappingMaintenanceWindowPatchingModeEnumLowerCase = map[string]MaintenanceWindowPatchingModeEnum{
	"rolling":    MaintenanceWindowPatchingModeRolling,
	"nonrolling": MaintenanceWindowPatchingModeNonrolling,
}

// GetMaintenanceWindowPatchingModeEnumValues Enumerates the set of values for MaintenanceWindowPatchingModeEnum
func GetMaintenanceWindowPatchingModeEnumValues() []MaintenanceWindowPatchingModeEnum {
	values := make([]MaintenanceWindowPatchingModeEnum, 0)
	for _, v := range mappingMaintenanceWindowPatchingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowPatchingModeEnumStringValues Enumerates the set of values in String for MaintenanceWindowPatchingModeEnum
func GetMaintenanceWindowPatchingModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingMaintenanceWindowPatchingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowPatchingModeEnum(val string) (MaintenanceWindowPatchingModeEnum, bool) {
	enum, ok := mappingMaintenanceWindowPatchingModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
