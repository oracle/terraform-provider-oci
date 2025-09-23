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

// TaskScheduleSummary A summary of the task schedule properties.
type TaskScheduleSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to identify this task schedule.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The name of the task schedule.
	Name *string `mandatory:"true" json:"name"`

	// The current state of the task schedule.
	LifecycleState TaskScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the task creator.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// Recurrence specification for the task schedule execution (formatted according to RFC-5545 (https://icalendar.org/RFC-Specifications/iCalendar-RFC-5545/)). To run daily for 10 occurrences starts on September 2, 2024 09:00 EDT, it should be 'DTSTART;TZID=America/New_York:20240902T090000
	// RRULE:FREQ=DAILY;COUNT=10'. To run every 3 hours from 9:00 AM to 5:00 PM on August 5, 2024 EDT, it should be 'DTSTART;TZID=America/New_York:20240805T090000 RRULE:FREQ=HOURLY;INTERVAL=3;UNTIL=20240805T170000Z'.
	ExecutionRecurrences *string `mandatory:"true" json:"executionRecurrences"`

	TaskDetails TaskDetails `mandatory:"true" json:"taskDetails"`

	// The date and time the task schedule was created (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the task schedule was last updated (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastUpdated *common.SDKTime `mandatory:"true" json:"timeLastUpdated"`

	// The date and time the task schedule ran last (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastRun *common.SDKTime `mandatory:"false" json:"timeLastRun"`

	// The date and time the task schedule will run next (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeNextRun *common.SDKTime `mandatory:"false" json:"timeNextRun"`
}

func (m TaskScheduleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskScheduleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTaskScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TaskScheduleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeLastRun          *common.SDKTime                `json:"timeLastRun"`
		TimeNextRun          *common.SDKTime                `json:"timeNextRun"`
		Id                   *string                        `json:"id"`
		FleetId              *string                        `json:"fleetId"`
		Name                 *string                        `json:"name"`
		LifecycleState       TaskScheduleLifecycleStateEnum `json:"lifecycleState"`
		CreatedBy            *string                        `json:"createdBy"`
		ExecutionRecurrences *string                        `json:"executionRecurrences"`
		TaskDetails          taskdetails                    `json:"taskDetails"`
		TimeCreated          *common.SDKTime                `json:"timeCreated"`
		TimeLastUpdated      *common.SDKTime                `json:"timeLastUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeLastRun = model.TimeLastRun

	m.TimeNextRun = model.TimeNextRun

	m.Id = model.Id

	m.FleetId = model.FleetId

	m.Name = model.Name

	m.LifecycleState = model.LifecycleState

	m.CreatedBy = model.CreatedBy

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

	m.TimeCreated = model.TimeCreated

	m.TimeLastUpdated = model.TimeLastUpdated

	return
}
