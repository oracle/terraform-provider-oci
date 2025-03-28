// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Job The details of the job.
type Job interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the job resides.
	GetCompartmentId() *string

	// The display name of the job.
	GetName() *string

	// The schedule type of the job.
	GetScheduleType() JobScheduleTypeEnum

	// The lifecycle state of the job.
	GetLifecycleState() JobLifecycleStateEnum

	// The date and time when the job was created.
	GetTimeCreated() *common.SDKTime

	// The date and time when the job was last updated.
	GetTimeUpdated() *common.SDKTime

	// The description of the job.
	GetDescription() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the job has to be executed.
	GetManagedDatabaseGroupId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database where the job has to be executed.
	GetManagedDatabaseId() *string

	// The details of the Managed Databases where the job has to be executed.
	GetManagedDatabasesDetails() []JobDatabase

	// The subtype of the Oracle Database where the job has to be executed. Applicable only when managedDatabaseGroupId is provided.
	GetDatabaseSubType() DatabaseSubTypeEnum

	// The job timeout duration, which is expressed like "1h 10m 15s".
	GetTimeout() *string

	GetResultLocation() JobExecutionResultLocation

	GetScheduleDetails() *JobScheduleDetails

	// The error message that is returned if the job submission fails. Null is returned in all other scenarios.
	GetSubmissionErrorMessage() *string
}

type job struct {
	JsonData                []byte
	Description             *string                    `mandatory:"false" json:"description"`
	ManagedDatabaseGroupId  *string                    `mandatory:"false" json:"managedDatabaseGroupId"`
	ManagedDatabaseId       *string                    `mandatory:"false" json:"managedDatabaseId"`
	ManagedDatabasesDetails []JobDatabase              `mandatory:"false" json:"managedDatabasesDetails"`
	DatabaseSubType         DatabaseSubTypeEnum        `mandatory:"false" json:"databaseSubType,omitempty"`
	Timeout                 *string                    `mandatory:"false" json:"timeout"`
	ResultLocation          jobexecutionresultlocation `mandatory:"false" json:"resultLocation"`
	ScheduleDetails         *JobScheduleDetails        `mandatory:"false" json:"scheduleDetails"`
	SubmissionErrorMessage  *string                    `mandatory:"false" json:"submissionErrorMessage"`
	Id                      *string                    `mandatory:"true" json:"id"`
	CompartmentId           *string                    `mandatory:"true" json:"compartmentId"`
	Name                    *string                    `mandatory:"true" json:"name"`
	ScheduleType            JobScheduleTypeEnum        `mandatory:"true" json:"scheduleType"`
	LifecycleState          JobLifecycleStateEnum      `mandatory:"true" json:"lifecycleState"`
	TimeCreated             *common.SDKTime            `mandatory:"true" json:"timeCreated"`
	TimeUpdated             *common.SDKTime            `mandatory:"true" json:"timeUpdated"`
	JobType                 string                     `json:"jobType"`
}

// UnmarshalJSON unmarshals json
func (m *job) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerjob job
	s := struct {
		Model Unmarshalerjob
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.Name = s.Model.Name
	m.ScheduleType = s.Model.ScheduleType
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.Description = s.Model.Description
	m.ManagedDatabaseGroupId = s.Model.ManagedDatabaseGroupId
	m.ManagedDatabaseId = s.Model.ManagedDatabaseId
	m.ManagedDatabasesDetails = s.Model.ManagedDatabasesDetails
	m.DatabaseSubType = s.Model.DatabaseSubType
	m.Timeout = s.Model.Timeout
	m.ResultLocation = s.Model.ResultLocation
	m.ScheduleDetails = s.Model.ScheduleDetails
	m.SubmissionErrorMessage = s.Model.SubmissionErrorMessage
	m.JobType = s.Model.JobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *job) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.JobType {
	case "SQL":
		mm := SqlJob{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Job: %s.", m.JobType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m job) GetDescription() *string {
	return m.Description
}

// GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m job) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

// GetManagedDatabaseId returns ManagedDatabaseId
func (m job) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

// GetManagedDatabasesDetails returns ManagedDatabasesDetails
func (m job) GetManagedDatabasesDetails() []JobDatabase {
	return m.ManagedDatabasesDetails
}

// GetDatabaseSubType returns DatabaseSubType
func (m job) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

// GetTimeout returns Timeout
func (m job) GetTimeout() *string {
	return m.Timeout
}

// GetResultLocation returns ResultLocation
func (m job) GetResultLocation() jobexecutionresultlocation {
	return m.ResultLocation
}

// GetScheduleDetails returns ScheduleDetails
func (m job) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

// GetSubmissionErrorMessage returns SubmissionErrorMessage
func (m job) GetSubmissionErrorMessage() *string {
	return m.SubmissionErrorMessage
}

// GetId returns Id
func (m job) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m job) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m job) GetName() *string {
	return m.Name
}

// GetScheduleType returns ScheduleType
func (m job) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

// GetLifecycleState returns LifecycleState
func (m job) GetLifecycleState() JobLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m job) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m job) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m job) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m job) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// JobScheduleTypeEnum Enum with underlying type: string
type JobScheduleTypeEnum string

// Set of constants representing the allowable values for JobScheduleTypeEnum
const (
	JobScheduleTypeImmediate JobScheduleTypeEnum = "IMMEDIATE"
	JobScheduleTypeLater     JobScheduleTypeEnum = "LATER"
)

var mappingJobScheduleTypeEnum = map[string]JobScheduleTypeEnum{
	"IMMEDIATE": JobScheduleTypeImmediate,
	"LATER":     JobScheduleTypeLater,
}

var mappingJobScheduleTypeEnumLowerCase = map[string]JobScheduleTypeEnum{
	"immediate": JobScheduleTypeImmediate,
	"later":     JobScheduleTypeLater,
}

// GetJobScheduleTypeEnumValues Enumerates the set of values for JobScheduleTypeEnum
func GetJobScheduleTypeEnumValues() []JobScheduleTypeEnum {
	values := make([]JobScheduleTypeEnum, 0)
	for _, v := range mappingJobScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJobScheduleTypeEnumStringValues Enumerates the set of values in String for JobScheduleTypeEnum
func GetJobScheduleTypeEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"LATER",
	}
}

// GetMappingJobScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobScheduleTypeEnum(val string) (JobScheduleTypeEnum, bool) {
	enum, ok := mappingJobScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// JobLifecycleStateEnum Enum with underlying type: string
type JobLifecycleStateEnum string

// Set of constants representing the allowable values for JobLifecycleStateEnum
const (
	JobLifecycleStateActive   JobLifecycleStateEnum = "ACTIVE"
	JobLifecycleStateInactive JobLifecycleStateEnum = "INACTIVE"
)

var mappingJobLifecycleStateEnum = map[string]JobLifecycleStateEnum{
	"ACTIVE":   JobLifecycleStateActive,
	"INACTIVE": JobLifecycleStateInactive,
}

var mappingJobLifecycleStateEnumLowerCase = map[string]JobLifecycleStateEnum{
	"active":   JobLifecycleStateActive,
	"inactive": JobLifecycleStateInactive,
}

// GetJobLifecycleStateEnumValues Enumerates the set of values for JobLifecycleStateEnum
func GetJobLifecycleStateEnumValues() []JobLifecycleStateEnum {
	values := make([]JobLifecycleStateEnum, 0)
	for _, v := range mappingJobLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStateEnumStringValues Enumerates the set of values in String for JobLifecycleStateEnum
func GetJobLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingJobLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStateEnum(val string) (JobLifecycleStateEnum, bool) {
	enum, ok := mappingJobLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
