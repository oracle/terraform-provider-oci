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

// CryptoNnePosture Native network encryption posture details.
type CryptoNnePosture struct {

	// NNE enablement status.
	Status CryptoFeatureStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Indicates if weak NNE options are allowed.
	AreWeakOptionsAllowed CryptoObservedBooleanValueEnum `mandatory:"false" json:"areWeakOptionsAllowed,omitempty"`

	// Configured network encryption algorithm.
	EncryptionConfigured []string `mandatory:"false" json:"encryptionConfigured"`

	// NNE integrity algorithm(s).
	Integrity []string `mandatory:"false" json:"integrity"`

	// Observed server-side encryption requirement.
	ServerEncryption CryptoNnePostureServerEncryptionEnum `mandatory:"false" json:"serverEncryption,omitempty"`

	// NNE server integrity algorithm(s).
	ServerIntegrity []string `mandatory:"false" json:"serverIntegrity"`

	// Observed NNE key exchange setting.
	KeyExchange *string `mandatory:"false" json:"keyExchange"`

	// Quantum-readiness classification for NNE posture.
	QuantumReadiness CryptoQuantumReadinessEnum `mandatory:"false" json:"quantumReadiness,omitempty"`

	// FIPS mode configured for NNE when the target uses legacy per-feature FIPS configuration.
	FipsModeConfigured CryptoPostureFipsModeConfiguredEnum `mandatory:"false" json:"fipsModeConfigured,omitempty"`
}

func (m CryptoNnePosture) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoNnePosture) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoFeatureStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoFeatureStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoObservedBooleanValueEnum(string(m.AreWeakOptionsAllowed)); !ok && m.AreWeakOptionsAllowed != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AreWeakOptionsAllowed: %s. Supported values are: %s.", m.AreWeakOptionsAllowed, strings.Join(GetCryptoObservedBooleanValueEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoNnePostureServerEncryptionEnum(string(m.ServerEncryption)); !ok && m.ServerEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServerEncryption: %s. Supported values are: %s.", m.ServerEncryption, strings.Join(GetCryptoNnePostureServerEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoQuantumReadinessEnum(string(m.QuantumReadiness)); !ok && m.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", m.QuantumReadiness, strings.Join(GetCryptoQuantumReadinessEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoPostureFipsModeConfiguredEnum(string(m.FipsModeConfigured)); !ok && m.FipsModeConfigured != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FipsModeConfigured: %s. Supported values are: %s.", m.FipsModeConfigured, strings.Join(GetCryptoPostureFipsModeConfiguredEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoNnePostureServerEncryptionEnum Enum with underlying type: string
type CryptoNnePostureServerEncryptionEnum string

// Set of constants representing the allowable values for CryptoNnePostureServerEncryptionEnum
const (
	CryptoNnePostureServerEncryptionRequired     CryptoNnePostureServerEncryptionEnum = "REQUIRED"
	CryptoNnePostureServerEncryptionRequested    CryptoNnePostureServerEncryptionEnum = "REQUESTED"
	CryptoNnePostureServerEncryptionAccepted     CryptoNnePostureServerEncryptionEnum = "ACCEPTED"
	CryptoNnePostureServerEncryptionRejected     CryptoNnePostureServerEncryptionEnum = "REJECTED"
	CryptoNnePostureServerEncryptionNotSupported CryptoNnePostureServerEncryptionEnum = "NOT_SUPPORTED"
)

var mappingCryptoNnePostureServerEncryptionEnum = map[string]CryptoNnePostureServerEncryptionEnum{
	"REQUIRED":      CryptoNnePostureServerEncryptionRequired,
	"REQUESTED":     CryptoNnePostureServerEncryptionRequested,
	"ACCEPTED":      CryptoNnePostureServerEncryptionAccepted,
	"REJECTED":      CryptoNnePostureServerEncryptionRejected,
	"NOT_SUPPORTED": CryptoNnePostureServerEncryptionNotSupported,
}

var mappingCryptoNnePostureServerEncryptionEnumLowerCase = map[string]CryptoNnePostureServerEncryptionEnum{
	"required":      CryptoNnePostureServerEncryptionRequired,
	"requested":     CryptoNnePostureServerEncryptionRequested,
	"accepted":      CryptoNnePostureServerEncryptionAccepted,
	"rejected":      CryptoNnePostureServerEncryptionRejected,
	"not_supported": CryptoNnePostureServerEncryptionNotSupported,
}

// GetCryptoNnePostureServerEncryptionEnumValues Enumerates the set of values for CryptoNnePostureServerEncryptionEnum
func GetCryptoNnePostureServerEncryptionEnumValues() []CryptoNnePostureServerEncryptionEnum {
	values := make([]CryptoNnePostureServerEncryptionEnum, 0)
	for _, v := range mappingCryptoNnePostureServerEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoNnePostureServerEncryptionEnumStringValues Enumerates the set of values in String for CryptoNnePostureServerEncryptionEnum
func GetCryptoNnePostureServerEncryptionEnumStringValues() []string {
	return []string{
		"REQUIRED",
		"REQUESTED",
		"ACCEPTED",
		"REJECTED",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoNnePostureServerEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoNnePostureServerEncryptionEnum(val string) (CryptoNnePostureServerEncryptionEnum, bool) {
	enum, ok := mappingCryptoNnePostureServerEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
