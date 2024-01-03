// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SnapshotSchedule The snapshot schedule is a structure within a parent file system snapshot policy. It contains data about
// the frequency of snapshot creation and the retention time of the taken snapshots.
type SnapshotSchedule struct {

	// The frequency of scheduled snapshots.
	Period SnapshotSchedulePeriodEnum `mandatory:"true" json:"period"`

	// Time zone used for scheduling the snapshot.
	TimeZone SnapshotScheduleTimeZoneEnum `mandatory:"true" json:"timeZone"`

	// A name prefix to be applied to snapshots created by this schedule.
	// Example: `compliance1`
	SchedulePrefix *string `mandatory:"false" json:"schedulePrefix"`

	// The starting point used to begin the scheduling of the snapshots based upon recurrence string
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// If no `timeScheduleStart` is provided, the value will be set to the time when the schedule was created.
	TimeScheduleStart *common.SDKTime `mandatory:"false" json:"timeScheduleStart"`

	// The number of seconds to retain snapshots created with this schedule.
	// Snapshot expiration time will not be set if this value is empty.
	RetentionDurationInSeconds *int64 `mandatory:"false" json:"retentionDurationInSeconds"`

	// The hour of the day to create a DAILY, WEEKLY, MONTHLY, or YEARLY snapshot.
	// If not set, the system chooses a value at creation time.
	HourOfDay *int `mandatory:"false" json:"hourOfDay"`

	// The day of the week to create a scheduled snapshot.
	// Used for WEEKLY snapshot schedules.
	// If not set, the system chooses a value at creation time.
	DayOfWeek SnapshotScheduleDayOfWeekEnum `mandatory:"false" json:"dayOfWeek,omitempty"`

	// The day of the month to create a scheduled snapshot.
	// If the day does not exist for the month, snapshot creation will be skipped.
	// Used for MONTHLY and YEARLY snapshot schedules.
	// If not set, the system chooses a value at creation time.
	DayOfMonth *int `mandatory:"false" json:"dayOfMonth"`

	// The month to create a scheduled snapshot.
	// Used only for YEARLY snapshot schedules.
	// If not set, the system chooses a value at creation time.
	Month SnapshotScheduleMonthEnum `mandatory:"false" json:"month,omitempty"`
}

