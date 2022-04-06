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

// DayOfWeek Day of the week.
type DayOfWeek struct {

	// Name of the day of the week.
	Name DayOfWeekNameEnum `mandatory:"true" json:"name"`
}

func (m DayOfWeek) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DayOfWeek) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDayOfWeekNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetDayOfWeekNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DayOfWeekNameEnum Enum with underlying type: string
type DayOfWeekNameEnum string

// Set of constants representing the allowable values for DayOfWeekNameEnum
const (
	DayOfWeekNameMonday    DayOfWeekNameEnum = "MONDAY"
	DayOfWeekNameTuesday   DayOfWeekNameEnum = "TUESDAY"
	DayOfWeekNameWednesday DayOfWeekNameEnum = "WEDNESDAY"
	DayOfWeekNameThursday  DayOfWeekNameEnum = "THURSDAY"
	DayOfWeekNameFriday    DayOfWeekNameEnum = "FRIDAY"
	DayOfWeekNameSaturday  DayOfWeekNameEnum = "SATURDAY"
	DayOfWeekNameSunday    DayOfWeekNameEnum = "SUNDAY"
)

var mappingDayOfWeekNameEnum = map[string]DayOfWeekNameEnum{
	"MONDAY":    DayOfWeekNameMonday,
	"TUESDAY":   DayOfWeekNameTuesday,
	"WEDNESDAY": DayOfWeekNameWednesday,
	"THURSDAY":  DayOfWeekNameThursday,
	"FRIDAY":    DayOfWeekNameFriday,
	"SATURDAY":  DayOfWeekNameSaturday,
	"SUNDAY":    DayOfWeekNameSunday,
}

var mappingDayOfWeekNameEnumLowerCase = map[string]DayOfWeekNameEnum{
	"monday":    DayOfWeekNameMonday,
	"tuesday":   DayOfWeekNameTuesday,
	"wednesday": DayOfWeekNameWednesday,
	"thursday":  DayOfWeekNameThursday,
	"friday":    DayOfWeekNameFriday,
	"saturday":  DayOfWeekNameSaturday,
	"sunday":    DayOfWeekNameSunday,
}

// GetDayOfWeekNameEnumValues Enumerates the set of values for DayOfWeekNameEnum
func GetDayOfWeekNameEnumValues() []DayOfWeekNameEnum {
	values := make([]DayOfWeekNameEnum, 0)
	for _, v := range mappingDayOfWeekNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDayOfWeekNameEnumStringValues Enumerates the set of values in String for DayOfWeekNameEnum
func GetDayOfWeekNameEnumStringValues() []string {
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

// GetMappingDayOfWeekNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDayOfWeekNameEnum(val string) (DayOfWeekNameEnum, bool) {
	enum, ok := mappingDayOfWeekNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
