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

// DropSqlsInSqlTuningSetDetails Drops the selected list of Sql statements from the current Sql tuning set.
// The basicFilter parameter specifies the Sql predicate to filter the Sql from the Sql tuning set defined on attributes of the SQLSET_ROW.
// If a valid filter criteria is specified, then, Sql statements matching this filter criteria will be deleted from the current Sql tuning set.
// If filter criteria is not specified, then, all Sql statements will be deleted from the current Sql tuning set.
// It takes either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type DropSqlsInSqlTuningSetDetails struct {

	// The name of the Sql tuning set.
	Name *string `mandatory:"true" json:"name"`

	CredentialDetails SqlTuningSetAdminCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`

	// Flag to indicate whether to drop the Sql statements or just display the plsql used to drop the Sql statements.
	ShowSqlOnly *int `mandatory:"false" json:"showSqlOnly"`

	// The owner of the Sql tuning set.
	Owner *string `mandatory:"false" json:"owner"`

	// Specifies the Sql predicate to filter the Sql from the Sql tuning set defined on attributes of the SQLSET_ROW.
	// User could use any combination of the following columns with appropriate values as Sql predicate
	// Refer to the documentation https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_SQLTUNE.html#GUID-1F4AFB03-7B29-46FC-B3F2-CB01EC36326C
	BasicFilter *string `mandatory:"false" json:"basicFilter"`
}

func (m DropSqlsInSqlTuningSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DropSqlsInSqlTuningSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DropSqlsInSqlTuningSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails  sqltuningsetadmincredentialdetails `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails          `json:"databaseCredential"`
		ShowSqlOnly        *int                               `json:"showSqlOnly"`
		Owner              *string                            `json:"owner"`
		BasicFilter        *string                            `json:"basicFilter"`
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

	m.ShowSqlOnly = model.ShowSqlOnly

	m.Owner = model.Owner

	m.BasicFilter = model.BasicFilter

	m.Name = model.Name

	return
}
