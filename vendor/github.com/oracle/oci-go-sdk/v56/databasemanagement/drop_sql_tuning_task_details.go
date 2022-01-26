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

// DropSqlTuningTaskDetails Request to drop a SQL tuning task.
type DropSqlTuningTaskDetails struct {

	// The identifier of the task being dropped. This is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint
	// ListSqlTuningAdvisorTasks
	TaskId *int64 `mandatory:"true" json:"taskId"`

	CredentialDetails SqlTuningTaskCredentialDetails `mandatory:"true" json:"credentialDetails"`
}

func (m DropSqlTuningTaskDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DropSqlTuningTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskId            *int64                         `json:"taskId"`
		CredentialDetails sqltuningtaskcredentialdetails `json:"credentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskId = model.TaskId

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
