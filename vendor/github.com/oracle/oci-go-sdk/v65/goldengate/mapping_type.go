// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"strings"
)

// MappingTypeEnum Enum with underlying type: string
type MappingTypeEnum string

// Set of constants representing the allowable values for MappingTypeEnum
const (
	MappingTypeInclude MappingTypeEnum = "INCLUDE"
	MappingTypeExclude MappingTypeEnum = "EXCLUDE"
)

var mappingMappingTypeEnum = map[string]MappingTypeEnum{
	"INCLUDE": MappingTypeInclude,
	"EXCLUDE": MappingTypeExclude,
}

var mappingMappingTypeEnumLowerCase = map[string]MappingTypeEnum{
	"include": MappingTypeInclude,
	"exclude": MappingTypeExclude,
}

// GetMappingTypeEnumValues Enumerates the set of values for MappingTypeEnum
func GetMappingTypeEnumValues() []MappingTypeEnum {
	values := make([]MappingTypeEnum, 0)
	for _, v := range mappingMappingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMappingTypeEnumStringValues Enumerates the set of values in String for MappingTypeEnum
func GetMappingTypeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingMappingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMappingTypeEnum(val string) (MappingTypeEnum, bool) {
	enum, ok := mappingMappingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
