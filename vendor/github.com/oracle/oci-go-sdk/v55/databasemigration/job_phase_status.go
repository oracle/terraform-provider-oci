// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// JobPhaseStatusEnum Enum with underlying type: string
type JobPhaseStatusEnum string

// Set of constants representing the allowable values for JobPhaseStatusEnum
const (
	JobPhaseStatusPending   JobPhaseStatusEnum = "PENDING"
	JobPhaseStatusStarted   JobPhaseStatusEnum = "STARTED"
	JobPhaseStatusCompleted JobPhaseStatusEnum = "COMPLETED"
	JobPhaseStatusFailed    JobPhaseStatusEnum = "FAILED"
)

var mappingJobPhaseStatus = map[string]JobPhaseStatusEnum{
	"PENDING":   JobPhaseStatusPending,
	"STARTED":   JobPhaseStatusStarted,
	"COMPLETED": JobPhaseStatusCompleted,
	"FAILED":    JobPhaseStatusFailed,
}

// GetJobPhaseStatusEnumValues Enumerates the set of values for JobPhaseStatusEnum
func GetJobPhaseStatusEnumValues() []JobPhaseStatusEnum {
	values := make([]JobPhaseStatusEnum, 0)
	for _, v := range mappingJobPhaseStatus {
		values = append(values, v)
	}
	return values
}
