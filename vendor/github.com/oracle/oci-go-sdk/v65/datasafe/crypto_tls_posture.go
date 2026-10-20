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

// CryptoTlsPosture TLS posture details.
type CryptoTlsPosture struct {

	// TLS enablement status.
	Status CryptoFeatureStatusEnum `mandatory:"false" json:"status,omitempty"`

	// TLS versions configured on target.
	Versions []string `mandatory:"false" json:"versions"`

	// TLS cipher suites configured on target.
	CipherSuitesConfigured []string `mandatory:"false" json:"cipherSuitesConfigured"`

	// Indicates if weak TLS cipher suites are allowed.
	AreWeakCipherSuitesAllowed CryptoObservedBooleanValueEnum `mandatory:"false" json:"areWeakCipherSuitesAllowed,omitempty"`

	// Whether TLS client authentication is configured.
	IsMtlsConfigured CryptoObservedBooleanValueEnum `mandatory:"false" json:"isMtlsConfigured,omitempty"`

	// Certificate revocation checking mode.
	RevocationMode *string `mandatory:"false" json:"revocationMode"`

	// FIPS mode configured for TLS when the target uses legacy per-feature FIPS configuration.
	FipsModeConfigured CryptoPostureFipsModeConfiguredEnum `mandatory:"false" json:"fipsModeConfigured,omitempty"`

	// TLS wallet location observed on target.
	WalletLocation *string `mandatory:"false" json:"walletLocation"`

	// Quantum-readiness classification for TLS posture.
	QuantumReadiness CryptoQuantumReadinessEnum `mandatory:"false" json:"quantumReadiness,omitempty"`
}

func (m CryptoTlsPosture) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoTlsPosture) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoFeatureStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoFeatureStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoObservedBooleanValueEnum(string(m.AreWeakCipherSuitesAllowed)); !ok && m.AreWeakCipherSuitesAllowed != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AreWeakCipherSuitesAllowed: %s. Supported values are: %s.", m.AreWeakCipherSuitesAllowed, strings.Join(GetCryptoObservedBooleanValueEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoObservedBooleanValueEnum(string(m.IsMtlsConfigured)); !ok && m.IsMtlsConfigured != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsMtlsConfigured: %s. Supported values are: %s.", m.IsMtlsConfigured, strings.Join(GetCryptoObservedBooleanValueEnumStringValues(), ",")))
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
