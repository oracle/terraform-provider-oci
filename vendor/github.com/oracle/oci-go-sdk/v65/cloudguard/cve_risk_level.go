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

// CveRiskLevelEnum Enum with underlying type: string
type CveRiskLevelEnum string

// Set of constants representing the allowable values for CveRiskLevelEnum
const (
	CveRiskLevelCritical CveRiskLevelEnum = "CRITICAL"
	CveRiskLevelHigh     CveRiskLevelEnum = "HIGH"
	CveRiskLevelMedium   CveRiskLevelEnum = "MEDIUM"
	CveRiskLevelLow      CveRiskLevelEnum = "LOW"
	CveRiskLevelMinor    CveRiskLevelEnum = "MINOR"
	CveRiskLevelNone     CveRiskLevelEnum = "NONE"
)

var mappingCveRiskLevelEnum = map[string]CveRiskLevelEnum{
	"CRITICAL": CveRiskLevelCritical,
	"HIGH":     CveRiskLevelHigh,
	"MEDIUM":   CveRiskLevelMedium,
	"LOW":      CveRiskLevelLow,
	"MINOR":    CveRiskLevelMinor,
	"NONE":     CveRiskLevelNone,
}

var mappingCveRiskLevelEnumLowerCase = map[string]CveRiskLevelEnum{
	"critical": CveRiskLevelCritical,
	"high":     CveRiskLevelHigh,
	"medium":   CveRiskLevelMedium,
	"low":      CveRiskLevelLow,
	"minor":    CveRiskLevelMinor,
	"none":     CveRiskLevelNone,
}

// GetCveRiskLevelEnumValues Enumerates the set of values for CveRiskLevelEnum
func GetCveRiskLevelEnumValues() []CveRiskLevelEnum {
	values := make([]CveRiskLevelEnum, 0)
	for _, v := range mappingCveRiskLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetCveRiskLevelEnumStringValues Enumerates the set of values in String for CveRiskLevelEnum
func GetCveRiskLevelEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"MINOR",
		"NONE",
	}
}

// GetMappingCveRiskLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCveRiskLevelEnum(val string) (CveRiskLevelEnum, bool) {
	enum, ok := mappingCveRiskLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
