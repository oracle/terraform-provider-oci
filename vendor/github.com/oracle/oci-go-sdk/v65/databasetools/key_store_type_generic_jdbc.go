// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// KeyStoreTypeGenericJdbcEnum Enum with underlying type: string
type KeyStoreTypeGenericJdbcEnum string

// Set of constants representing the allowable values for KeyStoreTypeGenericJdbcEnum
const (
	KeyStoreTypeGenericJdbcJavaKeyStore         KeyStoreTypeGenericJdbcEnum = "JAVA_KEY_STORE"
	KeyStoreTypeGenericJdbcJavaTrustStore       KeyStoreTypeGenericJdbcEnum = "JAVA_TRUST_STORE"
	KeyStoreTypeGenericJdbcPkcs12               KeyStoreTypeGenericJdbcEnum = "PKCS12"
	KeyStoreTypeGenericJdbcSso                  KeyStoreTypeGenericJdbcEnum = "SSO"
	KeyStoreTypeGenericJdbcClientCertificatePem KeyStoreTypeGenericJdbcEnum = "CLIENT_CERTIFICATE_PEM"
	KeyStoreTypeGenericJdbcClientPrivateKeyPem  KeyStoreTypeGenericJdbcEnum = "CLIENT_PRIVATE_KEY_PEM"
	KeyStoreTypeGenericJdbcCaCertificatePem     KeyStoreTypeGenericJdbcEnum = "CA_CERTIFICATE_PEM"
)

var mappingKeyStoreTypeGenericJdbcEnum = map[string]KeyStoreTypeGenericJdbcEnum{
	"JAVA_KEY_STORE":         KeyStoreTypeGenericJdbcJavaKeyStore,
	"JAVA_TRUST_STORE":       KeyStoreTypeGenericJdbcJavaTrustStore,
	"PKCS12":                 KeyStoreTypeGenericJdbcPkcs12,
	"SSO":                    KeyStoreTypeGenericJdbcSso,
	"CLIENT_CERTIFICATE_PEM": KeyStoreTypeGenericJdbcClientCertificatePem,
	"CLIENT_PRIVATE_KEY_PEM": KeyStoreTypeGenericJdbcClientPrivateKeyPem,
	"CA_CERTIFICATE_PEM":     KeyStoreTypeGenericJdbcCaCertificatePem,
}

var mappingKeyStoreTypeGenericJdbcEnumLowerCase = map[string]KeyStoreTypeGenericJdbcEnum{
	"java_key_store":         KeyStoreTypeGenericJdbcJavaKeyStore,
	"java_trust_store":       KeyStoreTypeGenericJdbcJavaTrustStore,
	"pkcs12":                 KeyStoreTypeGenericJdbcPkcs12,
	"sso":                    KeyStoreTypeGenericJdbcSso,
	"client_certificate_pem": KeyStoreTypeGenericJdbcClientCertificatePem,
	"client_private_key_pem": KeyStoreTypeGenericJdbcClientPrivateKeyPem,
	"ca_certificate_pem":     KeyStoreTypeGenericJdbcCaCertificatePem,
}

// GetKeyStoreTypeGenericJdbcEnumValues Enumerates the set of values for KeyStoreTypeGenericJdbcEnum
func GetKeyStoreTypeGenericJdbcEnumValues() []KeyStoreTypeGenericJdbcEnum {
	values := make([]KeyStoreTypeGenericJdbcEnum, 0)
	for _, v := range mappingKeyStoreTypeGenericJdbcEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreTypeGenericJdbcEnumStringValues Enumerates the set of values in String for KeyStoreTypeGenericJdbcEnum
func GetKeyStoreTypeGenericJdbcEnumStringValues() []string {
	return []string{
		"JAVA_KEY_STORE",
		"JAVA_TRUST_STORE",
		"PKCS12",
		"SSO",
		"CLIENT_CERTIFICATE_PEM",
		"CLIENT_PRIVATE_KEY_PEM",
		"CA_CERTIFICATE_PEM",
	}
}

// GetMappingKeyStoreTypeGenericJdbcEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreTypeGenericJdbcEnum(val string) (KeyStoreTypeGenericJdbcEnum, bool) {
	enum, ok := mappingKeyStoreTypeGenericJdbcEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
