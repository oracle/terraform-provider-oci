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

// HorizontalScalingScheduleDetails Details of a horizontal scaling schedule.
type HorizontalScalingScheduleDetails interface {
}

type horizontalscalingscheduledetails struct {
	JsonData     []byte
	ScheduleType string `json:"scheduleType"`
}

// UnmarshalJSON unmarshals json
func (m *horizontalscalingscheduledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhorizontalscalingscheduledetails horizontalscalingscheduledetails
	s := struct {
		Model Unmarshalerhorizontalscalingscheduledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ScheduleType = s.Model.ScheduleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *horizontalscalingscheduledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScheduleType {
	case "DAY_BASED":
		mm := DayBasedHorizontalScalingScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for HorizontalScalingScheduleDetails: %s.", m.ScheduleType)
		return *m, nil
	}
}

func (m horizontalscalingscheduledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m horizontalscalingscheduledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
