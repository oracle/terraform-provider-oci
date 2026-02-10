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

// AssessmentLifecycleStatesEnum Enum with underlying type: string
type AssessmentLifecycleStatesEnum string

// Set of constants representing the allowable values for AssessmentLifecycleStatesEnum
const (
	AssessmentLifecycleStatesCreating       AssessmentLifecycleStatesEnum = "CREATING"
	AssessmentLifecycleStatesUpdating       AssessmentLifecycleStatesEnum = "UPDATING"
	AssessmentLifecycleStatesActive         AssessmentLifecycleStatesEnum = "ACTIVE"
	AssessmentLifecycleStatesSucceeded      AssessmentLifecycleStatesEnum = "SUCCEEDED"
	AssessmentLifecycleStatesInProgress     AssessmentLifecycleStatesEnum = "IN_PROGRESS"
	AssessmentLifecycleStatesNeedsAttention AssessmentLifecycleStatesEnum = "NEEDS_ATTENTION"
	AssessmentLifecycleStatesDeleting       AssessmentLifecycleStatesEnum = "DELETING"
	AssessmentLifecycleStatesDeleted        AssessmentLifecycleStatesEnum = "DELETED"
	AssessmentLifecycleStatesFailed         AssessmentLifecycleStatesEnum = "FAILED"
)

var mappingAssessmentLifecycleStatesEnum = map[string]AssessmentLifecycleStatesEnum{
	"CREATING":        AssessmentLifecycleStatesCreating,
	"UPDATING":        AssessmentLifecycleStatesUpdating,
	"ACTIVE":          AssessmentLifecycleStatesActive,
	"SUCCEEDED":       AssessmentLifecycleStatesSucceeded,
	"IN_PROGRESS":     AssessmentLifecycleStatesInProgress,
	"NEEDS_ATTENTION": AssessmentLifecycleStatesNeedsAttention,
	"DELETING":        AssessmentLifecycleStatesDeleting,
	"DELETED":         AssessmentLifecycleStatesDeleted,
	"FAILED":          AssessmentLifecycleStatesFailed,
}

var mappingAssessmentLifecycleStatesEnumLowerCase = map[string]AssessmentLifecycleStatesEnum{
	"creating":        AssessmentLifecycleStatesCreating,
	"updating":        AssessmentLifecycleStatesUpdating,
	"active":          AssessmentLifecycleStatesActive,
	"succeeded":       AssessmentLifecycleStatesSucceeded,
	"in_progress":     AssessmentLifecycleStatesInProgress,
	"needs_attention": AssessmentLifecycleStatesNeedsAttention,
	"deleting":        AssessmentLifecycleStatesDeleting,
	"deleted":         AssessmentLifecycleStatesDeleted,
	"failed":          AssessmentLifecycleStatesFailed,
}

// GetAssessmentLifecycleStatesEnumValues Enumerates the set of values for AssessmentLifecycleStatesEnum
func GetAssessmentLifecycleStatesEnumValues() []AssessmentLifecycleStatesEnum {
	values := make([]AssessmentLifecycleStatesEnum, 0)
	for _, v := range mappingAssessmentLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssessmentLifecycleStatesEnumStringValues Enumerates the set of values in String for AssessmentLifecycleStatesEnum
func GetAssessmentLifecycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"SUCCEEDED",
		"IN_PROGRESS",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAssessmentLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssessmentLifecycleStatesEnum(val string) (AssessmentLifecycleStatesEnum, bool) {
	enum, ok := mappingAssessmentLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
