// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RotateSecretDetails Details of the secret to be rotated by the Secrets in Vault Service
type RotateSecretDetails struct {

	// The OCID of the Vault in which the secret exists.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the secret.
	SecretId *string `mandatory:"false" json:"secretId"`

	// The OCID of the compartment where you want to create the secret.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name of the secret. Avoid entering confidential information.
	SecretName *string `mandatory:"false" json:"secretName"`
}

func (m RotateSecretDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RotateSecretDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
