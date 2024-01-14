// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"strings"
)

// ConfigSeverityEnum Enum with underlying type: string
type ConfigSeverityEnum string

// Set of constants representing the allowable values for ConfigSeverityEnum
const (
	ConfigSeverityUnset    ConfigSeverityEnum = "UNSET"
	ConfigSeverityNone     ConfigSeverityEnum = "NONE"
	ConfigSeverityLow      ConfigSeverityEnum = "LOW"
	ConfigSeverityMedium   ConfigSeverityEnum = "MEDIUM"
	ConfigSeverityHigh     ConfigSeverityEnum = "HIGH"
	ConfigSeverityCritical ConfigSeverityEnum = "CRITICAL"
)

var mappingConfigSeverityEnum = map[string]ConfigSeverityEnum{
	"UNSET":    ConfigSeverityUnset,
	"NONE":     ConfigSeverityNone,
	"LOW":      ConfigSeverityLow,
	"MEDIUM":   ConfigSeverityMedium,
	"HIGH":     ConfigSeverityHigh,
	"CRITICAL": ConfigSeverityCritical,
}

var mappingConfigSeverityEnumLowerCase = map[string]ConfigSeverityEnum{
	"unset":    ConfigSeverityUnset,
	"none":     ConfigSeverityNone,
	"low":      ConfigSeverityLow,
	"medium":   ConfigSeverityMedium,
	"high":     ConfigSeverityHigh,
	"critical": ConfigSeverityCritical,
}

// GetConfigSeverityEnumValues Enumerates the set of values for ConfigSeverityEnum
func GetConfigSeverityEnumValues() []ConfigSeverityEnum {
	values := make([]ConfigSeverityEnum, 0)
	for _, v := range mappingConfigSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigSeverityEnumStringValues Enumerates the set of values in String for ConfigSeverityEnum
func GetConfigSeverityEnumStringValues() []string {
	return []string{
		"UNSET",
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

// GetMappingConfigSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigSeverityEnum(val string) (ConfigSeverityEnum, bool) {
	enum, ok := mappingConfigSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
