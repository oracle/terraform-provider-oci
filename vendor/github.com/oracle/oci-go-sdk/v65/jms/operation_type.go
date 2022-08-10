// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet                   OperationTypeEnum = "CREATE_FLEET"
	OperationTypeDeleteFleet                   OperationTypeEnum = "DELETE_FLEET"
	OperationTypeMoveFleet                     OperationTypeEnum = "MOVE_FLEET"
	OperationTypeUpdateFleet                   OperationTypeEnum = "UPDATE_FLEET"
	OperationTypeUpdateFleetAgentConfiguration OperationTypeEnum = "UPDATE_FLEET_AGENT_CONFIGURATION"
	OperationTypeDeleteJavaInstallation        OperationTypeEnum = "DELETE_JAVA_INSTALLATION"
	OperationTypeCreateJavaInstallation        OperationTypeEnum = "CREATE_JAVA_INSTALLATION"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_FLEET":                     OperationTypeCreateFleet,
	"DELETE_FLEET":                     OperationTypeDeleteFleet,
	"MOVE_FLEET":                       OperationTypeMoveFleet,
	"UPDATE_FLEET":                     OperationTypeUpdateFleet,
	"UPDATE_FLEET_AGENT_CONFIGURATION": OperationTypeUpdateFleetAgentConfiguration,
	"DELETE_JAVA_INSTALLATION":         OperationTypeDeleteJavaInstallation,
	"CREATE_JAVA_INSTALLATION":         OperationTypeCreateJavaInstallation,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_fleet":                     OperationTypeCreateFleet,
	"delete_fleet":                     OperationTypeDeleteFleet,
	"move_fleet":                       OperationTypeMoveFleet,
	"update_fleet":                     OperationTypeUpdateFleet,
	"update_fleet_agent_configuration": OperationTypeUpdateFleetAgentConfiguration,
	"delete_java_installation":         OperationTypeDeleteJavaInstallation,
	"create_java_installation":         OperationTypeCreateJavaInstallation,
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
		"CREATE_FLEET",
		"DELETE_FLEET",
		"MOVE_FLEET",
		"UPDATE_FLEET",
		"UPDATE_FLEET_AGENT_CONFIGURATION",
		"DELETE_JAVA_INSTALLATION",
		"CREATE_JAVA_INSTALLATION",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
