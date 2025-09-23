// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateTaskScheduleDetails Attributes to update a task schedule.
type UpdateTaskScheduleDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// Recurrence specification for the task schedule execution (formatted according to RFC-5545 (https://icalendar.org/RFC-Specifications/iCalendar-RFC-5545/)). To run daily for 10 occurrences starts on September 2, 2024 09:00 EDT, it should be 'DTSTART;TZID=America/New_York:20240902T090000
	// RRULE:FREQ=DAILY;COUNT=10'. To run every 3 hours from 9:00 AM to 5:00 PM on August 5, 2024 EDT, it should be 'DTSTART;TZID=America/New_York:20240805T090000 RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=20240805T170000Z'.
	ExecutionRecurrences *string `mandatory:"true" json:"executionRecurrences"`

	TaskDetails TaskDetails `mandatory:"true" json:"taskDetails"`
}

func (m UpdateTaskScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTaskScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTaskScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FleetId              *string     `json:"fleetId"`
		ExecutionRecurrences *string     `json:"executionRecurrences"`
		TaskDetails          taskdetails `json:"taskDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FleetId = model.FleetId

	m.ExecutionRecurrences = model.ExecutionRecurrences

	nn, e = model.TaskDetails.UnmarshalPolymorphicJSON(model.TaskDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TaskDetails = nn.(TaskDetails)
	} else {
		m.TaskDetails = nil
	}

	return
}
