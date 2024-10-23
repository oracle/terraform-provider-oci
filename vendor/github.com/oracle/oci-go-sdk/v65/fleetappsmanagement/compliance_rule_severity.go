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

// ComplianceRuleSeverityEnum Enum with underlying type: string
type ComplianceRuleSeverityEnum string

// Set of constants representing the allowable values for ComplianceRuleSeverityEnum
const (
	ComplianceRuleSeverityCritical ComplianceRuleSeverityEnum = "CRITICAL"
	ComplianceRuleSeverityHigh     ComplianceRuleSeverityEnum = "HIGH"
	ComplianceRuleSeverityMedium   ComplianceRuleSeverityEnum = "MEDIUM"
	ComplianceRuleSeverityLow      ComplianceRuleSeverityEnum = "LOW"
)

var mappingComplianceRuleSeverityEnum = map[string]ComplianceRuleSeverityEnum{
	"CRITICAL": ComplianceRuleSeverityCritical,
	"HIGH":     ComplianceRuleSeverityHigh,
	"MEDIUM":   ComplianceRuleSeverityMedium,
	"LOW":      ComplianceRuleSeverityLow,
}

var mappingComplianceRuleSeverityEnumLowerCase = map[string]ComplianceRuleSeverityEnum{
	"critical": ComplianceRuleSeverityCritical,
	"high":     ComplianceRuleSeverityHigh,
	"medium":   ComplianceRuleSeverityMedium,
	"low":      ComplianceRuleSeverityLow,
}

// GetComplianceRuleSeverityEnumValues Enumerates the set of values for ComplianceRuleSeverityEnum
func GetComplianceRuleSeverityEnumValues() []ComplianceRuleSeverityEnum {
	values := make([]ComplianceRuleSeverityEnum, 0)
	for _, v := range mappingComplianceRuleSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceRuleSeverityEnumStringValues Enumerates the set of values in String for ComplianceRuleSeverityEnum
func GetComplianceRuleSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingComplianceRuleSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceRuleSeverityEnum(val string) (ComplianceRuleSeverityEnum, bool) {
	enum, ok := mappingComplianceRuleSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
