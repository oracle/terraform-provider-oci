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

// ManagedDatabasePasswordCredential User provides a password to be used to connect to the database.
type ManagedDatabasePasswordCredential struct {

	// The user name used to connect to the database.
	Username *string `mandatory:"true" json:"username"`

	// The database user's password encoded using BASE64 scheme.
	Password *string `mandatory:"true" json:"password"`

	// The role of the database user.
	Role ManagedDatabaseCredentialRoleEnum `mandatory:"true" json:"role"`
}

// GetUsername returns Username
func (m ManagedDatabasePasswordCredential) GetUsername() *string {
	return m.Username
}

// GetRole returns Role
func (m ManagedDatabasePasswordCredential) GetRole() ManagedDatabaseCredentialRoleEnum {
	return m.Role
}

func (m ManagedDatabasePasswordCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedDatabasePasswordCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingManagedDatabaseCredentialRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetManagedDatabaseCredentialRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedDatabasePasswordCredential) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedDatabasePasswordCredential ManagedDatabasePasswordCredential
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeManagedDatabasePasswordCredential
	}{
		"PASSWORD",
		(MarshalTypeManagedDatabasePasswordCredential)(m),
	}

	return json.Marshal(&s)
}
