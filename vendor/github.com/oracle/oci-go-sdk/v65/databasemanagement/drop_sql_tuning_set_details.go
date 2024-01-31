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

// DropSqlTuningSetDetails The details required to drop a Sql tuning set.
// It takes either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type DropSqlTuningSetDetails struct {

	// A unique Sql tuning set name.
	Name *string `mandatory:"true" json:"name"`

	CredentialDetails SqlTuningSetAdminCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`

	// Owner of the Sql tuning set.
	Owner *string `mandatory:"false" json:"owner"`

	// Flag to indicate whether to drop  the Sql tuning set or just display the plsql used to drop Sql tuning set.
	ShowSqlOnly *int `mandatory:"false" json:"showSqlOnly"`
}

func (m DropSqlTuningSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DropSqlTuningSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DropSqlTuningSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails  sqltuningsetadmincredentialdetails `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails          `json:"databaseCredential"`
		Owner              *string                            `json:"owner"`
		ShowSqlOnly        *int                               `json:"showSqlOnly"`
		Name               *string                            `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningSetAdminCredentialDetails)
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

	m.Owner = model.Owner

	m.ShowSqlOnly = model.ShowSqlOnly

	m.Name = model.Name

	return
}
