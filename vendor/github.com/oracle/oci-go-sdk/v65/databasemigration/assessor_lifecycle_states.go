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

// AssessorLifecycleStatesEnum Enum with underlying type: string
type AssessorLifecycleStatesEnum string

// Set of constants representing the allowable values for AssessorLifecycleStatesEnum
const (
	AssessorLifecycleStatesAccepted       AssessorLifecycleStatesEnum = "ACCEPTED"
	AssessorLifecycleStatesInProgress     AssessorLifecycleStatesEnum = "IN_PROGRESS"
	AssessorLifecycleStatesSucceeded      AssessorLifecycleStatesEnum = "SUCCEEDED"
	AssessorLifecycleStatesNeedsAttention AssessorLifecycleStatesEnum = "NEEDS_ATTENTION"
	AssessorLifecycleStatesFailed         AssessorLifecycleStatesEnum = "FAILED"
)

var mappingAssessorLifecycleStatesEnum = map[string]AssessorLifecycleStatesEnum{
	"ACCEPTED":        AssessorLifecycleStatesAccepted,
	"IN_PROGRESS":     AssessorLifecycleStatesInProgress,
	"SUCCEEDED":       AssessorLifecycleStatesSucceeded,
	"NEEDS_ATTENTION": AssessorLifecycleStatesNeedsAttention,
	"FAILED":          AssessorLifecycleStatesFailed,
}

var mappingAssessorLifecycleStatesEnumLowerCase = map[string]AssessorLifecycleStatesEnum{
	"accepted":        AssessorLifecycleStatesAccepted,
	"in_progress":     AssessorLifecycleStatesInProgress,
	"succeeded":       AssessorLifecycleStatesSucceeded,
	"needs_attention": AssessorLifecycleStatesNeedsAttention,
	"failed":          AssessorLifecycleStatesFailed,
}

// GetAssessorLifecycleStatesEnumValues Enumerates the set of values for AssessorLifecycleStatesEnum
func GetAssessorLifecycleStatesEnumValues() []AssessorLifecycleStatesEnum {
	values := make([]AssessorLifecycleStatesEnum, 0)
	for _, v := range mappingAssessorLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessorLifecycleStatesEnumStringValues Enumerates the set of values in String for AssessorLifecycleStatesEnum
func GetAssessorLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingAssessorLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessorLifecycleStatesEnum(val string) (AssessorLifecycleStatesEnum, bool) {
	enum, ok := mappingAssessorLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
