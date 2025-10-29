// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"strings"
)

// AutomaticDrConfigurationMemberTypeEnum Enum with underlying type: string
type AutomaticDrConfigurationMemberTypeEnum string

// Set of constants representing the allowable values for AutomaticDrConfigurationMemberTypeEnum
const (
	AutomaticDrConfigurationMemberTypeDatabase                    AutomaticDrConfigurationMemberTypeEnum = "DATABASE"
	AutomaticDrConfigurationMemberTypeAutonomousDatabase          AutomaticDrConfigurationMemberTypeEnum = "AUTONOMOUS_DATABASE"
	AutomaticDrConfigurationMemberTypeAutonomousContainerDatabase AutomaticDrConfigurationMemberTypeEnum = "AUTONOMOUS_CONTAINER_DATABASE"
)

var mappingAutomaticDrConfigurationMemberTypeEnum = map[string]AutomaticDrConfigurationMemberTypeEnum{
	"DATABASE":                      AutomaticDrConfigurationMemberTypeDatabase,
	"AUTONOMOUS_DATABASE":           AutomaticDrConfigurationMemberTypeAutonomousDatabase,
	"AUTONOMOUS_CONTAINER_DATABASE": AutomaticDrConfigurationMemberTypeAutonomousContainerDatabase,
}

var mappingAutomaticDrConfigurationMemberTypeEnumLowerCase = map[string]AutomaticDrConfigurationMemberTypeEnum{
	"database":                      AutomaticDrConfigurationMemberTypeDatabase,
	"autonomous_database":           AutomaticDrConfigurationMemberTypeAutonomousDatabase,
	"autonomous_container_database": AutomaticDrConfigurationMemberTypeAutonomousContainerDatabase,
}

// GetAutomaticDrConfigurationMemberTypeEnumValues Enumerates the set of values for AutomaticDrConfigurationMemberTypeEnum
func GetAutomaticDrConfigurationMemberTypeEnumValues() []AutomaticDrConfigurationMemberTypeEnum {
	values := make([]AutomaticDrConfigurationMemberTypeEnum, 0)
	for _, v := range mappingAutomaticDrConfigurationMemberTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutomaticDrConfigurationMemberTypeEnumStringValues Enumerates the set of values in String for AutomaticDrConfigurationMemberTypeEnum
func GetAutomaticDrConfigurationMemberTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"AUTONOMOUS_DATABASE",
		"AUTONOMOUS_CONTAINER_DATABASE",
	}
}

// GetMappingAutomaticDrConfigurationMemberTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutomaticDrConfigurationMemberTypeEnum(val string) (AutomaticDrConfigurationMemberTypeEnum, bool) {
	enum, ok := mappingAutomaticDrConfigurationMemberTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
