// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

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

var mappingRiskLevel = map[string]RiskLevelEnum{
	"CRITICAL": RiskLevelCritical,
	"HIGH":     RiskLevelHigh,
	"MEDIUM":   RiskLevelMedium,
	"LOW":      RiskLevelLow,
	"MINOR":    RiskLevelMinor,
}

// GetRiskLevelEnumValues Enumerates the set of values for RiskLevelEnum
func GetRiskLevelEnumValues() []RiskLevelEnum {
	values := make([]RiskLevelEnum, 0)
	for _, v := range mappingRiskLevel {
		values = append(values, v)
	}
	return values
}
