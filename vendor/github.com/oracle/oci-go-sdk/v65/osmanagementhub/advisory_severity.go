// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// AdvisorySeverityEnum Enum with underlying type: string
type AdvisorySeverityEnum string

// Set of constants representing the allowable values for AdvisorySeverityEnum
const (
	AdvisorySeverityLow       AdvisorySeverityEnum = "LOW"
	AdvisorySeverityModerate  AdvisorySeverityEnum = "MODERATE"
	AdvisorySeverityImportant AdvisorySeverityEnum = "IMPORTANT"
	AdvisorySeverityCritical  AdvisorySeverityEnum = "CRITICAL"
)

var mappingAdvisorySeverityEnum = map[string]AdvisorySeverityEnum{
	"LOW":       AdvisorySeverityLow,
	"MODERATE":  AdvisorySeverityModerate,
	"IMPORTANT": AdvisorySeverityImportant,
	"CRITICAL":  AdvisorySeverityCritical,
}

var mappingAdvisorySeverityEnumLowerCase = map[string]AdvisorySeverityEnum{
	"low":       AdvisorySeverityLow,
	"moderate":  AdvisorySeverityModerate,
	"important": AdvisorySeverityImportant,
	"critical":  AdvisorySeverityCritical,
}

// GetAdvisorySeverityEnumValues Enumerates the set of values for AdvisorySeverityEnum
func GetAdvisorySeverityEnumValues() []AdvisorySeverityEnum {
	values := make([]AdvisorySeverityEnum, 0)
	for _, v := range mappingAdvisorySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAdvisorySeverityEnumStringValues Enumerates the set of values in String for AdvisorySeverityEnum
func GetAdvisorySeverityEnumStringValues() []string {
	return []string{
		"LOW",
		"MODERATE",
		"IMPORTANT",
		"CRITICAL",
	}
}

// GetMappingAdvisorySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdvisorySeverityEnum(val string) (AdvisorySeverityEnum, bool) {
	enum, ok := mappingAdvisorySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
