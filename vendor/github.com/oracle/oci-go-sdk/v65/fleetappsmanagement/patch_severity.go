// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// PatchSeverityEnum Enum with underlying type: string
type PatchSeverityEnum string

// Set of constants representing the allowable values for PatchSeverityEnum
const (
	PatchSeverityCritical PatchSeverityEnum = "CRITICAL"
	PatchSeverityHigh     PatchSeverityEnum = "HIGH"
	PatchSeverityMedium   PatchSeverityEnum = "MEDIUM"
	PatchSeverityLow      PatchSeverityEnum = "LOW"
)

var mappingPatchSeverityEnum = map[string]PatchSeverityEnum{
	"CRITICAL": PatchSeverityCritical,
	"HIGH":     PatchSeverityHigh,
	"MEDIUM":   PatchSeverityMedium,
	"LOW":      PatchSeverityLow,
}

var mappingPatchSeverityEnumLowerCase = map[string]PatchSeverityEnum{
	"critical": PatchSeverityCritical,
	"high":     PatchSeverityHigh,
	"medium":   PatchSeverityMedium,
	"low":      PatchSeverityLow,
}

// GetPatchSeverityEnumValues Enumerates the set of values for PatchSeverityEnum
func GetPatchSeverityEnumValues() []PatchSeverityEnum {
	values := make([]PatchSeverityEnum, 0)
	for _, v := range mappingPatchSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSeverityEnumStringValues Enumerates the set of values in String for PatchSeverityEnum
func GetPatchSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingPatchSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSeverityEnum(val string) (PatchSeverityEnum, bool) {
	enum, ok := mappingPatchSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
