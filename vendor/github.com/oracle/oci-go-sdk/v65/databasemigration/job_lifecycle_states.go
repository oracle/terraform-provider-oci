// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// JobLifecycleStatesEnum Enum with underlying type: string
type JobLifecycleStatesEnum string

// Set of constants representing the allowable values for JobLifecycleStatesEnum
const (
	JobLifecycleStatesAccepted   JobLifecycleStatesEnum = "ACCEPTED"
	JobLifecycleStatesInProgress JobLifecycleStatesEnum = "IN_PROGRESS"
	JobLifecycleStatesUnknown    JobLifecycleStatesEnum = "UNKNOWN"
	JobLifecycleStatesTerminated JobLifecycleStatesEnum = "TERMINATED"
	JobLifecycleStatesFailed     JobLifecycleStatesEnum = "FAILED"
	JobLifecycleStatesSucceeded  JobLifecycleStatesEnum = "SUCCEEDED"
	JobLifecycleStatesWaiting    JobLifecycleStatesEnum = "WAITING"
	JobLifecycleStatesCanceling  JobLifecycleStatesEnum = "CANCELING"
	JobLifecycleStatesCanceled   JobLifecycleStatesEnum = "CANCELED"
)

var mappingJobLifecycleStatesEnum = map[string]JobLifecycleStatesEnum{
	"ACCEPTED":    JobLifecycleStatesAccepted,
	"IN_PROGRESS": JobLifecycleStatesInProgress,
	"UNKNOWN":     JobLifecycleStatesUnknown,
	"TERMINATED":  JobLifecycleStatesTerminated,
	"FAILED":      JobLifecycleStatesFailed,
	"SUCCEEDED":   JobLifecycleStatesSucceeded,
	"WAITING":     JobLifecycleStatesWaiting,
	"CANCELING":   JobLifecycleStatesCanceling,
	"CANCELED":    JobLifecycleStatesCanceled,
}

var mappingJobLifecycleStatesEnumLowerCase = map[string]JobLifecycleStatesEnum{
	"accepted":    JobLifecycleStatesAccepted,
	"in_progress": JobLifecycleStatesInProgress,
	"unknown":     JobLifecycleStatesUnknown,
	"terminated":  JobLifecycleStatesTerminated,
	"failed":      JobLifecycleStatesFailed,
	"succeeded":   JobLifecycleStatesSucceeded,
	"waiting":     JobLifecycleStatesWaiting,
	"canceling":   JobLifecycleStatesCanceling,
	"canceled":    JobLifecycleStatesCanceled,
}

// GetJobLifecycleStatesEnumValues Enumerates the set of values for JobLifecycleStatesEnum
func GetJobLifecycleStatesEnumValues() []JobLifecycleStatesEnum {
	values := make([]JobLifecycleStatesEnum, 0)
	for _, v := range mappingJobLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetJobLifecycleStatesEnumStringValues Enumerates the set of values in String for JobLifecycleStatesEnum
func GetJobLifecycleStatesEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"UNKNOWN",
		"TERMINATED",
		"FAILED",
		"SUCCEEDED",
		"WAITING",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingJobLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobLifecycleStatesEnum(val string) (JobLifecycleStatesEnum, bool) {
	enum, ok := mappingJobLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
