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

// CryptoKeyCacheStatusEnum Enum with underlying type: string
type CryptoKeyCacheStatusEnum string

// Set of constants representing the allowable values for CryptoKeyCacheStatusEnum
const (
	CryptoKeyCacheStatusEnabled       CryptoKeyCacheStatusEnum = "ENABLED"
	CryptoKeyCacheStatusDisabled      CryptoKeyCacheStatusEnum = "DISABLED"
	CryptoKeyCacheStatusNotApplicable CryptoKeyCacheStatusEnum = "NOT_APPLICABLE"
	CryptoKeyCacheStatusNotSupported  CryptoKeyCacheStatusEnum = "NOT_SUPPORTED"
)

var mappingCryptoKeyCacheStatusEnum = map[string]CryptoKeyCacheStatusEnum{
	"ENABLED":        CryptoKeyCacheStatusEnabled,
	"DISABLED":       CryptoKeyCacheStatusDisabled,
	"NOT_APPLICABLE": CryptoKeyCacheStatusNotApplicable,
	"NOT_SUPPORTED":  CryptoKeyCacheStatusNotSupported,
}

var mappingCryptoKeyCacheStatusEnumLowerCase = map[string]CryptoKeyCacheStatusEnum{
	"enabled":        CryptoKeyCacheStatusEnabled,
	"disabled":       CryptoKeyCacheStatusDisabled,
	"not_applicable": CryptoKeyCacheStatusNotApplicable,
	"not_supported":  CryptoKeyCacheStatusNotSupported,
}

// GetCryptoKeyCacheStatusEnumValues Enumerates the set of values for CryptoKeyCacheStatusEnum
func GetCryptoKeyCacheStatusEnumValues() []CryptoKeyCacheStatusEnum {
	values := make([]CryptoKeyCacheStatusEnum, 0)
	for _, v := range mappingCryptoKeyCacheStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoKeyCacheStatusEnumStringValues Enumerates the set of values in String for CryptoKeyCacheStatusEnum
func GetCryptoKeyCacheStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoKeyCacheStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoKeyCacheStatusEnum(val string) (CryptoKeyCacheStatusEnum, bool) {
	enum, ok := mappingCryptoKeyCacheStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
