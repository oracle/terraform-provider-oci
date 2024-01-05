// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateJobDetails The details required to update a job.
type UpdateJobDetails interface {

	// The description of the job.
	GetDescription() *string

	// The job timeout duration, which is expressed like "1h 10m 15s".
	GetTimeout() *string

	GetResultLocation() JobExecutionResultLocation

	GetScheduleDetails() *JobScheduleDetails
}

type updatejobdetails struct {
	JsonData        []byte
	Description     *string                    `mandatory:"false" json:"description"`
	Timeout         *string                    `mandatory:"false" json:"timeout"`
	ResultLocation  jobexecutionresultlocation `mandatory:"false" json:"resultLocation"`
	ScheduleDetails *JobScheduleDetails        `mandatory:"false" json:"scheduleDetails"`
	JobType         string                     `json:"jobType"`
}

// UnmarshalJSON unmarshals json
func (m *updatejobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatejobdetails updatejobdetails
	s := struct {
		Model Unmarshalerupdatejobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.Timeout = s.Model.Timeout
	m.ResultLocation = s.Model.ResultLocation
	m.ScheduleDetails = s.Model.ScheduleDetails
	m.JobType = s.Model.JobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatejobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobType {
	case "SQL":
		mm := UpdateSqlJobDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateJobDetails: %s.", m.JobType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m updatejobdetails) GetDescription() *string {
	return m.Description
}

// GetTimeout returns Timeout
func (m updatejobdetails) GetTimeout() *string {
	return m.Timeout
}

// GetResultLocation returns ResultLocation
func (m updatejobdetails) GetResultLocation() jobexecutionresultlocation {
	return m.ResultLocation
}

// GetScheduleDetails returns ScheduleDetails
func (m updatejobdetails) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

func (m updatejobdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatejobdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
