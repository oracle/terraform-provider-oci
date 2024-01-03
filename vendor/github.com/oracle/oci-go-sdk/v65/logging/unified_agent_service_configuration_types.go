// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"strings"
)

// UnifiedAgentServiceConfigurationTypesEnum Enum with underlying type: string
type UnifiedAgentServiceConfigurationTypesEnum string

// Set of constants representing the allowable values for UnifiedAgentServiceConfigurationTypesEnum
const (
	UnifiedAgentServiceConfigurationTypesLogging    UnifiedAgentServiceConfigurationTypesEnum = "LOGGING"
	UnifiedAgentServiceConfigurationTypesMonitoring UnifiedAgentServiceConfigurationTypesEnum = "MONITORING"
)

var mappingUnifiedAgentServiceConfigurationTypesEnum = map[string]UnifiedAgentServiceConfigurationTypesEnum{
	"LOGGING":    UnifiedAgentServiceConfigurationTypesLogging,
	"MONITORING": UnifiedAgentServiceConfigurationTypesMonitoring,
}

var mappingUnifiedAgentServiceConfigurationTypesEnumLowerCase = map[string]UnifiedAgentServiceConfigurationTypesEnum{
	"logging":    UnifiedAgentServiceConfigurationTypesLogging,
	"monitoring": UnifiedAgentServiceConfigurationTypesMonitoring,
}

// GetUnifiedAgentServiceConfigurationTypesEnumValues Enumerates the set of values for UnifiedAgentServiceConfigurationTypesEnum
func GetUnifiedAgentServiceConfigurationTypesEnumValues() []UnifiedAgentServiceConfigurationTypesEnum {
	values := make([]UnifiedAgentServiceConfigurationTypesEnum, 0)
	for _, v := range mappingUnifiedAgentServiceConfigurationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentServiceConfigurationTypesEnumStringValues Enumerates the set of values in String for UnifiedAgentServiceConfigurationTypesEnum
func GetUnifiedAgentServiceConfigurationTypesEnumStringValues() []string {
	return []string{
		"LOGGING",
		"MONITORING",
	}
}

// GetMappingUnifiedAgentServiceConfigurationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentServiceConfigurationTypesEnum(val string) (UnifiedAgentServiceConfigurationTypesEnum, bool) {
	enum, ok := mappingUnifiedAgentServiceConfigurationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
