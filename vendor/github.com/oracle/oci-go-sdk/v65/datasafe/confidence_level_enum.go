// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ConfidenceLevelEnumEnum Enum with underlying type: string
type ConfidenceLevelEnumEnum string

// Set of constants representing the allowable values for ConfidenceLevelEnumEnum
const (
	ConfidenceLevelEnumNone   ConfidenceLevelEnumEnum = "NONE"
	ConfidenceLevelEnumHigh   ConfidenceLevelEnumEnum = "HIGH"
	ConfidenceLevelEnumMedium ConfidenceLevelEnumEnum = "MEDIUM"
	ConfidenceLevelEnumLow    ConfidenceLevelEnumEnum = "LOW"
)

var mappingConfidenceLevelEnumEnum = map[string]ConfidenceLevelEnumEnum{
	"NONE":   ConfidenceLevelEnumNone,
	"HIGH":   ConfidenceLevelEnumHigh,
	"MEDIUM": ConfidenceLevelEnumMedium,
	"LOW":    ConfidenceLevelEnumLow,
}

var mappingConfidenceLevelEnumEnumLowerCase = map[string]ConfidenceLevelEnumEnum{
	"none":   ConfidenceLevelEnumNone,
	"high":   ConfidenceLevelEnumHigh,
	"medium": ConfidenceLevelEnumMedium,
	"low":    ConfidenceLevelEnumLow,
}

// GetConfidenceLevelEnumEnumValues Enumerates the set of values for ConfidenceLevelEnumEnum
func GetConfidenceLevelEnumEnumValues() []ConfidenceLevelEnumEnum {
	values := make([]ConfidenceLevelEnumEnum, 0)
	for _, v := range mappingConfidenceLevelEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetConfidenceLevelEnumEnumStringValues Enumerates the set of values in String for ConfidenceLevelEnumEnum
func GetConfidenceLevelEnumEnumStringValues() []string {
	return []string{
		"NONE",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingConfidenceLevelEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfidenceLevelEnumEnum(val string) (ConfidenceLevelEnumEnum, bool) {
	enum, ok := mappingConfidenceLevelEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
