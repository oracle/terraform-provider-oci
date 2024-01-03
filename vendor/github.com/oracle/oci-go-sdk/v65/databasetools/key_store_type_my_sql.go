// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// KeyStoreTypeMySqlEnum Enum with underlying type: string
type KeyStoreTypeMySqlEnum string

// Set of constants representing the allowable values for KeyStoreTypeMySqlEnum
const (
	KeyStoreTypeMySqlClientCertificatePem KeyStoreTypeMySqlEnum = "CLIENT_CERTIFICATE_PEM"
	KeyStoreTypeMySqlClientPrivateKeyPem  KeyStoreTypeMySqlEnum = "CLIENT_PRIVATE_KEY_PEM"
	KeyStoreTypeMySqlCaCertificatePem     KeyStoreTypeMySqlEnum = "CA_CERTIFICATE_PEM"
)

var mappingKeyStoreTypeMySqlEnum = map[string]KeyStoreTypeMySqlEnum{
	"CLIENT_CERTIFICATE_PEM": KeyStoreTypeMySqlClientCertificatePem,
	"CLIENT_PRIVATE_KEY_PEM": KeyStoreTypeMySqlClientPrivateKeyPem,
	"CA_CERTIFICATE_PEM":     KeyStoreTypeMySqlCaCertificatePem,
}

var mappingKeyStoreTypeMySqlEnumLowerCase = map[string]KeyStoreTypeMySqlEnum{
	"client_certificate_pem": KeyStoreTypeMySqlClientCertificatePem,
	"client_private_key_pem": KeyStoreTypeMySqlClientPrivateKeyPem,
	"ca_certificate_pem":     KeyStoreTypeMySqlCaCertificatePem,
}

// GetKeyStoreTypeMySqlEnumValues Enumerates the set of values for KeyStoreTypeMySqlEnum
func GetKeyStoreTypeMySqlEnumValues() []KeyStoreTypeMySqlEnum {
	values := make([]KeyStoreTypeMySqlEnum, 0)
	for _, v := range mappingKeyStoreTypeMySqlEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreTypeMySqlEnumStringValues Enumerates the set of values in String for KeyStoreTypeMySqlEnum
func GetKeyStoreTypeMySqlEnumStringValues() []string {
	return []string{
		"CLIENT_CERTIFICATE_PEM",
		"CLIENT_PRIVATE_KEY_PEM",
		"CA_CERTIFICATE_PEM",
	}
}

// GetMappingKeyStoreTypeMySqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreTypeMySqlEnum(val string) (KeyStoreTypeMySqlEnum, bool) {
	enum, ok := mappingKeyStoreTypeMySqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
