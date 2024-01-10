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

// TablespaceAdminPasswordCredentialDetails User provides a password to be used to connect to the database.
type TablespaceAdminPasswordCredentialDetails struct {

	// The user to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The database user's password encoded using BASE64 scheme.
	Password *string `mandatory:"true" json:"password"`

	// The role of the database user.
	Role TablespaceAdminCredentialDetailsRoleEnum `mandatory:"true" json:"role"`
}

// GetUsername returns Username
func (m TablespaceAdminPasswordCredentialDetails) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m TablespaceAdminPasswordCredentialDetails) GetRole() TablespaceAdminCredentialDetailsRoleEnum {
	return m.Role
}

func (m TablespaceAdminPasswordCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TablespaceAdminPasswordCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTablespaceAdminCredentialDetailsRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetTablespaceAdminCredentialDetailsRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TablespaceAdminPasswordCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTablespaceAdminPasswordCredentialDetails TablespaceAdminPasswordCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"tablespaceAdminCredentialType"`
		MarshalTypeTablespaceAdminPasswordCredentialDetails
	}{
		"PASSWORD",
		(MarshalTypeTablespaceAdminPasswordCredentialDetails)(m),
	}

	return json.Marshal(&s)
}
