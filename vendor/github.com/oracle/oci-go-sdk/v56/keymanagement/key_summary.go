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

// KeySummary The representation of KeySummary
type KeySummary struct {

	// The OCID of the compartment that contains the key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the key. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the key.
	Id *string `mandatory:"true" json:"id"`

	// The key's current lifecycle state.
	// Example: `ENABLED`
	LifecycleState KeySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the key was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the vault that contains the key.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed.
	// A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside
	// the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists
	// on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,
	// a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported.
	ProtectionMode KeySummaryProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The algorithm used by a key's key versions to encrypt or decrypt data.
	Algorithm KeySummaryAlgorithmEnum `mandatory:"false" json:"algorithm,omitempty"`
}

func (m KeySummary) String() string {
	return common.PointerString(m)
}

// KeySummaryLifecycleStateEnum Enum with underlying type: string
type KeySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for KeySummaryLifecycleStateEnum
const (
	KeySummaryLifecycleStateCreating           KeySummaryLifecycleStateEnum = "CREATING"
	KeySummaryLifecycleStateEnabling           KeySummaryLifecycleStateEnum = "ENABLING"
	KeySummaryLifecycleStateEnabled            KeySummaryLifecycleStateEnum = "ENABLED"
	KeySummaryLifecycleStateDisabling          KeySummaryLifecycleStateEnum = "DISABLING"
	KeySummaryLifecycleStateDisabled           KeySummaryLifecycleStateEnum = "DISABLED"
	KeySummaryLifecycleStateDeleting           KeySummaryLifecycleStateEnum = "DELETING"
	KeySummaryLifecycleStateDeleted            KeySummaryLifecycleStateEnum = "DELETED"
	KeySummaryLifecycleStatePendingDeletion    KeySummaryLifecycleStateEnum = "PENDING_DELETION"
	KeySummaryLifecycleStateSchedulingDeletion KeySummaryLifecycleStateEnum = "SCHEDULING_DELETION"
	KeySummaryLifecycleStateCancellingDeletion KeySummaryLifecycleStateEnum = "CANCELLING_DELETION"
	KeySummaryLifecycleStateUpdating           KeySummaryLifecycleStateEnum = "UPDATING"
	KeySummaryLifecycleStateBackupInProgress   KeySummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	KeySummaryLifecycleStateRestoring          KeySummaryLifecycleStateEnum = "RESTORING"
)

var mappingKeySummaryLifecycleState = map[string]KeySummaryLifecycleStateEnum{
	"CREATING":            KeySummaryLifecycleStateCreating,
	"ENABLING":            KeySummaryLifecycleStateEnabling,
	"ENABLED":             KeySummaryLifecycleStateEnabled,
	"DISABLING":           KeySummaryLifecycleStateDisabling,
	"DISABLED":            KeySummaryLifecycleStateDisabled,
	"DELETING":            KeySummaryLifecycleStateDeleting,
	"DELETED":             KeySummaryLifecycleStateDeleted,
	"PENDING_DELETION":    KeySummaryLifecycleStatePendingDeletion,
	"SCHEDULING_DELETION": KeySummaryLifecycleStateSchedulingDeletion,
	"CANCELLING_DELETION": KeySummaryLifecycleStateCancellingDeletion,
	"UPDATING":            KeySummaryLifecycleStateUpdating,
	"BACKUP_IN_PROGRESS":  KeySummaryLifecycleStateBackupInProgress,
	"RESTORING":           KeySummaryLifecycleStateRestoring,
}

// GetKeySummaryLifecycleStateEnumValues Enumerates the set of values for KeySummaryLifecycleStateEnum
func GetKeySummaryLifecycleStateEnumValues() []KeySummaryLifecycleStateEnum {
	values := make([]KeySummaryLifecycleStateEnum, 0)
	for _, v := range mappingKeySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// KeySummaryProtectionModeEnum Enum with underlying type: string
type KeySummaryProtectionModeEnum string

// Set of constants representing the allowable values for KeySummaryProtectionModeEnum
const (
	KeySummaryProtectionModeHsm      KeySummaryProtectionModeEnum = "HSM"
	KeySummaryProtectionModeSoftware KeySummaryProtectionModeEnum = "SOFTWARE"
)

var mappingKeySummaryProtectionMode = map[string]KeySummaryProtectionModeEnum{
	"HSM":      KeySummaryProtectionModeHsm,
	"SOFTWARE": KeySummaryProtectionModeSoftware,
}

// GetKeySummaryProtectionModeEnumValues Enumerates the set of values for KeySummaryProtectionModeEnum
func GetKeySummaryProtectionModeEnumValues() []KeySummaryProtectionModeEnum {
	values := make([]KeySummaryProtectionModeEnum, 0)
	for _, v := range mappingKeySummaryProtectionMode {
		values = append(values, v)
	}
	return values
}

// KeySummaryAlgorithmEnum Enum with underlying type: string
type KeySummaryAlgorithmEnum string

// Set of constants representing the allowable values for KeySummaryAlgorithmEnum
const (
	KeySummaryAlgorithmAes   KeySummaryAlgorithmEnum = "AES"
	KeySummaryAlgorithmRsa   KeySummaryAlgorithmEnum = "RSA"
	KeySummaryAlgorithmEcdsa KeySummaryAlgorithmEnum = "ECDSA"
)

var mappingKeySummaryAlgorithm = map[string]KeySummaryAlgorithmEnum{
	"AES":   KeySummaryAlgorithmAes,
	"RSA":   KeySummaryAlgorithmRsa,
	"ECDSA": KeySummaryAlgorithmEcdsa,
}

// GetKeySummaryAlgorithmEnumValues Enumerates the set of values for KeySummaryAlgorithmEnum
func GetKeySummaryAlgorithmEnumValues() []KeySummaryAlgorithmEnum {
	values := make([]KeySummaryAlgorithmEnum, 0)
	for _, v := range mappingKeySummaryAlgorithm {
		values = append(values, v)
	}
	return values
}
