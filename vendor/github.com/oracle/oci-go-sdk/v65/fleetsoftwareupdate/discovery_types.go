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

// DiscoveryTypesEnum Enum with underlying type: string
type DiscoveryTypesEnum string

// Set of constants representing the allowable values for DiscoveryTypesEnum
const (
	DiscoveryTypesDb DiscoveryTypesEnum = "DB"
	DiscoveryTypesGi DiscoveryTypesEnum = "GI"
)

var mappingDiscoveryTypesEnum = map[string]DiscoveryTypesEnum{
	"DB": DiscoveryTypesDb,
	"GI": DiscoveryTypesGi,
}

var mappingDiscoveryTypesEnumLowerCase = map[string]DiscoveryTypesEnum{
	"db": DiscoveryTypesDb,
	"gi": DiscoveryTypesGi,
}

// GetDiscoveryTypesEnumValues Enumerates the set of values for DiscoveryTypesEnum
func GetDiscoveryTypesEnumValues() []DiscoveryTypesEnum {
	values := make([]DiscoveryTypesEnum, 0)
	for _, v := range mappingDiscoveryTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryTypesEnumStringValues Enumerates the set of values in String for DiscoveryTypesEnum
func GetDiscoveryTypesEnumStringValues() []string {
	return []string{
		"DB",
		"GI",
	}
}

// GetMappingDiscoveryTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryTypesEnum(val string) (DiscoveryTypesEnum, bool) {
	enum, ok := mappingDiscoveryTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
