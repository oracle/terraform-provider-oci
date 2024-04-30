// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PreferredCredentialSummary The summary of preferred credentials.
type PreferredCredentialSummary struct {

	// The name of the preferred credential.
	CredentialName *string `mandatory:"true" json:"credentialName"`

	// The status of the preferred credential.
	Status PreferredCredentialStatusEnum `mandatory:"true" json:"status"`

	// Indicates whether the preferred credential is accessible.
	IsAccessible *bool `mandatory:"true" json:"isAccessible"`

	// The user name used to connect to the database.
	UserName *string `mandatory:"false" json:"userName"`

	// The role of the database user.
	Role RoleEnum `mandatory:"false" json:"role,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Vault service secret that contains the database user password.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Named Credential that contains the database user password metadata.
	NamedCredentialId *string `mandatory:"false" json:"namedCredentialId"`
}

func (m PreferredCredentialSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PreferredCredentialSummary) ValidateEnumValue() (bool, error) {
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
