// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// EntityTypeEnum Enum with underlying type: string
type EntityTypeEnum string

// Set of constants representing the allowable values for EntityTypeEnum
const (
	EntityTypeComposite  EntityTypeEnum = "COMPOSITE"
	EntityTypeEnumValues EntityTypeEnum = "ENUM_VALUES"
)

var mappingEntityTypeEnum = map[string]EntityTypeEnum{
	"COMPOSITE":   EntityTypeComposite,
	"ENUM_VALUES": EntityTypeEnumValues,
}

var mappingEntityTypeEnumLowerCase = map[string]EntityTypeEnum{
	"composite":   EntityTypeComposite,
	"enum_values": EntityTypeEnumValues,
}

// GetEntityTypeEnumValues Enumerates the set of values for EntityTypeEnum
func GetEntityTypeEnumValues() []EntityTypeEnum {
	values := make([]EntityTypeEnum, 0)
	for _, v := range mappingEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityTypeEnumStringValues Enumerates the set of values in String for EntityTypeEnum
func GetEntityTypeEnumStringValues() []string {
	return []string{
		"COMPOSITE",
		"ENUM_VALUES",
	}
}

// GetMappingEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityTypeEnum(val string) (EntityTypeEnum, bool) {
	enum, ok := mappingEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
