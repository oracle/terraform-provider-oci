// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CrossConnectStatus The status of the cross-connect.
type CrossConnectStatus struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cross-connect.
	CrossConnectId *string `mandatory:"true" json:"crossConnectId"`

	// Indicates whether Oracle's side of the interface is up or down.
	InterfaceState CrossConnectStatusInterfaceStateEnum `mandatory:"false" json:"interfaceState,omitempty"`

	// The light level of the cross-connect (in dBm).
	// Example: `14.0`
	LightLevelIndBm *float32 `mandatory:"false" json:"lightLevelIndBm"`

	// Status indicator corresponding to the light level.
	//   * **NO_LIGHT:** No measurable light
	//   * **LOW_WARN:** There's measurable light but it's too low
	//   * **HIGH_WARN:** Light level is too high
	//   * **BAD:** There's measurable light but the signal-to-noise ratio is bad
	//   * **GOOD:** Good light level
	LightLevelIndicator CrossConnectStatusLightLevelIndicatorEnum `mandatory:"false" json:"lightLevelIndicator,omitempty"`

	// Encryption status of this cross connect.
	// Possible values:
	// * **UP:** Traffic is encrypted over this cross-connect
	// * **DOWN:** Traffic is not encrypted over this cross-connect
	// * **CIPHER_MISMATCH:** The MACsec encryption cipher doesn't match the cipher on the CPE
	// * **CKN_MISMATCH:** The MACsec Connectivity association Key Name (CKN) doesn't match the CKN on the CPE
	// * **CAK_MISMATCH:** The MACsec Connectivity Association Key (CAK) doesn't match the CAK on the CPE
	EncryptionStatus CrossConnectStatusEncryptionStatusEnum `mandatory:"false" json:"encryptionStatus,omitempty"`

	// The light levels of the cross-connect (in dBm).
	// Example: `[14.0, -14.0, 2.1, -10.1]`
	LightLevelsInDBm []float32 `mandatory:"false" json:"lightLevelsInDBm"`
}

func (m CrossConnectStatus) String() string {
	return common.PointerString(m)
}

// CrossConnectStatusInterfaceStateEnum Enum with underlying type: string
type CrossConnectStatusInterfaceStateEnum string

// Set of constants representing the allowable values for CrossConnectStatusInterfaceStateEnum
const (
	CrossConnectStatusInterfaceStateUp   CrossConnectStatusInterfaceStateEnum = "UP"
	CrossConnectStatusInterfaceStateDown CrossConnectStatusInterfaceStateEnum = "DOWN"
)

var mappingCrossConnectStatusInterfaceState = map[string]CrossConnectStatusInterfaceStateEnum{
	"UP":   CrossConnectStatusInterfaceStateUp,
	"DOWN": CrossConnectStatusInterfaceStateDown,
}

// GetCrossConnectStatusInterfaceStateEnumValues Enumerates the set of values for CrossConnectStatusInterfaceStateEnum
func GetCrossConnectStatusInterfaceStateEnumValues() []CrossConnectStatusInterfaceStateEnum {
	values := make([]CrossConnectStatusInterfaceStateEnum, 0)
	for _, v := range mappingCrossConnectStatusInterfaceState {
		values = append(values, v)
	}
	return values
}

// CrossConnectStatusLightLevelIndicatorEnum Enum with underlying type: string
type CrossConnectStatusLightLevelIndicatorEnum string

// Set of constants representing the allowable values for CrossConnectStatusLightLevelIndicatorEnum
const (
	CrossConnectStatusLightLevelIndicatorNoLight  CrossConnectStatusLightLevelIndicatorEnum = "NO_LIGHT"
	CrossConnectStatusLightLevelIndicatorLowWarn  CrossConnectStatusLightLevelIndicatorEnum = "LOW_WARN"
	CrossConnectStatusLightLevelIndicatorHighWarn CrossConnectStatusLightLevelIndicatorEnum = "HIGH_WARN"
	CrossConnectStatusLightLevelIndicatorBad      CrossConnectStatusLightLevelIndicatorEnum = "BAD"
	CrossConnectStatusLightLevelIndicatorGood     CrossConnectStatusLightLevelIndicatorEnum = "GOOD"
)

var mappingCrossConnectStatusLightLevelIndicator = map[string]CrossConnectStatusLightLevelIndicatorEnum{
	"NO_LIGHT":  CrossConnectStatusLightLevelIndicatorNoLight,
	"LOW_WARN":  CrossConnectStatusLightLevelIndicatorLowWarn,
	"HIGH_WARN": CrossConnectStatusLightLevelIndicatorHighWarn,
	"BAD":       CrossConnectStatusLightLevelIndicatorBad,
	"GOOD":      CrossConnectStatusLightLevelIndicatorGood,
}

// GetCrossConnectStatusLightLevelIndicatorEnumValues Enumerates the set of values for CrossConnectStatusLightLevelIndicatorEnum
func GetCrossConnectStatusLightLevelIndicatorEnumValues() []CrossConnectStatusLightLevelIndicatorEnum {
	values := make([]CrossConnectStatusLightLevelIndicatorEnum, 0)
	for _, v := range mappingCrossConnectStatusLightLevelIndicator {
		values = append(values, v)
	}
	return values
}

// CrossConnectStatusEncryptionStatusEnum Enum with underlying type: string
type CrossConnectStatusEncryptionStatusEnum string

// Set of constants representing the allowable values for CrossConnectStatusEncryptionStatusEnum
const (
	CrossConnectStatusEncryptionStatusUp             CrossConnectStatusEncryptionStatusEnum = "UP"
	CrossConnectStatusEncryptionStatusDown           CrossConnectStatusEncryptionStatusEnum = "DOWN"
	CrossConnectStatusEncryptionStatusCipherMismatch CrossConnectStatusEncryptionStatusEnum = "CIPHER_MISMATCH"
	CrossConnectStatusEncryptionStatusCknMismatch    CrossConnectStatusEncryptionStatusEnum = "CKN_MISMATCH"
	CrossConnectStatusEncryptionStatusCakMismatch    CrossConnectStatusEncryptionStatusEnum = "CAK_MISMATCH"
)

var mappingCrossConnectStatusEncryptionStatus = map[string]CrossConnectStatusEncryptionStatusEnum{
	"UP":              CrossConnectStatusEncryptionStatusUp,
	"DOWN":            CrossConnectStatusEncryptionStatusDown,
	"CIPHER_MISMATCH": CrossConnectStatusEncryptionStatusCipherMismatch,
	"CKN_MISMATCH":    CrossConnectStatusEncryptionStatusCknMismatch,
	"CAK_MISMATCH":    CrossConnectStatusEncryptionStatusCakMismatch,
}

// GetCrossConnectStatusEncryptionStatusEnumValues Enumerates the set of values for CrossConnectStatusEncryptionStatusEnum
func GetCrossConnectStatusEncryptionStatusEnumValues() []CrossConnectStatusEncryptionStatusEnum {
	values := make([]CrossConnectStatusEncryptionStatusEnum, 0)
	for _, v := range mappingCrossConnectStatusEncryptionStatus {
		values = append(values, v)
	}
	return values
}
