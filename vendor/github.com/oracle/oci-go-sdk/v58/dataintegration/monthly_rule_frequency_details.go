// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MonthlyRuleFrequencyDetails Frequency Details model for monthly frequency based on week of month and day of week.
type MonthlyRuleFrequencyDetails struct {

	// This hold the repeatability aspect of a schedule. i.e. in a monhtly frequency, a task can be scheduled for every month, once in two months, once in tree months etc.
	Interval *int `mandatory:"false" json:"interval"`

	Time *Time `mandatory:"false" json:"time"`

	// This holds the week of the month in which the schedule should be triggered.
	WeekOfMonth MonthlyRuleFrequencyDetailsWeekOfMonthEnum `mandatory:"false" json:"weekOfMonth,omitempty"`

	// This holds the day of the week on which the schedule should be triggered.
	DayOfWeek MonthlyRuleFrequencyDetailsDayOfWeekEnum `mandatory:"false" json:"dayOfWeek,omitempty"`

	// the frequency of the schedule.
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
}

//GetFrequency returns Frequency
func (m MonthlyRuleFrequencyDetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m MonthlyRuleFrequencyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonthlyRuleFrequencyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum(string(m.WeekOfMonth)); !ok && m.WeekOfMonth != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WeekOfMonth: %s. Supported values are: %s.", m.WeekOfMonth, strings.Join(GetMonthlyRuleFrequencyDetailsWeekOfMonthEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMonthlyRuleFrequencyDetailsDayOfWeekEnum(string(m.DayOfWeek)); !ok && m.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", m.DayOfWeek, strings.Join(GetMonthlyRuleFrequencyDetailsDayOfWeekEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAbstractFrequencyDetailsFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetAbstractFrequencyDetailsFrequencyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MonthlyRuleFrequencyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMonthlyRuleFrequencyDetails MonthlyRuleFrequencyDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeMonthlyRuleFrequencyDetails
	}{
		"MONTHLY_RULE",
		(MarshalTypeMonthlyRuleFrequencyDetails)(m),
	}

	return json.Marshal(&s)
}

// MonthlyRuleFrequencyDetailsWeekOfMonthEnum Enum with underlying type: string
type MonthlyRuleFrequencyDetailsWeekOfMonthEnum string

// Set of constants representing the allowable values for MonthlyRuleFrequencyDetailsWeekOfMonthEnum
const (
	MonthlyRuleFrequencyDetailsWeekOfMonthFirst  MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "FIRST"
	MonthlyRuleFrequencyDetailsWeekOfMonthSecond MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "SECOND"
	MonthlyRuleFrequencyDetailsWeekOfMonthThird  MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "THIRD"
	MonthlyRuleFrequencyDetailsWeekOfMonthFourth MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "FOURTH"
	MonthlyRuleFrequencyDetailsWeekOfMonthFifth  MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "FIFTH"
	MonthlyRuleFrequencyDetailsWeekOfMonthLast   MonthlyRuleFrequencyDetailsWeekOfMonthEnum = "LAST"
)

var mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum = map[string]MonthlyRuleFrequencyDetailsWeekOfMonthEnum{
	"FIRST":  MonthlyRuleFrequencyDetailsWeekOfMonthFirst,
	"SECOND": MonthlyRuleFrequencyDetailsWeekOfMonthSecond,
	"THIRD":  MonthlyRuleFrequencyDetailsWeekOfMonthThird,
	"FOURTH": MonthlyRuleFrequencyDetailsWeekOfMonthFourth,
	"FIFTH":  MonthlyRuleFrequencyDetailsWeekOfMonthFifth,
	"LAST":   MonthlyRuleFrequencyDetailsWeekOfMonthLast,
}

// GetMonthlyRuleFrequencyDetailsWeekOfMonthEnumValues Enumerates the set of values for MonthlyRuleFrequencyDetailsWeekOfMonthEnum
func GetMonthlyRuleFrequencyDetailsWeekOfMonthEnumValues() []MonthlyRuleFrequencyDetailsWeekOfMonthEnum {
	values := make([]MonthlyRuleFrequencyDetailsWeekOfMonthEnum, 0)
	for _, v := range mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum {
		values = append(values, v)
	}
	return values
}

