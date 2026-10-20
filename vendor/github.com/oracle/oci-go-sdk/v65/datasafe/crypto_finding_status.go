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

// CryptoFindingStatusEnum Enum with underlying type: string
type CryptoFindingStatusEnum string

// Set of constants representing the allowable values for CryptoFindingStatusEnum
const (
	CryptoFindingStatusPass          CryptoFindingStatusEnum = "PASS"
	CryptoFindingStatusFail          CryptoFindingStatusEnum = "FAIL"
	CryptoFindingStatusError         CryptoFindingStatusEnum = "ERROR"
	CryptoFindingStatusEvaluate      CryptoFindingStatusEnum = "EVALUATE"
	CryptoFindingStatusNotAvailable  CryptoFindingStatusEnum = "NOT_AVAILABLE"
	CryptoFindingStatusNotApplicable CryptoFindingStatusEnum = "NOT_APPLICABLE"
	CryptoFindingStatusNotSupported  CryptoFindingStatusEnum = "NOT_SUPPORTED"
)

var mappingCryptoFindingStatusEnum = map[string]CryptoFindingStatusEnum{
	"PASS":           CryptoFindingStatusPass,
	"FAIL":           CryptoFindingStatusFail,
	"ERROR":          CryptoFindingStatusError,
	"EVALUATE":       CryptoFindingStatusEvaluate,
	"NOT_AVAILABLE":  CryptoFindingStatusNotAvailable,
	"NOT_APPLICABLE": CryptoFindingStatusNotApplicable,
	"NOT_SUPPORTED":  CryptoFindingStatusNotSupported,
}

var mappingCryptoFindingStatusEnumLowerCase = map[string]CryptoFindingStatusEnum{
	"pass":           CryptoFindingStatusPass,
	"fail":           CryptoFindingStatusFail,
	"error":          CryptoFindingStatusError,
	"evaluate":       CryptoFindingStatusEvaluate,
	"not_available":  CryptoFindingStatusNotAvailable,
	"not_applicable": CryptoFindingStatusNotApplicable,
	"not_supported":  CryptoFindingStatusNotSupported,
}

// GetCryptoFindingStatusEnumValues Enumerates the set of values for CryptoFindingStatusEnum
func GetCryptoFindingStatusEnumValues() []CryptoFindingStatusEnum {
	values := make([]CryptoFindingStatusEnum, 0)
	for _, v := range mappingCryptoFindingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoFindingStatusEnumStringValues Enumerates the set of values in String for CryptoFindingStatusEnum
func GetCryptoFindingStatusEnumStringValues() []string {
	return []string{
		"PASS",
		"FAIL",
		"ERROR",
		"EVALUATE",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoFindingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoFindingStatusEnum(val string) (CryptoFindingStatusEnum, bool) {
	enum, ok := mappingCryptoFindingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
