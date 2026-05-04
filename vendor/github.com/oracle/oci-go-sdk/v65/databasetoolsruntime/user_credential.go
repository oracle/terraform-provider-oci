// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserCredential Credential information
type UserCredential struct {

	// Name of the credential
	Key *string `mandatory:"true" json:"key"`

	// Name of the user that will be used to log in to the remote database or the remote or local operating system
	UserName *string `mandatory:"true" json:"userName"`

	// Indicates whether this credential is enabled (TRUE) or not (FALSE)
	Enabled *string `mandatory:"true" json:"enabled"`

	// Owner of the credential
	Owner *string `mandatory:"true" json:"owner"`

	// For a Windows target, the Windows domain to use when logging in
	WindowsDomain *string `mandatory:"false" json:"windowsDomain"`

	// Indicates whether this refers to a public synonym or not.
	KeyType CredentialKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	RelatedResource *CredentialRelatedResource `mandatory:"false" json:"relatedResource"`
}

func (m UserCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCredentialKeyTypeEnum(string(m.KeyType)); !ok && m.KeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyType: %s. Supported values are: %s.", m.KeyType, strings.Join(GetCredentialKeyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
