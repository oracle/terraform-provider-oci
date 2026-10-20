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

// CryptoFeatureStatusEnum Enum with underlying type: string
type CryptoFeatureStatusEnum string

// Set of constants representing the allowable values for CryptoFeatureStatusEnum
const (
	CryptoFeatureStatusEnabled       CryptoFeatureStatusEnum = "ENABLED"
	CryptoFeatureStatusDisabled      CryptoFeatureStatusEnum = "DISABLED"
	CryptoFeatureStatusNotConfigured CryptoFeatureStatusEnum = "NOT_CONFIGURED"
	CryptoFeatureStatusNotAvailable  CryptoFeatureStatusEnum = "NOT_AVAILABLE"
	CryptoFeatureStatusNotApplicable CryptoFeatureStatusEnum = "NOT_APPLICABLE"
	CryptoFeatureStatusNotSupported  CryptoFeatureStatusEnum = "NOT_SUPPORTED"
)

var mappingCryptoFeatureStatusEnum = map[string]CryptoFeatureStatusEnum{
	"ENABLED":        CryptoFeatureStatusEnabled,
	"DISABLED":       CryptoFeatureStatusDisabled,
	"NOT_CONFIGURED": CryptoFeatureStatusNotConfigured,
	"NOT_AVAILABLE":  CryptoFeatureStatusNotAvailable,
	"NOT_APPLICABLE": CryptoFeatureStatusNotApplicable,
	"NOT_SUPPORTED":  CryptoFeatureStatusNotSupported,
}

var mappingCryptoFeatureStatusEnumLowerCase = map[string]CryptoFeatureStatusEnum{
	"enabled":        CryptoFeatureStatusEnabled,
	"disabled":       CryptoFeatureStatusDisabled,
	"not_configured": CryptoFeatureStatusNotConfigured,
	"not_available":  CryptoFeatureStatusNotAvailable,
	"not_applicable": CryptoFeatureStatusNotApplicable,
	"not_supported":  CryptoFeatureStatusNotSupported,
}

// GetCryptoFeatureStatusEnumValues Enumerates the set of values for CryptoFeatureStatusEnum
func GetCryptoFeatureStatusEnumValues() []CryptoFeatureStatusEnum {
	values := make([]CryptoFeatureStatusEnum, 0)
	for _, v := range mappingCryptoFeatureStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoFeatureStatusEnumStringValues Enumerates the set of values in String for CryptoFeatureStatusEnum
func GetCryptoFeatureStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NOT_CONFIGURED",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoFeatureStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoFeatureStatusEnum(val string) (CryptoFeatureStatusEnum, bool) {
	enum, ok := mappingCryptoFeatureStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
