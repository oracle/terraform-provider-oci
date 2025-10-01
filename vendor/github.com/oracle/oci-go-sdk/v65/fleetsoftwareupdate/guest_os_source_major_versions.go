// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GuestOsSourceMajorVersionsEnum Enum with underlying type: string
type GuestOsSourceMajorVersionsEnum string

// Set of constants representing the allowable values for GuestOsSourceMajorVersionsEnum
const (
	GuestOsSourceMajorVersionsExaOl5 GuestOsSourceMajorVersionsEnum = "EXA_OL_5"
	GuestOsSourceMajorVersionsExaOl6 GuestOsSourceMajorVersionsEnum = "EXA_OL_6"
	GuestOsSourceMajorVersionsExaOl7 GuestOsSourceMajorVersionsEnum = "EXA_OL_7"
	GuestOsSourceMajorVersionsExaOl8 GuestOsSourceMajorVersionsEnum = "EXA_OL_8"
)

var mappingGuestOsSourceMajorVersionsEnum = map[string]GuestOsSourceMajorVersionsEnum{
	"EXA_OL_5": GuestOsSourceMajorVersionsExaOl5,
	"EXA_OL_6": GuestOsSourceMajorVersionsExaOl6,
	"EXA_OL_7": GuestOsSourceMajorVersionsExaOl7,
	"EXA_OL_8": GuestOsSourceMajorVersionsExaOl8,
}

var mappingGuestOsSourceMajorVersionsEnumLowerCase = map[string]GuestOsSourceMajorVersionsEnum{
	"exa_ol_5": GuestOsSourceMajorVersionsExaOl5,
	"exa_ol_6": GuestOsSourceMajorVersionsExaOl6,
	"exa_ol_7": GuestOsSourceMajorVersionsExaOl7,
	"exa_ol_8": GuestOsSourceMajorVersionsExaOl8,
}

// GetGuestOsSourceMajorVersionsEnumValues Enumerates the set of values for GuestOsSourceMajorVersionsEnum
func GetGuestOsSourceMajorVersionsEnumValues() []GuestOsSourceMajorVersionsEnum {
	values := make([]GuestOsSourceMajorVersionsEnum, 0)
	for _, v := range mappingGuestOsSourceMajorVersionsEnum {
		values = append(values, v)
	}
	return values
}

// GetGuestOsSourceMajorVersionsEnumStringValues Enumerates the set of values in String for GuestOsSourceMajorVersionsEnum
func GetGuestOsSourceMajorVersionsEnumStringValues() []string {
	return []string{
		"EXA_OL_5",
		"EXA_OL_6",
		"EXA_OL_7",
		"EXA_OL_8",
	}
}

// GetMappingGuestOsSourceMajorVersionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGuestOsSourceMajorVersionsEnum(val string) (GuestOsSourceMajorVersionsEnum, bool) {
	enum, ok := mappingGuestOsSourceMajorVersionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
