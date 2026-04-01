// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// EnvironmentTypeEnum Enum with underlying type: string
type EnvironmentTypeEnum string

// Set of constants representing the allowable values for EnvironmentTypeEnum
const (
	EnvironmentTypeSource      EnvironmentTypeEnum = "SOURCE"
	EnvironmentTypeDestination EnvironmentTypeEnum = "DESTINATION"
)

var mappingEnvironmentTypeEnum = map[string]EnvironmentTypeEnum{
	"SOURCE":      EnvironmentTypeSource,
	"DESTINATION": EnvironmentTypeDestination,
}

var mappingEnvironmentTypeEnumLowerCase = map[string]EnvironmentTypeEnum{
	"source":      EnvironmentTypeSource,
	"destination": EnvironmentTypeDestination,
}

// GetEnvironmentTypeEnumValues Enumerates the set of values for EnvironmentTypeEnum
func GetEnvironmentTypeEnumValues() []EnvironmentTypeEnum {
	values := make([]EnvironmentTypeEnum, 0)
	for _, v := range mappingEnvironmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEnvironmentTypeEnumStringValues Enumerates the set of values in String for EnvironmentTypeEnum
func GetEnvironmentTypeEnumStringValues() []string {
	return []string{
		"SOURCE",
		"DESTINATION",
	}
}

// GetMappingEnvironmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnvironmentTypeEnum(val string) (EnvironmentTypeEnum, bool) {
	enum, ok := mappingEnvironmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
