// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ScheduleIntervalTrigger The interval schedule trigger.
type ScheduleIntervalTrigger struct {

	// The interval of frequency.
	Interval *int `mandatory:"true" json:"interval"`

	// The schedule starting date time, if null, System set the time when schedule is created.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The schedule end date time, if null, the schedule will never expire.
	// Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// when true and timeStart is null, system generate a random start time between now and now + interval;
	// isRandomStartTime can be true if timeStart is null.
	IsRandomStartTime *bool `mandatory:"false" json:"isRandomStartTime"`

	// The type of frequency
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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
	ScheduleIntervalTriggerFrequencyHourly ScheduleIntervalTriggerFrequencyEnum = "HOURLY"
	ScheduleIntervalTriggerFrequencyDaily  ScheduleIntervalTriggerFrequencyEnum = "DAILY"
)

var mappingScheduleIntervalTriggerFrequencyEnum = map[string]ScheduleIntervalTriggerFrequencyEnum{
	"HOURLY": ScheduleIntervalTriggerFrequencyHourly,
	"DAILY":  ScheduleIntervalTriggerFrequencyDaily,
}

var mappingScheduleIntervalTriggerFrequencyEnumLowerCase = map[string]ScheduleIntervalTriggerFrequencyEnum{
	"hourly": ScheduleIntervalTriggerFrequencyHourly,
	"daily":  ScheduleIntervalTriggerFrequencyDaily,
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
		"HOURLY",
		"DAILY",
	}
}

// GetMappingScheduleIntervalTriggerFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleIntervalTriggerFrequencyEnum(val string) (ScheduleIntervalTriggerFrequencyEnum, bool) {
	enum, ok := mappingScheduleIntervalTriggerFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
