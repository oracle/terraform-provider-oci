// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ExecutionSchedule Specifies the execution schedule for a policy.
type ExecutionSchedule interface {

	// Specifies the time zone the schedule is in.
	GetTimezone() ExecutionScheduleTimezoneEnum
}

type executionschedule struct {
	JsonData []byte
	Timezone ExecutionScheduleTimezoneEnum `mandatory:"true" json:"timezone"`
	Type     string                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *executionschedule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutionschedule executionschedule
	s := struct {
		Model Unmarshalerexecutionschedule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Timezone = s.Model.Timezone
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executionschedule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "cron":
		mm := CronExecutionSchedule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimezone returns Timezone
func (m executionschedule) GetTimezone() ExecutionScheduleTimezoneEnum {
	return m.Timezone
}

func (m executionschedule) String() string {
	return common.PointerString(m)
}

// ExecutionScheduleTimezoneEnum Enum with underlying type: string
type ExecutionScheduleTimezoneEnum string

// Set of constants representing the allowable values for ExecutionScheduleTimezoneEnum
const (
	ExecutionScheduleTimezoneUtc ExecutionScheduleTimezoneEnum = "UTC"
)

var mappingExecutionScheduleTimezone = map[string]ExecutionScheduleTimezoneEnum{
	"UTC": ExecutionScheduleTimezoneUtc,
}

// GetExecutionScheduleTimezoneEnumValues Enumerates the set of values for ExecutionScheduleTimezoneEnum
func GetExecutionScheduleTimezoneEnumValues() []ExecutionScheduleTimezoneEnum {
	values := make([]ExecutionScheduleTimezoneEnum, 0)
	for _, v := range mappingExecutionScheduleTimezone {
		values = append(values, v)
	}
	return values
}
