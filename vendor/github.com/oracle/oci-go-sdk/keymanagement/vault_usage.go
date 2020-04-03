// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Key Management Service API
//
// API for managing and performing operations with keys and vaults.
//

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// VaultUsage The representation of VaultUsage
type VaultUsage struct {

	// The number of keys in this vault, across all compartments, excluding keys in a `DELETED` state.
	KeyCount *int `mandatory:"true" json:"keyCount"`

	// The number of key versions in this vault, across all compartments, excluding key versions in a `DELETED` state.
	KeyVersionCount *int `mandatory:"true" json:"keyVersionCount"`
}

func (m VaultUsage) String() string {
	return common.PointerString(m)
}