// GetMonthlyRuleFrequencyDetailsWeekOfMonthEnumStringValues Enumerates the set of values in String for MonthlyRuleFrequencyDetailsWeekOfMonthEnum
func GetMonthlyRuleFrequencyDetailsWeekOfMonthEnumStringValues() []string {
	return []string{
		"FIRST",
		"SECOND",
		"THIRD",
		"FOURTH",
		"FIFTH",
		"LAST",
	}
}

// GetMappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum(val string) (MonthlyRuleFrequencyDetailsWeekOfMonthEnum, bool) {
	mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnumIgnoreCase := make(map[string]MonthlyRuleFrequencyDetailsWeekOfMonthEnum)
	for k, v := range mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnum {
		mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMonthlyRuleFrequencyDetailsWeekOfMonthEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// MonthlyRuleFrequencyDetailsDayOfWeekEnum Enum with underlying type: string
type MonthlyRuleFrequencyDetailsDayOfWeekEnum string

// Set of constants representing the allowable values for MonthlyRuleFrequencyDetailsDayOfWeekEnum
const (
	MonthlyRuleFrequencyDetailsDayOfWeekSunday    MonthlyRuleFrequencyDetailsDayOfWeekEnum = "SUNDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekMonday    MonthlyRuleFrequencyDetailsDayOfWeekEnum = "MONDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekTuesday   MonthlyRuleFrequencyDetailsDayOfWeekEnum = "TUESDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekWednesday MonthlyRuleFrequencyDetailsDayOfWeekEnum = "WEDNESDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekThursday  MonthlyRuleFrequencyDetailsDayOfWeekEnum = "THURSDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekFriday    MonthlyRuleFrequencyDetailsDayOfWeekEnum = "FRIDAY"
	MonthlyRuleFrequencyDetailsDayOfWeekSaturday  MonthlyRuleFrequencyDetailsDayOfWeekEnum = "SATURDAY"
)

var mappingMonthlyRuleFrequencyDetailsDayOfWeekEnum = map[string]MonthlyRuleFrequencyDetailsDayOfWeekEnum{
	"SUNDAY":    MonthlyRuleFrequencyDetailsDayOfWeekSunday,
	"MONDAY":    MonthlyRuleFrequencyDetailsDayOfWeekMonday,
	"TUESDAY":   MonthlyRuleFrequencyDetailsDayOfWeekTuesday,
	"WEDNESDAY": MonthlyRuleFrequencyDetailsDayOfWeekWednesday,
	"THURSDAY":  MonthlyRuleFrequencyDetailsDayOfWeekThursday,
	"FRIDAY":    MonthlyRuleFrequencyDetailsDayOfWeekFriday,
	"SATURDAY":  MonthlyRuleFrequencyDetailsDayOfWeekSaturday,
}

// GetMonthlyRuleFrequencyDetailsDayOfWeekEnumValues Enumerates the set of values for MonthlyRuleFrequencyDetailsDayOfWeekEnum
func GetMonthlyRuleFrequencyDetailsDayOfWeekEnumValues() []MonthlyRuleFrequencyDetailsDayOfWeekEnum {
	values := make([]MonthlyRuleFrequencyDetailsDayOfWeekEnum, 0)
	for _, v := range mappingMonthlyRuleFrequencyDetailsDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetMonthlyRuleFrequencyDetailsDayOfWeekEnumStringValues Enumerates the set of values in String for MonthlyRuleFrequencyDetailsDayOfWeekEnum
func GetMonthlyRuleFrequencyDetailsDayOfWeekEnumStringValues() []string {
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

// GetMappingMonthlyRuleFrequencyDetailsDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonthlyRuleFrequencyDetailsDayOfWeekEnum(val string) (MonthlyRuleFrequencyDetailsDayOfWeekEnum, bool) {
	mappingMonthlyRuleFrequencyDetailsDayOfWeekEnumIgnoreCase := make(map[string]MonthlyRuleFrequencyDetailsDayOfWeekEnum)
	for k, v := range mappingMonthlyRuleFrequencyDetailsDayOfWeekEnum {
		mappingMonthlyRuleFrequencyDetailsDayOfWeekEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMonthlyRuleFrequencyDetailsDayOfWeekEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
