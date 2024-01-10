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

// AddScheduleBasedVerticalScalingPolicyDetails Details of a schedule based vertical autoscaling policy.
// In a schedule-based autoscaling policy, an autoscaling action is triggered at the scheduled execution time.
type AddScheduleBasedVerticalScalingPolicyDetails struct {

	// The time zone of the execution schedule, in IANA time zone database name format
	Timezone *string `mandatory:"false" json:"timezone"`

	// Details of a vertical scaling schedule.
	ScheduleDetails []VerticalScalingScheduleDetails `mandatory:"false" json:"scheduleDetails"`
}

func (m AddScheduleBasedVerticalScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddScheduleBasedVerticalScalingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AddScheduleBasedVerticalScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddScheduleBasedVerticalScalingPolicyDetails AddScheduleBasedVerticalScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeAddScheduleBasedVerticalScalingPolicyDetails
	}{
		"SCHEDULE_BASED_VERTICAL_SCALING_POLICY",
		(MarshalTypeAddScheduleBasedVerticalScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AddScheduleBasedVerticalScalingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Timezone        *string                          `json:"timezone"`
		ScheduleDetails []verticalscalingscheduledetails `json:"scheduleDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Timezone = model.Timezone

	m.ScheduleDetails = make([]VerticalScalingScheduleDetails, len(model.ScheduleDetails))
	for i, n := range model.ScheduleDetails {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ScheduleDetails[i] = nn.(VerticalScalingScheduleDetails)
		} else {
			m.ScheduleDetails[i] = nil
		}
	}
	return
}
