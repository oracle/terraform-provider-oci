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

// EncryptDataDetails The representation of EncryptDataDetails
type EncryptDataDetails struct {

	// The OCID of the key to encrypt with.
	KeyId *string `mandatory:"true" json:"keyId"`

	// The plaintext data to encrypt.
	Plaintext *string `mandatory:"true" json:"plaintext"`

	// Information that can be used to provide an encryption context for the
	// encrypted data. The length of the string representation of the associated data
	// must be fewer than 4096 characters.
	AssociatedData map[string]string `mandatory:"false" json:"associatedData"`

	// Information that provides context for audit logging. You can provide this additional
	// data as key-value pairs to include in the audit logs when audit logging is enabled.
	LoggingContext map[string]string `mandatory:"false" json:"loggingContext"`

	// The OCID of the key version used to encrypt the ciphertext.
	KeyVersionId *string `mandatory:"false" json:"keyVersionId"`

	// The encryption algorithm to use to encrypt and decrypt data with a customer-managed key.
	// `AES_256_GCM` indicates that the key is a symmetric key that uses the Advanced Encryption Standard (AES) algorithm and
	// that the mode of encryption is the Galois/Counter Mode (GCM). `RSA_OAEP_SHA_1` indicates that the
	// key is an asymmetric key that uses the RSA encryption algorithm and uses Optimal Asymmetric Encryption Padding (OAEP).
	// `RSA_OAEP_SHA_256` indicates that the key is an asymmetric key that uses the RSA encryption algorithm with a SHA-256 hash
	// and uses OAEP.
	EncryptionAlgorithm EncryptDataDetailsEncryptionAlgorithmEnum `mandatory:"false" json:"encryptionAlgorithm,omitempty"`
}

func (m EncryptDataDetails) String() string {
	return common.PointerString(m)
}

// EncryptDataDetailsEncryptionAlgorithmEnum Enum with underlying type: string
type EncryptDataDetailsEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for EncryptDataDetailsEncryptionAlgorithmEnum
const (
	EncryptDataDetailsEncryptionAlgorithmAes256Gcm     EncryptDataDetailsEncryptionAlgorithmEnum = "AES_256_GCM"
	EncryptDataDetailsEncryptionAlgorithmRsaOaepSha1   EncryptDataDetailsEncryptionAlgorithmEnum = "RSA_OAEP_SHA_1"
	EncryptDataDetailsEncryptionAlgorithmRsaOaepSha256 EncryptDataDetailsEncryptionAlgorithmEnum = "RSA_OAEP_SHA_256"
)

var mappingEncryptDataDetailsEncryptionAlgorithm = map[string]EncryptDataDetailsEncryptionAlgorithmEnum{
	"AES_256_GCM":      EncryptDataDetailsEncryptionAlgorithmAes256Gcm,
	"RSA_OAEP_SHA_1":   EncryptDataDetailsEncryptionAlgorithmRsaOaepSha1,
	"RSA_OAEP_SHA_256": EncryptDataDetailsEncryptionAlgorithmRsaOaepSha256,
}

// GetEncryptDataDetailsEncryptionAlgorithmEnumValues Enumerates the set of values for EncryptDataDetailsEncryptionAlgorithmEnum
func GetEncryptDataDetailsEncryptionAlgorithmEnumValues() []EncryptDataDetailsEncryptionAlgorithmEnum {
	values := make([]EncryptDataDetailsEncryptionAlgorithmEnum, 0)
	for _, v := range mappingEncryptDataDetailsEncryptionAlgorithm {
		values = append(values, v)
	}
	return values
}
