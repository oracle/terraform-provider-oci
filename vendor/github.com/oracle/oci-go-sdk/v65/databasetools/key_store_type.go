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

// KeyStoreTypeEnum Enum with underlying type: string
type KeyStoreTypeEnum string

// Set of constants representing the allowable values for KeyStoreTypeEnum
const (
	KeyStoreTypeJavaKeyStore   KeyStoreTypeEnum = "JAVA_KEY_STORE"
	KeyStoreTypeJavaTrustStore KeyStoreTypeEnum = "JAVA_TRUST_STORE"
	KeyStoreTypePkcs12         KeyStoreTypeEnum = "PKCS12"
	KeyStoreTypeSso            KeyStoreTypeEnum = "SSO"
)

var mappingKeyStoreTypeEnum = map[string]KeyStoreTypeEnum{
	"JAVA_KEY_STORE":   KeyStoreTypeJavaKeyStore,
	"JAVA_TRUST_STORE": KeyStoreTypeJavaTrustStore,
	"PKCS12":           KeyStoreTypePkcs12,
	"SSO":              KeyStoreTypeSso,
}

var mappingKeyStoreTypeEnumLowerCase = map[string]KeyStoreTypeEnum{
	"java_key_store":   KeyStoreTypeJavaKeyStore,
	"java_trust_store": KeyStoreTypeJavaTrustStore,
	"pkcs12":           KeyStoreTypePkcs12,
	"sso":              KeyStoreTypeSso,
}

// GetKeyStoreTypeEnumValues Enumerates the set of values for KeyStoreTypeEnum
func GetKeyStoreTypeEnumValues() []KeyStoreTypeEnum {
	values := make([]KeyStoreTypeEnum, 0)
	for _, v := range mappingKeyStoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetKeyStoreTypeEnumStringValues Enumerates the set of values in String for KeyStoreTypeEnum
func GetKeyStoreTypeEnumStringValues() []string {
	return []string{
		"JAVA_KEY_STORE",
		"JAVA_TRUST_STORE",
		"PKCS12",
		"SSO",
	}
}

// GetMappingKeyStoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKeyStoreTypeEnum(val string) (KeyStoreTypeEnum, bool) {
	enum, ok := mappingKeyStoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
