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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SqlTuningTaskSecretCredentialDetails The OCID of the Secret provided by the user to retrieve the password to connect to the database.
type SqlTuningTaskSecretCredentialDetails struct {

	// The user name used to connect to the database.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningTaskSecretCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningTaskCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlTuningTaskCredentialDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
