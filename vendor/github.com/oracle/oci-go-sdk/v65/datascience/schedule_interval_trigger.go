// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleIntervalTrigger The interval schedule definition.
type ScheduleIntervalTrigger struct {

	// The interval of frequency.
	Interval *int `mandatory:"true" json:"interval"`

	// The schedule starting date time in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The schedule end date time in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// the count that the schedule repeats
	Count *int `mandatory:"false" json:"count"`

	// if timeStart missing, by default it is Now, this flag allows to generate a random start time
	IsRandomStartTime *bool `mandatory:"false" json:"isRandomStartTime"`

	// The frequency
	Frequency ScheduleIntervalTriggerFrequencyEnum `mandatory:"true" json:"frequency"`
}

// GetTimeStart returns TimeStart
func (m ScheduleIntervalTrigger) GetTimeStart() *common.SDKTime {
	return m.TimeStart
}

// GetTimeEnd returns TimeEnd
func (m ScheduleIntervalTrigger) GetTimeEnd() *common.SDKTime {
	return m.TimeEnd
}

func (m ScheduleIntervalTrigger) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleIntervalTrigger) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleIntervalTriggerFrequencyEnum(string(m.Frequency)); !ok && m.Frequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Frequency: %s. Supported values are: %s.", m.Frequency, strings.Join(GetScheduleIntervalTriggerFrequencyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleIntervalTrigger) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleIntervalTrigger ScheduleIntervalTrigger
	s := struct {
		DiscriminatorParam string `json:"triggerType"`
		MarshalTypeScheduleIntervalTrigger
	}{
		"INTERVAL",
		(MarshalTypeScheduleIntervalTrigger)(m),
	}

	return json.Marshal(&s)
}

// ScheduleIntervalTriggerFrequencyEnum Enum with underlying type: string
type ScheduleIntervalTriggerFrequencyEnum string

// Set of constants representing the allowable values for ScheduleIntervalTriggerFrequencyEnum
const (
	ScheduleIntervalTriggerFrequencyMinutely ScheduleIntervalTriggerFrequencyEnum = "MINUTELY"
	ScheduleIntervalTriggerFrequencyHourly   ScheduleIntervalTriggerFrequencyEnum = "HOURLY"
	ScheduleIntervalTriggerFrequencyDaily    ScheduleIntervalTriggerFrequencyEnum = "DAILY"
	ScheduleIntervalTriggerFrequencyWeekly   ScheduleIntervalTriggerFrequencyEnum = "WEEKLY"
	ScheduleIntervalTriggerFrequencyMonthly  ScheduleIntervalTriggerFrequencyEnum = "MONTHLY"
)

var mappingScheduleIntervalTriggerFrequencyEnum = map[string]ScheduleIntervalTriggerFrequencyEnum{
	"MINUTELY": ScheduleIntervalTriggerFrequencyMinutely,
	"HOURLY":   ScheduleIntervalTriggerFrequencyHourly,
	"DAILY":    ScheduleIntervalTriggerFrequencyDaily,
	"WEEKLY":   ScheduleIntervalTriggerFrequencyWeekly,
	"MONTHLY":  ScheduleIntervalTriggerFrequencyMonthly,
}

var mappingScheduleIntervalTriggerFrequencyEnumLowerCase = map[string]ScheduleIntervalTriggerFrequencyEnum{
	"minutely": ScheduleIntervalTriggerFrequencyMinutely,
	"hourly":   ScheduleIntervalTriggerFrequencyHourly,
	"daily":    ScheduleIntervalTriggerFrequencyDaily,
	"weekly":   ScheduleIntervalTriggerFrequencyWeekly,
	"monthly":  ScheduleIntervalTriggerFrequencyMonthly,
}

// GetScheduleIntervalTriggerFrequencyEnumValues Enumerates the set of values for ScheduleIntervalTriggerFrequencyEnum
func GetScheduleIntervalTriggerFrequencyEnumValues() []ScheduleIntervalTriggerFrequencyEnum {
	values := make([]ScheduleIntervalTriggerFrequencyEnum, 0)
	for _, v := range mappingScheduleIntervalTriggerFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleIntervalTriggerFrequencyEnumStringValues Enumerates the set of values in String for ScheduleIntervalTriggerFrequencyEnum
func GetScheduleIntervalTriggerFrequencyEnumStringValues() []string {
	return []string{
		"MINUTELY",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingScheduleIntervalTriggerFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleIntervalTriggerFrequencyEnum(val string) (ScheduleIntervalTriggerFrequencyEnum, bool) {
	enum, ok := mappingScheduleIntervalTriggerFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
