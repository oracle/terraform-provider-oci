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

// SqlJob The details of the SQL job.
type SqlJob struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the job.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the job resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the job.
	Name *string `mandatory:"true" json:"name"`

	// The date and time when the job was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the job was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group where the job has to be executed.
	ManagedDatabaseGroupId *string `mandatory:"false" json:"managedDatabaseGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database where the job has to be executed.
	ManagedDatabaseId *string `mandatory:"false" json:"managedDatabaseId"`

	// The details of the Managed Databases where the job has to be executed.
	ManagedDatabasesDetails []JobDatabase `mandatory:"false" json:"managedDatabasesDetails"`

	// The job timeout duration, which is expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	ScheduleDetails *JobScheduleDetails `mandatory:"false" json:"scheduleDetails"`

	// The error message that is returned if the job submission fails. Null is returned in all other scenarios.
	SubmissionErrorMessage *string `mandatory:"false" json:"submissionErrorMessage"`

	// The SQL text to be executed in the job. This is a mandatory field for the EXECUTE_SQL operationType.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// The database user name used to execute the SQL job. If the job is being executed on a Managed Database Group,
	// then the user name should exist on all the databases in the group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// The type of SQL. This is a mandatory field for the EXECUTE_SQL operationType.
	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// The SQL operation type.
	OperationType SqlJobOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The role of the database user. Indicates whether the database user is a normal user or sysdba.
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`

	// The subtype of the Oracle Database where the job has to be executed. Applicable only when managedDatabaseGroupId is provided.
	DatabaseSubType DatabaseSubTypeEnum `mandatory:"false" json:"databaseSubType,omitempty"`

	// The schedule type of the job.
	ScheduleType JobScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// The lifecycle state of the job.
	LifecycleState JobLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m SqlJob) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m SqlJob) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetName returns Name
func (m SqlJob) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m SqlJob) GetDescription() *string {
	return m.Description
}

//GetManagedDatabaseGroupId returns ManagedDatabaseGroupId
func (m SqlJob) GetManagedDatabaseGroupId() *string {
	return m.ManagedDatabaseGroupId
}

//GetManagedDatabaseId returns ManagedDatabaseId
func (m SqlJob) GetManagedDatabaseId() *string {
	return m.ManagedDatabaseId
}

//GetManagedDatabasesDetails returns ManagedDatabasesDetails
func (m SqlJob) GetManagedDatabasesDetails() []JobDatabase {
	return m.ManagedDatabasesDetails
}

//GetDatabaseSubType returns DatabaseSubType
func (m SqlJob) GetDatabaseSubType() DatabaseSubTypeEnum {
	return m.DatabaseSubType
}

//GetScheduleType returns ScheduleType
func (m SqlJob) GetScheduleType() JobScheduleTypeEnum {
	return m.ScheduleType
}

//GetLifecycleState returns LifecycleState
func (m SqlJob) GetLifecycleState() JobLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeout returns Timeout
func (m SqlJob) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m SqlJob) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

//GetScheduleDetails returns ScheduleDetails
func (m SqlJob) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

//GetSubmissionErrorMessage returns SubmissionErrorMessage
func (m SqlJob) GetSubmissionErrorMessage() *string {
	return m.SubmissionErrorMessage
}

