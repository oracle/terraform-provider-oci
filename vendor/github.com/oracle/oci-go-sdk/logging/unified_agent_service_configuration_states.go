// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

// UnifiedAgentServiceConfigurationStatesEnum Enum with underlying type: string
type UnifiedAgentServiceConfigurationStatesEnum string

// Set of constants representing the allowable values for UnifiedAgentServiceConfigurationStatesEnum
const (
	UnifiedAgentServiceConfigurationStatesValid   UnifiedAgentServiceConfigurationStatesEnum = "VALID"
	UnifiedAgentServiceConfigurationStatesInvalid UnifiedAgentServiceConfigurationStatesEnum = "INVALID"
)

var mappingUnifiedAgentServiceConfigurationStates = map[string]UnifiedAgentServiceConfigurationStatesEnum{
	"VALID":   UnifiedAgentServiceConfigurationStatesValid,
	"INVALID": UnifiedAgentServiceConfigurationStatesInvalid,
}

// GetUnifiedAgentServiceConfigurationStatesEnumValues Enumerates the set of values for UnifiedAgentServiceConfigurationStatesEnum
func GetUnifiedAgentServiceConfigurationStatesEnumValues() []UnifiedAgentServiceConfigurationStatesEnum {
	values := make([]UnifiedAgentServiceConfigurationStatesEnum, 0)
	for _, v := range mappingUnifiedAgentServiceConfigurationStates {
		values = append(values, v)
	}
	return values
}
