// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted       OperationStatusEnum = "ACCEPTED"
	OperationStatusInProgress     OperationStatusEnum = "IN_PROGRESS"
	OperationStatusWaiting        OperationStatusEnum = "WAITING"
	OperationStatusCanceling      OperationStatusEnum = "CANCELING"
	OperationStatusCanceled       OperationStatusEnum = "CANCELED"
	OperationStatusSucceeded      OperationStatusEnum = "SUCCEEDED"
	OperationStatusFailed         OperationStatusEnum = "FAILED"
	OperationStatusNeedsAttention OperationStatusEnum = "NEEDS_ATTENTION"
)

var mappingOperationStatusEnum = map[string]OperationStatusEnum{
	"ACCEPTED":        OperationStatusAccepted,
	"IN_PROGRESS":     OperationStatusInProgress,
	"WAITING":         OperationStatusWaiting,
	"CANCELING":       OperationStatusCanceling,
	"CANCELED":        OperationStatusCanceled,
	"SUCCEEDED":       OperationStatusSucceeded,
	"FAILED":          OperationStatusFailed,
	"NEEDS_ATTENTION": OperationStatusNeedsAttention,
}

var mappingOperationStatusEnumLowerCase = map[string]OperationStatusEnum{
	"accepted":        OperationStatusAccepted,
	"in_progress":     OperationStatusInProgress,
	"waiting":         OperationStatusWaiting,
	"canceling":       OperationStatusCanceling,
	"canceled":        OperationStatusCanceled,
	"succeeded":       OperationStatusSucceeded,
	"failed":          OperationStatusFailed,
	"needs_attention": OperationStatusNeedsAttention,
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
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationStatusEnum(val string) (OperationStatusEnum, bool) {
	enum, ok := mappingOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
