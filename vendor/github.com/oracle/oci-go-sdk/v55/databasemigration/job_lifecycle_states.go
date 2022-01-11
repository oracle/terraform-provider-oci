// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

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

var mappingJobLifecycleStates = map[string]JobLifecycleStatesEnum{
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

// GetJobLifecycleStatesEnumValues Enumerates the set of values for JobLifecycleStatesEnum
func GetJobLifecycleStatesEnumValues() []JobLifecycleStatesEnum {
	values := make([]JobLifecycleStatesEnum, 0)
	for _, v := range mappingJobLifecycleStates {
		values = append(values, v)
	}
	return values
}
