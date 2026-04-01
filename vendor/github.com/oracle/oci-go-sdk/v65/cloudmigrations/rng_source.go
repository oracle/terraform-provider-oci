// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// RngSourceEnum Enum with underlying type: string
type RngSourceEnum string

// Set of constants representing the allowable values for RngSourceEnum
const (
	RngSourceHwrng   RngSourceEnum = "HWRNG"
	RngSourceRandom  RngSourceEnum = "RANDOM"
	RngSourceUrandom RngSourceEnum = "URANDOM"
)

var mappingRngSourceEnum = map[string]RngSourceEnum{
	"HWRNG":   RngSourceHwrng,
	"RANDOM":  RngSourceRandom,
	"URANDOM": RngSourceUrandom,
}

var mappingRngSourceEnumLowerCase = map[string]RngSourceEnum{
	"hwrng":   RngSourceHwrng,
	"random":  RngSourceRandom,
	"urandom": RngSourceUrandom,
}

// GetRngSourceEnumValues Enumerates the set of values for RngSourceEnum
func GetRngSourceEnumValues() []RngSourceEnum {
	values := make([]RngSourceEnum, 0)
	for _, v := range mappingRngSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetRngSourceEnumStringValues Enumerates the set of values in String for RngSourceEnum
func GetRngSourceEnumStringValues() []string {
	return []string{
		"HWRNG",
		"RANDOM",
		"URANDOM",
	}
}

// GetMappingRngSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRngSourceEnum(val string) (RngSourceEnum, bool) {
	enum, ok := mappingRngSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
