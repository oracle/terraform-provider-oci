// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// RestConfigurationTypeEnum Enum with underlying type: string
type RestConfigurationTypeEnum string

// Set of constants representing the allowable values for RestConfigurationTypeEnum
const (
	RestConfigurationTypeDisabled     RestConfigurationTypeEnum = "DISABLED"
	RestConfigurationTypeDbsystemOnly RestConfigurationTypeEnum = "DBSYSTEM_ONLY"
)

var mappingRestConfigurationTypeEnum = map[string]RestConfigurationTypeEnum{
	"DISABLED":      RestConfigurationTypeDisabled,
	"DBSYSTEM_ONLY": RestConfigurationTypeDbsystemOnly,
}

var mappingRestConfigurationTypeEnumLowerCase = map[string]RestConfigurationTypeEnum{
	"disabled":      RestConfigurationTypeDisabled,
	"dbsystem_only": RestConfigurationTypeDbsystemOnly,
}

// GetRestConfigurationTypeEnumValues Enumerates the set of values for RestConfigurationTypeEnum
func GetRestConfigurationTypeEnumValues() []RestConfigurationTypeEnum {
	values := make([]RestConfigurationTypeEnum, 0)
	for _, v := range mappingRestConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRestConfigurationTypeEnumStringValues Enumerates the set of values in String for RestConfigurationTypeEnum
func GetRestConfigurationTypeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"DBSYSTEM_ONLY",
	}
}

// GetMappingRestConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRestConfigurationTypeEnum(val string) (RestConfigurationTypeEnum, bool) {
	enum, ok := mappingRestConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
