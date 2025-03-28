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

// RiskLevelEnum Enum with underlying type: string
type RiskLevelEnum string

// Set of constants representing the allowable values for RiskLevelEnum
const (
	RiskLevelCritical RiskLevelEnum = "CRITICAL"
	RiskLevelHigh     RiskLevelEnum = "HIGH"
	RiskLevelMedium   RiskLevelEnum = "MEDIUM"
	RiskLevelLow      RiskLevelEnum = "LOW"
	RiskLevelMinor    RiskLevelEnum = "MINOR"
)

var mappingRiskLevelEnum = map[string]RiskLevelEnum{
	"CRITICAL": RiskLevelCritical,
	"HIGH":     RiskLevelHigh,
	"MEDIUM":   RiskLevelMedium,
	"LOW":      RiskLevelLow,
	"MINOR":    RiskLevelMinor,
}

var mappingRiskLevelEnumLowerCase = map[string]RiskLevelEnum{
	"critical": RiskLevelCritical,
	"high":     RiskLevelHigh,
	"medium":   RiskLevelMedium,
	"low":      RiskLevelLow,
	"minor":    RiskLevelMinor,
}

// GetRiskLevelEnumValues Enumerates the set of values for RiskLevelEnum
func GetRiskLevelEnumValues() []RiskLevelEnum {
	values := make([]RiskLevelEnum, 0)
	for _, v := range mappingRiskLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRiskLevelEnumStringValues Enumerates the set of values in String for RiskLevelEnum
func GetRiskLevelEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
	}
}

// GetMappingRiskLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRiskLevelEnum(val string) (RiskLevelEnum, bool) {
	enum, ok := mappingRiskLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
