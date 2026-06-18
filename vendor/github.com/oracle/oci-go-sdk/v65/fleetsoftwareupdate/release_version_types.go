// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ReleaseVersionTypesEnum Enum with underlying type: string
type ReleaseVersionTypesEnum string

// Set of constants representing the allowable values for ReleaseVersionTypesEnum
const (
	ReleaseVersionTypesExadataReleaseVersion ReleaseVersionTypesEnum = "EXADATA_RELEASE_VERSION"
)

var mappingReleaseVersionTypesEnum = map[string]ReleaseVersionTypesEnum{
	"EXADATA_RELEASE_VERSION": ReleaseVersionTypesExadataReleaseVersion,
}

var mappingReleaseVersionTypesEnumLowerCase = map[string]ReleaseVersionTypesEnum{
	"exadata_release_version": ReleaseVersionTypesExadataReleaseVersion,
}

// GetReleaseVersionTypesEnumValues Enumerates the set of values for ReleaseVersionTypesEnum
func GetReleaseVersionTypesEnumValues() []ReleaseVersionTypesEnum {
	values := make([]ReleaseVersionTypesEnum, 0)
	for _, v := range mappingReleaseVersionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetReleaseVersionTypesEnumStringValues Enumerates the set of values in String for ReleaseVersionTypesEnum
func GetReleaseVersionTypesEnumStringValues() []string {
	return []string{
		"EXADATA_RELEASE_VERSION",
	}
}

// GetMappingReleaseVersionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReleaseVersionTypesEnum(val string) (ReleaseVersionTypesEnum, bool) {
	enum, ok := mappingReleaseVersionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
