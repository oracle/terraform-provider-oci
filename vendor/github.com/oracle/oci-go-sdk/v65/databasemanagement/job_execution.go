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

// JobExecution The details of a job execution.
type JobExecution struct {

	// The identifier of the job execution.
	Id *string `mandatory:"true" json:"id"`

	// The name of the job execution.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the parent job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database associated with the job execution.
	ManagedDatabaseId *string `mandatory:"true" json:"managedDatabaseId"`

	// The name of the Managed Database associated with the job execution.
	ManagedDatabaseName *string `mandatory:"true" json:"managedDatabaseName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the parent job.
	JobId *string `mandatory:"true" json:"jobId"`

	// The name of the parent job.
	JobName *string `mandatory:"true" json:"jobName"`

	// The identifier of the associated job run.
	JobRunId *string `mandatory:"true" json:"jobRunId"`

	// The status of the job execution.
	Status JobExecutionStatusEnum `mandatory:"true" json:"status"`

	// The date and time when the job execution was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the parent job has to be executed.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// The type of Oracle Database installation.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// A list of the supported infrastructure that can be used to deploy the database.
	DeploymentType DeploymentTypeEnum `mandatory:"false" json:"deploymentType,omitempty"`

	// Indicates whether the Oracle Database is part of a cluster.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The workload type of the Autonomous Database.
	WorkloadType WorkloadTypeEnum `mandatory:"false" json:"workloadType,omitempty"`

	// The error message that is returned if the job execution fails. Null is returned if the job is
	// still running or if the job execution is successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	ResultDetails JobExecutionResultDetails `mandatory:"false" json:"resultDetails"`

	// The date and time when the job execution completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The database user name used to execute the SQL job.
	UserName *string `mandatory:"false" json:"userName"`

	// The SQL text executed as part of the job.
	SqlText *string `mandatory:"false" json:"sqlText"`

	InBinds *JobInBindsDetails `mandatory:"false" json:"inBinds"`

	OutBinds *JobOutBindsDetails `mandatory:"false" json:"outBinds"`

	ScheduleDetails *JobScheduleDetails `mandatory:"false" json:"scheduleDetails"`
}

func (m JobExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JobExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobExecutionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobExecutionStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkloadTypeEnum(string(m.WorkloadType)); !ok && m.WorkloadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkloadType: %s. Supported values are: %s.", m.WorkloadType, strings.Join(GetWorkloadTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *JobExecution) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ManagedDatabaseGroupId *string                   `json:"managedDatabaseGroupId"`
		DatabaseType           DatabaseTypeEnum          `json:"databaseType"`
		DatabaseSubType        DatabaseSubTypeEnum       `json:"databaseSubType"`
		DeploymentType         DeploymentTypeEnum        `json:"deploymentType"`
		IsCluster              *bool                     `json:"isCluster"`
		WorkloadType           WorkloadTypeEnum          `json:"workloadType"`
		ErrorMessage           *string                   `json:"errorMessage"`
		ResultDetails          jobexecutionresultdetails `json:"resultDetails"`
		TimeCompleted          *common.SDKTime           `json:"timeCompleted"`
		UserName               *string                   `json:"userName"`
		SqlText                *string                   `json:"sqlText"`
		InBinds                *JobInBindsDetails        `json:"inBinds"`
		OutBinds               *JobOutBindsDetails       `json:"outBinds"`
		ScheduleDetails        *JobScheduleDetails       `json:"scheduleDetails"`
		Id                     *string                   `json:"id"`
		Name                   *string                   `json:"name"`
		CompartmentId          *string                   `json:"compartmentId"`
		ManagedDatabaseId      *string                   `json:"managedDatabaseId"`
		ManagedDatabaseName    *string                   `json:"managedDatabaseName"`
		JobId                  *string                   `json:"jobId"`
		JobName                *string                   `json:"jobName"`
		JobRunId               *string                   `json:"jobRunId"`
		Status                 JobExecutionStatusEnum    `json:"status"`
		TimeCreated            *common.SDKTime           `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ManagedDatabaseGroupId = model.ManagedDatabaseGroupId

	m.DatabaseType = model.DatabaseType

	m.DatabaseSubType = model.DatabaseSubType

	m.DeploymentType = model.DeploymentType

	m.IsCluster = model.IsCluster

	m.WorkloadType = model.WorkloadType

	m.ErrorMessage = model.ErrorMessage

	nn, e = model.ResultDetails.UnmarshalPolymorphicJSON(model.ResultDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultDetails = nn.(JobExecutionResultDetails)
	} else {
		m.ResultDetails = nil
	}

	m.TimeCompleted = model.TimeCompleted

	m.UserName = model.UserName

	m.SqlText = model.SqlText

	m.InBinds = model.InBinds

	m.OutBinds = model.OutBinds

	m.ScheduleDetails = model.ScheduleDetails

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	m.ManagedDatabaseId = model.ManagedDatabaseId

	m.ManagedDatabaseName = model.ManagedDatabaseName

	m.JobId = model.JobId

	m.JobName = model.JobName

	m.JobRunId = model.JobRunId

	m.Status = model.Status

	m.TimeCreated = model.TimeCreated

	return
}

// JobExecutionStatusEnum Enum with underlying type: string
type JobExecutionStatusEnum string

// Set of constants representing the allowable values for JobExecutionStatusEnum
const (
	JobExecutionStatusSucceeded  JobExecutionStatusEnum = "SUCCEEDED"
	JobExecutionStatusFailed     JobExecutionStatusEnum = "FAILED"
	JobExecutionStatusInProgress JobExecutionStatusEnum = "IN_PROGRESS"
)

var mappingJobExecutionStatusEnum = map[string]JobExecutionStatusEnum{
	"SUCCEEDED":   JobExecutionStatusSucceeded,
	"FAILED":      JobExecutionStatusFailed,
	"IN_PROGRESS": JobExecutionStatusInProgress,
}

var mappingJobExecutionStatusEnumLowerCase = map[string]JobExecutionStatusEnum{
	"succeeded":   JobExecutionStatusSucceeded,
	"failed":      JobExecutionStatusFailed,
	"in_progress": JobExecutionStatusInProgress,
}

// GetJobExecutionStatusEnumValues Enumerates the set of values for JobExecutionStatusEnum
func GetJobExecutionStatusEnumValues() []JobExecutionStatusEnum {
	values := make([]JobExecutionStatusEnum, 0)
	for _, v := range mappingJobExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJobExecutionStatusEnumStringValues Enumerates the set of values in String for JobExecutionStatusEnum
func GetJobExecutionStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingJobExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobExecutionStatusEnum(val string) (JobExecutionStatusEnum, bool) {
	enum, ok := mappingJobExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
