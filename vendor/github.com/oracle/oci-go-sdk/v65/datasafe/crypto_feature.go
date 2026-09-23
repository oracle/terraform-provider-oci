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

// CryptoFeatureEnum Enum with underlying type: string
type CryptoFeatureEnum string

// Set of constants representing the allowable values for CryptoFeatureEnum
const (
	CryptoFeatureTde CryptoFeatureEnum = "TDE"
	CryptoFeatureTls CryptoFeatureEnum = "TLS"
	CryptoFeatureNne CryptoFeatureEnum = "NNE"
)

var mappingCryptoFeatureEnum = map[string]CryptoFeatureEnum{
	"TDE": CryptoFeatureTde,
	"TLS": CryptoFeatureTls,
	"NNE": CryptoFeatureNne,
}

var mappingCryptoFeatureEnumLowerCase = map[string]CryptoFeatureEnum{
	"tde": CryptoFeatureTde,
	"tls": CryptoFeatureTls,
	"nne": CryptoFeatureNne,
}

// GetCryptoFeatureEnumValues Enumerates the set of values for CryptoFeatureEnum
func GetCryptoFeatureEnumValues() []CryptoFeatureEnum {
	values := make([]CryptoFeatureEnum, 0)
	for _, v := range mappingCryptoFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoFeatureEnumStringValues Enumerates the set of values in String for CryptoFeatureEnum
func GetCryptoFeatureEnumStringValues() []string {
	return []string{
		"TDE",
		"TLS",
		"NNE",
	}
}

// GetMappingCryptoFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoFeatureEnum(val string) (CryptoFeatureEnum, bool) {
	enum, ok := mappingCryptoFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
