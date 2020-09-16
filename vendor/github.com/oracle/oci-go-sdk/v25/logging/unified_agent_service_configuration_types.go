// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

// UnifiedAgentServiceConfigurationTypesEnum Enum with underlying type: string
type UnifiedAgentServiceConfigurationTypesEnum string

// Set of constants representing the allowable values for UnifiedAgentServiceConfigurationTypesEnum
const (
	UnifiedAgentServiceConfigurationTypesLogging UnifiedAgentServiceConfigurationTypesEnum = "LOGGING"
)

var mappingUnifiedAgentServiceConfigurationTypes = map[string]UnifiedAgentServiceConfigurationTypesEnum{
	"LOGGING": UnifiedAgentServiceConfigurationTypesLogging,
}

// GetUnifiedAgentServiceConfigurationTypesEnumValues Enumerates the set of values for UnifiedAgentServiceConfigurationTypesEnum
func GetUnifiedAgentServiceConfigurationTypesEnumValues() []UnifiedAgentServiceConfigurationTypesEnum {
	values := make([]UnifiedAgentServiceConfigurationTypesEnum, 0)
	for _, v := range mappingUnifiedAgentServiceConfigurationTypes {
		values = append(values, v)
	}
	return values
}
