// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyEncryptionCredentialDetails Details for the Credentials that use key encryption.
type KeyEncryptionCredentialDetails struct {

	// The value corresponding to the credential.
	Value *string `mandatory:"true" json:"value"`

	// OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	KeyId *string `mandatory:"true" json:"keyId"`

	// OCID for the Vault that will be used to fetch the key to encrypt/decrypt the value given.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The Vault Key version.
	KeyVersion *string `mandatory:"false" json:"keyVersion"`
}

func (m KeyEncryptionCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KeyEncryptionCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KeyEncryptionCredentialDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKeyEncryptionCredentialDetails KeyEncryptionCredentialDetails
	s := struct {
		DiscriminatorParam string `json:"credentialType"`
		MarshalTypeKeyEncryptionCredentialDetails
	}{
		"KEY_ENCRYPTION",
		(MarshalTypeKeyEncryptionCredentialDetails)(m),
	}

	return json.Marshal(&s)
}
