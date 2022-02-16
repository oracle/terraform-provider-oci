// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateJobDetails The details required to create a job.
type CreateJobDetails interface {

	// The name of the job. Valid characters are uppercase or lowercase letters,
	// numbers, and "_". The name of the job cannot be modified. It must be unique
	// in the compartment and must begin with an alphabetic character.
	GetName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the job resides.
	GetCompartmentId() *string

	// The schedule type of the job.
	GetScheduleType() JobScheduleTypeEnum

	// The description of the job.
	GetDescription() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the job has to be executed.
	GetManagedDatabaseGroupId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database where the job has to be executed.
	GetManagedDatabaseId() *string

	// The subtype of the Oracle Database where the job has to be executed. Only applicable when managedDatabaseGroupId is provided.
	GetDatabaseSubType() DatabaseSubTypeEnum

	// The job timeout duration, which is expressed like "1h 10m 15s".
	GetTimeout() *string

	GetResultLocation() JobExecutionResultLocation

	GetScheduleDetails() *JobScheduleDetails
}

type createjobdetails struct {
	JsonData               []byte
	Name                   *string                    `mandatory:"true" json:"name"`
	CompartmentId          *string                    `mandatory:"true" json:"compartmentId"`
	ScheduleType           JobScheduleTypeEnum        `mandatory:"true" json:"scheduleType"`
	Description            *string                    `mandatory:"false" json:"description"`
	ManagedDatabaseGroupId *string                    `mandatory:"false" json:"managedDatabaseGroupId"`
	ManagedDatabaseId      *string                    `mandatory:"false" json:"managedDatabaseId"`
	DatabaseSubType        DatabaseSubTypeEnum        `mandatory:"false" json:"databaseSubType,omitempty"`
	Timeout                *string                    `mandatory:"false" json:"timeout"`
	ResultLocation         JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`
	ScheduleDetails        *JobScheduleDetails        `mandatory:"false" json:"scheduleDetails"`
	JobType                string                     `json:"jobType"`
}

// UnmarshalJSON unmarshals json
func (m *createjobdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatejobdetails createjobdetails
	s := struct {
		Model Unmarshalercreatejobdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.CompartmentId = s.Model.CompartmentId
	m.ScheduleType = s.Model.ScheduleType
	m.Description = s.Model.Description
	m.ManagedDatabaseGroupId = s.Model.ManagedDatabaseGroupId
	m.ManagedDatabaseId = s.Model.ManagedDatabaseId
	m.DatabaseSubType = s.Model.DatabaseSubType
	m.Timeout = s.Model.Timeout
	m.ResultLocation = s.Model.ResultLocation
	m.ScheduleDetails = s.Model.ScheduleDetails
	m.JobType = s.Model.JobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createjobdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobType {
	case "SQL":
		mm := CreateSqlJobDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createjobdetails) GetName() *string {
	return m.Name
}

//GetCompartmentId returns CompartmentId
func (m createjobdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetScheduleType returns ScheduleType
func (m createjobdetails) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

//GetDescription returns Description
func (m createjobdetails) GetDescription() *string {
	return m.Description
}

//GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m createjobdetails) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

//GetManagedDatabaseId returns ManagedDatabaseId
func (m createjobdetails) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

//GetDatabaseSubType returns DatabaseSubType
func (m createjobdetails) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

//GetTimeout returns Timeout
func (m createjobdetails) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m createjobdetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

//GetScheduleDetails returns ScheduleDetails
func (m createjobdetails) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

func (m createjobdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createjobdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
