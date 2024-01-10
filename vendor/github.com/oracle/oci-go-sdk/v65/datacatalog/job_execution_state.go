// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// JobExecutionStateEnum Enum with underlying type: string
type JobExecutionStateEnum string

// Set of constants representing the allowable values for JobExecutionStateEnum
const (
	JobExecutionStateCreated               JobExecutionStateEnum = "CREATED"
	JobExecutionStateInProgress            JobExecutionStateEnum = "IN_PROGRESS"
	JobExecutionStateInactive              JobExecutionStateEnum = "INACTIVE"
	JobExecutionStateFailed                JobExecutionStateEnum = "FAILED"
	JobExecutionStateSucceeded             JobExecutionStateEnum = "SUCCEEDED"
	JobExecutionStateCanceled              JobExecutionStateEnum = "CANCELED"
	JobExecutionStateSucceededWithWarnings JobExecutionStateEnum = "SUCCEEDED_WITH_WARNINGS"
)

var mappingJobExecutionStateEnum = map[string]JobExecutionStateEnum{
	"CREATED":                 JobExecutionStateCreated,
	"IN_PROGRESS":             JobExecutionStateInProgress,
	"INACTIVE":                JobExecutionStateInactive,
	"FAILED":                  JobExecutionStateFailed,
	"SUCCEEDED":               JobExecutionStateSucceeded,
	"CANCELED":                JobExecutionStateCanceled,
	"SUCCEEDED_WITH_WARNINGS": JobExecutionStateSucceededWithWarnings,
}

var mappingJobExecutionStateEnumLowerCase = map[string]JobExecutionStateEnum{
	"created":                 JobExecutionStateCreated,
	"in_progress":             JobExecutionStateInProgress,
	"inactive":                JobExecutionStateInactive,
	"failed":                  JobExecutionStateFailed,
	"succeeded":               JobExecutionStateSucceeded,
	"canceled":                JobExecutionStateCanceled,
	"succeeded_with_warnings": JobExecutionStateSucceededWithWarnings,
}

// GetJobExecutionStateEnumValues Enumerates the set of values for JobExecutionStateEnum
func GetJobExecutionStateEnumValues() []JobExecutionStateEnum {
	values := make([]JobExecutionStateEnum, 0)
	for _, v := range mappingJobExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJobExecutionStateEnumStringValues Enumerates the set of values in String for JobExecutionStateEnum
func GetJobExecutionStateEnumStringValues() []string {
	return []string{
		"CREATED",
		"IN_PROGRESS",
		"INACTIVE",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"SUCCEEDED_WITH_WARNINGS",
	}
}

// GetMappingJobExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobExecutionStateEnum(val string) (JobExecutionStateEnum, bool) {
	enum, ok := mappingJobExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
