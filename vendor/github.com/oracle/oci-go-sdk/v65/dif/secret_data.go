// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecretData Maps a kubernetes secret data key to a value sourced from Oci Vault. At deploy time, the platform fetches the secret value by OCID and injects it under the specified key in the secrets data. No plaintext secret values are required in the payload.
type SecretData struct {

	// Data key in the kubernetes secret.
	Key *string `mandatory:"true" json:"key"`

	// OCID of the Oci vault secret that provides the value for this key. The latest active secret version is used at deploy time unless otherwise configured.
	SecretId *string `mandatory:"true" json:"secretId"`
}

func (m SecretData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecretData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
