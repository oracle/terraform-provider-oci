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

// AssessorCheckStatesEnum Enum with underlying type: string
type AssessorCheckStatesEnum string

// Set of constants representing the allowable values for AssessorCheckStatesEnum
const (
	AssessorCheckStatesPendingApproval   AssessorCheckStatesEnum = "PENDING_APPROVAL"
	AssessorCheckStatesPendingResolution AssessorCheckStatesEnum = "PENDING_RESOLUTION"
	AssessorCheckStatesPendingValidation AssessorCheckStatesEnum = "PENDING_VALIDATION"
	AssessorCheckStatesNotAcknowledged   AssessorCheckStatesEnum = "NOT_ACKNOWLEDGED"
	AssessorCheckStatesApproved          AssessorCheckStatesEnum = "APPROVED"
	AssessorCheckStatesAcknowledged      AssessorCheckStatesEnum = "ACKNOWLEDGED"
	AssessorCheckStatesValidated         AssessorCheckStatesEnum = "VALIDATED"
	AssessorCheckStatesPassed            AssessorCheckStatesEnum = "PASSED"
	AssessorCheckStatesPending           AssessorCheckStatesEnum = "PENDING"
	AssessorCheckStatesStarted           AssessorCheckStatesEnum = "STARTED"
	AssessorCheckStatesCompleted         AssessorCheckStatesEnum = "COMPLETED"
	AssessorCheckStatesFailed            AssessorCheckStatesEnum = "FAILED"
)

var mappingAssessorCheckStatesEnum = map[string]AssessorCheckStatesEnum{
	"PENDING_APPROVAL":   AssessorCheckStatesPendingApproval,
	"PENDING_RESOLUTION": AssessorCheckStatesPendingResolution,
	"PENDING_VALIDATION": AssessorCheckStatesPendingValidation,
	"NOT_ACKNOWLEDGED":   AssessorCheckStatesNotAcknowledged,
	"APPROVED":           AssessorCheckStatesApproved,
	"ACKNOWLEDGED":       AssessorCheckStatesAcknowledged,
	"VALIDATED":          AssessorCheckStatesValidated,
	"PASSED":             AssessorCheckStatesPassed,
	"PENDING":            AssessorCheckStatesPending,
	"STARTED":            AssessorCheckStatesStarted,
	"COMPLETED":          AssessorCheckStatesCompleted,
	"FAILED":             AssessorCheckStatesFailed,
}

var mappingAssessorCheckStatesEnumLowerCase = map[string]AssessorCheckStatesEnum{
	"pending_approval":   AssessorCheckStatesPendingApproval,
	"pending_resolution": AssessorCheckStatesPendingResolution,
	"pending_validation": AssessorCheckStatesPendingValidation,
	"not_acknowledged":   AssessorCheckStatesNotAcknowledged,
	"approved":           AssessorCheckStatesApproved,
	"acknowledged":       AssessorCheckStatesAcknowledged,
	"validated":          AssessorCheckStatesValidated,
	"passed":             AssessorCheckStatesPassed,
	"pending":            AssessorCheckStatesPending,
	"started":            AssessorCheckStatesStarted,
	"completed":          AssessorCheckStatesCompleted,
	"failed":             AssessorCheckStatesFailed,
}

// GetAssessorCheckStatesEnumValues Enumerates the set of values for AssessorCheckStatesEnum
func GetAssessorCheckStatesEnumValues() []AssessorCheckStatesEnum {
	values := make([]AssessorCheckStatesEnum, 0)
	for _, v := range mappingAssessorCheckStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessorCheckStatesEnumStringValues Enumerates the set of values in String for AssessorCheckStatesEnum
func GetAssessorCheckStatesEnumStringValues() []string {
	return []string{
		"PENDING_APPROVAL",
		"PENDING_RESOLUTION",
		"PENDING_VALIDATION",
		"NOT_ACKNOWLEDGED",
		"APPROVED",
		"ACKNOWLEDGED",
		"VALIDATED",
		"PASSED",
		"PENDING",
		"STARTED",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingAssessorCheckStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessorCheckStatesEnum(val string) (AssessorCheckStatesEnum, bool) {
	enum, ok := mappingAssessorCheckStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
