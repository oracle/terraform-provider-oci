// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// JobExecutionStateEnum Enum with underlying type: string
type JobExecutionStateEnum string

// Set of constants representing the allowable values for JobExecutionStateEnum
const (
	JobExecutionStateCreated    JobExecutionStateEnum = "CREATED"
	JobExecutionStateInProgress JobExecutionStateEnum = "IN_PROGRESS"
	JobExecutionStateInactive   JobExecutionStateEnum = "INACTIVE"
	JobExecutionStateFailed     JobExecutionStateEnum = "FAILED"
	JobExecutionStateSucceeded  JobExecutionStateEnum = "SUCCEEDED"
	JobExecutionStateCanceled   JobExecutionStateEnum = "CANCELED"
)

var mappingJobExecutionState = map[string]JobExecutionStateEnum{
	"CREATED":     JobExecutionStateCreated,
	"IN_PROGRESS": JobExecutionStateInProgress,
	"INACTIVE":    JobExecutionStateInactive,
	"FAILED":      JobExecutionStateFailed,
	"SUCCEEDED":   JobExecutionStateSucceeded,
	"CANCELED":    JobExecutionStateCanceled,
}

// GetJobExecutionStateEnumValues Enumerates the set of values for JobExecutionStateEnum
func GetJobExecutionStateEnumValues() []JobExecutionStateEnum {
	values := make([]JobExecutionStateEnum, 0)
	for _, v := range mappingJobExecutionState {
		values = append(values, v)
	}
	return values
}
