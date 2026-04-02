// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFunction                  OperationTypeEnum = "CREATE_FUNCTION"
	OperationTypeUpdateFunction                  OperationTypeEnum = "UPDATE_FUNCTION"
	OperationTypeDeleteFunction                  OperationTypeEnum = "DELETE_FUNCTION"
	OperationTypeCreateApplication               OperationTypeEnum = "CREATE_APPLICATION"
	OperationTypeUpdateApplication               OperationTypeEnum = "UPDATE_APPLICATION"
	OperationTypeDeleteApplication               OperationTypeEnum = "DELETE_APPLICATION"
	OperationTypeCreateFunctionsRuntime          OperationTypeEnum = "CREATE_FUNCTIONS_RUNTIME"
	OperationTypeUpdateFunctionsRuntime          OperationTypeEnum = "UPDATE_FUNCTIONS_RUNTIME"
	OperationTypeDeleteFunctionsRuntime          OperationTypeEnum = "DELETE_FUNCTIONS_RUNTIME"
	OperationTypeRollbackFunctionsRuntimeVersion OperationTypeEnum = "ROLLBACK_FUNCTIONS_RUNTIME_VERSION"
	OperationTypeCreateFunctionsRuntimeVersion   OperationTypeEnum = "CREATE_FUNCTIONS_RUNTIME_VERSION"
	OperationTypeUpdateFunctionsRuntimeVersion   OperationTypeEnum = "UPDATE_FUNCTIONS_RUNTIME_VERSION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_FUNCTION":                    OperationTypeCreateFunction,
	"UPDATE_FUNCTION":                    OperationTypeUpdateFunction,
	"DELETE_FUNCTION":                    OperationTypeDeleteFunction,
	"CREATE_APPLICATION":                 OperationTypeCreateApplication,
	"UPDATE_APPLICATION":                 OperationTypeUpdateApplication,
	"DELETE_APPLICATION":                 OperationTypeDeleteApplication,
	"CREATE_FUNCTIONS_RUNTIME":           OperationTypeCreateFunctionsRuntime,
	"UPDATE_FUNCTIONS_RUNTIME":           OperationTypeUpdateFunctionsRuntime,
	"DELETE_FUNCTIONS_RUNTIME":           OperationTypeDeleteFunctionsRuntime,
	"ROLLBACK_FUNCTIONS_RUNTIME_VERSION": OperationTypeRollbackFunctionsRuntimeVersion,
	"CREATE_FUNCTIONS_RUNTIME_VERSION":   OperationTypeCreateFunctionsRuntimeVersion,
	"UPDATE_FUNCTIONS_RUNTIME_VERSION":   OperationTypeUpdateFunctionsRuntimeVersion,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_function":                    OperationTypeCreateFunction,
	"update_function":                    OperationTypeUpdateFunction,
	"delete_function":                    OperationTypeDeleteFunction,
	"create_application":                 OperationTypeCreateApplication,
	"update_application":                 OperationTypeUpdateApplication,
	"delete_application":                 OperationTypeDeleteApplication,
	"create_functions_runtime":           OperationTypeCreateFunctionsRuntime,
	"update_functions_runtime":           OperationTypeUpdateFunctionsRuntime,
	"delete_functions_runtime":           OperationTypeDeleteFunctionsRuntime,
	"rollback_functions_runtime_version": OperationTypeRollbackFunctionsRuntimeVersion,
	"create_functions_runtime_version":   OperationTypeCreateFunctionsRuntimeVersion,
	"update_functions_runtime_version":   OperationTypeUpdateFunctionsRuntimeVersion,
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
		"CREATE_FUNCTION",
		"UPDATE_FUNCTION",
		"DELETE_FUNCTION",
		"CREATE_APPLICATION",
		"UPDATE_APPLICATION",
		"DELETE_APPLICATION",
		"CREATE_FUNCTIONS_RUNTIME",
		"UPDATE_FUNCTIONS_RUNTIME",
		"DELETE_FUNCTIONS_RUNTIME",
		"ROLLBACK_FUNCTIONS_RUNTIME_VERSION",
		"CREATE_FUNCTIONS_RUNTIME_VERSION",
		"UPDATE_FUNCTIONS_RUNTIME_VERSION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
