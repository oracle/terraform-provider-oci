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

// ComplianceLevelEnum Enum with underlying type: string
type ComplianceLevelEnum string

// Set of constants representing the allowable values for ComplianceLevelEnum
const (
	ComplianceLevelFleet  ComplianceLevelEnum = "FLEET"
	ComplianceLevelTarget ComplianceLevelEnum = "TARGET"
)

var mappingComplianceLevelEnum = map[string]ComplianceLevelEnum{
	"FLEET":  ComplianceLevelFleet,
	"TARGET": ComplianceLevelTarget,
}

var mappingComplianceLevelEnumLowerCase = map[string]ComplianceLevelEnum{
	"fleet":  ComplianceLevelFleet,
	"target": ComplianceLevelTarget,
}

// GetComplianceLevelEnumValues Enumerates the set of values for ComplianceLevelEnum
func GetComplianceLevelEnumValues() []ComplianceLevelEnum {
	values := make([]ComplianceLevelEnum, 0)
	for _, v := range mappingComplianceLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceLevelEnumStringValues Enumerates the set of values in String for ComplianceLevelEnum
func GetComplianceLevelEnumStringValues() []string {
	return []string{
		"FLEET",
		"TARGET",
	}
}

// GetMappingComplianceLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceLevelEnum(val string) (ComplianceLevelEnum, bool) {
	enum, ok := mappingComplianceLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
