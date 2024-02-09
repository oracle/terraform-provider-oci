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

// EntityAttributeTypeEnum Enum with underlying type: string
type EntityAttributeTypeEnum string

// Set of constants representing the allowable values for EntityAttributeTypeEnum
const (
	EntityAttributeTypeText               EntityAttributeTypeEnum = "TEXT"
	EntityAttributeTypeNumber             EntityAttributeTypeEnum = "NUMBER"
	EntityAttributeTypeDateTime           EntityAttributeTypeEnum = "DATE_TIME"
	EntityAttributeTypeDate               EntityAttributeTypeEnum = "DATE"
	EntityAttributeTypeEntity             EntityAttributeTypeEnum = "ENTITY"
	EntityAttributeTypeCompositeEntity    EntityAttributeTypeEnum = "COMPOSITE_ENTITY"
	EntityAttributeTypeAttributeReference EntityAttributeTypeEnum = "ATTRIBUTE_REFERENCE"
	EntityAttributeTypeBoolean            EntityAttributeTypeEnum = "BOOLEAN"
)

var mappingEntityAttributeTypeEnum = map[string]EntityAttributeTypeEnum{
	"TEXT":                EntityAttributeTypeText,
	"NUMBER":              EntityAttributeTypeNumber,
	"DATE_TIME":           EntityAttributeTypeDateTime,
	"DATE":                EntityAttributeTypeDate,
	"ENTITY":              EntityAttributeTypeEntity,
	"COMPOSITE_ENTITY":    EntityAttributeTypeCompositeEntity,
	"ATTRIBUTE_REFERENCE": EntityAttributeTypeAttributeReference,
	"BOOLEAN":             EntityAttributeTypeBoolean,
}

var mappingEntityAttributeTypeEnumLowerCase = map[string]EntityAttributeTypeEnum{
	"text":                EntityAttributeTypeText,
	"number":              EntityAttributeTypeNumber,
	"date_time":           EntityAttributeTypeDateTime,
	"date":                EntityAttributeTypeDate,
	"entity":              EntityAttributeTypeEntity,
	"composite_entity":    EntityAttributeTypeCompositeEntity,
	"attribute_reference": EntityAttributeTypeAttributeReference,
	"boolean":             EntityAttributeTypeBoolean,
}

// GetEntityAttributeTypeEnumValues Enumerates the set of values for EntityAttributeTypeEnum
func GetEntityAttributeTypeEnumValues() []EntityAttributeTypeEnum {
	values := make([]EntityAttributeTypeEnum, 0)
	for _, v := range mappingEntityAttributeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEntityAttributeTypeEnumStringValues Enumerates the set of values in String for EntityAttributeTypeEnum
func GetEntityAttributeTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"NUMBER",
		"DATE_TIME",
		"DATE",
		"ENTITY",
		"COMPOSITE_ENTITY",
		"ATTRIBUTE_REFERENCE",
		"BOOLEAN",
	}
}

// GetMappingEntityAttributeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEntityAttributeTypeEnum(val string) (EntityAttributeTypeEnum, bool) {
	enum, ok := mappingEntityAttributeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
