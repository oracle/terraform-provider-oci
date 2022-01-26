// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// VaultUsage The representation of VaultUsage
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
