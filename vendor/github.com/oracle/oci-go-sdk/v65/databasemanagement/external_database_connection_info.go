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

// ExternalDatabaseConnectionInfo The details required to connect to an external Oracle Database.
// It takes either connectionCredentials or databaseCredential. It's recommended to provide databaseCredential
type ExternalDatabaseConnectionInfo struct {
	ConnectionString *DatabaseConnectionString `mandatory:"true" json:"connectionString"`

	ConnectionCredentials DatabaseConnectionCredentials `mandatory:"false" json:"connectionCredentials"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`
}

func (m ExternalDatabaseConnectionInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseConnectionInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDatabaseConnectionInfo) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseConnectionInfo ExternalDatabaseConnectionInfo
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeExternalDatabaseConnectionInfo
	}{
		"DATABASE",
		(MarshalTypeExternalDatabaseConnectionInfo)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseConnectionInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionCredentials databaseconnectioncredentials `json:"connectionCredentials"`
		DatabaseCredential    databasecredentialdetails     `json:"databaseCredential"`
		ConnectionString      *DatabaseConnectionString     `json:"connectionString"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConnectionCredentials.UnmarshalPolymorphicJSON(model.ConnectionCredentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionCredentials = nn.(DatabaseConnectionCredentials)
	} else {
		m.ConnectionCredentials = nil
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

	m.ConnectionString = model.ConnectionString

	return
}
