// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProject OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeMoveProject   OperationTypeEnum = "MOVE_PROJECT"
	OperationTypeCreateModel   OperationTypeEnum = "CREATE_MODEL"
	OperationTypeUpdateModel   OperationTypeEnum = "UPDATE_MODEL"
	OperationTypeDeleteModel   OperationTypeEnum = "DELETE_MODEL"
	OperationTypeMoveModel     OperationTypeEnum = "MOVE_MODEL"
	OperationTypeComposeModel  OperationTypeEnum = "COMPOSE_MODEL"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PROJECT": OperationTypeCreateProject,
	"UPDATE_PROJECT": OperationTypeUpdateProject,
	"DELETE_PROJECT": OperationTypeDeleteProject,
	"MOVE_PROJECT":   OperationTypeMoveProject,
	"CREATE_MODEL":   OperationTypeCreateModel,
	"UPDATE_MODEL":   OperationTypeUpdateModel,
	"DELETE_MODEL":   OperationTypeDeleteModel,
	"MOVE_MODEL":     OperationTypeMoveModel,
	"COMPOSE_MODEL":  OperationTypeComposeModel,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_project": OperationTypeCreateProject,
	"update_project": OperationTypeUpdateProject,
	"delete_project": OperationTypeDeleteProject,
	"move_project":   OperationTypeMoveProject,
	"create_model":   OperationTypeCreateModel,
	"update_model":   OperationTypeUpdateModel,
	"delete_model":   OperationTypeDeleteModel,
	"move_model":     OperationTypeMoveModel,
	"compose_model":  OperationTypeComposeModel,
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
		"CREATE_PROJECT",
		"UPDATE_PROJECT",
		"DELETE_PROJECT",
		"MOVE_PROJECT",
		"CREATE_MODEL",
		"UPDATE_MODEL",
		"DELETE_MODEL",
		"MOVE_MODEL",
		"COMPOSE_MODEL",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
