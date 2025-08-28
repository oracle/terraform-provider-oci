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

// CloudDatabaseConnectionInfo The details required to connect to a cloud Oracle Database.
type CloudDatabaseConnectionInfo struct {
	ConnectionString *DatabaseConnectionString `mandatory:"true" json:"connectionString"`

	ConnectionCredentials DatabaseConnectionCredentials `mandatory:"false" json:"connectionCredentials"`
}

func (m CloudDatabaseConnectionInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudDatabaseConnectionInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudDatabaseConnectionInfo) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudDatabaseConnectionInfo CloudDatabaseConnectionInfo
	s := struct {
		DiscriminatorParam string `json:"componentType"`
		MarshalTypeCloudDatabaseConnectionInfo
	}{
		"DATABASE",
		(MarshalTypeCloudDatabaseConnectionInfo)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CloudDatabaseConnectionInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionCredentials databaseconnectioncredentials `json:"connectionCredentials"`
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

	m.ConnectionString = model.ConnectionString

	return
}
