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

// ImportKeyDetails The representation of ImportKeyDetails
type ImportKeyDetails struct {

	// The OCID of the compartment that contains this key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the key. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	KeyShape *KeyShape `mandatory:"true" json:"keyShape"`

	WrappedImportKey *WrappedImportKey `mandatory:"true" json:"wrappedImportKey"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "foo-value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed.
	// A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside
	// the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists
	// on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,
	// a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported.
	ProtectionMode ImportKeyDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`
}

func (m ImportKeyDetails) String() string {
	return common.PointerString(m)
}

// ImportKeyDetailsProtectionModeEnum Enum with underlying type: string
type ImportKeyDetailsProtectionModeEnum string

// Set of constants representing the allowable values for ImportKeyDetailsProtectionModeEnum
const (
	ImportKeyDetailsProtectionModeHsm      ImportKeyDetailsProtectionModeEnum = "HSM"
	ImportKeyDetailsProtectionModeSoftware ImportKeyDetailsProtectionModeEnum = "SOFTWARE"
)

var mappingImportKeyDetailsProtectionMode = map[string]ImportKeyDetailsProtectionModeEnum{
	"HSM":      ImportKeyDetailsProtectionModeHsm,
	"SOFTWARE": ImportKeyDetailsProtectionModeSoftware,
}

// GetImportKeyDetailsProtectionModeEnumValues Enumerates the set of values for ImportKeyDetailsProtectionModeEnum
func GetImportKeyDetailsProtectionModeEnumValues() []ImportKeyDetailsProtectionModeEnum {
	values := make([]ImportKeyDetailsProtectionModeEnum, 0)
	for _, v := range mappingImportKeyDetailsProtectionMode {
		values = append(values, v)
	}
	return values
}
