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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateSqlJobDetails The details specific to the SQL job request.
type UpdateSqlJobDetails struct {

	// The description of the job.
	Description *string `mandatory:"false" json:"description"`

	// The job timeout duration, which is expressed like "1h 10m 15s".
	Timeout *string `mandatory:"false" json:"timeout"`

	ResultLocation JobExecutionResultLocation `mandatory:"false" json:"resultLocation"`

	ScheduleDetails *JobScheduleDetails `mandatory:"false" json:"scheduleDetails"`

	// The SQL text to be executed as part of the job.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// The database user name used to execute the SQL job. If the job is being executed on a
	// Managed Database Group, then the user name should exist on all the databases in the
	// group with the same password.
	UserName *string `mandatory:"false" json:"userName"`

	// The password for the database user name used to execute the SQL job.
	Password *string `mandatory:"false" json:"password"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	SecretId *string `mandatory:"false" json:"secretId"`

	SqlType SqlJobSqlTypeEnum `mandatory:"false" json:"sqlType,omitempty"`

	// The role of the database user. Indicates whether the database user is a normal user or sysdba.
	Role SqlJobRoleEnum `mandatory:"false" json:"role,omitempty"`
}

//GetDescription returns Description
func (m UpdateSqlJobDetails) GetDescription() *string {
	return m.Description
}

//GetTimeout returns Timeout
func (m UpdateSqlJobDetails) GetTimeout() *string {
	return m.Timeout
}

//GetResultLocation returns ResultLocation
func (m UpdateSqlJobDetails) GetResultLocation() JobExecutionResultLocation {
	return m.ResultLocation
}

//GetScheduleDetails returns ScheduleDetails
func (m UpdateSqlJobDetails) GetScheduleDetails() *JobScheduleDetails {
	return m.ScheduleDetails
}

func (m UpdateSqlJobDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateSqlJobDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateSqlJobDetails UpdateSqlJobDetails
	s := struct {
		DiscriminatorParam string `json:"jobType"`
		MarshalTypeUpdateSqlJobDetails
	}{
		"SQL",
		(MarshalTypeUpdateSqlJobDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateSqlJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description     *string                    `json:"description"`
		Timeout         *string                    `json:"timeout"`
		ResultLocation  jobexecutionresultlocation `json:"resultLocation"`
		ScheduleDetails *JobScheduleDetails        `json:"scheduleDetails"`
		SqlText         *string                    `json:"sqlText"`
		SqlType         SqlJobSqlTypeEnum          `json:"sqlType"`
		UserName        *string                    `json:"userName"`
		Password        *string                    `json:"password"`
		SecretId        *string                    `json:"secretId"`
		Role            SqlJobRoleEnum             `json:"role"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

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

	m.SqlType = model.SqlType

	m.UserName = model.UserName

	m.Password = model.Password

	m.SecretId = model.SecretId

	m.Role = model.Role

	return
}
