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

// CloneSqlTuningTaskDetails The request to clone and run a SQL tuning task. The new task uses the same inputs as the one being cloned.
// It takes either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type CloneSqlTuningTaskDetails struct {

	// The name of the SQL tuning task. The name is unique per user in a database, and it is case-sensitive.
	TaskName *string `mandatory:"true" json:"taskName"`

	// The identifier of the SQL tuning task being cloned. This is not the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint
	// ListSqlTuningAdvisorTasks.
	OriginalTaskId *int64 `mandatory:"true" json:"originalTaskId"`

	// The description of the SQL tuning task.
	TaskDescription *string `mandatory:"false" json:"taskDescription"`

	CredentialDetails SqlTuningTaskCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`
}

func (m CloneSqlTuningTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloneSqlTuningTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CloneSqlTuningTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TaskDescription    *string                        `json:"taskDescription"`
		CredentialDetails  sqltuningtaskcredentialdetails `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails      `json:"databaseCredential"`
		TaskName           *string                        `json:"taskName"`
		OriginalTaskId     *int64                         `json:"originalTaskId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TaskDescription = model.TaskDescription

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningTaskCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	nn, e = model.DatabaseCredential.UnmarshalPolymorphicJSON(model.DatabaseCredential.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseCredential = nn.(DatabaseCredentialDetails)
	} else {
		m.DatabaseCredential = nil
	}

	m.TaskName = model.TaskName

	m.OriginalTaskId = model.OriginalTaskId

	return
}
