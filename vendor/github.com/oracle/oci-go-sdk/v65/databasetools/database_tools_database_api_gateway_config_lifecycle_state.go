// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum Enum with underlying type: string
type DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
const (
	DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = "ACTIVE"
	DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = "DELETED"
)

var mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum = map[string]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum{
	"ACTIVE":  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive,
	"DELETED": DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted,
}

var mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumLowerCase = map[string]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum{
	"active":  DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive,
	"deleted": DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted,
}

// GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumValues Enumerates the set of values for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
func GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumValues() []DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum {
	values := make([]DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum
func GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum(val string) (DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum, bool) {
	enum, ok := mappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
