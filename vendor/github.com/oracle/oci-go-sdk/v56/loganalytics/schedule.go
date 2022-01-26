// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Schedule Schedule for scheduled task.
type Schedule interface {

	// Schedule misfire retry policy.
	GetMisfirePolicy() ScheduleMisfirePolicyEnum

	// The date and time the scheduled task should execute first time after create or update;
	// thereafter the task will execute as specified in the schedule.
	GetTimeOfFirstExecution() *common.SDKTime
}

type schedule struct {
	JsonData             []byte
	MisfirePolicy        ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
	TimeOfFirstExecution *common.SDKTime           `mandatory:"false" json:"timeOfFirstExecution"`
	Type                 string                    `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *schedule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerschedule schedule
	s := struct {
		Model Unmarshalerschedule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MisfirePolicy = s.Model.MisfirePolicy
	m.TimeOfFirstExecution = s.Model.TimeOfFirstExecution
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *schedule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "CRON":
		mm := CronSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIXED_FREQUENCY":
		mm := FixedFrequencySchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMisfirePolicy returns MisfirePolicy
func (m schedule) GetMisfirePolicy() ScheduleMisfirePolicyEnum {
	return m.MisfirePolicy
}

//GetTimeOfFirstExecution returns TimeOfFirstExecution
func (m schedule) GetTimeOfFirstExecution() *common.SDKTime {
	return m.TimeOfFirstExecution
}

func (m schedule) String() string {
	return common.PointerString(m)
}

// ScheduleMisfirePolicyEnum Enum with underlying type: string
type ScheduleMisfirePolicyEnum string

// Set of constants representing the allowable values for ScheduleMisfirePolicyEnum
const (
	ScheduleMisfirePolicyRetryOnce         ScheduleMisfirePolicyEnum = "RETRY_ONCE"
	ScheduleMisfirePolicyRetryIndefinitely ScheduleMisfirePolicyEnum = "RETRY_INDEFINITELY"
	ScheduleMisfirePolicySkip              ScheduleMisfirePolicyEnum = "SKIP"
)

var mappingScheduleMisfirePolicy = map[string]ScheduleMisfirePolicyEnum{
	"RETRY_ONCE":         ScheduleMisfirePolicyRetryOnce,
	"RETRY_INDEFINITELY": ScheduleMisfirePolicyRetryIndefinitely,
	"SKIP":               ScheduleMisfirePolicySkip,
}

// GetScheduleMisfirePolicyEnumValues Enumerates the set of values for ScheduleMisfirePolicyEnum
func GetScheduleMisfirePolicyEnumValues() []ScheduleMisfirePolicyEnum {
	values := make([]ScheduleMisfirePolicyEnum, 0)
	for _, v := range mappingScheduleMisfirePolicy {
		values = append(values, v)
	}
	return values
}

// ScheduleTypeEnum Enum with underlying type: string
type ScheduleTypeEnum string

// Set of constants representing the allowable values for ScheduleTypeEnum
const (
	ScheduleTypeFixedFrequency ScheduleTypeEnum = "FIXED_FREQUENCY"
	ScheduleTypeCron           ScheduleTypeEnum = "CRON"
)

var mappingScheduleType = map[string]ScheduleTypeEnum{
	"FIXED_FREQUENCY": ScheduleTypeFixedFrequency,
	"CRON":            ScheduleTypeCron,
}

// GetScheduleTypeEnumValues Enumerates the set of values for ScheduleTypeEnum
func GetScheduleTypeEnumValues() []ScheduleTypeEnum {
	values := make([]ScheduleTypeEnum, 0)
	for _, v := range mappingScheduleType {
		values = append(values, v)
	}
	return values
}
