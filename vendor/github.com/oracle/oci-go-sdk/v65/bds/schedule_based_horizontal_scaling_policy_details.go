// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleBasedHorizontalScalingPolicyDetails Details of a schedule based horizontal autoscaling policy.
// In a schedule-based autoscaling policy, an autoscaling action is triggered at the scheduled execution time.
type ScheduleBasedHorizontalScalingPolicyDetails struct {

	// The time zone of the execution schedule, in IANA time zone database name format
	Timezone *string `mandatory:"false" json:"timezone"`

	// Details of a horizontal scaling schedule.
	ScheduleDetails []HorizontalScalingScheduleDetails `mandatory:"false" json:"scheduleDetails"`

	// The type of autoscaling trigger.
	TriggerType AutoScalePolicyDetailsTriggerTypeEnum `mandatory:"true" json:"triggerType"`

	// The type of autoscaling action to take.
	ActionType AutoScalePolicyDetailsActionTypeEnum `mandatory:"true" json:"actionType"`
}

// GetTriggerType returns TriggerType
func (m ScheduleBasedHorizontalScalingPolicyDetails) GetTriggerType() AutoScalePolicyDetailsTriggerTypeEnum {
	return m.TriggerType
}

// GetActionType returns ActionType
func (m ScheduleBasedHorizontalScalingPolicyDetails) GetActionType() AutoScalePolicyDetailsActionTypeEnum {
	return m.ActionType
}

func (m ScheduleBasedHorizontalScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleBasedHorizontalScalingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutoScalePolicyDetailsTriggerTypeEnum(string(m.TriggerType)); !ok && m.TriggerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggerType: %s. Supported values are: %s.", m.TriggerType, strings.Join(GetAutoScalePolicyDetailsTriggerTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoScalePolicyDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetAutoScalePolicyDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduleBasedHorizontalScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduleBasedHorizontalScalingPolicyDetails ScheduleBasedHorizontalScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeScheduleBasedHorizontalScalingPolicyDetails
	}{
		"SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY",
		(MarshalTypeScheduleBasedHorizontalScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ScheduleBasedHorizontalScalingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Timezone        *string                               `json:"timezone"`
		ScheduleDetails []horizontalscalingscheduledetails    `json:"scheduleDetails"`
		TriggerType     AutoScalePolicyDetailsTriggerTypeEnum `json:"triggerType"`
		ActionType      AutoScalePolicyDetailsActionTypeEnum  `json:"actionType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Timezone = model.Timezone

	m.ScheduleDetails = make([]HorizontalScalingScheduleDetails, len(model.ScheduleDetails))
	for i, n := range model.ScheduleDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ScheduleDetails[i] = nn.(HorizontalScalingScheduleDetails)
		} else {
			m.ScheduleDetails[i] = nil
		}
	}
	m.TriggerType = model.TriggerType

	m.ActionType = model.ActionType

	return
}
