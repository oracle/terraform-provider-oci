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
	"github.com/oracle/oci-go-sdk/v56/common"
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

var mappingMonthlyRuleFrequencyDetailsWeekOfMonth = map[string]MonthlyRuleFrequencyDetailsWeekOfMonthEnum{
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
	for _, v := range mappingMonthlyRuleFrequencyDetailsWeekOfMonth {
		values = append(values, v)
	}
	return values
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

var mappingMonthlyRuleFrequencyDetailsDayOfWeek = map[string]MonthlyRuleFrequencyDetailsDayOfWeekEnum{
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
	for _, v := range mappingMonthlyRuleFrequencyDetailsDayOfWeek {
		values = append(values, v)
	}
	return values
}