func (m SnapshotSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapshotSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSnapshotSchedulePeriodEnum(string(m.Period)); !ok && m.Period != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Period: %s. Supported values are: %s.", m.Period, strings.Join(GetSnapshotSchedulePeriodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSnapshotScheduleTimeZoneEnum(string(m.TimeZone)); !ok && m.TimeZone != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TimeZone: %s. Supported values are: %s.", m.TimeZone, strings.Join(GetSnapshotScheduleTimeZoneEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSnapshotScheduleDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetSnapshotScheduleDayOfWeekEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSnapshotScheduleMonthEnum(string(m.Month)); !ok && m.Month != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Month: %s. Supported values are: %s.", m.Month, strings.Join(GetSnapshotScheduleMonthEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SnapshotSchedulePeriodEnum Enum with underlying type: string
type SnapshotSchedulePeriodEnum string

// Set of constants representing the allowable values for SnapshotSchedulePeriodEnum
const (
	SnapshotSchedulePeriodHourly  SnapshotSchedulePeriodEnum = "HOURLY"
	SnapshotSchedulePeriodDaily   SnapshotSchedulePeriodEnum = "DAILY"
	SnapshotSchedulePeriodWeekly  SnapshotSchedulePeriodEnum = "WEEKLY"
	SnapshotSchedulePeriodMonthly SnapshotSchedulePeriodEnum = "MONTHLY"
	SnapshotSchedulePeriodYearly  SnapshotSchedulePeriodEnum = "YEARLY"
)

var mappingSnapshotSchedulePeriodEnum = map[string]SnapshotSchedulePeriodEnum{
	"HOURLY":  SnapshotSchedulePeriodHourly,
	"DAILY":   SnapshotSchedulePeriodDaily,
	"WEEKLY":  SnapshotSchedulePeriodWeekly,
	"MONTHLY": SnapshotSchedulePeriodMonthly,
	"YEARLY":  SnapshotSchedulePeriodYearly,
}

var mappingSnapshotSchedulePeriodEnumLowerCase = map[string]SnapshotSchedulePeriodEnum{
	"hourly":  SnapshotSchedulePeriodHourly,
	"daily":   SnapshotSchedulePeriodDaily,
	"weekly":  SnapshotSchedulePeriodWeekly,
	"monthly": SnapshotSchedulePeriodMonthly,
	"yearly":  SnapshotSchedulePeriodYearly,
}

// GetSnapshotSchedulePeriodEnumValues Enumerates the set of values for SnapshotSchedulePeriodEnum
func GetSnapshotSchedulePeriodEnumValues() []SnapshotSchedulePeriodEnum {
	values := make([]SnapshotSchedulePeriodEnum, 0)
	for _, v := range mappingSnapshotSchedulePeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotSchedulePeriodEnumStringValues Enumerates the set of values in String for SnapshotSchedulePeriodEnum
func GetSnapshotSchedulePeriodEnumStringValues() []string {
	return []string{
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
}

// GetMappingSnapshotSchedulePeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotSchedulePeriodEnum(val string) (SnapshotSchedulePeriodEnum, bool) {
	enum, ok := mappingSnapshotSchedulePeriodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnapshotScheduleTimeZoneEnum Enum with underlying type: string
type SnapshotScheduleTimeZoneEnum string

// Set of constants representing the allowable values for SnapshotScheduleTimeZoneEnum
const (
	SnapshotScheduleTimeZoneUtc                    SnapshotScheduleTimeZoneEnum = "UTC"
	SnapshotScheduleTimeZoneRegionalDataCenterTime SnapshotScheduleTimeZoneEnum = "REGIONAL_DATA_CENTER_TIME"
)

var mappingSnapshotScheduleTimeZoneEnum = map[string]SnapshotScheduleTimeZoneEnum{
	"UTC":                       SnapshotScheduleTimeZoneUtc,
	"REGIONAL_DATA_CENTER_TIME": SnapshotScheduleTimeZoneRegionalDataCenterTime,
}

var mappingSnapshotScheduleTimeZoneEnumLowerCase = map[string]SnapshotScheduleTimeZoneEnum{
	"utc":                       SnapshotScheduleTimeZoneUtc,
	"regional_data_center_time": SnapshotScheduleTimeZoneRegionalDataCenterTime,
}

// GetSnapshotScheduleTimeZoneEnumValues Enumerates the set of values for SnapshotScheduleTimeZoneEnum
func GetSnapshotScheduleTimeZoneEnumValues() []SnapshotScheduleTimeZoneEnum {
	values := make([]SnapshotScheduleTimeZoneEnum, 0)
	for _, v := range mappingSnapshotScheduleTimeZoneEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotScheduleTimeZoneEnumStringValues Enumerates the set of values in String for SnapshotScheduleTimeZoneEnum
func GetSnapshotScheduleTimeZoneEnumStringValues() []string {
	return []string{
		"UTC",
		"REGIONAL_DATA_CENTER_TIME",
	}
}

// GetMappingSnapshotScheduleTimeZoneEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotScheduleTimeZoneEnum(val string) (SnapshotScheduleTimeZoneEnum, bool) {
	enum, ok := mappingSnapshotScheduleTimeZoneEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnapshotScheduleDayOfWeekEnum Enum with underlying type: string
type SnapshotScheduleDayOfWeekEnum string

// Set of constants representing the allowable values for SnapshotScheduleDayOfWeekEnum
const (
	SnapshotScheduleDayOfWeekMonday    SnapshotScheduleDayOfWeekEnum = "MONDAY"
	SnapshotScheduleDayOfWeekTuesday   SnapshotScheduleDayOfWeekEnum = "TUESDAY"
	SnapshotScheduleDayOfWeekWednesday SnapshotScheduleDayOfWeekEnum = "WEDNESDAY"
	SnapshotScheduleDayOfWeekThursday  SnapshotScheduleDayOfWeekEnum = "THURSDAY"
	SnapshotScheduleDayOfWeekFriday    SnapshotScheduleDayOfWeekEnum = "FRIDAY"
	SnapshotScheduleDayOfWeekSaturday  SnapshotScheduleDayOfWeekEnum = "SATURDAY"
	SnapshotScheduleDayOfWeekSunday    SnapshotScheduleDayOfWeekEnum = "SUNDAY"
)

var mappingSnapshotScheduleDayOfWeekEnum = map[string]SnapshotScheduleDayOfWeekEnum{
	"MONDAY":    SnapshotScheduleDayOfWeekMonday,
	"TUESDAY":   SnapshotScheduleDayOfWeekTuesday,
	"WEDNESDAY": SnapshotScheduleDayOfWeekWednesday,
	"THURSDAY":  SnapshotScheduleDayOfWeekThursday,
	"FRIDAY":    SnapshotScheduleDayOfWeekFriday,
	"SATURDAY":  SnapshotScheduleDayOfWeekSaturday,
	"SUNDAY":    SnapshotScheduleDayOfWeekSunday,
}

var mappingSnapshotScheduleDayOfWeekEnumLowerCase = map[string]SnapshotScheduleDayOfWeekEnum{
	"monday":    SnapshotScheduleDayOfWeekMonday,
	"tuesday":   SnapshotScheduleDayOfWeekTuesday,
	"wednesday": SnapshotScheduleDayOfWeekWednesday,
	"thursday":  SnapshotScheduleDayOfWeekThursday,
	"friday":    SnapshotScheduleDayOfWeekFriday,
	"saturday":  SnapshotScheduleDayOfWeekSaturday,
	"sunday":    SnapshotScheduleDayOfWeekSunday,
}

// GetSnapshotScheduleDayOfWeekEnumValues Enumerates the set of values for SnapshotScheduleDayOfWeekEnum
func GetSnapshotScheduleDayOfWeekEnumValues() []SnapshotScheduleDayOfWeekEnum {
	values := make([]SnapshotScheduleDayOfWeekEnum, 0)
	for _, v := range mappingSnapshotScheduleDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotScheduleDayOfWeekEnumStringValues Enumerates the set of values in String for SnapshotScheduleDayOfWeekEnum
func GetSnapshotScheduleDayOfWeekEnumStringValues() []string {
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

// GetMappingSnapshotScheduleDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotScheduleDayOfWeekEnum(val string) (SnapshotScheduleDayOfWeekEnum, bool) {
	enum, ok := mappingSnapshotScheduleDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnapshotScheduleMonthEnum Enum with underlying type: string
type SnapshotScheduleMonthEnum string

// Set of constants representing the allowable values for SnapshotScheduleMonthEnum
const (
	SnapshotScheduleMonthJanuary   SnapshotScheduleMonthEnum = "JANUARY"
	SnapshotScheduleMonthFebruary  SnapshotScheduleMonthEnum = "FEBRUARY"
	SnapshotScheduleMonthMarch     SnapshotScheduleMonthEnum = "MARCH"
	SnapshotScheduleMonthApril     SnapshotScheduleMonthEnum = "APRIL"
	SnapshotScheduleMonthMay       SnapshotScheduleMonthEnum = "MAY"
	SnapshotScheduleMonthJune      SnapshotScheduleMonthEnum = "JUNE"
	SnapshotScheduleMonthJuly      SnapshotScheduleMonthEnum = "JULY"
	SnapshotScheduleMonthAugust    SnapshotScheduleMonthEnum = "AUGUST"
	SnapshotScheduleMonthSeptember SnapshotScheduleMonthEnum = "SEPTEMBER"
	SnapshotScheduleMonthOctober   SnapshotScheduleMonthEnum = "OCTOBER"
	SnapshotScheduleMonthNovember  SnapshotScheduleMonthEnum = "NOVEMBER"
	SnapshotScheduleMonthDecember  SnapshotScheduleMonthEnum = "DECEMBER"
)

var mappingSnapshotScheduleMonthEnum = map[string]SnapshotScheduleMonthEnum{
	"JANUARY":   SnapshotScheduleMonthJanuary,
	"FEBRUARY":  SnapshotScheduleMonthFebruary,
	"MARCH":     SnapshotScheduleMonthMarch,
	"APRIL":     SnapshotScheduleMonthApril,
	"MAY":       SnapshotScheduleMonthMay,
	"JUNE":      SnapshotScheduleMonthJune,
	"JULY":      SnapshotScheduleMonthJuly,
	"AUGUST":    SnapshotScheduleMonthAugust,
	"SEPTEMBER": SnapshotScheduleMonthSeptember,
	"OCTOBER":   SnapshotScheduleMonthOctober,
	"NOVEMBER":  SnapshotScheduleMonthNovember,
	"DECEMBER":  SnapshotScheduleMonthDecember,
}

var mappingSnapshotScheduleMonthEnumLowerCase = map[string]SnapshotScheduleMonthEnum{
	"january":   SnapshotScheduleMonthJanuary,
	"february":  SnapshotScheduleMonthFebruary,
	"march":     SnapshotScheduleMonthMarch,
	"april":     SnapshotScheduleMonthApril,
	"may":       SnapshotScheduleMonthMay,
	"june":      SnapshotScheduleMonthJune,
	"july":      SnapshotScheduleMonthJuly,
	"august":    SnapshotScheduleMonthAugust,
	"september": SnapshotScheduleMonthSeptember,
	"october":   SnapshotScheduleMonthOctober,
	"november":  SnapshotScheduleMonthNovember,
	"december":  SnapshotScheduleMonthDecember,
}

// GetSnapshotScheduleMonthEnumValues Enumerates the set of values for SnapshotScheduleMonthEnum
func GetSnapshotScheduleMonthEnumValues() []SnapshotScheduleMonthEnum {
	values := make([]SnapshotScheduleMonthEnum, 0)
	for _, v := range mappingSnapshotScheduleMonthEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotScheduleMonthEnumStringValues Enumerates the set of values in String for SnapshotScheduleMonthEnum
func GetSnapshotScheduleMonthEnumStringValues() []string {
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

// GetMappingSnapshotScheduleMonthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotScheduleMonthEnum(val string) (SnapshotScheduleMonthEnum, bool) {
	enum, ok := mappingSnapshotScheduleMonthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
