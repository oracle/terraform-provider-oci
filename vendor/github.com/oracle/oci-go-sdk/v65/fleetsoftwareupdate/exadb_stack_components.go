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

// ExadbStackComponentsEnum Enum with underlying type: string
type ExadbStackComponentsEnum string

// Set of constants representing the allowable values for ExadbStackComponentsEnum
const (
	ExadbStackComponentsGi      ExadbStackComponentsEnum = "GI"
	ExadbStackComponentsGuestOs ExadbStackComponentsEnum = "GUEST_OS"
)

var mappingExadbStackComponentsEnum = map[string]ExadbStackComponentsEnum{
	"GI":       ExadbStackComponentsGi,
	"GUEST_OS": ExadbStackComponentsGuestOs,
}

var mappingExadbStackComponentsEnumLowerCase = map[string]ExadbStackComponentsEnum{
	"gi":       ExadbStackComponentsGi,
	"guest_os": ExadbStackComponentsGuestOs,
}

// GetExadbStackComponentsEnumValues Enumerates the set of values for ExadbStackComponentsEnum
func GetExadbStackComponentsEnumValues() []ExadbStackComponentsEnum {
	values := make([]ExadbStackComponentsEnum, 0)
	for _, v := range mappingExadbStackComponentsEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbStackComponentsEnumStringValues Enumerates the set of values in String for ExadbStackComponentsEnum
func GetExadbStackComponentsEnumStringValues() []string {
	return []string{
		"GI",
		"GUEST_OS",
	}
}

// GetMappingExadbStackComponentsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbStackComponentsEnum(val string) (ExadbStackComponentsEnum, bool) {
	enum, ok := mappingExadbStackComponentsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
