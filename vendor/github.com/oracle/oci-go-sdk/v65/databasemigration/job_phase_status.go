// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// JobPhaseStatusEnum Enum with underlying type: string
type JobPhaseStatusEnum string

// Set of constants representing the allowable values for JobPhaseStatusEnum
const (
	JobPhaseStatusPending   JobPhaseStatusEnum = "PENDING"
	JobPhaseStatusStarted   JobPhaseStatusEnum = "STARTED"
	JobPhaseStatusCompleted JobPhaseStatusEnum = "COMPLETED"
	JobPhaseStatusFailed    JobPhaseStatusEnum = "FAILED"
)

var mappingJobPhaseStatusEnum = map[string]JobPhaseStatusEnum{
	"PENDING":   JobPhaseStatusPending,
	"STARTED":   JobPhaseStatusStarted,
	"COMPLETED": JobPhaseStatusCompleted,
	"FAILED":    JobPhaseStatusFailed,
}

var mappingJobPhaseStatusEnumLowerCase = map[string]JobPhaseStatusEnum{
	"pending":   JobPhaseStatusPending,
	"started":   JobPhaseStatusStarted,
	"completed": JobPhaseStatusCompleted,
	"failed":    JobPhaseStatusFailed,
}

// GetJobPhaseStatusEnumValues Enumerates the set of values for JobPhaseStatusEnum
func GetJobPhaseStatusEnumValues() []JobPhaseStatusEnum {
	values := make([]JobPhaseStatusEnum, 0)
	for _, v := range mappingJobPhaseStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJobPhaseStatusEnumStringValues Enumerates the set of values in String for JobPhaseStatusEnum
func GetJobPhaseStatusEnumStringValues() []string {
	return []string{
		"PENDING",
		"STARTED",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingJobPhaseStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobPhaseStatusEnum(val string) (JobPhaseStatusEnum, bool) {
	enum, ok := mappingJobPhaseStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
