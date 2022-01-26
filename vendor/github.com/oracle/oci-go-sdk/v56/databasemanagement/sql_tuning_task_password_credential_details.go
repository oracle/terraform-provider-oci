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

// SqlTuningTaskPasswordCredentialDetails User provides a password to be used to connect to the database.
type SqlTuningTaskPasswordCredentialDetails struct {

	// The user to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The database user's password encoded using BASE64 scheme.
	Password *string `mandatory:"true" json:"password"`

	// The role of the database user.
	Role SqlTuningTaskCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
}

//GetUsername returns Username
func (m SqlTuningTaskPasswordCredentialDetails) GetUsername() *string {
	return m.Username
}

//GetRole returns Role
func (m SqlTuningTaskPasswordCredentialDetails) GetRole() SqlTuningTaskCredentialDetailsRoleEnum {
	return m.Role
}

func (m SqlTuningTaskPasswordCredentialDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SqlTuningTaskPasswordCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlTuningTaskPasswordCredentialDetails SqlTuningTaskPasswordCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"sqlTuningTaskCredentialType"`
		MarshalTypeSqlTuningTaskPasswordCredentialDetails
	}{
		"PASSWORD",
		(MarshalTypeSqlTuningTaskPasswordCredentialDetails)(m),
	}

	return json.Marshal(&s)
}
