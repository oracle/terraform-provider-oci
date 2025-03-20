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

// ComplianceRecordSortByEnum Enum with underlying type: string
type ComplianceRecordSortByEnum string

// Set of constants representing the allowable values for ComplianceRecordSortByEnum
const (
	ComplianceRecordSortByTimeCreated ComplianceRecordSortByEnum = "timeCreated"
	ComplianceRecordSortByDisplayName ComplianceRecordSortByEnum = "displayName"
)

var mappingComplianceRecordSortByEnum = map[string]ComplianceRecordSortByEnum{
	"timeCreated": ComplianceRecordSortByTimeCreated,
	"displayName": ComplianceRecordSortByDisplayName,
}

var mappingComplianceRecordSortByEnumLowerCase = map[string]ComplianceRecordSortByEnum{
	"timecreated": ComplianceRecordSortByTimeCreated,
	"displayname": ComplianceRecordSortByDisplayName,
}

// GetComplianceRecordSortByEnumValues Enumerates the set of values for ComplianceRecordSortByEnum
func GetComplianceRecordSortByEnumValues() []ComplianceRecordSortByEnum {
	values := make([]ComplianceRecordSortByEnum, 0)
	for _, v := range mappingComplianceRecordSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceRecordSortByEnumStringValues Enumerates the set of values in String for ComplianceRecordSortByEnum
func GetComplianceRecordSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingComplianceRecordSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceRecordSortByEnum(val string) (ComplianceRecordSortByEnum, bool) {
	enum, ok := mappingComplianceRecordSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
