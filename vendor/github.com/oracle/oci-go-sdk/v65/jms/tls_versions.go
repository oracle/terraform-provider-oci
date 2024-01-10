// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// TlsVersionsEnum Enum with underlying type: string
type TlsVersionsEnum string

// Set of constants representing the allowable values for TlsVersionsEnum
const (
	TlsVersionsTls10 TlsVersionsEnum = "TLS_1_0"
	TlsVersionsTls11 TlsVersionsEnum = "TLS_1_1"
)

var mappingTlsVersionsEnum = map[string]TlsVersionsEnum{
	"TLS_1_0": TlsVersionsTls10,
	"TLS_1_1": TlsVersionsTls11,
}

var mappingTlsVersionsEnumLowerCase = map[string]TlsVersionsEnum{
	"tls_1_0": TlsVersionsTls10,
	"tls_1_1": TlsVersionsTls11,
}

// GetTlsVersionsEnumValues Enumerates the set of values for TlsVersionsEnum
func GetTlsVersionsEnumValues() []TlsVersionsEnum {
	values := make([]TlsVersionsEnum, 0)
	for _, v := range mappingTlsVersionsEnum {
		values = append(values, v)
	}
	return values
}

// GetTlsVersionsEnumStringValues Enumerates the set of values in String for TlsVersionsEnum
func GetTlsVersionsEnumStringValues() []string {
	return []string{
		"TLS_1_0",
		"TLS_1_1",
	}
}

// GetMappingTlsVersionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTlsVersionsEnum(val string) (TlsVersionsEnum, bool) {
	enum, ok := mappingTlsVersionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
