// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MaintenanceWindow The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window.
type MaintenanceWindow struct {

	// The maintenance window scheduling preference.
	Preference MaintenancePreferenceEnum `mandatory:"false" json:"preference,omitempty"`

	// Cloud Database Infrastructure node patching method.
	// *IMPORTANT*: Non-rolling infrastructure patching involves system down time.
	// See Oracle-Managed Database Infrastructure Maintenance Updates (https://docs.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle)
	// for more information.
	PatchingMode BasePatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`

	// If true, enables the configuration of a custom action timeout (waiting period) between Database Infrastructure server patching operations.
	IsCustomActionTimeoutEnabled *bool `mandatory:"false" json:"isCustomActionTimeoutEnabled"`

	// Determines the amount of time the system will wait before the start of each Database Infrastructure server patching operation.
	// Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive).
	CustomActionTimeoutInMins *int `mandatory:"false" json:"customActionTimeoutInMins"`

	// If true, enables the monthly patching option.
	IsMonthlyPatchingEnabled *bool `mandatory:"false" json:"isMonthlyPatchingEnabled"`

	// Months during the year when maintenance should be performed.
	Months []MaintenanceWindowMonthsEnum `mandatory:"false" json:"months,omitempty"`

	// Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and
	// 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates,
	// not days of the week.For example, to allow maintenance during the 2nd week of the month
	// (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for
	// the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction
	// with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and
	// hours that maintenance will be performed.
	WeeksOfMonth []int `mandatory:"false" json:"weeksOfMonth"`

	// Days during the week when maintenance should be performed.
	DaysOfWeek []MaintenanceWindowDaysOfWeekEnum `mandatory:"false" json:"daysOfWeek,omitempty"`

	// The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are - 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	HoursOfDay []int `mandatory:"false" json:"hoursOfDay"`

	// Lead time window allows user to set a lead time to prepare for a down time.
	// The lead time is in weeks and valid value is between 1 to 4.
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

	if _, ok := GetMappingMaintenancePreferenceEnum(string(m.Preference)); !ok && m.Preference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Preference: %s. Supported values are: %s.", m.Preference, strings.Join(GetMaintenancePreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBasePatchingModeEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetBasePatchingModeEnumStringValues(), ",")))
	}
	for _, val := range m.Months {
		if _, ok := GetMappingMaintenanceWindowMonthsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Months: %s. Supported values are: %s.", val, strings.Join(GetMaintenanceWindowMonthsEnumStringValues(), ",")))
		}
	}

	for _, val := range m.DaysOfWeek {
		if _, ok := GetMappingMaintenanceWindowDaysOfWeekEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DaysOfWeek: %s. Supported values are: %s.", val, strings.Join(GetMaintenanceWindowDaysOfWeekEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceWindowMonthsEnum Enum with underlying type: string
type MaintenanceWindowMonthsEnum string

// Set of constants representing the allowable values for MaintenanceWindowMonthsEnum
const (
	MaintenanceWindowMonthsJanuary   MaintenanceWindowMonthsEnum = "JANUARY"
	MaintenanceWindowMonthsFebruary  MaintenanceWindowMonthsEnum = "FEBRUARY"
	MaintenanceWindowMonthsMarch     MaintenanceWindowMonthsEnum = "MARCH"
	MaintenanceWindowMonthsApril     MaintenanceWindowMonthsEnum = "APRIL"
	MaintenanceWindowMonthsMay       MaintenanceWindowMonthsEnum = "MAY"
	MaintenanceWindowMonthsJune      MaintenanceWindowMonthsEnum = "JUNE"
	MaintenanceWindowMonthsJuly      MaintenanceWindowMonthsEnum = "JULY"
	MaintenanceWindowMonthsAugust    MaintenanceWindowMonthsEnum = "AUGUST"
	MaintenanceWindowMonthsSeptember MaintenanceWindowMonthsEnum = "SEPTEMBER"
	MaintenanceWindowMonthsOctober   MaintenanceWindowMonthsEnum = "OCTOBER"
	MaintenanceWindowMonthsNovember  MaintenanceWindowMonthsEnum = "NOVEMBER"
	MaintenanceWindowMonthsDecember  MaintenanceWindowMonthsEnum = "DECEMBER"
)

var mappingMaintenanceWindowMonthsEnum = map[string]MaintenanceWindowMonthsEnum{
	"JANUARY":   MaintenanceWindowMonthsJanuary,
	"FEBRUARY":  MaintenanceWindowMonthsFebruary,
	"MARCH":     MaintenanceWindowMonthsMarch,
	"APRIL":     MaintenanceWindowMonthsApril,
	"MAY":       MaintenanceWindowMonthsMay,
	"JUNE":      MaintenanceWindowMonthsJune,
	"JULY":      MaintenanceWindowMonthsJuly,
	"AUGUST":    MaintenanceWindowMonthsAugust,
	"SEPTEMBER": MaintenanceWindowMonthsSeptember,
	"OCTOBER":   MaintenanceWindowMonthsOctober,
	"NOVEMBER":  MaintenanceWindowMonthsNovember,
	"DECEMBER":  MaintenanceWindowMonthsDecember,
}

var mappingMaintenanceWindowMonthsEnumLowerCase = map[string]MaintenanceWindowMonthsEnum{
	"january":   MaintenanceWindowMonthsJanuary,
	"february":  MaintenanceWindowMonthsFebruary,
	"march":     MaintenanceWindowMonthsMarch,
	"april":     MaintenanceWindowMonthsApril,
	"may":       MaintenanceWindowMonthsMay,
	"june":      MaintenanceWindowMonthsJune,
	"july":      MaintenanceWindowMonthsJuly,
	"august":    MaintenanceWindowMonthsAugust,
	"september": MaintenanceWindowMonthsSeptember,
	"october":   MaintenanceWindowMonthsOctober,
	"november":  MaintenanceWindowMonthsNovember,
	"december":  MaintenanceWindowMonthsDecember,
}

// GetMaintenanceWindowMonthsEnumValues Enumerates the set of values for MaintenanceWindowMonthsEnum
func GetMaintenanceWindowMonthsEnumValues() []MaintenanceWindowMonthsEnum {
	values := make([]MaintenanceWindowMonthsEnum, 0)
	for _, v := range mappingMaintenanceWindowMonthsEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowMonthsEnumStringValues Enumerates the set of values in String for MaintenanceWindowMonthsEnum
func GetMaintenanceWindowMonthsEnumStringValues() []string {
	return []string{
		"JANUARY",
		"FEBRUARY",
		"MARCH",
		"APRIL",
		"MAY",
		"JUNE",
		"JULY",
		"AUGUST",
		"SEPTEMBER",
		"OCTOBER",
		"NOVEMBER",
		"DECEMBER",
	}
}

// GetMappingMaintenanceWindowMonthsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowMonthsEnum(val string) (MaintenanceWindowMonthsEnum, bool) {
	enum, ok := mappingMaintenanceWindowMonthsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MaintenanceWindowDaysOfWeekEnum Enum with underlying type: string
type MaintenanceWindowDaysOfWeekEnum string

// Set of constants representing the allowable values for MaintenanceWindowDaysOfWeekEnum
const (
	MaintenanceWindowDaysOfWeekMonday    MaintenanceWindowDaysOfWeekEnum = "MONDAY"
	MaintenanceWindowDaysOfWeekTuesday   MaintenanceWindowDaysOfWeekEnum = "TUESDAY"
	MaintenanceWindowDaysOfWeekWednesday MaintenanceWindowDaysOfWeekEnum = "WEDNESDAY"
	MaintenanceWindowDaysOfWeekThursday  MaintenanceWindowDaysOfWeekEnum = "THURSDAY"
	MaintenanceWindowDaysOfWeekFriday    MaintenanceWindowDaysOfWeekEnum = "FRIDAY"
	MaintenanceWindowDaysOfWeekSaturday  MaintenanceWindowDaysOfWeekEnum = "SATURDAY"
	MaintenanceWindowDaysOfWeekSunday    MaintenanceWindowDaysOfWeekEnum = "SUNDAY"
)

var mappingMaintenanceWindowDaysOfWeekEnum = map[string]MaintenanceWindowDaysOfWeekEnum{
	"MONDAY":    MaintenanceWindowDaysOfWeekMonday,
	"TUESDAY":   MaintenanceWindowDaysOfWeekTuesday,
	"WEDNESDAY": MaintenanceWindowDaysOfWeekWednesday,
	"THURSDAY":  MaintenanceWindowDaysOfWeekThursday,
	"FRIDAY":    MaintenanceWindowDaysOfWeekFriday,
	"SATURDAY":  MaintenanceWindowDaysOfWeekSaturday,
	"SUNDAY":    MaintenanceWindowDaysOfWeekSunday,
}

var mappingMaintenanceWindowDaysOfWeekEnumLowerCase = map[string]MaintenanceWindowDaysOfWeekEnum{
	"monday":    MaintenanceWindowDaysOfWeekMonday,
	"tuesday":   MaintenanceWindowDaysOfWeekTuesday,
	"wednesday": MaintenanceWindowDaysOfWeekWednesday,
	"thursday":  MaintenanceWindowDaysOfWeekThursday,
	"friday":    MaintenanceWindowDaysOfWeekFriday,
	"saturday":  MaintenanceWindowDaysOfWeekSaturday,
	"sunday":    MaintenanceWindowDaysOfWeekSunday,
}

// GetMaintenanceWindowDaysOfWeekEnumValues Enumerates the set of values for MaintenanceWindowDaysOfWeekEnum
func GetMaintenanceWindowDaysOfWeekEnumValues() []MaintenanceWindowDaysOfWeekEnum {
	values := make([]MaintenanceWindowDaysOfWeekEnum, 0)
	for _, v := range mappingMaintenanceWindowDaysOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowDaysOfWeekEnumStringValues Enumerates the set of values in String for MaintenanceWindowDaysOfWeekEnum
func GetMaintenanceWindowDaysOfWeekEnumStringValues() []string {
	return []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
}

// GetMappingMaintenanceWindowDaysOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowDaysOfWeekEnum(val string) (MaintenanceWindowDaysOfWeekEnum, bool) {
	enum, ok := mappingMaintenanceWindowDaysOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
