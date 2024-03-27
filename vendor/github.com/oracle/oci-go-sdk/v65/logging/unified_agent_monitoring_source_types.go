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

// UnifiedAgentMonitoringSourceTypesEnum Enum with underlying type: string
type UnifiedAgentMonitoringSourceTypesEnum string

// Set of constants representing the allowable values for UnifiedAgentMonitoringSourceTypesEnum
const (
	UnifiedAgentMonitoringSourceTypesKubernetes UnifiedAgentMonitoringSourceTypesEnum = "KUBERNETES"
	UnifiedAgentMonitoringSourceTypesTail       UnifiedAgentMonitoringSourceTypesEnum = "TAIL"
	UnifiedAgentMonitoringSourceTypesUrl        UnifiedAgentMonitoringSourceTypesEnum = "URL"
)

var mappingUnifiedAgentMonitoringSourceTypesEnum = map[string]UnifiedAgentMonitoringSourceTypesEnum{
	"KUBERNETES": UnifiedAgentMonitoringSourceTypesKubernetes,
	"TAIL":       UnifiedAgentMonitoringSourceTypesTail,
	"URL":        UnifiedAgentMonitoringSourceTypesUrl,
}

var mappingUnifiedAgentMonitoringSourceTypesEnumLowerCase = map[string]UnifiedAgentMonitoringSourceTypesEnum{
	"kubernetes": UnifiedAgentMonitoringSourceTypesKubernetes,
	"tail":       UnifiedAgentMonitoringSourceTypesTail,
	"url":        UnifiedAgentMonitoringSourceTypesUrl,
}

// GetUnifiedAgentMonitoringSourceTypesEnumValues Enumerates the set of values for UnifiedAgentMonitoringSourceTypesEnum
func GetUnifiedAgentMonitoringSourceTypesEnumValues() []UnifiedAgentMonitoringSourceTypesEnum {
	values := make([]UnifiedAgentMonitoringSourceTypesEnum, 0)
	for _, v := range mappingUnifiedAgentMonitoringSourceTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetUnifiedAgentMonitoringSourceTypesEnumStringValues Enumerates the set of values in String for UnifiedAgentMonitoringSourceTypesEnum
func GetUnifiedAgentMonitoringSourceTypesEnumStringValues() []string {
	return []string{
		"KUBERNETES",
		"TAIL",
		"URL",
	}
}

// GetMappingUnifiedAgentMonitoringSourceTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnifiedAgentMonitoringSourceTypesEnum(val string) (UnifiedAgentMonitoringSourceTypesEnum, bool) {
	enum, ok := mappingUnifiedAgentMonitoringSourceTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
