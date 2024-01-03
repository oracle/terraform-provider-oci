// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateModel    OperationTypeEnum = "CREATE_MODEL"
	OperationTypeUpdateModel    OperationTypeEnum = "UPDATE_MODEL"
	OperationTypeDeleteModel    OperationTypeEnum = "DELETE_MODEL"
	OperationTypeCreateProject  OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject  OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject  OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeCreateEndpoint OperationTypeEnum = "CREATE_ENDPOINT"
	OperationTypeUpdateEndpoint OperationTypeEnum = "UPDATE_ENDPOINT"
	OperationTypeDeleteEndpoint OperationTypeEnum = "DELETE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MODEL":    OperationTypeCreateModel,
	"UPDATE_MODEL":    OperationTypeUpdateModel,
	"DELETE_MODEL":    OperationTypeDeleteModel,
	"CREATE_PROJECT":  OperationTypeCreateProject,
	"UPDATE_PROJECT":  OperationTypeUpdateProject,
	"DELETE_PROJECT":  OperationTypeDeleteProject,
	"CREATE_ENDPOINT": OperationTypeCreateEndpoint,
	"UPDATE_ENDPOINT": OperationTypeUpdateEndpoint,
	"DELETE_ENDPOINT": OperationTypeDeleteEndpoint,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_model":    OperationTypeCreateModel,
	"update_model":    OperationTypeUpdateModel,
	"delete_model":    OperationTypeDeleteModel,
	"create_project":  OperationTypeCreateProject,
	"update_project":  OperationTypeUpdateProject,
	"delete_project":  OperationTypeDeleteProject,
	"create_endpoint": OperationTypeCreateEndpoint,
	"update_endpoint": OperationTypeUpdateEndpoint,
	"delete_endpoint": OperationTypeDeleteEndpoint,
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
		"CREATE_MODEL",
		"UPDATE_MODEL",
		"DELETE_MODEL",
		"CREATE_PROJECT",
		"UPDATE_PROJECT",
		"DELETE_PROJECT",
		"CREATE_ENDPOINT",
		"UPDATE_ENDPOINT",
		"DELETE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
