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

// DistributedDbProtectionModeEnum Enum with underlying type: string
type DistributedDbProtectionModeEnum string

// Set of constants representing the allowable values for DistributedDbProtectionModeEnum
const (
	DistributedDbProtectionModeMaximumAvailability DistributedDbProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DistributedDbProtectionModeMaximumPerformance  DistributedDbProtectionModeEnum = "MAXIMUM_PERFORMANCE"
	DistributedDbProtectionModeMaximumProtection   DistributedDbProtectionModeEnum = "MAXIMUM_PROTECTION"
)

var mappingDistributedDbProtectionModeEnum = map[string]DistributedDbProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DistributedDbProtectionModeMaximumAvailability,
	"MAXIMUM_PERFORMANCE":  DistributedDbProtectionModeMaximumPerformance,
	"MAXIMUM_PROTECTION":   DistributedDbProtectionModeMaximumProtection,
}

var mappingDistributedDbProtectionModeEnumLowerCase = map[string]DistributedDbProtectionModeEnum{
	"maximum_availability": DistributedDbProtectionModeMaximumAvailability,
	"maximum_performance":  DistributedDbProtectionModeMaximumPerformance,
	"maximum_protection":   DistributedDbProtectionModeMaximumProtection,
}

// GetDistributedDbProtectionModeEnumValues Enumerates the set of values for DistributedDbProtectionModeEnum
func GetDistributedDbProtectionModeEnumValues() []DistributedDbProtectionModeEnum {
	values := make([]DistributedDbProtectionModeEnum, 0)
	for _, v := range mappingDistributedDbProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDbProtectionModeEnumStringValues Enumerates the set of values in String for DistributedDbProtectionModeEnum
func GetDistributedDbProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
		"MAXIMUM_PROTECTION",
	}
}

// GetMappingDistributedDbProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDbProtectionModeEnum(val string) (DistributedDbProtectionModeEnum, bool) {
	enum, ok := mappingDistributedDbProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
