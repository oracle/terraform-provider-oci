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

// CreateSqlJobDetails The details specific to the SQL job request.
type CreateSqlJobDetails struct {

	// The name of the job. Valid characters are uppercase or lowercase letters,
	// numbers, and "_". The name of the job cannot be modified. It must be unique
	// in the compartment and must begin with an alphabetic character.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the job has to be executed.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database where the job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`

	// The job timeout duration, which is expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	ScheduleDetails *JobScheduleDetails `mandatory:"false" json:"scheduleDetails"`

	// The SQL text to be executed as part of the job.
	SqlText *string `mandatory:"false" json:"sqlText"`

	InBinds *JobInBindsDetails `mandatory:"false" json:"inBinds"`

	OutBinds *JobOutBindsDetails `mandatory:"false" json:"outBinds"`

	// The database user name used to execute the SQL job. If the job is being executed on a
	// Managed Database Group, then the user name should exist on all the databases in the
	// group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// The password for the database user name used to execute the SQL job.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The schedule type of the job.
	ScheduleType JobScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// The subtype of the Oracle Database where the job has to be executed. Only applicable when managedDatabaseGroupId is provided.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// The SQL operation type.
	OperationType SqlJobOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The role of the database user. Indicates whether the database user is a normal user or sysdba.
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`
}

// GetName returns Name
func (m CreateSqlJobDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateSqlJobDetails) GetDescription() *string {
	return m.Description
}

// GetCompartmentId returns CompartmentId
func (m CreateSqlJobDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m CreateSqlJobDetails) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

// GetManagedDatabaseId returns ManagedDatabaseId
func (m CreateSqlJobDetails) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

// GetDatabaseSubType returns DatabaseSubType
func (m CreateSqlJobDetails) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

// GetScheduleType returns ScheduleType
func (m CreateSqlJobDetails) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

// GetTimeout returns Timeout
func (m CreateSqlJobDetails) GetTimeout() *string {
	return m.Timeout
}

// GetResultLocation returns ResultLocation
func (m CreateSqlJobDetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

// GetScheduleDetails returns ScheduleDetails
func (m CreateSqlJobDetails) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

func (m CreateSqlJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSqlJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobSqlTypeEnum(string(m.SqlType)); !ok && m.SqlType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlType: %s. Supported values are: %s.", m.SqlType, strings.Join(GetSqlJobSqlTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetSqlJobOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlJobRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateSqlJobDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSqlJobDetails CreateSqlJobDetails
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeCreateSqlJobDetails
	}{
		"SQL",
		(MarshalTypeCreateSqlJobDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSqlJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description            *string                    `json:"description"`
		ManagedDatabaseGroupId *string                    `json:"managedDatabaseGroupId"`
		ManagedDatabaseId      *string                    `json:"managedDatabaseId"`
		DatabaseSubType        DatabaseSubTypeEnum        `json:"databaseSubType"`
		Timeout                *string                    `json:"timeout"`
		ResultLocation         jobexecutionresultlocation `json:"resultLocation"`
		ScheduleDetails        *JobScheduleDetails        `json:"scheduleDetails"`
		SqlText                *string                    `json:"sqlText"`
		InBinds                *JobInBindsDetails         `json:"inBinds"`
		OutBinds               *JobOutBindsDetails        `json:"outBinds"`
		SqlType                SqlJobSqlTypeEnum          `json:"sqlType"`
		UserName               *string                    `json:"userName"`
		Password               *string                    `json:"password"`
		SecretId               *string                    `json:"secretId"`
		Role                   SqlJobRoleEnum             `json:"role"`
		Name                   *string                    `json:"name"`
		CompartmentId          *string                    `json:"compartmentId"`
		ScheduleType           JobScheduleTypeEnum        `json:"scheduleType"`
		OperationType          SqlJobOperationTypeEnum    `json:"operationType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ManagedDatabaseGroupId = model.ManagedDatabaseGroupId

	m.ManagedDatabaseId = model.ManagedDatabaseId

	m.DatabaseSubType = model.DatabaseSubType

	m.Timeout = model.Timeout

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(JobExecutionResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.ScheduleDetails = model.ScheduleDetails

	m.SqlText = model.SqlText

	m.InBinds = model.InBinds

	m.OutBinds = model.OutBinds

	m.SqlType = model.SqlType

	m.UserName = model.UserName

	m.Password = model.Password

	m.SecretId = model.SecretId

	m.Role = model.Role

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.ScheduleType = model.ScheduleType

	m.OperationType = model.OperationType

	return
}
