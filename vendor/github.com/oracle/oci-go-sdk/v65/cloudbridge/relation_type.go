// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// RelationTypeEnum Enum with underlying type: string
type RelationTypeEnum string

// Set of constants representing the allowable values for RelationTypeEnum
const (
	RelationTypeAssociation RelationTypeEnum = "ASSOCIATION"
	RelationTypeDependency  RelationTypeEnum = "DEPENDENCY"
	RelationTypeComposition RelationTypeEnum = "COMPOSITION"
)

var mappingRelationTypeEnum = map[string]RelationTypeEnum{
	"ASSOCIATION": RelationTypeAssociation,
	"DEPENDENCY":  RelationTypeDependency,
	"COMPOSITION": RelationTypeComposition,
}

var mappingRelationTypeEnumLowerCase = map[string]RelationTypeEnum{
	"association": RelationTypeAssociation,
	"dependency":  RelationTypeDependency,
	"composition": RelationTypeComposition,
}

// GetRelationTypeEnumValues Enumerates the set of values for RelationTypeEnum
func GetRelationTypeEnumValues() []RelationTypeEnum {
	values := make([]RelationTypeEnum, 0)
	for _, v := range mappingRelationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRelationTypeEnumStringValues Enumerates the set of values in String for RelationTypeEnum
func GetRelationTypeEnumStringValues() []string {
	return []string{
		"ASSOCIATION",
		"DEPENDENCY",
		"COMPOSITION",
	}
}

// GetMappingRelationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRelationTypeEnum(val string) (RelationTypeEnum, bool) {
	enum, ok := mappingRelationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
