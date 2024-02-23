// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// AvailabilityEnum Enum with underlying type: string
type AvailabilityEnum string

// Set of constants representing the allowable values for AvailabilityEnum
const (
	AvailabilityAvailable        AvailabilityEnum = "AVAILABLE"
	AvailabilitySelected         AvailabilityEnum = "SELECTED"
	AvailabilityRestricted       AvailabilityEnum = "RESTRICTED"
	AvailabilityUnavailable      AvailabilityEnum = "UNAVAILABLE"
	AvailabilityAvailableAtOci   AvailabilityEnum = "AVAILABLE_AT_OCI"
	AvailabilitySelectedAtOci    AvailabilityEnum = "SELECTED_AT_OCI"
	AvailabilityRestrictedAtOci  AvailabilityEnum = "RESTRICTED_AT_OCI"
	AvailabilityUnavailableAtOci AvailabilityEnum = "UNAVAILABLE_AT_OCI"
)

var mappingAvailabilityEnum = map[string]AvailabilityEnum{
	"AVAILABLE":          AvailabilityAvailable,
	"SELECTED":           AvailabilitySelected,
	"RESTRICTED":         AvailabilityRestricted,
	"UNAVAILABLE":        AvailabilityUnavailable,
	"AVAILABLE_AT_OCI":   AvailabilityAvailableAtOci,
	"SELECTED_AT_OCI":    AvailabilitySelectedAtOci,
	"RESTRICTED_AT_OCI":  AvailabilityRestrictedAtOci,
	"UNAVAILABLE_AT_OCI": AvailabilityUnavailableAtOci,
}

var mappingAvailabilityEnumLowerCase = map[string]AvailabilityEnum{
	"available":          AvailabilityAvailable,
	"selected":           AvailabilitySelected,
	"restricted":         AvailabilityRestricted,
	"unavailable":        AvailabilityUnavailable,
	"available_at_oci":   AvailabilityAvailableAtOci,
	"selected_at_oci":    AvailabilitySelectedAtOci,
	"restricted_at_oci":  AvailabilityRestrictedAtOci,
	"unavailable_at_oci": AvailabilityUnavailableAtOci,
}

// GetAvailabilityEnumValues Enumerates the set of values for AvailabilityEnum
func GetAvailabilityEnumValues() []AvailabilityEnum {
	values := make([]AvailabilityEnum, 0)
	for _, v := range mappingAvailabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailabilityEnumStringValues Enumerates the set of values in String for AvailabilityEnum
func GetAvailabilityEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SELECTED",
		"RESTRICTED",
		"UNAVAILABLE",
		"AVAILABLE_AT_OCI",
		"SELECTED_AT_OCI",
		"RESTRICTED_AT_OCI",
		"UNAVAILABLE_AT_OCI",
	}
}

// GetMappingAvailabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailabilityEnum(val string) (AvailabilityEnum, bool) {
	enum, ok := mappingAvailabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
