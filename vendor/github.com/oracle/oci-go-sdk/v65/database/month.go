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

// Month Month of the year.
type Month struct {

	// Name of the month of the year.
	Name MonthNameEnum `mandatory:"true" json:"name"`
}

func (m Month) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Month) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonthNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetMonthNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingMonthNameEnum = map[string]MonthNameEnum{
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

var mappingMonthNameEnumLowerCase = map[string]MonthNameEnum{
	"january":   MonthNameJanuary,
	"february":  MonthNameFebruary,
	"march":     MonthNameMarch,
	"april":     MonthNameApril,
	"may":       MonthNameMay,
	"june":      MonthNameJune,
	"july":      MonthNameJuly,
	"august":    MonthNameAugust,
	"september": MonthNameSeptember,
	"october":   MonthNameOctober,
	"november":  MonthNameNovember,
	"december":  MonthNameDecember,
}

// GetMonthNameEnumValues Enumerates the set of values for MonthNameEnum
func GetMonthNameEnumValues() []MonthNameEnum {
	values := make([]MonthNameEnum, 0)
	for _, v := range mappingMonthNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMonthNameEnumStringValues Enumerates the set of values in String for MonthNameEnum
func GetMonthNameEnumStringValues() []string {
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

// GetMappingMonthNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonthNameEnum(val string) (MonthNameEnum, bool) {
	enum, ok := mappingMonthNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
