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

// CryptoAnalysisResultModeEnum Enum with underlying type: string
type CryptoAnalysisResultModeEnum string

// Set of constants representing the allowable values for CryptoAnalysisResultModeEnum
const (
	CryptoAnalysisResultModeJfr             CryptoAnalysisResultModeEnum = "JFR"
	CryptoAnalysisResultModeManagedInstance CryptoAnalysisResultModeEnum = "MANAGED_INSTANCE"
)

var mappingCryptoAnalysisResultModeEnum = map[string]CryptoAnalysisResultModeEnum{
	"JFR":              CryptoAnalysisResultModeJfr,
	"MANAGED_INSTANCE": CryptoAnalysisResultModeManagedInstance,
}

var mappingCryptoAnalysisResultModeEnumLowerCase = map[string]CryptoAnalysisResultModeEnum{
	"jfr":              CryptoAnalysisResultModeJfr,
	"managed_instance": CryptoAnalysisResultModeManagedInstance,
}

// GetCryptoAnalysisResultModeEnumValues Enumerates the set of values for CryptoAnalysisResultModeEnum
func GetCryptoAnalysisResultModeEnumValues() []CryptoAnalysisResultModeEnum {
	values := make([]CryptoAnalysisResultModeEnum, 0)
	for _, v := range mappingCryptoAnalysisResultModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAnalysisResultModeEnumStringValues Enumerates the set of values in String for CryptoAnalysisResultModeEnum
func GetCryptoAnalysisResultModeEnumStringValues() []string {
	return []string{
		"JFR",
		"MANAGED_INSTANCE",
	}
}

// GetMappingCryptoAnalysisResultModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAnalysisResultModeEnum(val string) (CryptoAnalysisResultModeEnum, bool) {
	enum, ok := mappingCryptoAnalysisResultModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
