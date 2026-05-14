// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"strings"
)

// CaIssuanceModeEnum Enum with underlying type: string
type CaIssuanceModeEnum string

// Set of constants representing the allowable values for CaIssuanceModeEnum
const (
	CaIssuanceModeShortLived CaIssuanceModeEnum = "SHORT_LIVED"
	CaIssuanceModeLongLived  CaIssuanceModeEnum = "LONG_LIVED"
)

var mappingCaIssuanceModeEnum = map[string]CaIssuanceModeEnum{
	"SHORT_LIVED": CaIssuanceModeShortLived,
	"LONG_LIVED":  CaIssuanceModeLongLived,
}

var mappingCaIssuanceModeEnumLowerCase = map[string]CaIssuanceModeEnum{
	"short_lived": CaIssuanceModeShortLived,
	"long_lived":  CaIssuanceModeLongLived,
}

// GetCaIssuanceModeEnumValues Enumerates the set of values for CaIssuanceModeEnum
func GetCaIssuanceModeEnumValues() []CaIssuanceModeEnum {
	values := make([]CaIssuanceModeEnum, 0)
	for _, v := range mappingCaIssuanceModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCaIssuanceModeEnumStringValues Enumerates the set of values in String for CaIssuanceModeEnum
func GetCaIssuanceModeEnumStringValues() []string {
	return []string{
		"SHORT_LIVED",
		"LONG_LIVED",
	}
}

// GetMappingCaIssuanceModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCaIssuanceModeEnum(val string) (CaIssuanceModeEnum, bool) {
	enum, ok := mappingCaIssuanceModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
