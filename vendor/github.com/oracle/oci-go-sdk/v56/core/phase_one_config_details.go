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

// PhaseOneConfigDetails Phase 1 Configuration Details
type PhaseOneConfigDetails struct {

	// Indicates whether custom phase one configuration is enabled.
	IsCustomPhaseOneConfig *bool `mandatory:"false" json:"isCustomPhaseOneConfig"`

	// Phase one authentication algorithm supported during tunnel negotiation.
	AuthenticationAlgorithm PhaseOneConfigDetailsAuthenticationAlgorithmEnum `mandatory:"false" json:"authenticationAlgorithm,omitempty"`

	// Phase one encryption algorithm supported during tunnel negotiation.
	EncryptionAlgorithm PhaseOneConfigDetailsEncryptionAlgorithmEnum `mandatory:"false" json:"encryptionAlgorithm,omitempty"`

	// Phase One Diffie Hellman group supported during tunnel negotiation.
	DiffieHelmanGroup PhaseOneConfigDetailsDiffieHelmanGroupEnum `mandatory:"false" json:"diffieHelmanGroup,omitempty"`

	// IKE session key lifetime in seconds for IPSec phase one.
	LifetimeInSeconds *int `mandatory:"false" json:"lifetimeInSeconds"`
}

func (m PhaseOneConfigDetails) String() string {
	return common.PointerString(m)
}

// PhaseOneConfigDetailsAuthenticationAlgorithmEnum Enum with underlying type: string
type PhaseOneConfigDetailsAuthenticationAlgorithmEnum string

// Set of constants representing the allowable values for PhaseOneConfigDetailsAuthenticationAlgorithmEnum
const (
	PhaseOneConfigDetailsAuthenticationAlgorithmSha2384 PhaseOneConfigDetailsAuthenticationAlgorithmEnum = "SHA2_384"
	PhaseOneConfigDetailsAuthenticationAlgorithmSha2256 PhaseOneConfigDetailsAuthenticationAlgorithmEnum = "SHA2_256"
	PhaseOneConfigDetailsAuthenticationAlgorithmSha196  PhaseOneConfigDetailsAuthenticationAlgorithmEnum = "SHA1_96"
)

var mappingPhaseOneConfigDetailsAuthenticationAlgorithm = map[string]PhaseOneConfigDetailsAuthenticationAlgorithmEnum{
	"SHA2_384": PhaseOneConfigDetailsAuthenticationAlgorithmSha2384,
	"SHA2_256": PhaseOneConfigDetailsAuthenticationAlgorithmSha2256,
	"SHA1_96":  PhaseOneConfigDetailsAuthenticationAlgorithmSha196,
}

// GetPhaseOneConfigDetailsAuthenticationAlgorithmEnumValues Enumerates the set of values for PhaseOneConfigDetailsAuthenticationAlgorithmEnum
func GetPhaseOneConfigDetailsAuthenticationAlgorithmEnumValues() []PhaseOneConfigDetailsAuthenticationAlgorithmEnum {
	values := make([]PhaseOneConfigDetailsAuthenticationAlgorithmEnum, 0)
	for _, v := range mappingPhaseOneConfigDetailsAuthenticationAlgorithm {
		values = append(values, v)
	}
	return values
}

// PhaseOneConfigDetailsEncryptionAlgorithmEnum Enum with underlying type: string
type PhaseOneConfigDetailsEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for PhaseOneConfigDetailsEncryptionAlgorithmEnum
const (
	PhaseOneConfigDetailsEncryptionAlgorithm256Cbc PhaseOneConfigDetailsEncryptionAlgorithmEnum = "AES_256_CBC"
	PhaseOneConfigDetailsEncryptionAlgorithm192Cbc PhaseOneConfigDetailsEncryptionAlgorithmEnum = "AES_192_CBC"
	PhaseOneConfigDetailsEncryptionAlgorithm128Cbc PhaseOneConfigDetailsEncryptionAlgorithmEnum = "AES_128_CBC"
)

var mappingPhaseOneConfigDetailsEncryptionAlgorithm = map[string]PhaseOneConfigDetailsEncryptionAlgorithmEnum{
	"AES_256_CBC": PhaseOneConfigDetailsEncryptionAlgorithm256Cbc,
	"AES_192_CBC": PhaseOneConfigDetailsEncryptionAlgorithm192Cbc,
	"AES_128_CBC": PhaseOneConfigDetailsEncryptionAlgorithm128Cbc,
}

// GetPhaseOneConfigDetailsEncryptionAlgorithmEnumValues Enumerates the set of values for PhaseOneConfigDetailsEncryptionAlgorithmEnum
func GetPhaseOneConfigDetailsEncryptionAlgorithmEnumValues() []PhaseOneConfigDetailsEncryptionAlgorithmEnum {
	values := make([]PhaseOneConfigDetailsEncryptionAlgorithmEnum, 0)
	for _, v := range mappingPhaseOneConfigDetailsEncryptionAlgorithm {
		values = append(values, v)
	}
	return values
}

// PhaseOneConfigDetailsDiffieHelmanGroupEnum Enum with underlying type: string
type PhaseOneConfigDetailsDiffieHelmanGroupEnum string

// Set of constants representing the allowable values for PhaseOneConfigDetailsDiffieHelmanGroupEnum
const (
	PhaseOneConfigDetailsDiffieHelmanGroupGroup2  PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP2"
	PhaseOneConfigDetailsDiffieHelmanGroupGroup5  PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP5"
	PhaseOneConfigDetailsDiffieHelmanGroupGroup14 PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP14"
	PhaseOneConfigDetailsDiffieHelmanGroupGroup19 PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP19"
	PhaseOneConfigDetailsDiffieHelmanGroupGroup20 PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP20"
	PhaseOneConfigDetailsDiffieHelmanGroupGroup24 PhaseOneConfigDetailsDiffieHelmanGroupEnum = "GROUP24"
)

var mappingPhaseOneConfigDetailsDiffieHelmanGroup = map[string]PhaseOneConfigDetailsDiffieHelmanGroupEnum{
	"GROUP2":  PhaseOneConfigDetailsDiffieHelmanGroupGroup2,
	"GROUP5":  PhaseOneConfigDetailsDiffieHelmanGroupGroup5,
	"GROUP14": PhaseOneConfigDetailsDiffieHelmanGroupGroup14,
	"GROUP19": PhaseOneConfigDetailsDiffieHelmanGroupGroup19,
	"GROUP20": PhaseOneConfigDetailsDiffieHelmanGroupGroup20,
	"GROUP24": PhaseOneConfigDetailsDiffieHelmanGroupGroup24,
}

// GetPhaseOneConfigDetailsDiffieHelmanGroupEnumValues Enumerates the set of values for PhaseOneConfigDetailsDiffieHelmanGroupEnum
func GetPhaseOneConfigDetailsDiffieHelmanGroupEnumValues() []PhaseOneConfigDetailsDiffieHelmanGroupEnum {
	values := make([]PhaseOneConfigDetailsDiffieHelmanGroupEnum, 0)
	for _, v := range mappingPhaseOneConfigDetailsDiffieHelmanGroup {
		values = append(values, v)
	}
	return values
}
