// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// AssessorCheckActionNamesEnum Enum with underlying type: string
type AssessorCheckActionNamesEnum string

// Set of constants representing the allowable values for AssessorCheckActionNamesEnum
const (
	AssessorCheckActionNamesApprove       AssessorCheckActionNamesEnum = "APPROVE"
	AssessorCheckActionNamesDisapprove    AssessorCheckActionNamesEnum = "DISAPPROVE"
	AssessorCheckActionNamesAcknowledge   AssessorCheckActionNamesEnum = "ACKNOWLEDGE"
	AssessorCheckActionNamesUnacknowledge AssessorCheckActionNamesEnum = "UNACKNOWLEDGE"
)

var mappingAssessorCheckActionNamesEnum = map[string]AssessorCheckActionNamesEnum{
	"APPROVE":       AssessorCheckActionNamesApprove,
	"DISAPPROVE":    AssessorCheckActionNamesDisapprove,
	"ACKNOWLEDGE":   AssessorCheckActionNamesAcknowledge,
	"UNACKNOWLEDGE": AssessorCheckActionNamesUnacknowledge,
}

var mappingAssessorCheckActionNamesEnumLowerCase = map[string]AssessorCheckActionNamesEnum{
	"approve":       AssessorCheckActionNamesApprove,
	"disapprove":    AssessorCheckActionNamesDisapprove,
	"acknowledge":   AssessorCheckActionNamesAcknowledge,
	"unacknowledge": AssessorCheckActionNamesUnacknowledge,
}

// GetAssessorCheckActionNamesEnumValues Enumerates the set of values for AssessorCheckActionNamesEnum
func GetAssessorCheckActionNamesEnumValues() []AssessorCheckActionNamesEnum {
	values := make([]AssessorCheckActionNamesEnum, 0)
	for _, v := range mappingAssessorCheckActionNamesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessorCheckActionNamesEnumStringValues Enumerates the set of values in String for AssessorCheckActionNamesEnum
func GetAssessorCheckActionNamesEnumStringValues() []string {
	return []string{
		"APPROVE",
		"DISAPPROVE",
		"ACKNOWLEDGE",
		"UNACKNOWLEDGE",
	}
}

// GetMappingAssessorCheckActionNamesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessorCheckActionNamesEnum(val string) (AssessorCheckActionNamesEnum, bool) {
	enum, ok := mappingAssessorCheckActionNamesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
