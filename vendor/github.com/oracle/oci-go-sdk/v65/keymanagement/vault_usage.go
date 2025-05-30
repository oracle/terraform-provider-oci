// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VaultUsage The details of the number of Keys and KeyVersions usage in a Vault.
type VaultUsage struct {

	// The number of keys in this vault that persist on a hardware security module (HSM), across all compartments, excluding keys in a `DELETED` state.
	KeyCount *int `mandatory:"true" json:"keyCount"`

	// The number of key versions in this vault that persist on a hardware security module (HSM), across all compartments, excluding key versions in a `DELETED` state.
	KeyVersionCount *int `mandatory:"true" json:"keyVersionCount"`

	// The number of keys in this vault that persist on the server, across all compartments, excluding keys in a `DELETED` state.
	SoftwareKeyCount *int `mandatory:"false" json:"softwareKeyCount"`

	// The number of key versions in this vault that persist on the server, across all compartments, excluding key versions in a `DELETED` state.
	SoftwareKeyVersionCount *int `mandatory:"false" json:"softwareKeyVersionCount"`
}

func (m VaultUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VaultUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
