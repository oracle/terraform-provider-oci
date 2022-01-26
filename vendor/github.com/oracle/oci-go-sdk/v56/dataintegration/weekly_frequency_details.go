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

// WeeklyFrequencyDetails Frequency Details model for weekly frequency based on day of week.
type WeeklyFrequencyDetails struct {
	Time *Time `mandatory:"false" json:"time"`

	// A list of days of the week to be scheduled. i.e. execute on Monday and Thursday.
	Days []WeeklyFrequencyDetailsDaysEnum `mandatory:"false" json:"days,omitempty"`

	// the frequency of the schedule.
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
}

//GetFrequency returns Frequency
func (m WeeklyFrequencyDetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m WeeklyFrequencyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m WeeklyFrequencyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeWeeklyFrequencyDetails WeeklyFrequencyDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeWeeklyFrequencyDetails
	}{
		"WEEKLY",
		(MarshalTypeWeeklyFrequencyDetails)(m),
	}

	return json.Marshal(&s)
}

// WeeklyFrequencyDetailsDaysEnum Enum with underlying type: string
type WeeklyFrequencyDetailsDaysEnum string

// Set of constants representing the allowable values for WeeklyFrequencyDetailsDaysEnum
const (
	WeeklyFrequencyDetailsDaysSunday    WeeklyFrequencyDetailsDaysEnum = "SUNDAY"
	WeeklyFrequencyDetailsDaysMonday    WeeklyFrequencyDetailsDaysEnum = "MONDAY"
	WeeklyFrequencyDetailsDaysTuesday   WeeklyFrequencyDetailsDaysEnum = "TUESDAY"
	WeeklyFrequencyDetailsDaysWednesday WeeklyFrequencyDetailsDaysEnum = "WEDNESDAY"
	WeeklyFrequencyDetailsDaysThursday  WeeklyFrequencyDetailsDaysEnum = "THURSDAY"
	WeeklyFrequencyDetailsDaysFriday    WeeklyFrequencyDetailsDaysEnum = "FRIDAY"
	WeeklyFrequencyDetailsDaysSaturday  WeeklyFrequencyDetailsDaysEnum = "SATURDAY"
)

var mappingWeeklyFrequencyDetailsDays = map[string]WeeklyFrequencyDetailsDaysEnum{
	"SUNDAY":    WeeklyFrequencyDetailsDaysSunday,
	"MONDAY":    WeeklyFrequencyDetailsDaysMonday,
	"TUESDAY":   WeeklyFrequencyDetailsDaysTuesday,
	"WEDNESDAY": WeeklyFrequencyDetailsDaysWednesday,
	"THURSDAY":  WeeklyFrequencyDetailsDaysThursday,
	"FRIDAY":    WeeklyFrequencyDetailsDaysFriday,
	"SATURDAY":  WeeklyFrequencyDetailsDaysSaturday,
}

// GetWeeklyFrequencyDetailsDaysEnumValues Enumerates the set of values for WeeklyFrequencyDetailsDaysEnum
func GetWeeklyFrequencyDetailsDaysEnumValues() []WeeklyFrequencyDetailsDaysEnum {
	values := make([]WeeklyFrequencyDetailsDaysEnum, 0)
	for _, v := range mappingWeeklyFrequencyDetailsDays {
		values = append(values, v)
	}
	return values
}
