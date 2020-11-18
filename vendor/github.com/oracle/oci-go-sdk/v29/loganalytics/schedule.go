// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v29/common"
)

// Schedule Schedule for scheduled task.
type Schedule interface {

	// Schedule misfire retry policy.
	GetMisfirePolicy() ScheduleMisfirePolicyEnum
}

type schedule struct {
	JsonData      []byte
	MisfirePolicy ScheduleMisfirePolicyEnum `mandatory:"false" json:"misfirePolicy,omitempty"`
	Type          string                    `json:"type"`
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
