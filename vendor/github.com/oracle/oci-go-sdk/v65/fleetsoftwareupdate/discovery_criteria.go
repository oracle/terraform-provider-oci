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

// DiscoveryCriteriaEnum Enum with underlying type: string
type DiscoveryCriteriaEnum string

// Set of constants representing the allowable values for DiscoveryCriteriaEnum
const (
	DiscoveryCriteriaSearchQuery DiscoveryCriteriaEnum = "SEARCH_QUERY"
	DiscoveryCriteriaFilters     DiscoveryCriteriaEnum = "FILTERS"
)

var mappingDiscoveryCriteriaEnum = map[string]DiscoveryCriteriaEnum{
	"SEARCH_QUERY": DiscoveryCriteriaSearchQuery,
	"FILTERS":      DiscoveryCriteriaFilters,
}

var mappingDiscoveryCriteriaEnumLowerCase = map[string]DiscoveryCriteriaEnum{
	"search_query": DiscoveryCriteriaSearchQuery,
	"filters":      DiscoveryCriteriaFilters,
}

// GetDiscoveryCriteriaEnumValues Enumerates the set of values for DiscoveryCriteriaEnum
func GetDiscoveryCriteriaEnumValues() []DiscoveryCriteriaEnum {
	values := make([]DiscoveryCriteriaEnum, 0)
	for _, v := range mappingDiscoveryCriteriaEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryCriteriaEnumStringValues Enumerates the set of values in String for DiscoveryCriteriaEnum
func GetDiscoveryCriteriaEnumStringValues() []string {
	return []string{
		"SEARCH_QUERY",
		"FILTERS",
	}
}

// GetMappingDiscoveryCriteriaEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryCriteriaEnum(val string) (DiscoveryCriteriaEnum, bool) {
	enum, ok := mappingDiscoveryCriteriaEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
