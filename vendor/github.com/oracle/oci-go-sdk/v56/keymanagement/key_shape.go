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

// KeyShape The cryptographic properties of a key.
type KeyShape struct {

	// The algorithm used by a key's key versions to encrypt or decrypt.
	Algorithm KeyShapeAlgorithmEnum `mandatory:"true" json:"algorithm"`

	// The length of the key in bytes, expressed as an integer. Supported values include the following:
	//   - AES: 16, 24, or 32
	//   - RSA: 256, 384, or 512
	//   - ECDSA: 32, 48, or 66
	Length *int `mandatory:"true" json:"length"`

	// Supported curve IDs for ECDSA keys.
	CurveId KeyShapeCurveIdEnum `mandatory:"false" json:"curveId,omitempty"`
}

func (m KeyShape) String() string {
	return common.PointerString(m)
}

// KeyShapeAlgorithmEnum Enum with underlying type: string
type KeyShapeAlgorithmEnum string

// Set of constants representing the allowable values for KeyShapeAlgorithmEnum
const (
	KeyShapeAlgorithmAes   KeyShapeAlgorithmEnum = "AES"
	KeyShapeAlgorithmRsa   KeyShapeAlgorithmEnum = "RSA"
	KeyShapeAlgorithmEcdsa KeyShapeAlgorithmEnum = "ECDSA"
)

var mappingKeyShapeAlgorithm = map[string]KeyShapeAlgorithmEnum{
	"AES":   KeyShapeAlgorithmAes,
	"RSA":   KeyShapeAlgorithmRsa,
	"ECDSA": KeyShapeAlgorithmEcdsa,
}

// GetKeyShapeAlgorithmEnumValues Enumerates the set of values for KeyShapeAlgorithmEnum
func GetKeyShapeAlgorithmEnumValues() []KeyShapeAlgorithmEnum {
	values := make([]KeyShapeAlgorithmEnum, 0)
	for _, v := range mappingKeyShapeAlgorithm {
		values = append(values, v)
	}
	return values
}

// KeyShapeCurveIdEnum Enum with underlying type: string
type KeyShapeCurveIdEnum string

// Set of constants representing the allowable values for KeyShapeCurveIdEnum
const (
	KeyShapeCurveIdP256 KeyShapeCurveIdEnum = "NIST_P256"
	KeyShapeCurveIdP384 KeyShapeCurveIdEnum = "NIST_P384"
	KeyShapeCurveIdP521 KeyShapeCurveIdEnum = "NIST_P521"
)

var mappingKeyShapeCurveId = map[string]KeyShapeCurveIdEnum{
	"NIST_P256": KeyShapeCurveIdP256,
	"NIST_P384": KeyShapeCurveIdP384,
	"NIST_P521": KeyShapeCurveIdP521,
}

// GetKeyShapeCurveIdEnumValues Enumerates the set of values for KeyShapeCurveIdEnum
func GetKeyShapeCurveIdEnumValues() []KeyShapeCurveIdEnum {
	values := make([]KeyShapeCurveIdEnum, 0)
	for _, v := range mappingKeyShapeCurveId {
		values = append(values, v)
	}
	return values
}
