// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateBastion OperationTypeEnum = "CREATE_BASTION"
	OperationTypeUpdateBastion OperationTypeEnum = "UPDATE_BASTION"
	OperationTypeDeleteBastion OperationTypeEnum = "DELETE_BASTION"
	OperationTypeCreateSession OperationTypeEnum = "CREATE_SESSION"
	OperationTypeDeleteSession OperationTypeEnum = "DELETE_SESSION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_BASTION": OperationTypeCreateBastion,
	"UPDATE_BASTION": OperationTypeUpdateBastion,
	"DELETE_BASTION": OperationTypeDeleteBastion,
	"CREATE_SESSION": OperationTypeCreateSession,
	"DELETE_SESSION": OperationTypeDeleteSession,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_bastion": OperationTypeCreateBastion,
	"update_bastion": OperationTypeUpdateBastion,
	"delete_bastion": OperationTypeDeleteBastion,
	"create_session": OperationTypeCreateSession,
	"delete_session": OperationTypeDeleteSession,
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
		"CREATE_BASTION",
		"UPDATE_BASTION",
		"DELETE_BASTION",
		"CREATE_SESSION",
		"DELETE_SESSION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
