// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// RequestQueueLifecycleStateEnum Enum with underlying type: string
type RequestQueueLifecycleStateEnum string

// Set of constants representing the allowable values for RequestQueueLifecycleStateEnum
const (
	RequestQueueLifecycleStateWaiting    RequestQueueLifecycleStateEnum = "WAITING"
	RequestQueueLifecycleStateAccepted   RequestQueueLifecycleStateEnum = "ACCEPTED"
	RequestQueueLifecycleStateInProgress RequestQueueLifecycleStateEnum = "IN_PROGRESS"
	RequestQueueLifecycleStateSucceeded  RequestQueueLifecycleStateEnum = "SUCCEEDED"
	RequestQueueLifecycleStateFailed     RequestQueueLifecycleStateEnum = "FAILED"
)

var mappingRequestQueueLifecycleStateEnum = map[string]RequestQueueLifecycleStateEnum{
	"WAITING":     RequestQueueLifecycleStateWaiting,
	"ACCEPTED":    RequestQueueLifecycleStateAccepted,
	"IN_PROGRESS": RequestQueueLifecycleStateInProgress,
	"SUCCEEDED":   RequestQueueLifecycleStateSucceeded,
	"FAILED":      RequestQueueLifecycleStateFailed,
}

var mappingRequestQueueLifecycleStateEnumLowerCase = map[string]RequestQueueLifecycleStateEnum{
	"waiting":     RequestQueueLifecycleStateWaiting,
	"accepted":    RequestQueueLifecycleStateAccepted,
	"in_progress": RequestQueueLifecycleStateInProgress,
	"succeeded":   RequestQueueLifecycleStateSucceeded,
	"failed":      RequestQueueLifecycleStateFailed,
}

// GetRequestQueueLifecycleStateEnumValues Enumerates the set of values for RequestQueueLifecycleStateEnum
func GetRequestQueueLifecycleStateEnumValues() []RequestQueueLifecycleStateEnum {
	values := make([]RequestQueueLifecycleStateEnum, 0)
	for _, v := range mappingRequestQueueLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestQueueLifecycleStateEnumStringValues Enumerates the set of values in String for RequestQueueLifecycleStateEnum
func GetRequestQueueLifecycleStateEnumStringValues() []string {
	return []string{
		"WAITING",
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingRequestQueueLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestQueueLifecycleStateEnum(val string) (RequestQueueLifecycleStateEnum, bool) {
	enum, ok := mappingRequestQueueLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
