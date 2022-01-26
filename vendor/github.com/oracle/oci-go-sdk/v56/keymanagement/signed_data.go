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

// SignedData The representation of SignedData
type SignedData struct {

	// The OCID of the key used to sign the message.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The OCID of the key version used to sign the message.
	KeyVersionId *string `mandatory:"true" json:"keyVersionId"`

	// The base64-encoded binary data object denoting the cryptographic signature generated for the message or message digest.
	Signature *string `mandatory:"true" json:"signature"`

	// The algorithm to use to sign the message or message digest.
	// For RSA keys, supported signature schemes include PKCS #1 and RSASSA-PSS, along with
	// different hashing algorithms.
	// For ECDSA keys, ECDSA is the supported signature scheme with different hashing algorithms.
	// When you pass a message digest for signing, ensure that you specify the same hashing algorithm
	// as used when creating the message digest.
	SigningAlgorithm SignedDataSigningAlgorithmEnum `mandatory:"true" json:"signingAlgorithm"`
}

func (m SignedData) String() string {
	return common.PointerString(m)
}

// SignedDataSigningAlgorithmEnum Enum with underlying type: string
type SignedDataSigningAlgorithmEnum string

// Set of constants representing the allowable values for SignedDataSigningAlgorithmEnum
const (
	SignedDataSigningAlgorithmSha224RsaPkcsPss  SignedDataSigningAlgorithmEnum = "SHA_224_RSA_PKCS_PSS"
	SignedDataSigningAlgorithmSha256RsaPkcsPss  SignedDataSigningAlgorithmEnum = "SHA_256_RSA_PKCS_PSS"
	SignedDataSigningAlgorithmSha384RsaPkcsPss  SignedDataSigningAlgorithmEnum = "SHA_384_RSA_PKCS_PSS"
	SignedDataSigningAlgorithmSha512RsaPkcsPss  SignedDataSigningAlgorithmEnum = "SHA_512_RSA_PKCS_PSS"
	SignedDataSigningAlgorithmSha224RsaPkcs1V15 SignedDataSigningAlgorithmEnum = "SHA_224_RSA_PKCS1_V1_5"
	SignedDataSigningAlgorithmSha256RsaPkcs1V15 SignedDataSigningAlgorithmEnum = "SHA_256_RSA_PKCS1_V1_5"
	SignedDataSigningAlgorithmSha384RsaPkcs1V15 SignedDataSigningAlgorithmEnum = "SHA_384_RSA_PKCS1_V1_5"
	SignedDataSigningAlgorithmSha512RsaPkcs1V15 SignedDataSigningAlgorithmEnum = "SHA_512_RSA_PKCS1_V1_5"
	SignedDataSigningAlgorithmEcdsaSha256       SignedDataSigningAlgorithmEnum = "ECDSA_SHA_256"
	SignedDataSigningAlgorithmEcdsaSha384       SignedDataSigningAlgorithmEnum = "ECDSA_SHA_384"
	SignedDataSigningAlgorithmEcdsaSha512       SignedDataSigningAlgorithmEnum = "ECDSA_SHA_512"
)

var mappingSignedDataSigningAlgorithm = map[string]SignedDataSigningAlgorithmEnum{
	"SHA_224_RSA_PKCS_PSS":   SignedDataSigningAlgorithmSha224RsaPkcsPss,
	"SHA_256_RSA_PKCS_PSS":   SignedDataSigningAlgorithmSha256RsaPkcsPss,
	"SHA_384_RSA_PKCS_PSS":   SignedDataSigningAlgorithmSha384RsaPkcsPss,
	"SHA_512_RSA_PKCS_PSS":   SignedDataSigningAlgorithmSha512RsaPkcsPss,
	"SHA_224_RSA_PKCS1_V1_5": SignedDataSigningAlgorithmSha224RsaPkcs1V15,
	"SHA_256_RSA_PKCS1_V1_5": SignedDataSigningAlgorithmSha256RsaPkcs1V15,
	"SHA_384_RSA_PKCS1_V1_5": SignedDataSigningAlgorithmSha384RsaPkcs1V15,
	"SHA_512_RSA_PKCS1_V1_5": SignedDataSigningAlgorithmSha512RsaPkcs1V15,
	"ECDSA_SHA_256":          SignedDataSigningAlgorithmEcdsaSha256,
	"ECDSA_SHA_384":          SignedDataSigningAlgorithmEcdsaSha384,
	"ECDSA_SHA_512":          SignedDataSigningAlgorithmEcdsaSha512,
}

// GetSignedDataSigningAlgorithmEnumValues Enumerates the set of values for SignedDataSigningAlgorithmEnum
func GetSignedDataSigningAlgorithmEnumValues() []SignedDataSigningAlgorithmEnum {
	values := make([]SignedDataSigningAlgorithmEnum, 0)
	for _, v := range mappingSignedDataSigningAlgorithm {
		values = append(values, v)
	}
	return values
}
