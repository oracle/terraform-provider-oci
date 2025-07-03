// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database MultiCloud Data plane Integration
//
// 1. Oracle Azure Connector Resource: This is for installing Azure Arc Server in ExaCS VM Cluster.
//   There are two way to install Azure Arc Server (Azure Identity) in ExaCS VMCluster.
//     a. Using Bearer Access Token or
//     b. By providing Authentication token
// 2. Oracle Azure Blob Container Resource: This is for to capture Azure Container details
//    and same will be used in multiple ExaCS VMCluster to mount the Azure Container.
// 3. Oracle Azure Blob Mount Resource: This is for to mount Azure Container in ExaCS VMCluster
//    using Oracle Azure Connector and Oracle Azure Blob Container Resource.
//

package dbmulticloud

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
