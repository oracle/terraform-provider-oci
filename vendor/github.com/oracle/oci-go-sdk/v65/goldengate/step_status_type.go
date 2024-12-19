// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// StepStatusTypeEnum Enum with underlying type: string
type StepStatusTypeEnum string

// Set of constants representing the allowable values for StepStatusTypeEnum
const (
	StepStatusTypePending    StepStatusTypeEnum = "PENDING"
	StepStatusTypeInProgress StepStatusTypeEnum = "IN_PROGRESS"
	StepStatusTypeCompleted  StepStatusTypeEnum = "COMPLETED"
	StepStatusTypeFailed     StepStatusTypeEnum = "FAILED"
)

var mappingStepStatusTypeEnum = map[string]StepStatusTypeEnum{
	"PENDING":     StepStatusTypePending,
	"IN_PROGRESS": StepStatusTypeInProgress,
	"COMPLETED":   StepStatusTypeCompleted,
	"FAILED":      StepStatusTypeFailed,
}

var mappingStepStatusTypeEnumLowerCase = map[string]StepStatusTypeEnum{
	"pending":     StepStatusTypePending,
	"in_progress": StepStatusTypeInProgress,
	"completed":   StepStatusTypeCompleted,
	"failed":      StepStatusTypeFailed,
}

// GetStepStatusTypeEnumValues Enumerates the set of values for StepStatusTypeEnum
func GetStepStatusTypeEnumValues() []StepStatusTypeEnum {
	values := make([]StepStatusTypeEnum, 0)
	for _, v := range mappingStepStatusTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStepStatusTypeEnumStringValues Enumerates the set of values in String for StepStatusTypeEnum
func GetStepStatusTypeEnumStringValues() []string {
	return []string{
		"PENDING",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingStepStatusTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStepStatusTypeEnum(val string) (StepStatusTypeEnum, bool) {
	enum, ok := mappingStepStatusTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
