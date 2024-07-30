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

// FleetDiscoveryOperatorsEnum Enum with underlying type: string
type FleetDiscoveryOperatorsEnum string

// Set of constants representing the allowable values for FleetDiscoveryOperatorsEnum
const (
	FleetDiscoveryOperatorsAnd FleetDiscoveryOperatorsEnum = "AND"
	FleetDiscoveryOperatorsOr  FleetDiscoveryOperatorsEnum = "OR"
)

var mappingFleetDiscoveryOperatorsEnum = map[string]FleetDiscoveryOperatorsEnum{
	"AND": FleetDiscoveryOperatorsAnd,
	"OR":  FleetDiscoveryOperatorsOr,
}

var mappingFleetDiscoveryOperatorsEnumLowerCase = map[string]FleetDiscoveryOperatorsEnum{
	"and": FleetDiscoveryOperatorsAnd,
	"or":  FleetDiscoveryOperatorsOr,
}

// GetFleetDiscoveryOperatorsEnumValues Enumerates the set of values for FleetDiscoveryOperatorsEnum
func GetFleetDiscoveryOperatorsEnumValues() []FleetDiscoveryOperatorsEnum {
	values := make([]FleetDiscoveryOperatorsEnum, 0)
	for _, v := range mappingFleetDiscoveryOperatorsEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetDiscoveryOperatorsEnumStringValues Enumerates the set of values in String for FleetDiscoveryOperatorsEnum
func GetFleetDiscoveryOperatorsEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
	}
}

// GetMappingFleetDiscoveryOperatorsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetDiscoveryOperatorsEnum(val string) (FleetDiscoveryOperatorsEnum, bool) {
	enum, ok := mappingFleetDiscoveryOperatorsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
