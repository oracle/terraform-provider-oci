// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted  OperationStatusEnum = "ACCEPTED"
	OperationStatusFailed    OperationStatusEnum = "FAILED"
	OperationStatusSucceeded OperationStatusEnum = "SUCCEEDED"
)

var mappingOperationStatus = map[string]OperationStatusEnum{
	"ACCEPTED":  OperationStatusAccepted,
	"FAILED":    OperationStatusFailed,
	"SUCCEEDED": OperationStatusSucceeded,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatus {
		values = append(values, v)
	}
	return values
}
