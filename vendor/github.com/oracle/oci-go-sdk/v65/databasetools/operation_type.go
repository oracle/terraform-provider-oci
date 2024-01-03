// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateDatabaseToolsConnection      OperationTypeEnum = "CREATE_DATABASE_TOOLS_CONNECTION"
	OperationTypeUpdateDatabaseToolsConnection      OperationTypeEnum = "UPDATE_DATABASE_TOOLS_CONNECTION"
	OperationTypeDeleteDatabaseToolsConnection      OperationTypeEnum = "DELETE_DATABASE_TOOLS_CONNECTION"
	OperationTypeCreateDatabaseToolsServiceInstance OperationTypeEnum = "CREATE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeUpdateDatabaseToolsServiceInstance OperationTypeEnum = "UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeDeleteDatabaseToolsServiceInstance OperationTypeEnum = "DELETE_DATABASE_TOOLS_SERVICE_INSTANCE"
	OperationTypeCreateDatabaseToolsPrivateEndpoint OperationTypeEnum = "CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
	OperationTypeUpdateDatabaseToolsPrivateEndpoint OperationTypeEnum = "UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
	OperationTypeDeleteDatabaseToolsPrivateEndpoint OperationTypeEnum = "DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_DATABASE_TOOLS_CONNECTION":       OperationTypeCreateDatabaseToolsConnection,
	"UPDATE_DATABASE_TOOLS_CONNECTION":       OperationTypeUpdateDatabaseToolsConnection,
	"DELETE_DATABASE_TOOLS_CONNECTION":       OperationTypeDeleteDatabaseToolsConnection,
	"CREATE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeCreateDatabaseToolsServiceInstance,
	"UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeUpdateDatabaseToolsServiceInstance,
	"DELETE_DATABASE_TOOLS_SERVICE_INSTANCE": OperationTypeDeleteDatabaseToolsServiceInstance,
	"CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeCreateDatabaseToolsPrivateEndpoint,
	"UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeUpdateDatabaseToolsPrivateEndpoint,
	"DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT": OperationTypeDeleteDatabaseToolsPrivateEndpoint,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_database_tools_connection":       OperationTypeCreateDatabaseToolsConnection,
	"update_database_tools_connection":       OperationTypeUpdateDatabaseToolsConnection,
	"delete_database_tools_connection":       OperationTypeDeleteDatabaseToolsConnection,
	"create_database_tools_service_instance": OperationTypeCreateDatabaseToolsServiceInstance,
	"update_database_tools_service_instance": OperationTypeUpdateDatabaseToolsServiceInstance,
	"delete_database_tools_service_instance": OperationTypeDeleteDatabaseToolsServiceInstance,
	"create_database_tools_private_endpoint": OperationTypeCreateDatabaseToolsPrivateEndpoint,
	"update_database_tools_private_endpoint": OperationTypeUpdateDatabaseToolsPrivateEndpoint,
	"delete_database_tools_private_endpoint": OperationTypeDeleteDatabaseToolsPrivateEndpoint,
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
		"CREATE_DATABASE_TOOLS_CONNECTION",
		"UPDATE_DATABASE_TOOLS_CONNECTION",
		"DELETE_DATABASE_TOOLS_CONNECTION",
		"CREATE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"UPDATE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"DELETE_DATABASE_TOOLS_SERVICE_INSTANCE",
		"CREATE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
		"UPDATE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
		"DELETE_DATABASE_TOOLS_PRIVATE_ENDPOINT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
