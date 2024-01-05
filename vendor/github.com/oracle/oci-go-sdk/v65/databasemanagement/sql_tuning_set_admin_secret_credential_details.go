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

// SqlTuningSetAdminSecretCredentialDetails User provides a secret OCID, which will be used to retrieve the password to connect to the database.
type SqlTuningSetAdminSecretCredentialDetails struct {

	// The user to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Secret
	// where the database password is stored.
	SecretId *string `mandatory:"true" json:"secretId"`

	// The role of the database user.
	Role SqlTuningSetAdminCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
}

// GetUsername returns Username
func (m SqlTuningSetAdminSecretCredentialDetails) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m SqlTuningSetAdminSecretCredentialDetails) GetRole() SqlTuningSetAdminCredentialDetailsRoleEnum {
	return m.Role
}

func (m SqlTuningSetAdminSecretCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningSetAdminSecretCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningSetAdminCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetSqlTuningSetAdminCredentialDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlTuningSetAdminSecretCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlTuningSetAdminSecretCredentialDetails SqlTuningSetAdminSecretCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"sqlTuningSetAdminCredentialType"`
		MarshalTypeSqlTuningSetAdminSecretCredentialDetails
	}{
		"SECRET",
		(MarshalTypeSqlTuningSetAdminSecretCredentialDetails)(m),
	}

	return json.Marshal(&s)
}
