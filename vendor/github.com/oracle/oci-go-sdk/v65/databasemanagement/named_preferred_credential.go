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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedPreferredCredential The details of the 'NAMED_CREDENTIAL' preferred credential.
type NamedPreferredCredential struct {

	// The name of the preferred credential.
	CredentialName *string `mandatory:"false" json:"credentialName"`

	// Indicates whether the preferred credential is accessible.
	IsAccessible *bool `mandatory:"false" json:"isAccessible"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Named Credential that contains the database user password metadata.
	NamedCredentialId *string `mandatory:"false" json:"namedCredentialId"`

	// The status of the preferred credential.
	Status PreferredCredentialStatusEnum `mandatory:"false" json:"status,omitempty"`
}

// GetCredentialName returns CredentialName
func (m NamedPreferredCredential) GetCredentialName() *string {
	return m.CredentialName
}

// GetStatus returns Status
func (m NamedPreferredCredential) GetStatus() PreferredCredentialStatusEnum {
	return m.Status
}

// GetIsAccessible returns IsAccessible
func (m NamedPreferredCredential) GetIsAccessible() *bool {
	return m.IsAccessible
}

func (m NamedPreferredCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedPreferredCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPreferredCredentialStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPreferredCredentialStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NamedPreferredCredential) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNamedPreferredCredential NamedPreferredCredential
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNamedPreferredCredential
	}{
		"NAMED_CREDENTIAL",
		(MarshalTypeNamedPreferredCredential)(m),
	}

	return json.Marshal(&s)
}
