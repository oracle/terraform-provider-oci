// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProtectedDatabase     OperationTypeEnum = "CREATE_PROTECTED_DATABASE"
	OperationTypeUpdateProtectedDatabase     OperationTypeEnum = "UPDATE_PROTECTED_DATABASE"
	OperationTypeDeleteProtectedDatabase     OperationTypeEnum = "DELETE_PROTECTED_DATABASE"
	OperationTypeMoveProtectedDatabase       OperationTypeEnum = "MOVE_PROTECTED_DATABASE"
	OperationTypeCreateProtectionPolicy      OperationTypeEnum = "CREATE_PROTECTION_POLICY"
	OperationTypeUpdateProtectionPolicy      OperationTypeEnum = "UPDATE_PROTECTION_POLICY"
	OperationTypeDeleteProtectionPolicy      OperationTypeEnum = "DELETE_PROTECTION_POLICY"
	OperationTypeMoveProtectionPolicy        OperationTypeEnum = "MOVE_PROTECTION_POLICY"
	OperationTypeCreateRecoveryServiceSubnet OperationTypeEnum = "CREATE_RECOVERY_SERVICE_SUBNET"
	OperationTypeUpdateRecoveryServiceSubnet OperationTypeEnum = "UPDATE_RECOVERY_SERVICE_SUBNET"
	OperationTypeDeleteRecoveryServiceSubnet OperationTypeEnum = "DELETE_RECOVERY_SERVICE_SUBNET"
	OperationTypeMoveRecoveryServiceSubnet   OperationTypeEnum = "MOVE_RECOVERY_SERVICE_SUBNET"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_PROTECTED_DATABASE":      OperationTypeCreateProtectedDatabase,
	"UPDATE_PROTECTED_DATABASE":      OperationTypeUpdateProtectedDatabase,
	"DELETE_PROTECTED_DATABASE":      OperationTypeDeleteProtectedDatabase,
	"MOVE_PROTECTED_DATABASE":        OperationTypeMoveProtectedDatabase,
	"CREATE_PROTECTION_POLICY":       OperationTypeCreateProtectionPolicy,
	"UPDATE_PROTECTION_POLICY":       OperationTypeUpdateProtectionPolicy,
	"DELETE_PROTECTION_POLICY":       OperationTypeDeleteProtectionPolicy,
	"MOVE_PROTECTION_POLICY":         OperationTypeMoveProtectionPolicy,
	"CREATE_RECOVERY_SERVICE_SUBNET": OperationTypeCreateRecoveryServiceSubnet,
	"UPDATE_RECOVERY_SERVICE_SUBNET": OperationTypeUpdateRecoveryServiceSubnet,
	"DELETE_RECOVERY_SERVICE_SUBNET": OperationTypeDeleteRecoveryServiceSubnet,
	"MOVE_RECOVERY_SERVICE_SUBNET":   OperationTypeMoveRecoveryServiceSubnet,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_protected_database":      OperationTypeCreateProtectedDatabase,
	"update_protected_database":      OperationTypeUpdateProtectedDatabase,
	"delete_protected_database":      OperationTypeDeleteProtectedDatabase,
	"move_protected_database":        OperationTypeMoveProtectedDatabase,
	"create_protection_policy":       OperationTypeCreateProtectionPolicy,
	"update_protection_policy":       OperationTypeUpdateProtectionPolicy,
	"delete_protection_policy":       OperationTypeDeleteProtectionPolicy,
	"move_protection_policy":         OperationTypeMoveProtectionPolicy,
	"create_recovery_service_subnet": OperationTypeCreateRecoveryServiceSubnet,
	"update_recovery_service_subnet": OperationTypeUpdateRecoveryServiceSubnet,
	"delete_recovery_service_subnet": OperationTypeDeleteRecoveryServiceSubnet,
	"move_recovery_service_subnet":   OperationTypeMoveRecoveryServiceSubnet,
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
		"CREATE_PROTECTED_DATABASE",
		"UPDATE_PROTECTED_DATABASE",
		"DELETE_PROTECTED_DATABASE",
		"MOVE_PROTECTED_DATABASE",
		"CREATE_PROTECTION_POLICY",
		"UPDATE_PROTECTION_POLICY",
		"DELETE_PROTECTION_POLICY",
		"MOVE_PROTECTION_POLICY",
		"CREATE_RECOVERY_SERVICE_SUBNET",
		"UPDATE_RECOVERY_SERVICE_SUBNET",
		"DELETE_RECOVERY_SERVICE_SUBNET",
		"MOVE_RECOVERY_SERVICE_SUBNET",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
