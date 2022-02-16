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

// EditModesEnum Enum with underlying type: string
type EditModesEnum string

// Set of constants representing the allowable values for EditModesEnum
const (
	EditModesReadOnly   EditModesEnum = "READ_ONLY"
	EditModesWritable   EditModesEnum = "WRITABLE"
	EditModesExtensible EditModesEnum = "EXTENSIBLE"
)

var mappingEditModesEnum = map[string]EditModesEnum{
	"READ_ONLY":  EditModesReadOnly,
	"WRITABLE":   EditModesWritable,
	"EXTENSIBLE": EditModesExtensible,
}

// GetEditModesEnumValues Enumerates the set of values for EditModesEnum
func GetEditModesEnumValues() []EditModesEnum {
	values := make([]EditModesEnum, 0)
	for _, v := range mappingEditModesEnum {
		values = append(values, v)
	}
	return values
}

// GetEditModesEnumStringValues Enumerates the set of values in String for EditModesEnum
func GetEditModesEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"WRITABLE",
		"EXTENSIBLE",
	}
}

// GetMappingEditModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEditModesEnum(val string) (EditModesEnum, bool) {
	mappingEditModesEnumIgnoreCase := make(map[string]EditModesEnum)
	for k, v := range mappingEditModesEnum {
		mappingEditModesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingEditModesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
