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

// CryptoKeystoreTypeEnum Enum with underlying type: string
type CryptoKeystoreTypeEnum string

// Set of constants representing the allowable values for CryptoKeystoreTypeEnum
const (
	CryptoKeystoreTypeFile          CryptoKeystoreTypeEnum = "FILE"
	CryptoKeystoreTypeOkv           CryptoKeystoreTypeEnum = "OKV"
	CryptoKeystoreTypeHsm           CryptoKeystoreTypeEnum = "HSM"
	CryptoKeystoreTypeNotConfigured CryptoKeystoreTypeEnum = "NOT_CONFIGURED"
	CryptoKeystoreTypeNotApplicable CryptoKeystoreTypeEnum = "NOT_APPLICABLE"
	CryptoKeystoreTypeNotAvailable  CryptoKeystoreTypeEnum = "NOT_AVAILABLE"
	CryptoKeystoreTypeUnknown       CryptoKeystoreTypeEnum = "UNKNOWN"
	CryptoKeystoreTypeNotSupported  CryptoKeystoreTypeEnum = "NOT_SUPPORTED"
)

var mappingCryptoKeystoreTypeEnum = map[string]CryptoKeystoreTypeEnum{
	"FILE":           CryptoKeystoreTypeFile,
	"OKV":            CryptoKeystoreTypeOkv,
	"HSM":            CryptoKeystoreTypeHsm,
	"NOT_CONFIGURED": CryptoKeystoreTypeNotConfigured,
	"NOT_APPLICABLE": CryptoKeystoreTypeNotApplicable,
	"NOT_AVAILABLE":  CryptoKeystoreTypeNotAvailable,
	"UNKNOWN":        CryptoKeystoreTypeUnknown,
	"NOT_SUPPORTED":  CryptoKeystoreTypeNotSupported,
}

var mappingCryptoKeystoreTypeEnumLowerCase = map[string]CryptoKeystoreTypeEnum{
	"file":           CryptoKeystoreTypeFile,
	"okv":            CryptoKeystoreTypeOkv,
	"hsm":            CryptoKeystoreTypeHsm,
	"not_configured": CryptoKeystoreTypeNotConfigured,
	"not_applicable": CryptoKeystoreTypeNotApplicable,
	"not_available":  CryptoKeystoreTypeNotAvailable,
	"unknown":        CryptoKeystoreTypeUnknown,
	"not_supported":  CryptoKeystoreTypeNotSupported,
}

// GetCryptoKeystoreTypeEnumValues Enumerates the set of values for CryptoKeystoreTypeEnum
func GetCryptoKeystoreTypeEnumValues() []CryptoKeystoreTypeEnum {
	values := make([]CryptoKeystoreTypeEnum, 0)
	for _, v := range mappingCryptoKeystoreTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoKeystoreTypeEnumStringValues Enumerates the set of values in String for CryptoKeystoreTypeEnum
func GetCryptoKeystoreTypeEnumStringValues() []string {
	return []string{
		"FILE",
		"OKV",
		"HSM",
		"NOT_CONFIGURED",
		"NOT_APPLICABLE",
		"NOT_AVAILABLE",
		"UNKNOWN",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoKeystoreTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoKeystoreTypeEnum(val string) (CryptoKeystoreTypeEnum, bool) {
	enum, ok := mappingCryptoKeystoreTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
