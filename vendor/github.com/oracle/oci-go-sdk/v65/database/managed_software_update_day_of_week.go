// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedSoftwareUpdateDayOfWeek Day of the week.
type ManagedSoftwareUpdateDayOfWeek struct {

	// Day of the week.
	DayOfWeek ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum `mandatory:"true" json:"dayOfWeek"`
}

func (m ManagedSoftwareUpdateDayOfWeek) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedSoftwareUpdateDayOfWeek) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetManagedSoftwareUpdateDayOfWeekDayOfWeekEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum Enum with underlying type: string
type ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum string

// Set of constants representing the allowable values for ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum
const (
	ManagedSoftwareUpdateDayOfWeekDayOfWeekSunday    ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "SUNDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekMonday    ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "MONDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekTuesday   ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "TUESDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekWednesday ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "WEDNESDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekThursday  ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "THURSDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekFriday    ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "FRIDAY"
	ManagedSoftwareUpdateDayOfWeekDayOfWeekSaturday  ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = "SATURDAY"
)

var mappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnum = map[string]ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum{
	"SUNDAY":    ManagedSoftwareUpdateDayOfWeekDayOfWeekSunday,
	"MONDAY":    ManagedSoftwareUpdateDayOfWeekDayOfWeekMonday,
	"TUESDAY":   ManagedSoftwareUpdateDayOfWeekDayOfWeekTuesday,
	"WEDNESDAY": ManagedSoftwareUpdateDayOfWeekDayOfWeekWednesday,
	"THURSDAY":  ManagedSoftwareUpdateDayOfWeekDayOfWeekThursday,
	"FRIDAY":    ManagedSoftwareUpdateDayOfWeekDayOfWeekFriday,
	"SATURDAY":  ManagedSoftwareUpdateDayOfWeekDayOfWeekSaturday,
}

var mappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnumLowerCase = map[string]ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum{
	"sunday":    ManagedSoftwareUpdateDayOfWeekDayOfWeekSunday,
	"monday":    ManagedSoftwareUpdateDayOfWeekDayOfWeekMonday,
	"tuesday":   ManagedSoftwareUpdateDayOfWeekDayOfWeekTuesday,
	"wednesday": ManagedSoftwareUpdateDayOfWeekDayOfWeekWednesday,
	"thursday":  ManagedSoftwareUpdateDayOfWeekDayOfWeekThursday,
	"friday":    ManagedSoftwareUpdateDayOfWeekDayOfWeekFriday,
	"saturday":  ManagedSoftwareUpdateDayOfWeekDayOfWeekSaturday,
}

// GetManagedSoftwareUpdateDayOfWeekDayOfWeekEnumValues Enumerates the set of values for ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum
func GetManagedSoftwareUpdateDayOfWeekDayOfWeekEnumValues() []ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum {
	values := make([]ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum, 0)
	for _, v := range mappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedSoftwareUpdateDayOfWeekDayOfWeekEnumStringValues Enumerates the set of values in String for ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum
func GetManagedSoftwareUpdateDayOfWeekDayOfWeekEnumStringValues() []string {
	return []string{
		"SUNDAY",
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
	}
}

// GetMappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnum(val string) (ManagedSoftwareUpdateDayOfWeekDayOfWeekEnum, bool) {
	enum, ok := mappingManagedSoftwareUpdateDayOfWeekDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
