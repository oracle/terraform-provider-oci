// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling Management API
//
// Use Data Labeling Management API to create, list, edit & delete datasets.
//

package datalabelingservice

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted   OperationStatusEnum = "ACCEPTED"
	OperationStatusInProgress OperationStatusEnum = "IN_PROGRESS"
	OperationStatusWaiting    OperationStatusEnum = "WAITING"
	OperationStatusSucceeded  OperationStatusEnum = "SUCCEEDED"
	OperationStatusCanceling  OperationStatusEnum = "CANCELING"
	OperationStatusCanceled   OperationStatusEnum = "CANCELED"
	OperationStatusFailed     OperationStatusEnum = "FAILED"
)

var mappingOperationStatus = map[string]OperationStatusEnum{
	"ACCEPTED":    OperationStatusAccepted,
	"IN_PROGRESS": OperationStatusInProgress,
	"WAITING":     OperationStatusWaiting,
	"SUCCEEDED":   OperationStatusSucceeded,
	"CANCELING":   OperationStatusCanceling,
	"CANCELED":    OperationStatusCanceled,
	"FAILED":      OperationStatusFailed,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatus {
		values = append(values, v)
	}
	return values
}
