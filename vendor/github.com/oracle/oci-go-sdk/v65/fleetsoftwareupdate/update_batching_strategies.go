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

// UpdateBatchingStrategiesEnum Enum with underlying type: string
type UpdateBatchingStrategiesEnum string

// Set of constants representing the allowable values for UpdateBatchingStrategiesEnum
const (
	UpdateBatchingStrategiesSequential                UpdateBatchingStrategiesEnum = "SEQUENTIAL"
	UpdateBatchingStrategiesFiftyFifty                UpdateBatchingStrategiesEnum = "FIFTY_FIFTY"
	UpdateBatchingStrategiesServiceAvailabilityFactor UpdateBatchingStrategiesEnum = "SERVICE_AVAILABILITY_FACTOR"
	UpdateBatchingStrategiesNonRolling                UpdateBatchingStrategiesEnum = "NON_ROLLING"
	UpdateBatchingStrategiesNone                      UpdateBatchingStrategiesEnum = "NONE"
)

var mappingUpdateBatchingStrategiesEnum = map[string]UpdateBatchingStrategiesEnum{
	"SEQUENTIAL":                  UpdateBatchingStrategiesSequential,
	"FIFTY_FIFTY":                 UpdateBatchingStrategiesFiftyFifty,
	"SERVICE_AVAILABILITY_FACTOR": UpdateBatchingStrategiesServiceAvailabilityFactor,
	"NON_ROLLING":                 UpdateBatchingStrategiesNonRolling,
	"NONE":                        UpdateBatchingStrategiesNone,
}

var mappingUpdateBatchingStrategiesEnumLowerCase = map[string]UpdateBatchingStrategiesEnum{
	"sequential":                  UpdateBatchingStrategiesSequential,
	"fifty_fifty":                 UpdateBatchingStrategiesFiftyFifty,
	"service_availability_factor": UpdateBatchingStrategiesServiceAvailabilityFactor,
	"non_rolling":                 UpdateBatchingStrategiesNonRolling,
	"none":                        UpdateBatchingStrategiesNone,
}

// GetUpdateBatchingStrategiesEnumValues Enumerates the set of values for UpdateBatchingStrategiesEnum
func GetUpdateBatchingStrategiesEnumValues() []UpdateBatchingStrategiesEnum {
	values := make([]UpdateBatchingStrategiesEnum, 0)
	for _, v := range mappingUpdateBatchingStrategiesEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateBatchingStrategiesEnumStringValues Enumerates the set of values in String for UpdateBatchingStrategiesEnum
func GetUpdateBatchingStrategiesEnumStringValues() []string {
	return []string{
		"SEQUENTIAL",
		"FIFTY_FIFTY",
		"SERVICE_AVAILABILITY_FACTOR",
		"NON_ROLLING",
		"NONE",
	}
}

// GetMappingUpdateBatchingStrategiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateBatchingStrategiesEnum(val string) (UpdateBatchingStrategiesEnum, bool) {
	enum, ok := mappingUpdateBatchingStrategiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
