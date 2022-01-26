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

// SqlTuningTaskSecretCredentialDetails User provides a secret OCID, which will be used to retrieve the password to connect to the database.
type SqlTuningTaskSecretCredentialDetails struct {

	// The user to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret
	// where the database password is stored.
	PasswordSecretId *string `mandatory:"true" json:"passwordSecretId"`

	// The role of the database user.
	Role SqlTuningTaskCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
}

//GetUsername returns Username
func (m SqlTuningTaskSecretCredentialDetails) GetUsername() *string {
	return m.Username
}

//GetRole returns Role
func (m SqlTuningTaskSecretCredentialDetails) GetRole() SqlTuningTaskCredentialDetailsRoleEnum {
	return m.Role
}

func (m SqlTuningTaskSecretCredentialDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SqlTuningTaskSecretCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlTuningTaskSecretCredentialDetails SqlTuningTaskSecretCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"sqlTuningTaskCredentialType"`
		MarshalTypeSqlTuningTaskSecretCredentialDetails
	}{
		"SECRET",
		(MarshalTypeSqlTuningTaskSecretCredentialDetails)(m),
	}

	return json.Marshal(&s)
}
