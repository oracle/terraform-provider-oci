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

// CryptoQuantumReadinessEnum Enum with underlying type: string
type CryptoQuantumReadinessEnum string

// Set of constants representing the allowable values for CryptoQuantumReadinessEnum
const (
	CryptoQuantumReadinessResistant     CryptoQuantumReadinessEnum = "RESISTANT"
	CryptoQuantumReadinessNotResistant  CryptoQuantumReadinessEnum = "NOT_RESISTANT"
	CryptoQuantumReadinessNotAvailable  CryptoQuantumReadinessEnum = "NOT_AVAILABLE"
	CryptoQuantumReadinessNotApplicable CryptoQuantumReadinessEnum = "NOT_APPLICABLE"
	CryptoQuantumReadinessNotSupported  CryptoQuantumReadinessEnum = "NOT_SUPPORTED"
)

var mappingCryptoQuantumReadinessEnum = map[string]CryptoQuantumReadinessEnum{
	"RESISTANT":      CryptoQuantumReadinessResistant,
	"NOT_RESISTANT":  CryptoQuantumReadinessNotResistant,
	"NOT_AVAILABLE":  CryptoQuantumReadinessNotAvailable,
	"NOT_APPLICABLE": CryptoQuantumReadinessNotApplicable,
	"NOT_SUPPORTED":  CryptoQuantumReadinessNotSupported,
}

var mappingCryptoQuantumReadinessEnumLowerCase = map[string]CryptoQuantumReadinessEnum{
	"resistant":      CryptoQuantumReadinessResistant,
	"not_resistant":  CryptoQuantumReadinessNotResistant,
	"not_available":  CryptoQuantumReadinessNotAvailable,
	"not_applicable": CryptoQuantumReadinessNotApplicable,
	"not_supported":  CryptoQuantumReadinessNotSupported,
}

// GetCryptoQuantumReadinessEnumValues Enumerates the set of values for CryptoQuantumReadinessEnum
func GetCryptoQuantumReadinessEnumValues() []CryptoQuantumReadinessEnum {
	values := make([]CryptoQuantumReadinessEnum, 0)
	for _, v := range mappingCryptoQuantumReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoQuantumReadinessEnumStringValues Enumerates the set of values in String for CryptoQuantumReadinessEnum
func GetCryptoQuantumReadinessEnumStringValues() []string {
	return []string{
		"RESISTANT",
		"NOT_RESISTANT",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoQuantumReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoQuantumReadinessEnum(val string) (CryptoQuantumReadinessEnum, bool) {
	enum, ok := mappingCryptoQuantumReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
