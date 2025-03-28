// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// SeverityEnum Enum with underlying type: string
type SeverityEnum string

// Set of constants representing the allowable values for SeverityEnum
const (
	SeverityCritical SeverityEnum = "CRITICAL"
	SeverityHigh     SeverityEnum = "HIGH"
	SeverityMedium   SeverityEnum = "MEDIUM"
	SeverityLow      SeverityEnum = "LOW"
	SeverityMinor    SeverityEnum = "MINOR"
)

var mappingSeverityEnum = map[string]SeverityEnum{
	"CRITICAL": SeverityCritical,
	"HIGH":     SeverityHigh,
	"MEDIUM":   SeverityMedium,
	"LOW":      SeverityLow,
	"MINOR":    SeverityMinor,
}

var mappingSeverityEnumLowerCase = map[string]SeverityEnum{
	"critical": SeverityCritical,
	"high":     SeverityHigh,
	"medium":   SeverityMedium,
	"low":      SeverityLow,
	"minor":    SeverityMinor,
}

// GetSeverityEnumValues Enumerates the set of values for SeverityEnum
func GetSeverityEnumValues() []SeverityEnum {
	values := make([]SeverityEnum, 0)
	for _, v := range mappingSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetSeverityEnumStringValues Enumerates the set of values in String for SeverityEnum
func GetSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
	}
}

// GetMappingSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSeverityEnum(val string) (SeverityEnum, bool) {
	enum, ok := mappingSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
