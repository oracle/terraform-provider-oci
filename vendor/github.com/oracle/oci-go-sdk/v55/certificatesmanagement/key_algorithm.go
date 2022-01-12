// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// KeyAlgorithmEnum Enum with underlying type: string
type KeyAlgorithmEnum string

// Set of constants representing the allowable values for KeyAlgorithmEnum
const (
	KeyAlgorithmRsa2048   KeyAlgorithmEnum = "RSA2048"
	KeyAlgorithmRsa4096   KeyAlgorithmEnum = "RSA4096"
	KeyAlgorithmEcdsaP256 KeyAlgorithmEnum = "ECDSA_P256"
	KeyAlgorithmEcdsaP384 KeyAlgorithmEnum = "ECDSA_P384"
)

var mappingKeyAlgorithm = map[string]KeyAlgorithmEnum{
	"RSA2048":    KeyAlgorithmRsa2048,
	"RSA4096":    KeyAlgorithmRsa4096,
	"ECDSA_P256": KeyAlgorithmEcdsaP256,
	"ECDSA_P384": KeyAlgorithmEcdsaP384,
}

// GetKeyAlgorithmEnumValues Enumerates the set of values for KeyAlgorithmEnum
func GetKeyAlgorithmEnumValues() []KeyAlgorithmEnum {
	values := make([]KeyAlgorithmEnum, 0)
	for _, v := range mappingKeyAlgorithm {
		values = append(values, v)
	}
	return values
}
