// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateModel           OperationTypeEnum = "CREATE_MODEL"
	OperationTypeUpdateModel           OperationTypeEnum = "UPDATE_MODEL"
	OperationTypeDeleteModel           OperationTypeEnum = "DELETE_MODEL"
	OperationTypeCreatePrivateEndpoint OperationTypeEnum = "CREATE_PRIVATE_ENDPOINT"
	OperationTypeDeletePrivateEndpoint OperationTypeEnum = "DELETE_PRIVATE_ENDPOINT"
	OperationTypeUpdatePrivateEndpoint OperationTypeEnum = "UPDATE_PRIVATE_ENDPOINT"
	OperationTypeMovePrivateEndpoint   OperationTypeEnum = "MOVE_PRIVATE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MODEL":            OperationTypeCreateModel,
	"UPDATE_MODEL":            OperationTypeUpdateModel,
	"DELETE_MODEL":            OperationTypeDeleteModel,
	"CREATE_PRIVATE_ENDPOINT": OperationTypeCreatePrivateEndpoint,
	"DELETE_PRIVATE_ENDPOINT": OperationTypeDeletePrivateEndpoint,
	"UPDATE_PRIVATE_ENDPOINT": OperationTypeUpdatePrivateEndpoint,
	"MOVE_PRIVATE_ENDPOINT":   OperationTypeMovePrivateEndpoint,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_model":            OperationTypeCreateModel,
	"update_model":            OperationTypeUpdateModel,
	"delete_model":            OperationTypeDeleteModel,
	"create_private_endpoint": OperationTypeCreatePrivateEndpoint,
	"delete_private_endpoint": OperationTypeDeletePrivateEndpoint,
	"update_private_endpoint": OperationTypeUpdatePrivateEndpoint,
	"move_private_endpoint":   OperationTypeMovePrivateEndpoint,
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
		"CREATE_PRIVATE_ENDPOINT",
		"DELETE_PRIVATE_ENDPOINT",
		"UPDATE_PRIVATE_ENDPOINT",
		"MOVE_PRIVATE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
