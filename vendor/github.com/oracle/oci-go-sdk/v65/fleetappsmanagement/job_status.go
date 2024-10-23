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

// JobStatusEnum Enum with underlying type: string
type JobStatusEnum string

// Set of constants representing the allowable values for JobStatusEnum
const (
	JobStatusAccepted      JobStatusEnum = "ACCEPTED"
	JobStatusWaiting       JobStatusEnum = "WAITING"
	JobStatusInProgress    JobStatusEnum = "IN_PROGRESS"
	JobStatusFailed        JobStatusEnum = "FAILED"
	JobStatusSucceeded     JobStatusEnum = "SUCCEEDED"
	JobStatusCanceled      JobStatusEnum = "CANCELED"
	JobStatusSkipped       JobStatusEnum = "SKIPPED"
	JobStatusIgnored       JobStatusEnum = "IGNORED"
	JobStatusNotApplicable JobStatusEnum = "NOT_APPLICABLE"
	JobStatusAborted       JobStatusEnum = "ABORTED"
	JobStatusTimedOut      JobStatusEnum = "TIMED_OUT"
	JobStatusPaused        JobStatusEnum = "PAUSED"
)

var mappingJobStatusEnum = map[string]JobStatusEnum{
	"ACCEPTED":       JobStatusAccepted,
	"WAITING":        JobStatusWaiting,
	"IN_PROGRESS":    JobStatusInProgress,
	"FAILED":         JobStatusFailed,
	"SUCCEEDED":      JobStatusSucceeded,
	"CANCELED":       JobStatusCanceled,
	"SKIPPED":        JobStatusSkipped,
	"IGNORED":        JobStatusIgnored,
	"NOT_APPLICABLE": JobStatusNotApplicable,
	"ABORTED":        JobStatusAborted,
	"TIMED_OUT":      JobStatusTimedOut,
	"PAUSED":         JobStatusPaused,
}

var mappingJobStatusEnumLowerCase = map[string]JobStatusEnum{
	"accepted":       JobStatusAccepted,
	"waiting":        JobStatusWaiting,
	"in_progress":    JobStatusInProgress,
	"failed":         JobStatusFailed,
	"succeeded":      JobStatusSucceeded,
	"canceled":       JobStatusCanceled,
	"skipped":        JobStatusSkipped,
	"ignored":        JobStatusIgnored,
	"not_applicable": JobStatusNotApplicable,
	"aborted":        JobStatusAborted,
	"timed_out":      JobStatusTimedOut,
	"paused":         JobStatusPaused,
}

// GetJobStatusEnumValues Enumerates the set of values for JobStatusEnum
func GetJobStatusEnumValues() []JobStatusEnum {
	values := make([]JobStatusEnum, 0)
	for _, v := range mappingJobStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJobStatusEnumStringValues Enumerates the set of values in String for JobStatusEnum
func GetJobStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"WAITING",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"SKIPPED",
		"IGNORED",
		"NOT_APPLICABLE",
		"ABORTED",
		"TIMED_OUT",
		"PAUSED",
	}
}

// GetMappingJobStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobStatusEnum(val string) (JobStatusEnum, bool) {
	enum, ok := mappingJobStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
