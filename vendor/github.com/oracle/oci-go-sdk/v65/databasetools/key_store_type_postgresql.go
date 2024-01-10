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

// KeyStoreTypePostgresqlEnum Enum with underlying type: string
type KeyStoreTypePostgresqlEnum string

// Set of constants representing the allowable values for KeyStoreTypePostgresqlEnum
const (
	KeyStoreTypePostgresqlClientCertificatePem KeyStoreTypePostgresqlEnum = "CLIENT_CERTIFICATE_PEM"
	KeyStoreTypePostgresqlClientPrivateKeyPem  KeyStoreTypePostgresqlEnum = "CLIENT_PRIVATE_KEY_PEM"
	KeyStoreTypePostgresqlCaCertificatePem     KeyStoreTypePostgresqlEnum = "CA_CERTIFICATE_PEM"
)

var mappingKeyStoreTypePostgresqlEnum = map[string]KeyStoreTypePostgresqlEnum{
	"CLIENT_CERTIFICATE_PEM": KeyStoreTypePostgresqlClientCertificatePem,
	"CLIENT_PRIVATE_KEY_PEM": KeyStoreTypePostgresqlClientPrivateKeyPem,
	"CA_CERTIFICATE_PEM":     KeyStoreTypePostgresqlCaCertificatePem,
}

var mappingKeyStoreTypePostgresqlEnumLowerCase = map[string]KeyStoreTypePostgresqlEnum{
	"client_certificate_pem": KeyStoreTypePostgresqlClientCertificatePem,
	"client_private_key_pem": KeyStoreTypePostgresqlClientPrivateKeyPem,
	"ca_certificate_pem":     KeyStoreTypePostgresqlCaCertificatePem,
}

// GetKeyStoreTypePostgresqlEnumValues Enumerates the set of values for KeyStoreTypePostgresqlEnum
func GetKeyStoreTypePostgresqlEnumValues() []KeyStoreTypePostgresqlEnum {
	values := make([]KeyStoreTypePostgresqlEnum, 0)
	for _, v := range mappingKeyStoreTypePostgresqlEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreTypePostgresqlEnumStringValues Enumerates the set of values in String for KeyStoreTypePostgresqlEnum
func GetKeyStoreTypePostgresqlEnumStringValues() []string {
	return []string{
		"CLIENT_CERTIFICATE_PEM",
		"CLIENT_PRIVATE_KEY_PEM",
		"CA_CERTIFICATE_PEM",
	}
}

// GetMappingKeyStoreTypePostgresqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreTypePostgresqlEnum(val string) (KeyStoreTypePostgresqlEnum, bool) {
	enum, ok := mappingKeyStoreTypePostgresqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
