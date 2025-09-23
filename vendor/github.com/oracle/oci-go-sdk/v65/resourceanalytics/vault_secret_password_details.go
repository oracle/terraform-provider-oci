// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VaultSecretPasswordDetails Secret details for the ADW Admin password.
type VaultSecretPasswordDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the vault secret to use as the ADW admin password.
	SecretId *string `mandatory:"true" json:"secretId"`
}

func (m VaultSecretPasswordDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultSecretPasswordDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VaultSecretPasswordDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVaultSecretPasswordDetails VaultSecretPasswordDetails
	s := struct {
		DiscriminatorParam string `json:"passwordType"`
		MarshalTypeVaultSecretPasswordDetails
	}{
		"VAULT_SECRET",
		(MarshalTypeVaultSecretPasswordDetails)(m),
	}

	return json.Marshal(&s)
}
