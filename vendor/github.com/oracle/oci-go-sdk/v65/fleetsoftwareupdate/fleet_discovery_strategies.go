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

// FleetDiscoveryStrategiesEnum Enum with underlying type: string
type FleetDiscoveryStrategiesEnum string

// Set of constants representing the allowable values for FleetDiscoveryStrategiesEnum
const (
	FleetDiscoveryStrategiesSearchQuery      FleetDiscoveryStrategiesEnum = "SEARCH_QUERY"
	FleetDiscoveryStrategiesFilters          FleetDiscoveryStrategiesEnum = "FILTERS"
	FleetDiscoveryStrategiesTargetList       FleetDiscoveryStrategiesEnum = "TARGET_LIST"
	FleetDiscoveryStrategiesDiscoveryResults FleetDiscoveryStrategiesEnum = "DISCOVERY_RESULTS"
)

var mappingFleetDiscoveryStrategiesEnum = map[string]FleetDiscoveryStrategiesEnum{
	"SEARCH_QUERY":      FleetDiscoveryStrategiesSearchQuery,
	"FILTERS":           FleetDiscoveryStrategiesFilters,
	"TARGET_LIST":       FleetDiscoveryStrategiesTargetList,
	"DISCOVERY_RESULTS": FleetDiscoveryStrategiesDiscoveryResults,
}

var mappingFleetDiscoveryStrategiesEnumLowerCase = map[string]FleetDiscoveryStrategiesEnum{
	"search_query":      FleetDiscoveryStrategiesSearchQuery,
	"filters":           FleetDiscoveryStrategiesFilters,
	"target_list":       FleetDiscoveryStrategiesTargetList,
	"discovery_results": FleetDiscoveryStrategiesDiscoveryResults,
}

// GetFleetDiscoveryStrategiesEnumValues Enumerates the set of values for FleetDiscoveryStrategiesEnum
func GetFleetDiscoveryStrategiesEnumValues() []FleetDiscoveryStrategiesEnum {
	values := make([]FleetDiscoveryStrategiesEnum, 0)
	for _, v := range mappingFleetDiscoveryStrategiesEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetDiscoveryStrategiesEnumStringValues Enumerates the set of values in String for FleetDiscoveryStrategiesEnum
func GetFleetDiscoveryStrategiesEnumStringValues() []string {
	return []string{
		"SEARCH_QUERY",
		"FILTERS",
		"TARGET_LIST",
		"DISCOVERY_RESULTS",
	}
}

// GetMappingFleetDiscoveryStrategiesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetDiscoveryStrategiesEnum(val string) (FleetDiscoveryStrategiesEnum, bool) {
	enum, ok := mappingFleetDiscoveryStrategiesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
