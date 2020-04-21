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

// DayOfWeek Day of the week.
type DayOfWeek struct {

	// Name of the day of the week.
	Name DayOfWeekNameEnum `mandatory:"true" json:"name"`
}

func (m DayOfWeek) String() string {
	return common.PointerString(m)
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

var mappingDayOfWeekName = map[string]DayOfWeekNameEnum{
	"MONDAY":    DayOfWeekNameMonday,
	"TUESDAY":   DayOfWeekNameTuesday,
	"WEDNESDAY": DayOfWeekNameWednesday,
	"THURSDAY":  DayOfWeekNameThursday,
	"FRIDAY":    DayOfWeekNameFriday,
	"SATURDAY":  DayOfWeekNameSaturday,
	"SUNDAY":    DayOfWeekNameSunday,
}

// GetDayOfWeekNameEnumValues Enumerates the set of values for DayOfWeekNameEnum
func GetDayOfWeekNameEnumValues() []DayOfWeekNameEnum {
	values := make([]DayOfWeekNameEnum, 0)
	for _, v := range mappingDayOfWeekName {
		values = append(values, v)
	}
	return values
}
