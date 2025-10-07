// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AiDataPlatform Control Plane API
//
// Use the AiDataPlatform Control Plane API to manage Data Lakes.
//

package aidataplatform

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDataLake      OperationTypeEnum = "CREATE_DATA_LAKE"
	OperationTypeUpdateDataLake      OperationTypeEnum = "UPDATE_DATA_LAKE"
	OperationTypeDeleteDataLake      OperationTypeEnum = "DELETE_DATA_LAKE"
	OperationTypeMoveDataLake        OperationTypeEnum = "MOVE_DATA_LAKE"
	OperationTypeCreateWorkspace     OperationTypeEnum = "CREATE_WORKSPACE"
	OperationTypeUpdateWorkspace     OperationTypeEnum = "UPDATE_WORKSPACE"
	OperationTypeDeleteWorkspace     OperationTypeEnum = "DELETE_WORKSPACE"
	OperationTypeEvaluateIamPolicies OperationTypeEnum = "EVALUATE_IAM_POLICIES"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DATA_LAKE":      OperationTypeCreateDataLake,
	"UPDATE_DATA_LAKE":      OperationTypeUpdateDataLake,
	"DELETE_DATA_LAKE":      OperationTypeDeleteDataLake,
	"MOVE_DATA_LAKE":        OperationTypeMoveDataLake,
	"CREATE_WORKSPACE":      OperationTypeCreateWorkspace,
	"UPDATE_WORKSPACE":      OperationTypeUpdateWorkspace,
	"DELETE_WORKSPACE":      OperationTypeDeleteWorkspace,
	"EVALUATE_IAM_POLICIES": OperationTypeEvaluateIamPolicies,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_data_lake":      OperationTypeCreateDataLake,
	"update_data_lake":      OperationTypeUpdateDataLake,
	"delete_data_lake":      OperationTypeDeleteDataLake,
	"move_data_lake":        OperationTypeMoveDataLake,
	"create_workspace":      OperationTypeCreateWorkspace,
	"update_workspace":      OperationTypeUpdateWorkspace,
	"delete_workspace":      OperationTypeDeleteWorkspace,
	"evaluate_iam_policies": OperationTypeEvaluateIamPolicies,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_DATA_LAKE",
		"UPDATE_DATA_LAKE",
		"DELETE_DATA_LAKE",
		"MOVE_DATA_LAKE",
		"CREATE_WORKSPACE",
		"UPDATE_WORKSPACE",
		"DELETE_WORKSPACE",
		"EVALUATE_IAM_POLICIES",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
