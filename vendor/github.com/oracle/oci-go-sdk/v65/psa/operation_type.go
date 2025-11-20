// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreatePrivateServiceAccess OperationTypeEnum = "CREATE_PRIVATE_SERVICE_ACCESS"
	OperationTypeUpdatePrivateServiceAccess OperationTypeEnum = "UPDATE_PRIVATE_SERVICE_ACCESS"
	OperationTypeDeletePrivateServiceAccess OperationTypeEnum = "DELETE_PRIVATE_SERVICE_ACCESS"
	OperationTypeMovePrivateServiceAccess   OperationTypeEnum = "MOVE_PRIVATE_SERVICE_ACCESS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PRIVATE_SERVICE_ACCESS": OperationTypeCreatePrivateServiceAccess,
	"UPDATE_PRIVATE_SERVICE_ACCESS": OperationTypeUpdatePrivateServiceAccess,
	"DELETE_PRIVATE_SERVICE_ACCESS": OperationTypeDeletePrivateServiceAccess,
	"MOVE_PRIVATE_SERVICE_ACCESS":   OperationTypeMovePrivateServiceAccess,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_private_service_access": OperationTypeCreatePrivateServiceAccess,
	"update_private_service_access": OperationTypeUpdatePrivateServiceAccess,
	"delete_private_service_access": OperationTypeDeletePrivateServiceAccess,
	"move_private_service_access":   OperationTypeMovePrivateServiceAccess,
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
		"CREATE_PRIVATE_SERVICE_ACCESS",
		"UPDATE_PRIVATE_SERVICE_ACCESS",
		"DELETE_PRIVATE_SERVICE_ACCESS",
		"MOVE_PRIVATE_SERVICE_ACCESS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