//GetTimeCreated returns TimeCreated
func (m SqlJob) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m SqlJob) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m SqlJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlJobSqlTypeEnum(string(m.SqlType)); !ok && m.SqlType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlType: %s. Supported values are: %s.", m.SqlType, strings.Join(GetSqlJobSqlTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetSqlJobOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlJobRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlJobRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseSubTypeEnum(string(m.DatabaseSubType)); !ok && m.DatabaseSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseSubType: %s. Supported values are: %s.", m.DatabaseSubType, strings.Join(GetDatabaseSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetJobScheduleTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingJobLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetJobLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlJob) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlJob SqlJob
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeSqlJob
	}{
		"SQL",
		(MarshalTypeSqlJob)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SqlJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                    `json:"description"`
		ManagedDatabaseGroupId  *string                    `json:"managedDatabaseGroupId"`
		ManagedDatabaseId       *string                    `json:"managedDatabaseId"`
		ManagedDatabasesDetails []JobDatabase              `json:"managedDatabasesDetails"`
		DatabaseSubType         DatabaseSubTypeEnum        `json:"databaseSubType"`
		Timeout                 *string                    `json:"timeout"`
		ResultLocation          jobexecutionresultlocation `json:"resultLocation"`
		ScheduleDetails         *JobScheduleDetails        `json:"scheduleDetails"`
		SubmissionErrorMessage  *string                    `json:"submissionErrorMessage"`
		SqlType                 SqlJobSqlTypeEnum          `json:"sqlType"`
		SqlText                 *string                    `json:"sqlText"`
		UserName                *string                    `json:"userName"`
		Role                    SqlJobRoleEnum             `json:"role"`
		Id                      *string                    `json:"id"`
		CompartmentId           *string                    `json:"compartmentId"`
		Name                    *string                    `json:"name"`
		ScheduleType            JobScheduleTypeEnum        `json:"scheduleType"`
		LifecycleState          JobLifecycleStateEnum      `json:"lifecycleState"`
		TimeCreated             *common.SDKTime            `json:"timeCreated"`
		TimeUpdated             *common.SDKTime            `json:"timeUpdated"`
		OperationType           SqlJobOperationTypeEnum    `json:"operationType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ManagedDatabaseGroupId = model.ManagedDatabaseGroupId

	m.ManagedDatabaseId = model.ManagedDatabaseId

	m.ManagedDatabasesDetails = make([]JobDatabase, len(model.ManagedDatabasesDetails))
	for i, n := range model.ManagedDatabasesDetails {
		m.ManagedDatabasesDetails[i] = n
	}

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

	m.SubmissionErrorMessage = model.SubmissionErrorMessage

	m.SqlType = model.SqlType

	m.SqlText = model.SqlText

	m.UserName = model.UserName

	m.Role = model.Role

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.Name = model.Name

	m.ScheduleType = model.ScheduleType

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.OperationType = model.OperationType

	return
}

// SqlJobSqlTypeEnum Enum with underlying type: string
type SqlJobSqlTypeEnum string

// Set of constants representing the allowable values for SqlJobSqlTypeEnum
const (
	SqlJobSqlTypeQuery SqlJobSqlTypeEnum = "QUERY"
	SqlJobSqlTypeDml   SqlJobSqlTypeEnum = "DML"
	SqlJobSqlTypeDdl   SqlJobSqlTypeEnum = "DDL"
	SqlJobSqlTypePlsql SqlJobSqlTypeEnum = "PLSQL"
)

var mappingSqlJobSqlTypeEnum = map[string]SqlJobSqlTypeEnum{
	"QUERY": SqlJobSqlTypeQuery,
	"DML":   SqlJobSqlTypeDml,
	"DDL":   SqlJobSqlTypeDdl,
	"PLSQL": SqlJobSqlTypePlsql,
}

// GetSqlJobSqlTypeEnumValues Enumerates the set of values for SqlJobSqlTypeEnum
func GetSqlJobSqlTypeEnumValues() []SqlJobSqlTypeEnum {
	values := make([]SqlJobSqlTypeEnum, 0)
	for _, v := range mappingSqlJobSqlTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlJobSqlTypeEnumStringValues Enumerates the set of values in String for SqlJobSqlTypeEnum
func GetSqlJobSqlTypeEnumStringValues() []string {
	return []string{
		"QUERY",
		"DML",
		"DDL",
		"PLSQL",
	}
}

// GetMappingSqlJobSqlTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlJobSqlTypeEnum(val string) (SqlJobSqlTypeEnum, bool) {
	mappingSqlJobSqlTypeEnumIgnoreCase := make(map[string]SqlJobSqlTypeEnum)
	for k, v := range mappingSqlJobSqlTypeEnum {
		mappingSqlJobSqlTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSqlJobSqlTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SqlJobOperationTypeEnum Enum with underlying type: string
type SqlJobOperationTypeEnum string

// Set of constants representing the allowable values for SqlJobOperationTypeEnum
const (
	SqlJobOperationTypeExecuteSql SqlJobOperationTypeEnum = "EXECUTE_SQL"
)

var mappingSqlJobOperationTypeEnum = map[string]SqlJobOperationTypeEnum{
	"EXECUTE_SQL": SqlJobOperationTypeExecuteSql,
}

// GetSqlJobOperationTypeEnumValues Enumerates the set of values for SqlJobOperationTypeEnum
func GetSqlJobOperationTypeEnumValues() []SqlJobOperationTypeEnum {
	values := make([]SqlJobOperationTypeEnum, 0)
	for _, v := range mappingSqlJobOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlJobOperationTypeEnumStringValues Enumerates the set of values in String for SqlJobOperationTypeEnum
func GetSqlJobOperationTypeEnumStringValues() []string {
	return []string{
		"EXECUTE_SQL",
	}
}

// GetMappingSqlJobOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlJobOperationTypeEnum(val string) (SqlJobOperationTypeEnum, bool) {
	mappingSqlJobOperationTypeEnumIgnoreCase := make(map[string]SqlJobOperationTypeEnum)
	for k, v := range mappingSqlJobOperationTypeEnum {
		mappingSqlJobOperationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSqlJobOperationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SqlJobRoleEnum Enum with underlying type: string
type SqlJobRoleEnum string

// Set of constants representing the allowable values for SqlJobRoleEnum
const (
	SqlJobRoleNormal SqlJobRoleEnum = "NORMAL"
	SqlJobRoleSysdba SqlJobRoleEnum = "SYSDBA"
)

var mappingSqlJobRoleEnum = map[string]SqlJobRoleEnum{
	"NORMAL": SqlJobRoleNormal,
	"SYSDBA": SqlJobRoleSysdba,
}

// GetSqlJobRoleEnumValues Enumerates the set of values for SqlJobRoleEnum
func GetSqlJobRoleEnumValues() []SqlJobRoleEnum {
	values := make([]SqlJobRoleEnum, 0)
	for _, v := range mappingSqlJobRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlJobRoleEnumStringValues Enumerates the set of values in String for SqlJobRoleEnum
func GetSqlJobRoleEnumStringValues() []string {
	return []string{
		"NORMAL",
		"SYSDBA",
	}
}

// GetMappingSqlJobRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlJobRoleEnum(val string) (SqlJobRoleEnum, bool) {
	mappingSqlJobRoleEnumIgnoreCase := make(map[string]SqlJobRoleEnum)
	for k, v := range mappingSqlJobRoleEnum {
		mappingSqlJobRoleEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSqlJobRoleEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
