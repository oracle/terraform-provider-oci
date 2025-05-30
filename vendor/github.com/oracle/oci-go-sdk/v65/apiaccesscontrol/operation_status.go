// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

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
	OperationStatusNeedsAttention OperationStatusEnum = "NEEDS_ATTENTION"
	OperationStatusFailed         OperationStatusEnum = "FAILED"
	OperationStatusSucceeded      OperationStatusEnum = "SUCCEEDED"
	OperationStatusCanceling      OperationStatusEnum = "CANCELING"
	OperationStatusCanceled       OperationStatusEnum = "CANCELED"
)

var mappingOperationStatusEnum = map[string]OperationStatusEnum{
	"ACCEPTED":        OperationStatusAccepted,
	"IN_PROGRESS":     OperationStatusInProgress,
	"WAITING":         OperationStatusWaiting,
	"NEEDS_ATTENTION": OperationStatusNeedsAttention,
	"FAILED":          OperationStatusFailed,
	"SUCCEEDED":       OperationStatusSucceeded,
	"CANCELING":       OperationStatusCanceling,
	"CANCELED":        OperationStatusCanceled,
}

var mappingOperationStatusEnumLowerCase = map[string]OperationStatusEnum{
	"accepted":        OperationStatusAccepted,
	"in_progress":     OperationStatusInProgress,
	"waiting":         OperationStatusWaiting,
	"needs_attention": OperationStatusNeedsAttention,
	"failed":          OperationStatusFailed,
	"succeeded":       OperationStatusSucceeded,
	"canceling":       OperationStatusCanceling,
	"canceled":        OperationStatusCanceled,
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
		"NEEDS_ATTENTION",
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
