// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// KeyAlgorithmEnum Enum with underlying type: string
type KeyAlgorithmEnum string

// Set of constants representing the allowable values for KeyAlgorithmEnum
const (
	KeyAlgorithmRsa2048   KeyAlgorithmEnum = "RSA2048"
	KeyAlgorithmRsa4096   KeyAlgorithmEnum = "RSA4096"
	KeyAlgorithmEcdsaP256 KeyAlgorithmEnum = "ECDSA_P256"
	KeyAlgorithmEcdsaP384 KeyAlgorithmEnum = "ECDSA_P384"
)

var mappingKeyAlgorithmEnum = map[string]KeyAlgorithmEnum{
	"RSA2048":    KeyAlgorithmRsa2048,
	"RSA4096":    KeyAlgorithmRsa4096,
	"ECDSA_P256": KeyAlgorithmEcdsaP256,
	"ECDSA_P384": KeyAlgorithmEcdsaP384,
}

// GetKeyAlgorithmEnumValues Enumerates the set of values for KeyAlgorithmEnum
func GetKeyAlgorithmEnumValues() []KeyAlgorithmEnum {
	values := make([]KeyAlgorithmEnum, 0)
	for _, v := range mappingKeyAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyAlgorithmEnumStringValues Enumerates the set of values in String for KeyAlgorithmEnum
func GetKeyAlgorithmEnumStringValues() []string {
	return []string{
		"RSA2048",
		"RSA4096",
		"ECDSA_P256",
		"ECDSA_P384",
	}
}

// GetMappingKeyAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyAlgorithmEnum(val string) (KeyAlgorithmEnum, bool) {
	mappingKeyAlgorithmEnumIgnoreCase := make(map[string]KeyAlgorithmEnum)
	for k, v := range mappingKeyAlgorithmEnum {
		mappingKeyAlgorithmEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingKeyAlgorithmEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
