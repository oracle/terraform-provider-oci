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

// GiSourceMajorVersionsEnum Enum with underlying type: string
type GiSourceMajorVersionsEnum string

// Set of constants representing the allowable values for GiSourceMajorVersionsEnum
const (
	GiSourceMajorVersionsGi18 GiSourceMajorVersionsEnum = "GI_18"
	GiSourceMajorVersionsGi19 GiSourceMajorVersionsEnum = "GI_19"
)

var mappingGiSourceMajorVersionsEnum = map[string]GiSourceMajorVersionsEnum{
	"GI_18": GiSourceMajorVersionsGi18,
	"GI_19": GiSourceMajorVersionsGi19,
}

var mappingGiSourceMajorVersionsEnumLowerCase = map[string]GiSourceMajorVersionsEnum{
	"gi_18": GiSourceMajorVersionsGi18,
	"gi_19": GiSourceMajorVersionsGi19,
}

// GetGiSourceMajorVersionsEnumValues Enumerates the set of values for GiSourceMajorVersionsEnum
func GetGiSourceMajorVersionsEnumValues() []GiSourceMajorVersionsEnum {
	values := make([]GiSourceMajorVersionsEnum, 0)
	for _, v := range mappingGiSourceMajorVersionsEnum {
		values = append(values, v)
	}
	return values
}

// GetGiSourceMajorVersionsEnumStringValues Enumerates the set of values in String for GiSourceMajorVersionsEnum
func GetGiSourceMajorVersionsEnumStringValues() []string {
	return []string{
		"GI_18",
		"GI_19",
	}
}

// GetMappingGiSourceMajorVersionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiSourceMajorVersionsEnum(val string) (GiSourceMajorVersionsEnum, bool) {
	enum, ok := mappingGiSourceMajorVersionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
