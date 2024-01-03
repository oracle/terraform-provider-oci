// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// StatementLifecycleStateEnum Enum with underlying type: string
type StatementLifecycleStateEnum string

// Set of constants representing the allowable values for StatementLifecycleStateEnum
const (
	StatementLifecycleStateAccepted   StatementLifecycleStateEnum = "ACCEPTED"
	StatementLifecycleStateCancelling StatementLifecycleStateEnum = "CANCELLING"
	StatementLifecycleStateCancelled  StatementLifecycleStateEnum = "CANCELLED"
	StatementLifecycleStateFailed     StatementLifecycleStateEnum = "FAILED"
	StatementLifecycleStateInProgress StatementLifecycleStateEnum = "IN_PROGRESS"
	StatementLifecycleStateSucceeded  StatementLifecycleStateEnum = "SUCCEEDED"
)

var mappingStatementLifecycleStateEnum = map[string]StatementLifecycleStateEnum{
	"ACCEPTED":    StatementLifecycleStateAccepted,
	"CANCELLING":  StatementLifecycleStateCancelling,
	"CANCELLED":   StatementLifecycleStateCancelled,
	"FAILED":      StatementLifecycleStateFailed,
	"IN_PROGRESS": StatementLifecycleStateInProgress,
	"SUCCEEDED":   StatementLifecycleStateSucceeded,
}

var mappingStatementLifecycleStateEnumLowerCase = map[string]StatementLifecycleStateEnum{
	"accepted":    StatementLifecycleStateAccepted,
	"cancelling":  StatementLifecycleStateCancelling,
	"cancelled":   StatementLifecycleStateCancelled,
	"failed":      StatementLifecycleStateFailed,
	"in_progress": StatementLifecycleStateInProgress,
	"succeeded":   StatementLifecycleStateSucceeded,
}

// GetStatementLifecycleStateEnumValues Enumerates the set of values for StatementLifecycleStateEnum
func GetStatementLifecycleStateEnumValues() []StatementLifecycleStateEnum {
	values := make([]StatementLifecycleStateEnum, 0)
	for _, v := range mappingStatementLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetStatementLifecycleStateEnumStringValues Enumerates the set of values in String for StatementLifecycleStateEnum
func GetStatementLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"CANCELLING",
		"CANCELLED",
		"FAILED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

// GetMappingStatementLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatementLifecycleStateEnum(val string) (StatementLifecycleStateEnum, bool) {
	enum, ok := mappingStatementLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
