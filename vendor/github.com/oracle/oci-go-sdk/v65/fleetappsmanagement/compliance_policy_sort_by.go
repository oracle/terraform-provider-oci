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

// CompliancePolicySortByEnum Enum with underlying type: string
type CompliancePolicySortByEnum string

// Set of constants representing the allowable values for CompliancePolicySortByEnum
const (
	CompliancePolicySortByTimeCreated CompliancePolicySortByEnum = "timeCreated"
	CompliancePolicySortByDisplayName CompliancePolicySortByEnum = "displayName"
)

var mappingCompliancePolicySortByEnum = map[string]CompliancePolicySortByEnum{
	"timeCreated": CompliancePolicySortByTimeCreated,
	"displayName": CompliancePolicySortByDisplayName,
}

var mappingCompliancePolicySortByEnumLowerCase = map[string]CompliancePolicySortByEnum{
	"timecreated": CompliancePolicySortByTimeCreated,
	"displayname": CompliancePolicySortByDisplayName,
}

// GetCompliancePolicySortByEnumValues Enumerates the set of values for CompliancePolicySortByEnum
func GetCompliancePolicySortByEnumValues() []CompliancePolicySortByEnum {
	values := make([]CompliancePolicySortByEnum, 0)
	for _, v := range mappingCompliancePolicySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetCompliancePolicySortByEnumStringValues Enumerates the set of values in String for CompliancePolicySortByEnum
func GetCompliancePolicySortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingCompliancePolicySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompliancePolicySortByEnum(val string) (CompliancePolicySortByEnum, bool) {
	enum, ok := mappingCompliancePolicySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
