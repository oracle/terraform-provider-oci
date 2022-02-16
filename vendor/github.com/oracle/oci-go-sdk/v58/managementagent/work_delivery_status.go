// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// WorkDeliveryStatusEnum Enum with underlying type: string
type WorkDeliveryStatusEnum string

// Set of constants representing the allowable values for WorkDeliveryStatusEnum
const (
	WorkDeliveryStatusAccepted   WorkDeliveryStatusEnum = "ACCEPTED"
	WorkDeliveryStatusInProgress WorkDeliveryStatusEnum = "IN_PROGRESS"
	WorkDeliveryStatusFailed     WorkDeliveryStatusEnum = "FAILED"
	WorkDeliveryStatusSucceeded  WorkDeliveryStatusEnum = "SUCCEEDED"
	WorkDeliveryStatusCanceling  WorkDeliveryStatusEnum = "CANCELING"
	WorkDeliveryStatusCanceled   WorkDeliveryStatusEnum = "CANCELED"
)

var mappingWorkDeliveryStatusEnum = map[string]WorkDeliveryStatusEnum{
	"ACCEPTED":    WorkDeliveryStatusAccepted,
	"IN_PROGRESS": WorkDeliveryStatusInProgress,
	"FAILED":      WorkDeliveryStatusFailed,
	"SUCCEEDED":   WorkDeliveryStatusSucceeded,
	"CANCELING":   WorkDeliveryStatusCanceling,
	"CANCELED":    WorkDeliveryStatusCanceled,
}

// GetWorkDeliveryStatusEnumValues Enumerates the set of values for WorkDeliveryStatusEnum
func GetWorkDeliveryStatusEnumValues() []WorkDeliveryStatusEnum {
	values := make([]WorkDeliveryStatusEnum, 0)
	for _, v := range mappingWorkDeliveryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkDeliveryStatusEnumStringValues Enumerates the set of values in String for WorkDeliveryStatusEnum
func GetWorkDeliveryStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingWorkDeliveryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkDeliveryStatusEnum(val string) (WorkDeliveryStatusEnum, bool) {
	mappingWorkDeliveryStatusEnumIgnoreCase := make(map[string]WorkDeliveryStatusEnum)
	for k, v := range mappingWorkDeliveryStatusEnum {
		mappingWorkDeliveryStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingWorkDeliveryStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
