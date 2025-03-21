// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MaintenanceWindow The preferred day and time to perform maintenance.
type MaintenanceWindow struct {

	// Day of the week when the maintainence window starts.
	DayOfWeek MaintenanceWindowDayOfWeekEnum `mandatory:"false" json:"dayOfWeek,omitempty"`

	// The time to start the maintenance window. The format is 'HH:MM', 'HH:MM' represents the time in UTC.
	// Example: `22:00`
	TimeStart *string `mandatory:"false" json:"timeStart"`
}

func (m MaintenanceWindow) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaintenanceWindow) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMaintenanceWindowDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetMaintenanceWindowDayOfWeekEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaintenanceWindowDayOfWeekEnum Enum with underlying type: string
type MaintenanceWindowDayOfWeekEnum string

// Set of constants representing the allowable values for MaintenanceWindowDayOfWeekEnum
const (
	MaintenanceWindowDayOfWeekMonday    MaintenanceWindowDayOfWeekEnum = "MONDAY"
	MaintenanceWindowDayOfWeekTuesday   MaintenanceWindowDayOfWeekEnum = "TUESDAY"
	MaintenanceWindowDayOfWeekWednesday MaintenanceWindowDayOfWeekEnum = "WEDNESDAY"
	MaintenanceWindowDayOfWeekThursday  MaintenanceWindowDayOfWeekEnum = "THURSDAY"
	MaintenanceWindowDayOfWeekFriday    MaintenanceWindowDayOfWeekEnum = "FRIDAY"
	MaintenanceWindowDayOfWeekSaturday  MaintenanceWindowDayOfWeekEnum = "SATURDAY"
	MaintenanceWindowDayOfWeekSunday    MaintenanceWindowDayOfWeekEnum = "SUNDAY"
)

var mappingMaintenanceWindowDayOfWeekEnum = map[string]MaintenanceWindowDayOfWeekEnum{
	"MONDAY":    MaintenanceWindowDayOfWeekMonday,
	"TUESDAY":   MaintenanceWindowDayOfWeekTuesday,
	"WEDNESDAY": MaintenanceWindowDayOfWeekWednesday,
	"THURSDAY":  MaintenanceWindowDayOfWeekThursday,
	"FRIDAY":    MaintenanceWindowDayOfWeekFriday,
	"SATURDAY":  MaintenanceWindowDayOfWeekSaturday,
	"SUNDAY":    MaintenanceWindowDayOfWeekSunday,
}

var mappingMaintenanceWindowDayOfWeekEnumLowerCase = map[string]MaintenanceWindowDayOfWeekEnum{
	"monday":    MaintenanceWindowDayOfWeekMonday,
	"tuesday":   MaintenanceWindowDayOfWeekTuesday,
	"wednesday": MaintenanceWindowDayOfWeekWednesday,
	"thursday":  MaintenanceWindowDayOfWeekThursday,
	"friday":    MaintenanceWindowDayOfWeekFriday,
	"saturday":  MaintenanceWindowDayOfWeekSaturday,
	"sunday":    MaintenanceWindowDayOfWeekSunday,
}

// GetMaintenanceWindowDayOfWeekEnumValues Enumerates the set of values for MaintenanceWindowDayOfWeekEnum
func GetMaintenanceWindowDayOfWeekEnumValues() []MaintenanceWindowDayOfWeekEnum {
	values := make([]MaintenanceWindowDayOfWeekEnum, 0)
	for _, v := range mappingMaintenanceWindowDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowDayOfWeekEnumStringValues Enumerates the set of values in String for MaintenanceWindowDayOfWeekEnum
func GetMaintenanceWindowDayOfWeekEnumStringValues() []string {
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

// GetMappingMaintenanceWindowDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowDayOfWeekEnum(val string) (MaintenanceWindowDayOfWeekEnum, bool) {
	enum, ok := mappingMaintenanceWindowDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
