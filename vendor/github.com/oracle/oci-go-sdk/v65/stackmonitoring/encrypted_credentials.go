// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EncryptedCredentials Encrypted credentials [indicated by the type property in CredentialStore].
type EncryptedCredentials struct {

	// The master key should be created in OCI Vault owned by the client of this API.
	// The user should have permission to access the vault key.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The credential properties list. Credential property values will be encrypted format.
	Properties []CredentialProperty `mandatory:"true" json:"properties"`

	// The source type and source name combination, delimited with (.) separator.
	// {source type}.{source name} and source type max char limit is 63.
	Source *string `mandatory:"false" json:"source"`

	// The name of the credential, within the context of the source.
	Name *string `mandatory:"false" json:"name"`

	// The type of the credential ( ex. JMXCreds,DBCreds).
	Type *string `mandatory:"false" json:"type"`

	// The user-specified textual description of the credential.
	Description *string `mandatory:"false" json:"description"`
}

// GetSource returns Source
func (m EncryptedCredentials) GetSource() *string {
	return m.Source
}

// GetName returns Name
func (m EncryptedCredentials) GetName() *string {
	return m.Name
}

// GetType returns Type
func (m EncryptedCredentials) GetType() *string {
	return m.Type
}

// GetDescription returns Description
func (m EncryptedCredentials) GetDescription() *string {
	return m.Description
}

func (m EncryptedCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EncryptedCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EncryptedCredentials) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEncryptedCredentials EncryptedCredentials
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeEncryptedCredentials
	}{
		"ENCRYPTED",
		(MarshalTypeEncryptedCredentials)(m),
	}

	return json.Marshal(&s)
}
