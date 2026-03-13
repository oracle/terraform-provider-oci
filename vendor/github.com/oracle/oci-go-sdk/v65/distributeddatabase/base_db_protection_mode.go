// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbProtectionModeEnum Enum with underlying type: string
type BaseDbProtectionModeEnum string

// Set of constants representing the allowable values for BaseDbProtectionModeEnum
const (
	BaseDbProtectionModeMaximumAvailability BaseDbProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	BaseDbProtectionModeMaximumPerformance  BaseDbProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	BaseDbProtectionModeMaximumProtection   BaseDbProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingBaseDbProtectionModeEnum = map[string]BaseDbProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": BaseDbProtectionModeMaximumAvailability,
	"MAXIMUM_PERFORMANCE":  BaseDbProtectionModeMaximumPerformance,
	"MAXIMUM_PROTECTION":   BaseDbProtectionModeMaximumProtection,
}

var mappingBaseDbProtectionModeEnumLowerCase = map[string]BaseDbProtectionModeEnum{
	"maximum_availability": BaseDbProtectionModeMaximumAvailability,
	"maximum_performance":  BaseDbProtectionModeMaximumPerformance,
	"maximum_protection":   BaseDbProtectionModeMaximumProtection,
}

// GetBaseDbProtectionModeEnumValues Enumerates the set of values for BaseDbProtectionModeEnum
func GetBaseDbProtectionModeEnumValues() []BaseDbProtectionModeEnum {
	values := make([]BaseDbProtectionModeEnum, 0)
	for _, v := range mappingBaseDbProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbProtectionModeEnumStringValues Enumerates the set of values in String for BaseDbProtectionModeEnum
func GetBaseDbProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingBaseDbProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbProtectionModeEnum(val string) (BaseDbProtectionModeEnum, bool) {
	enum, ok := mappingBaseDbProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
