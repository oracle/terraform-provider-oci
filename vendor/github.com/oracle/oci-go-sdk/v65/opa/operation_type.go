// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Process Automation
//
// Process Automation helps you to rapidly design, automate, and manage business processes in the cloud. With the Process Automation design-time (Designer) and the runtime (Workspace) environments, you can easily create, develop, manage, test, and monitor process applications and their components.
//

package opa

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateOpaInstance           OperationTypeEnum = "CREATE_OPA_INSTANCE"
	OperationTypeUpdateOpaInstance           OperationTypeEnum = "UPDATE_OPA_INSTANCE"
	OperationTypeDeleteOpaInstance           OperationTypeEnum = "DELETE_OPA_INSTANCE"
	OperationTypeMoveOpaInstance             OperationTypeEnum = "MOVE_OPA_INSTANCE"
	OperationTypeCreateOpaInstanceAttachment OperationTypeEnum = "CREATE_OPA_INSTANCE_ATTACHMENT"
	OperationTypeDeleteOpaInstanceAttachment OperationTypeEnum = "DELETE_OPA_INSTANCE_ATTACHMENT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_OPA_INSTANCE":            OperationTypeCreateOpaInstance,
	"UPDATE_OPA_INSTANCE":            OperationTypeUpdateOpaInstance,
	"DELETE_OPA_INSTANCE":            OperationTypeDeleteOpaInstance,
	"MOVE_OPA_INSTANCE":              OperationTypeMoveOpaInstance,
	"CREATE_OPA_INSTANCE_ATTACHMENT": OperationTypeCreateOpaInstanceAttachment,
	"DELETE_OPA_INSTANCE_ATTACHMENT": OperationTypeDeleteOpaInstanceAttachment,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_opa_instance":            OperationTypeCreateOpaInstance,
	"update_opa_instance":            OperationTypeUpdateOpaInstance,
	"delete_opa_instance":            OperationTypeDeleteOpaInstance,
	"move_opa_instance":              OperationTypeMoveOpaInstance,
	"create_opa_instance_attachment": OperationTypeCreateOpaInstanceAttachment,
	"delete_opa_instance_attachment": OperationTypeDeleteOpaInstanceAttachment,
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
		"CREATE_OPA_INSTANCE",
		"UPDATE_OPA_INSTANCE",
		"DELETE_OPA_INSTANCE",
		"MOVE_OPA_INSTANCE",
		"CREATE_OPA_INSTANCE_ATTACHMENT",
		"DELETE_OPA_INSTANCE_ATTACHMENT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
