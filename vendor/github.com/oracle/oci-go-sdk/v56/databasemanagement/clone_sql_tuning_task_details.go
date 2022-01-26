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

// CloneSqlTuningTaskDetails Request to clone and run a SQL tuning task. The new task uses same inputs as the one being cloned.
type CloneSqlTuningTaskDetails struct {

	// The name of the SQL tuning task. The name is unique per user in a database, and it is case sensitive.
	TaskName *string `mandatory:"true" json:"taskName"`

	// The identifier of the task being cloned. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint
	// ListSqlTuningAdvisorTasks
	OriginalTaskId *int64 `mandatory:"true" json:"originalTaskId"`

	CredentialDetails SqlTuningTaskCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The description of the SQL tuning task.
	TaskDescription *string `mandatory:"false" json:"taskDescription"`
}

func (m CloneSqlTuningTaskDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CloneSqlTuningTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskDescription   *string                        `json:"taskDescription"`
		TaskName          *string                        `json:"taskName"`
		OriginalTaskId    *int64                         `json:"originalTaskId"`
		CredentialDetails sqltuningtaskcredentialdetails `json:"credentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskDescription = model.TaskDescription

	m.TaskName = model.TaskName

	m.OriginalTaskId = model.OriginalTaskId

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningTaskCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	return
}
