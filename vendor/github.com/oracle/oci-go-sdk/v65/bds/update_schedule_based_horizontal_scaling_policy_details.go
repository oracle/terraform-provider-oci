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

// UpdateScheduleBasedHorizontalScalingPolicyDetails Update details of a schedule based horizontal autoscaling policy.
// In a schedule-based autoscaling policy, an autoscaling action is triggered at the scheduled execution time.
type UpdateScheduleBasedHorizontalScalingPolicyDetails struct {

	// The time zone of the execution schedule, in IANA time zone database name format
	Timezone *string `mandatory:"false" json:"timezone"`

	// Details of a horizontal scaling schedule.
	ScheduleDetails []HorizontalScalingScheduleDetails `mandatory:"false" json:"scheduleDetails"`
}

func (m UpdateScheduleBasedHorizontalScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScheduleBasedHorizontalScalingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateScheduleBasedHorizontalScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateScheduleBasedHorizontalScalingPolicyDetails UpdateScheduleBasedHorizontalScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeUpdateScheduleBasedHorizontalScalingPolicyDetails
	}{
		"SCHEDULE_BASED_HORIZONTAL_SCALING_POLICY",
		(MarshalTypeUpdateScheduleBasedHorizontalScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateScheduleBasedHorizontalScalingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Timezone        *string                            `json:"timezone"`
		ScheduleDetails []horizontalscalingscheduledetails `json:"scheduleDetails"`
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
	return
}
