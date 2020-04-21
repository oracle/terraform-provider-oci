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

// Month Month of the year.
type Month struct {

	// Name of the month of the year.
	Name MonthNameEnum `mandatory:"true" json:"name"`
}

func (m Month) String() string {
	return common.PointerString(m)
}

// MonthNameEnum Enum with underlying type: string
type MonthNameEnum string

// Set of constants representing the allowable values for MonthNameEnum
const (
	MonthNameJanuary   MonthNameEnum = "JANUARY"
	MonthNameFebruary  MonthNameEnum = "FEBRUARY"
	MonthNameMarch     MonthNameEnum = "MARCH"
	MonthNameApril     MonthNameEnum = "APRIL"
	MonthNameMay       MonthNameEnum = "MAY"
	MonthNameJune      MonthNameEnum = "JUNE"
	MonthNameJuly      MonthNameEnum = "JULY"
	MonthNameAugust    MonthNameEnum = "AUGUST"
	MonthNameSeptember MonthNameEnum = "SEPTEMBER"
	MonthNameOctober   MonthNameEnum = "OCTOBER"
	MonthNameNovember  MonthNameEnum = "NOVEMBER"
	MonthNameDecember  MonthNameEnum = "DECEMBER"
)

var mappingMonthName = map[string]MonthNameEnum{
	"JANUARY":   MonthNameJanuary,
	"FEBRUARY":  MonthNameFebruary,
	"MARCH":     MonthNameMarch,
	"APRIL":     MonthNameApril,
	"MAY":       MonthNameMay,
	"JUNE":      MonthNameJune,
	"JULY":      MonthNameJuly,
	"AUGUST":    MonthNameAugust,
	"SEPTEMBER": MonthNameSeptember,
	"OCTOBER":   MonthNameOctober,
	"NOVEMBER":  MonthNameNovember,
	"DECEMBER":  MonthNameDecember,
}

// GetMonthNameEnumValues Enumerates the set of values for MonthNameEnum
func GetMonthNameEnumValues() []MonthNameEnum {
	values := make([]MonthNameEnum, 0)
	for _, v := range mappingMonthName {
		values = append(values, v)
	}
	return values
}
