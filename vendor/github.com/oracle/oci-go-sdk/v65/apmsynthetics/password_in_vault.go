// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PasswordInVault Vault secret OCID for password that can be used with monitor Resource Principal.
// Example, ocid1.vaultsecret.oc1.iad.amaaaaaagpihjxqadwyc4kjhpeis2bylhzmp5r2si6mz2h4eujevnmf3zoca.
type PasswordInVault struct {

	// Vault secret OCID.
	VaultSecretId *string `mandatory:"true" json:"vaultSecretId"`
}

func (m PasswordInVault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PasswordInVault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PasswordInVault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePasswordInVault PasswordInVault
	s := struct {
		DiscriminatorParam string `json:"passwordType"`
		MarshalTypePasswordInVault
	}{
		"VAULT_SECRET_ID",
		(MarshalTypePasswordInVault)(m),
	}

	return json.Marshal(&s)
}
