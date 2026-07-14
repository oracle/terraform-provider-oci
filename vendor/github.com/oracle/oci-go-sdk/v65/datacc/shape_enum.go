// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// ShapeEnumEnum Enum with underlying type: string
type ShapeEnumEnum string

// Set of constants representing the allowable values for ShapeEnumEnum
const (
	ShapeEnumSixSsds        ShapeEnumEnum = "SIX_SSDS"
	ShapeEnumTwelveSsds     ShapeEnumEnum = "TWELVE_SSDS"
	ShapeEnumEighteenSsds   ShapeEnumEnum = "EIGHTEEN_SSDS"
	ShapeEnumTwentyFourSsds ShapeEnumEnum = "TWENTY_FOUR_SSDS"
)

var mappingShapeEnumEnum = map[string]ShapeEnumEnum{
	"SIX_SSDS":         ShapeEnumSixSsds,
	"TWELVE_SSDS":      ShapeEnumTwelveSsds,
	"EIGHTEEN_SSDS":    ShapeEnumEighteenSsds,
	"TWENTY_FOUR_SSDS": ShapeEnumTwentyFourSsds,
}

var mappingShapeEnumEnumLowerCase = map[string]ShapeEnumEnum{
	"six_ssds":         ShapeEnumSixSsds,
	"twelve_ssds":      ShapeEnumTwelveSsds,
	"eighteen_ssds":    ShapeEnumEighteenSsds,
	"twenty_four_ssds": ShapeEnumTwentyFourSsds,
}

// GetShapeEnumEnumValues Enumerates the set of values for ShapeEnumEnum
func GetShapeEnumEnumValues() []ShapeEnumEnum {
	values := make([]ShapeEnumEnum, 0)
	for _, v := range mappingShapeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetShapeEnumEnumStringValues Enumerates the set of values in String for ShapeEnumEnum
func GetShapeEnumEnumStringValues() []string {
	return []string{
		"SIX_SSDS",
		"TWELVE_SSDS",
		"EIGHTEEN_SSDS",
		"TWENTY_FOUR_SSDS",
	}
}

// GetMappingShapeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShapeEnumEnum(val string) (ShapeEnumEnum, bool) {
	enum, ok := mappingShapeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
