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

// ReleaseVersionStatusEnum Enum with underlying type: string
type ReleaseVersionStatusEnum string

// Set of constants representing the allowable values for ReleaseVersionStatusEnum
const (
	ReleaseVersionStatusSupported  ReleaseVersionStatusEnum = "SUPPORTED"
	ReleaseVersionStatusDeprecated ReleaseVersionStatusEnum = "DEPRECATED"
)

var mappingReleaseVersionStatusEnum = map[string]ReleaseVersionStatusEnum{
	"SUPPORTED":  ReleaseVersionStatusSupported,
	"DEPRECATED": ReleaseVersionStatusDeprecated,
}

var mappingReleaseVersionStatusEnumLowerCase = map[string]ReleaseVersionStatusEnum{
	"supported":  ReleaseVersionStatusSupported,
	"deprecated": ReleaseVersionStatusDeprecated,
}

// GetReleaseVersionStatusEnumValues Enumerates the set of values for ReleaseVersionStatusEnum
func GetReleaseVersionStatusEnumValues() []ReleaseVersionStatusEnum {
	values := make([]ReleaseVersionStatusEnum, 0)
	for _, v := range mappingReleaseVersionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReleaseVersionStatusEnumStringValues Enumerates the set of values in String for ReleaseVersionStatusEnum
func GetReleaseVersionStatusEnumStringValues() []string {
	return []string{
		"SUPPORTED",
		"DEPRECATED",
	}
}

// GetMappingReleaseVersionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReleaseVersionStatusEnum(val string) (ReleaseVersionStatusEnum, bool) {
	enum, ok := mappingReleaseVersionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
