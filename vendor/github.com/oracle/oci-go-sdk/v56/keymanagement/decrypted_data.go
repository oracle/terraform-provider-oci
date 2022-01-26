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

// DecryptedData The representation of DecryptedData
type DecryptedData struct {

	// The decrypted data, expressed as a base64-encoded value.
	Plaintext *string `mandatory:"true" json:"plaintext"`

	// The checksum of the decrypted data.
	PlaintextChecksum *string `mandatory:"true" json:"plaintextChecksum"`

	// The OCID of the key used to encrypt the ciphertext.
	KeyId *string `mandatory:"false" json:"keyId"`

	// The OCID of the key version used to encrypt the ciphertext.
	KeyVersionId *string `mandatory:"false" json:"keyVersionId"`

	// The encryption algorithm to use to encrypt and decrypt data with a customer-managed key
	// `AES_256_GCM` indicates that the key is a symmetric key that uses the Advanced Encryption Standard (AES) algorithm and
	// that the mode of encryption is the Galois/Counter Mode (GCM). `RSA_OAEP_SHA_1` indicates that the
	// key is an asymmetric key that uses the RSA encryption algorithm and uses Optimal Asymmetric Encryption Padding (OAEP).
	// `RSA_OAEP_SHA_256` indicates that the key is an asymmetric key that uses the RSA encryption algorithm with a SHA-256 hash
	// and uses OAEP.
	EncryptionAlgorithm DecryptedDataEncryptionAlgorithmEnum `mandatory:"false" json:"encryptionAlgorithm,omitempty"`
}

func (m DecryptedData) String() string {
	return common.PointerString(m)
}

// DecryptedDataEncryptionAlgorithmEnum Enum with underlying type: string
type DecryptedDataEncryptionAlgorithmEnum string

// Set of constants representing the allowable values for DecryptedDataEncryptionAlgorithmEnum
const (
	DecryptedDataEncryptionAlgorithmAes256Gcm     DecryptedDataEncryptionAlgorithmEnum = "AES_256_GCM"
	DecryptedDataEncryptionAlgorithmRsaOaepSha1   DecryptedDataEncryptionAlgorithmEnum = "RSA_OAEP_SHA_1"
	DecryptedDataEncryptionAlgorithmRsaOaepSha256 DecryptedDataEncryptionAlgorithmEnum = "RSA_OAEP_SHA_256"
)

var mappingDecryptedDataEncryptionAlgorithm = map[string]DecryptedDataEncryptionAlgorithmEnum{
	"AES_256_GCM":      DecryptedDataEncryptionAlgorithmAes256Gcm,
	"RSA_OAEP_SHA_1":   DecryptedDataEncryptionAlgorithmRsaOaepSha1,
	"RSA_OAEP_SHA_256": DecryptedDataEncryptionAlgorithmRsaOaepSha256,
}

// GetDecryptedDataEncryptionAlgorithmEnumValues Enumerates the set of values for DecryptedDataEncryptionAlgorithmEnum
func GetDecryptedDataEncryptionAlgorithmEnumValues() []DecryptedDataEncryptionAlgorithmEnum {
	values := make([]DecryptedDataEncryptionAlgorithmEnum, 0)
	for _, v := range mappingDecryptedDataEncryptionAlgorithm {
		values = append(values, v)
	}
	return values
}
