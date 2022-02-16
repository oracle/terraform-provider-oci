// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"strings"
)

// UnifiedAgentServiceConfigurationStatesEnum Enum with underlying type: string
type UnifiedAgentServiceConfigurationStatesEnum string

// Set of constants representing the allowable values for UnifiedAgentServiceConfigurationStatesEnum
const (
	UnifiedAgentServiceConfigurationStatesValid   UnifiedAgentServiceConfigurationStatesEnum = "VALID"
	UnifiedAgentServiceConfigurationStatesInvalid UnifiedAgentServiceConfigurationStatesEnum = "INVALID"
)

var mappingUnifiedAgentServiceConfigurationStatesEnum = map[string]UnifiedAgentServiceConfigurationStatesEnum{
	"VALID":   UnifiedAgentServiceConfigurationStatesValid,
	"INVALID": UnifiedAgentServiceConfigurationStatesInvalid,
}

// GetUnifiedAgentServiceConfigurationStatesEnumValues Enumerates the set of values for UnifiedAgentServiceConfigurationStatesEnum
func GetUnifiedAgentServiceConfigurationStatesEnumValues() []UnifiedAgentServiceConfigurationStatesEnum {
	values := make([]UnifiedAgentServiceConfigurationStatesEnum, 0)
	for _, v := range mappingUnifiedAgentServiceConfigurationStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentServiceConfigurationStatesEnumStringValues Enumerates the set of values in String for UnifiedAgentServiceConfigurationStatesEnum
func GetUnifiedAgentServiceConfigurationStatesEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
	}
}

// GetMappingUnifiedAgentServiceConfigurationStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentServiceConfigurationStatesEnum(val string) (UnifiedAgentServiceConfigurationStatesEnum, bool) {
	mappingUnifiedAgentServiceConfigurationStatesEnumIgnoreCase := make(map[string]UnifiedAgentServiceConfigurationStatesEnum)
	for k, v := range mappingUnifiedAgentServiceConfigurationStatesEnum {
		mappingUnifiedAgentServiceConfigurationStatesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUnifiedAgentServiceConfigurationStatesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
