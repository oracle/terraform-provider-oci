// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyEncryptionCredentialDetails Details for Credentials using key encryption.
type KeyEncryptionCredentialDetails struct {

	// The value corresponding to the credential
	Value *string `mandatory:"true" json:"value"`

	// OCID for the Vault Key that will be used to encrypt/decrypt the value given.
	KeyId *string `mandatory:"true" json:"keyId"`

	// OCID for the Vault that will be used to fetch key to encrypt/decrypt the value given.
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
