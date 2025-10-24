// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateStack       OperationTypeEnum = "CREATE_STACK"
	OperationTypeUpdateStack       OperationTypeEnum = "UPDATE_STACK"
	OperationTypeDeleteStack       OperationTypeEnum = "DELETE_STACK"
	OperationTypeMoveStack         OperationTypeEnum = "MOVE_STACK"
	OperationTypeDeployArtifacts   OperationTypeEnum = "DEPLOY_ARTIFACTS"
	OperationTypeAddServiceInStack OperationTypeEnum = "ADD_SERVICE_IN_STACK"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_STACK":         OperationTypeCreateStack,
	"UPDATE_STACK":         OperationTypeUpdateStack,
	"DELETE_STACK":         OperationTypeDeleteStack,
	"MOVE_STACK":           OperationTypeMoveStack,
	"DEPLOY_ARTIFACTS":     OperationTypeDeployArtifacts,
	"ADD_SERVICE_IN_STACK": OperationTypeAddServiceInStack,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_stack":         OperationTypeCreateStack,
	"update_stack":         OperationTypeUpdateStack,
	"delete_stack":         OperationTypeDeleteStack,
	"move_stack":           OperationTypeMoveStack,
	"deploy_artifacts":     OperationTypeDeployArtifacts,
	"add_service_in_stack": OperationTypeAddServiceInStack,
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
		"CREATE_STACK",
		"UPDATE_STACK",
		"DELETE_STACK",
		"MOVE_STACK",
		"DEPLOY_ARTIFACTS",
		"ADD_SERVICE_IN_STACK",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
