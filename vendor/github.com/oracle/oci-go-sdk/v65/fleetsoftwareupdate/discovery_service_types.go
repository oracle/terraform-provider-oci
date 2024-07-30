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

// DiscoveryServiceTypesEnum Enum with underlying type: string
type DiscoveryServiceTypesEnum string

// Set of constants representing the allowable values for DiscoveryServiceTypesEnum
const (
	DiscoveryServiceTypesExacs DiscoveryServiceTypesEnum = "EXACS"
	DiscoveryServiceTypesExacc DiscoveryServiceTypesEnum = "EXACC"
)

var mappingDiscoveryServiceTypesEnum = map[string]DiscoveryServiceTypesEnum{
	"EXACS": DiscoveryServiceTypesExacs,
	"EXACC": DiscoveryServiceTypesExacc,
}

var mappingDiscoveryServiceTypesEnumLowerCase = map[string]DiscoveryServiceTypesEnum{
	"exacs": DiscoveryServiceTypesExacs,
	"exacc": DiscoveryServiceTypesExacc,
}

// GetDiscoveryServiceTypesEnumValues Enumerates the set of values for DiscoveryServiceTypesEnum
func GetDiscoveryServiceTypesEnumValues() []DiscoveryServiceTypesEnum {
	values := make([]DiscoveryServiceTypesEnum, 0)
	for _, v := range mappingDiscoveryServiceTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryServiceTypesEnumStringValues Enumerates the set of values in String for DiscoveryServiceTypesEnum
func GetDiscoveryServiceTypesEnumStringValues() []string {
	return []string{
		"EXACS",
		"EXACC",
	}
}

// GetMappingDiscoveryServiceTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryServiceTypesEnum(val string) (DiscoveryServiceTypesEnum, bool) {
	enum, ok := mappingDiscoveryServiceTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
