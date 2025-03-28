// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Process Automation
//
// Process Automation helps you to rapidly design, automate, and manage business processes in the cloud. With the Process Automation design-time (Designer) and the runtime (Workspace) environments, you can easily create, develop, manage, test, and monitor process applications and their components.
//

package opa

import (
	"strings"
)

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted   OperationStatusEnum = "ACCEPTED"
	OperationStatusInProgress OperationStatusEnum = "IN_PROGRESS"
	OperationStatusWaiting    OperationStatusEnum = "WAITING"
	OperationStatusFailed     OperationStatusEnum = "FAILED"
	OperationStatusSucceeded  OperationStatusEnum = "SUCCEEDED"
	OperationStatusCanceling  OperationStatusEnum = "CANCELING"
	OperationStatusCanceled   OperationStatusEnum = "CANCELED"
)

var mappingOperationStatusEnum = map[string]OperationStatusEnum{
	"ACCEPTED":    OperationStatusAccepted,
	"IN_PROGRESS": OperationStatusInProgress,
	"WAITING":     OperationStatusWaiting,
	"FAILED":      OperationStatusFailed,
	"SUCCEEDED":   OperationStatusSucceeded,
	"CANCELING":   OperationStatusCanceling,
	"CANCELED":    OperationStatusCanceled,
}

var mappingOperationStatusEnumLowerCase = map[string]OperationStatusEnum{
	"accepted":    OperationStatusAccepted,
	"in_progress": OperationStatusInProgress,
	"waiting":     OperationStatusWaiting,
	"failed":      OperationStatusFailed,
	"succeeded":   OperationStatusSucceeded,
	"canceling":   OperationStatusCanceling,
	"canceled":    OperationStatusCanceled,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationStatusEnumStringValues Enumerates the set of values in String for OperationStatusEnum
func GetOperationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationStatusEnum(val string) (OperationStatusEnum, bool) {
	enum, ok := mappingOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
