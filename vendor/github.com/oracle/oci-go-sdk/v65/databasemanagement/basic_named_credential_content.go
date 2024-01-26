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

// BasicNamedCredentialContent The details of the 'BASIC' named credential.
type BasicNamedCredentialContent struct {

	// The user name used to connect to the database.
	UserName *string `mandatory:"true" json:"userName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	PasswordSecretId *string `mandatory:"true" json:"passwordSecretId"`

	// The role of the database user.
	Role RoleEnum `mandatory:"true" json:"role"`

	// The mechanism used to access the password plain text value.
	PasswordSecretAccessMode PasswordSecretAccessModeEnum `mandatory:"true" json:"passwordSecretAccessMode"`
}

func (m BasicNamedCredentialContent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BasicNamedCredentialContent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPasswordSecretAccessModeEnum(string(m.PasswordSecretAccessMode)); !ok && m.PasswordSecretAccessMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PasswordSecretAccessMode: %s. Supported values are: %s.", m.PasswordSecretAccessMode, strings.Join(GetPasswordSecretAccessModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BasicNamedCredentialContent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBasicNamedCredentialContent BasicNamedCredentialContent
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeBasicNamedCredentialContent
	}{
		"BASIC",
		(MarshalTypeBasicNamedCredentialContent)(m),
	}

	return json.Marshal(&s)
}
