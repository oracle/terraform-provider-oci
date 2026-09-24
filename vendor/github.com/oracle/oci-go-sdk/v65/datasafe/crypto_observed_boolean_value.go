// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// CryptoObservedBooleanValueEnum Enum with underlying type: string
type CryptoObservedBooleanValueEnum string

// Set of constants representing the allowable values for CryptoObservedBooleanValueEnum
const (
	CryptoObservedBooleanValueTrue         CryptoObservedBooleanValueEnum = "TRUE"
	CryptoObservedBooleanValueFalse        CryptoObservedBooleanValueEnum = "FALSE"
	CryptoObservedBooleanValueNotSupported CryptoObservedBooleanValueEnum = "NOT_SUPPORTED"
)

var mappingCryptoObservedBooleanValueEnum = map[string]CryptoObservedBooleanValueEnum{
	"TRUE":          CryptoObservedBooleanValueTrue,
	"FALSE":         CryptoObservedBooleanValueFalse,
	"NOT_SUPPORTED": CryptoObservedBooleanValueNotSupported,
}

var mappingCryptoObservedBooleanValueEnumLowerCase = map[string]CryptoObservedBooleanValueEnum{
	"true":          CryptoObservedBooleanValueTrue,
	"false":         CryptoObservedBooleanValueFalse,
	"not_supported": CryptoObservedBooleanValueNotSupported,
}

// GetCryptoObservedBooleanValueEnumValues Enumerates the set of values for CryptoObservedBooleanValueEnum
func GetCryptoObservedBooleanValueEnumValues() []CryptoObservedBooleanValueEnum {
	values := make([]CryptoObservedBooleanValueEnum, 0)
	for _, v := range mappingCryptoObservedBooleanValueEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoObservedBooleanValueEnumStringValues Enumerates the set of values in String for CryptoObservedBooleanValueEnum
func GetCryptoObservedBooleanValueEnumStringValues() []string {
	return []string{
		"TRUE",
		"FALSE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoObservedBooleanValueEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoObservedBooleanValueEnum(val string) (CryptoObservedBooleanValueEnum, bool) {
	enum, ok := mappingCryptoObservedBooleanValueEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
