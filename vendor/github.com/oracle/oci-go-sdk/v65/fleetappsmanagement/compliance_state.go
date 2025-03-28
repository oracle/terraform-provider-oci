// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ComplianceStateEnum Enum with underlying type: string
type ComplianceStateEnum string

// Set of constants representing the allowable values for ComplianceStateEnum
const (
	ComplianceStateUnknown      ComplianceStateEnum = "UNKNOWN"
	ComplianceStateCompliant    ComplianceStateEnum = "COMPLIANT"
	ComplianceStateNonCompliant ComplianceStateEnum = "NON_COMPLIANT"
	ComplianceStateWarning      ComplianceStateEnum = "WARNING"
)

var mappingComplianceStateEnum = map[string]ComplianceStateEnum{
	"UNKNOWN":       ComplianceStateUnknown,
	"COMPLIANT":     ComplianceStateCompliant,
	"NON_COMPLIANT": ComplianceStateNonCompliant,
	"WARNING":       ComplianceStateWarning,
}

var mappingComplianceStateEnumLowerCase = map[string]ComplianceStateEnum{
	"unknown":       ComplianceStateUnknown,
	"compliant":     ComplianceStateCompliant,
	"non_compliant": ComplianceStateNonCompliant,
	"warning":       ComplianceStateWarning,
}

// GetComplianceStateEnumValues Enumerates the set of values for ComplianceStateEnum
func GetComplianceStateEnumValues() []ComplianceStateEnum {
	values := make([]ComplianceStateEnum, 0)
	for _, v := range mappingComplianceStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceStateEnumStringValues Enumerates the set of values in String for ComplianceStateEnum
func GetComplianceStateEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"COMPLIANT",
		"NON_COMPLIANT",
		"WARNING",
	}
}

// GetMappingComplianceStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceStateEnum(val string) (ComplianceStateEnum, bool) {
	enum, ok := mappingComplianceStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
