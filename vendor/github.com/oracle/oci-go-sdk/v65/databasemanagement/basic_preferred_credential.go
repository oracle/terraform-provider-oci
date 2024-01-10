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

// BasicPreferredCredential The details of the 'BASIC' preferred credential.
type BasicPreferredCredential struct {

	// The name of the preferred credential.
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// Indicates whether the preferred credential is accessible.
	IsAccessible *bool `mandatory:"false" json:"isAccessible"`

	// The user name used to connect to the database.
	UserName *string `mandatory:"false" json:"userName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The status of the preferred credential.
	Status PreferredCredentialStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The role of the database user.
	Role RoleEnum `mandatory:"false" json:"role,omitempty"`
}

// GetCredentialName returns CredentialName
func (m BasicPreferredCredential) GetCredentialName() *string {
	return m.CredentialName
}

// GetStatus returns Status
func (m BasicPreferredCredential) GetStatus() PreferredCredentialStatusEnum {
	return m.Status
}

// GetIsAccessible returns IsAccessible
func (m BasicPreferredCredential) GetIsAccessible() *bool {
	return m.IsAccessible
}

func (m BasicPreferredCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicPreferredCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPreferredCredentialStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPreferredCredentialStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicPreferredCredential) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicPreferredCredential BasicPreferredCredential
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeBasicPreferredCredential
	}{
		"BASIC",
		(MarshalTypeBasicPreferredCredential)(m),
	}

	return json.Marshal(&s)
}
