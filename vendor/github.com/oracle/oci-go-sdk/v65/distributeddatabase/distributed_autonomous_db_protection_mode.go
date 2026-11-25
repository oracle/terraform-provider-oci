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

// DistributedAutonomousDbProtectionModeEnum Enum with underlying type: string
type DistributedAutonomousDbProtectionModeEnum string

// Set of constants representing the allowable values for DistributedAutonomousDbProtectionModeEnum
const (
	DistributedAutonomousDbProtectionModeMaximumAvailability DistributedAutonomousDbProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	DistributedAutonomousDbProtectionModeMaximumPerformance  DistributedAutonomousDbProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingDistributedAutonomousDbProtectionModeEnum = map[string]DistributedAutonomousDbProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": DistributedAutonomousDbProtectionModeMaximumAvailability,
	"MAXIMUM_PERFORMANCE":  DistributedAutonomousDbProtectionModeMaximumPerformance,
}

var mappingDistributedAutonomousDbProtectionModeEnumLowerCase = map[string]DistributedAutonomousDbProtectionModeEnum{
	"maximum_availability": DistributedAutonomousDbProtectionModeMaximumAvailability,
	"maximum_performance":  DistributedAutonomousDbProtectionModeMaximumPerformance,
}

// GetDistributedAutonomousDbProtectionModeEnumValues Enumerates the set of values for DistributedAutonomousDbProtectionModeEnum
func GetDistributedAutonomousDbProtectionModeEnumValues() []DistributedAutonomousDbProtectionModeEnum {
	values := make([]DistributedAutonomousDbProtectionModeEnum, 0)
	for _, v := range mappingDistributedAutonomousDbProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDbProtectionModeEnumStringValues Enumerates the set of values in String for DistributedAutonomousDbProtectionModeEnum
func GetDistributedAutonomousDbProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingDistributedAutonomousDbProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDbProtectionModeEnum(val string) (DistributedAutonomousDbProtectionModeEnum, bool) {
	enum, ok := mappingDistributedAutonomousDbProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
