// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// BatchingStrategiesEnum Enum with underlying type: string
type BatchingStrategiesEnum string

// Set of constants representing the allowable values for BatchingStrategiesEnum
const (
	BatchingStrategiesSequential                BatchingStrategiesEnum = "SEQUENTIAL"
	BatchingStrategiesFiftyFifty                BatchingStrategiesEnum = "FIFTY_FIFTY"
	BatchingStrategiesServiceAvailabilityFactor BatchingStrategiesEnum = "SERVICE_AVAILABILITY_FACTOR"
	BatchingStrategiesNonRolling                BatchingStrategiesEnum = "NON_ROLLING"
)

var mappingBatchingStrategiesEnum = map[string]BatchingStrategiesEnum{
	"SEQUENTIAL":                  BatchingStrategiesSequential,
	"FIFTY_FIFTY":                 BatchingStrategiesFiftyFifty,
	"SERVICE_AVAILABILITY_FACTOR": BatchingStrategiesServiceAvailabilityFactor,
	"NON_ROLLING":                 BatchingStrategiesNonRolling,
}

var mappingBatchingStrategiesEnumLowerCase = map[string]BatchingStrategiesEnum{
	"sequential":                  BatchingStrategiesSequential,
	"fifty_fifty":                 BatchingStrategiesFiftyFifty,
	"service_availability_factor": BatchingStrategiesServiceAvailabilityFactor,
	"non_rolling":                 BatchingStrategiesNonRolling,
}

// GetBatchingStrategiesEnumValues Enumerates the set of values for BatchingStrategiesEnum
func GetBatchingStrategiesEnumValues() []BatchingStrategiesEnum {
	values := make([]BatchingStrategiesEnum, 0)
	for _, v := range mappingBatchingStrategiesEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchingStrategiesEnumStringValues Enumerates the set of values in String for BatchingStrategiesEnum
func GetBatchingStrategiesEnumStringValues() []string {
	return []string{
		"SEQUENTIAL",
		"FIFTY_FIFTY",
		"SERVICE_AVAILABILITY_FACTOR",
		"NON_ROLLING",
	}
}

// GetMappingBatchingStrategiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchingStrategiesEnum(val string) (BatchingStrategiesEnum, bool) {
	enum, ok := mappingBatchingStrategiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
