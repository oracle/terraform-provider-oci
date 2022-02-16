// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// DatatypesEnum Enum with underlying type: string
type DatatypesEnum string

// Set of constants representing the allowable values for DatatypesEnum
const (
	DatatypesLong    DatatypesEnum = "LONG"
	DatatypesDouble  DatatypesEnum = "DOUBLE"
	DatatypesInteger DatatypesEnum = "INTEGER"
	DatatypesString  DatatypesEnum = "STRING"
	DatatypesBoolean DatatypesEnum = "BOOLEAN"
)

var mappingDatatypesEnum = map[string]DatatypesEnum{
	"LONG":    DatatypesLong,
	"DOUBLE":  DatatypesDouble,
	"INTEGER": DatatypesInteger,
	"STRING":  DatatypesString,
	"BOOLEAN": DatatypesBoolean,
}

// GetDatatypesEnumValues Enumerates the set of values for DatatypesEnum
func GetDatatypesEnumValues() []DatatypesEnum {
	values := make([]DatatypesEnum, 0)
	for _, v := range mappingDatatypesEnum {
		values = append(values, v)
	}
	return values
}

// GetDatatypesEnumStringValues Enumerates the set of values in String for DatatypesEnum
func GetDatatypesEnumStringValues() []string {
	return []string{
		"LONG",
		"DOUBLE",
		"INTEGER",
		"STRING",
		"BOOLEAN",
	}
}

// GetMappingDatatypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatatypesEnum(val string) (DatatypesEnum, bool) {
	mappingDatatypesEnumIgnoreCase := make(map[string]DatatypesEnum)
	for k, v := range mappingDatatypesEnum {
		mappingDatatypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDatatypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
