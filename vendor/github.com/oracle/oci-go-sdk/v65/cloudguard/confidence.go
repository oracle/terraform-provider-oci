// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ConfidenceEnum Enum with underlying type: string
type ConfidenceEnum string

// Set of constants representing the allowable values for ConfidenceEnum
const (
	ConfidenceCritical ConfidenceEnum = "CRITICAL"
	ConfidenceHigh     ConfidenceEnum = "HIGH"
	ConfidenceMedium   ConfidenceEnum = "MEDIUM"
	ConfidenceLow      ConfidenceEnum = "LOW"
	ConfidenceMinor    ConfidenceEnum = "MINOR"
)

var mappingConfidenceEnum = map[string]ConfidenceEnum{
	"CRITICAL": ConfidenceCritical,
	"HIGH":     ConfidenceHigh,
	"MEDIUM":   ConfidenceMedium,
	"LOW":      ConfidenceLow,
	"MINOR":    ConfidenceMinor,
}

var mappingConfidenceEnumLowerCase = map[string]ConfidenceEnum{
	"critical": ConfidenceCritical,
	"high":     ConfidenceHigh,
	"medium":   ConfidenceMedium,
	"low":      ConfidenceLow,
	"minor":    ConfidenceMinor,
}

// GetConfidenceEnumValues Enumerates the set of values for ConfidenceEnum
func GetConfidenceEnumValues() []ConfidenceEnum {
	values := make([]ConfidenceEnum, 0)
	for _, v := range mappingConfidenceEnum {
		values = append(values, v)
	}
	return values
}

// GetConfidenceEnumStringValues Enumerates the set of values in String for ConfidenceEnum
func GetConfidenceEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
	}
}

// GetMappingConfidenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfidenceEnum(val string) (ConfidenceEnum, bool) {
	enum, ok := mappingConfidenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
