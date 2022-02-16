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

// StartSqlTuningTaskDetails The request to start a SQL tuning task.
type StartSqlTuningTaskDetails struct {

	// The name of the SQL tuning task. The name is unique per user in a database, and it is case-sensitive.
	TaskName *string `mandatory:"true" json:"taskName"`

	CredentialDetails SqlTuningTaskCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The time limit for running the SQL tuning task.
	TotalTimeLimitInMinutes *int `mandatory:"true" json:"totalTimeLimitInMinutes"`

	// The scope for the SQL tuning task. For LIMITED scope, the SQL profile recommendation
	// is excluded, so the task is executed faster. For COMPREHENSIVE scope, the SQL profile recommendation
	// is included.
	Scope StartSqlTuningTaskDetailsScopeEnum `mandatory:"true" json:"scope"`

	// The array of the details of SQL statement on which tuning is performed.
	SqlDetails []SqlTuningTaskSqlDetail `mandatory:"true" json:"sqlDetails"`

	// The start time of the period in which SQL statements are running.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The end time of the period in which SQL statements are running.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// The description of the SQL tuning task.
	TaskDescription *string `mandatory:"false" json:"taskDescription"`

	// The time limit per SQL statement (in minutes). This is for a task with the COMPREHENSIVE scope.
	// The time limit per SQL statement should not be more than the total time limit.
	StatementTimeLimitInMinutes *int `mandatory:"false" json:"statementTimeLimitInMinutes"`
}

func (m StartSqlTuningTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StartSqlTuningTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStartSqlTuningTaskDetailsScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetStartSqlTuningTaskDetailsScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *StartSqlTuningTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskDescription             *string                            `json:"taskDescription"`
		StatementTimeLimitInMinutes *int                               `json:"statementTimeLimitInMinutes"`
		TaskName                    *string                            `json:"taskName"`
		CredentialDetails           sqltuningtaskcredentialdetails     `json:"credentialDetails"`
		TotalTimeLimitInMinutes     *int                               `json:"totalTimeLimitInMinutes"`
		Scope                       StartSqlTuningTaskDetailsScopeEnum `json:"scope"`
		SqlDetails                  []SqlTuningTaskSqlDetail           `json:"sqlDetails"`
		TimeStarted                 *common.SDKTime                    `json:"timeStarted"`
		TimeEnded                   *common.SDKTime                    `json:"timeEnded"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskDescription = model.TaskDescription

	m.StatementTimeLimitInMinutes = model.StatementTimeLimitInMinutes

	m.TaskName = model.TaskName

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningTaskCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.TotalTimeLimitInMinutes = model.TotalTimeLimitInMinutes

	m.Scope = model.Scope

	m.SqlDetails = make([]SqlTuningTaskSqlDetail, len(model.SqlDetails))
	for i, n := range model.SqlDetails {
		m.SqlDetails[i] = n
	}

	m.TimeStarted = model.TimeStarted

	m.TimeEnded = model.TimeEnded

	return
}

// StartSqlTuningTaskDetailsScopeEnum Enum with underlying type: string
type StartSqlTuningTaskDetailsScopeEnum string

// Set of constants representing the allowable values for StartSqlTuningTaskDetailsScopeEnum
const (
	StartSqlTuningTaskDetailsScopeLimited       StartSqlTuningTaskDetailsScopeEnum = "LIMITED"
	StartSqlTuningTaskDetailsScopeComprehensive StartSqlTuningTaskDetailsScopeEnum = "COMPREHENSIVE"
)

var mappingStartSqlTuningTaskDetailsScopeEnum = map[string]StartSqlTuningTaskDetailsScopeEnum{
	"LIMITED":       StartSqlTuningTaskDetailsScopeLimited,
	"COMPREHENSIVE": StartSqlTuningTaskDetailsScopeComprehensive,
}

// GetStartSqlTuningTaskDetailsScopeEnumValues Enumerates the set of values for StartSqlTuningTaskDetailsScopeEnum
func GetStartSqlTuningTaskDetailsScopeEnumValues() []StartSqlTuningTaskDetailsScopeEnum {
	values := make([]StartSqlTuningTaskDetailsScopeEnum, 0)
	for _, v := range mappingStartSqlTuningTaskDetailsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetStartSqlTuningTaskDetailsScopeEnumStringValues Enumerates the set of values in String for StartSqlTuningTaskDetailsScopeEnum
func GetStartSqlTuningTaskDetailsScopeEnumStringValues() []string {
	return []string{
		"LIMITED",
		"COMPREHENSIVE",
	}
}

// GetMappingStartSqlTuningTaskDetailsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStartSqlTuningTaskDetailsScopeEnum(val string) (StartSqlTuningTaskDetailsScopeEnum, bool) {
	mappingStartSqlTuningTaskDetailsScopeEnumIgnoreCase := make(map[string]StartSqlTuningTaskDetailsScopeEnum)
	for k, v := range mappingStartSqlTuningTaskDetailsScopeEnum {
		mappingStartSqlTuningTaskDetailsScopeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStartSqlTuningTaskDetailsScopeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
