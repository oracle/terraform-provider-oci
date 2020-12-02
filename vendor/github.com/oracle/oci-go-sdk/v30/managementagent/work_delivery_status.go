// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

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

var mappingWorkDeliveryStatus = map[string]WorkDeliveryStatusEnum{
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
	for _, v := range mappingWorkDeliveryStatus {
		values = append(values, v)
	}
	return values
}
