// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AvailableMaintenanceScheduleStartTimeSummary Information about the list of available start times on a particular day of the week.
// User can choose their preferred day of the week and start time for creating request/input
// for Create or Update LustreFileSystem operation.
type AvailableMaintenanceScheduleStartTimeSummary struct {

	// Day of the week
	DayOfWeek AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum `mandatory:"true" json:"dayOfWeek"`

	// List of available start times. Each array item is of the format `HH:mm`
	StartTimes []string `mandatory:"true" json:"startTimes"`
}

func (m AvailableMaintenanceScheduleStartTimeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailableMaintenanceScheduleStartTimeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum Enum with underlying type: string
type AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum string

// Set of constants representing the allowable values for AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum
const (
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekMonday    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "MONDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekTuesday   AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "TUESDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekWednesday AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "WEDNESDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekThursday  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "THURSDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekFriday    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "FRIDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSaturday  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "SATURDAY"
	AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSunday    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = "SUNDAY"
)

var mappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum = map[string]AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum{
	"MONDAY":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekMonday,
	"TUESDAY":   AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekTuesday,
	"WEDNESDAY": AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekWednesday,
	"THURSDAY":  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekThursday,
	"FRIDAY":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekFriday,
	"SATURDAY":  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSaturday,
	"SUNDAY":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSunday,
}

var mappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumLowerCase = map[string]AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum{
	"monday":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekMonday,
	"tuesday":   AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekTuesday,
	"wednesday": AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekWednesday,
	"thursday":  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekThursday,
	"friday":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekFriday,
	"saturday":  AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSaturday,
	"sunday":    AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekSunday,
}

// GetAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumValues Enumerates the set of values for AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum
func GetAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumValues() []AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum {
	values := make([]AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum, 0)
	for _, v := range mappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumStringValues Enumerates the set of values in String for AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum
func GetAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumStringValues() []string {
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

// GetMappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum(val string) (AvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnum, bool) {
	enum, ok := mappingAvailableMaintenanceScheduleStartTimeSummaryDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
