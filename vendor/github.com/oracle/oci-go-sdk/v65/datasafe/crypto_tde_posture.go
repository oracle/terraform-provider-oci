// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoTdePosture Transparent data encryption posture details.
type CryptoTdePosture struct {

	// TDE enablement status.
	Status CryptoFeatureStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Configured TDE encryption algorithm.
	EncryptionConfigured []string `mandatory:"false" json:"encryptionConfigured"`

	// Observed redo log encryption algorithm.
	RedoEncryptionObserved *string `mandatory:"false" json:"redoEncryptionObserved"`

	// Observed DB credentials encryption algorithm.
	DbCredentialsEncryptionObserved *string `mandatory:"false" json:"dbCredentialsEncryptionObserved"`

	// Configured TDE integrity-related setting.
	IntegrityConfigured []string `mandatory:"false" json:"integrityConfigured"`

	// Number of encrypted tablespaces detected.
	EncryptedTablespacesCount *int `mandatory:"false" json:"encryptedTablespacesCount"`

	// Number of unencrypted tablespaces detected.
	UnencryptedTablespacesCount *int `mandatory:"false" json:"unencryptedTablespacesCount"`

	// The observed TDE master key identifier.
	MasterKeyId *string `mandatory:"false" json:"masterKeyId"`

	// The observed wallet location for TDE keys.
	WalletLocation *string `mandatory:"false" json:"walletLocation"`

	// The observed encryption algorithm used by the master key.
	MasterKeyEncryptionAlgorithm *string `mandatory:"false" json:"masterKeyEncryptionAlgorithm"`

	// The last observed rotation time for the TDE master key, in RFC3339 format.
	TimeMasterKeyLastRotation *common.SDKTime `mandatory:"false" json:"timeMasterKeyLastRotation"`

	// The observed TDE key store type.
	KeyStoreType CryptoKeystoreTypeEnum `mandatory:"false" json:"keyStoreType,omitempty"`

	// The observed TDE key cache status.
	KeyCacheStatus CryptoKeyCacheStatusEnum `mandatory:"false" json:"keyCacheStatus,omitempty"`

	// FIPS mode configured for TDE when the target uses legacy per-feature FIPS configuration.
	FipsModeConfigured CryptoPostureFipsModeConfiguredEnum `mandatory:"false" json:"fipsModeConfigured,omitempty"`

	// Quantum-readiness classification for TDE posture.
	QuantumReadiness CryptoQuantumReadinessEnum `mandatory:"false" json:"quantumReadiness,omitempty"`
}

func (m CryptoTdePosture) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoTdePosture) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoFeatureStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoFeatureStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoKeystoreTypeEnum(string(m.KeyStoreType)); !ok && m.KeyStoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyStoreType: %s. Supported values are: %s.", m.KeyStoreType, strings.Join(GetCryptoKeystoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoKeyCacheStatusEnum(string(m.KeyCacheStatus)); !ok && m.KeyCacheStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyCacheStatus: %s. Supported values are: %s.", m.KeyCacheStatus, strings.Join(GetCryptoKeyCacheStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoPostureFipsModeConfiguredEnum(string(m.FipsModeConfigured)); !ok && m.FipsModeConfigured != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FipsModeConfigured: %s. Supported values are: %s.", m.FipsModeConfigured, strings.Join(GetCryptoPostureFipsModeConfiguredEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoQuantumReadinessEnum(string(m.QuantumReadiness)); !ok && m.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", m.QuantumReadiness, strings.Join(GetCryptoQuantumReadinessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
