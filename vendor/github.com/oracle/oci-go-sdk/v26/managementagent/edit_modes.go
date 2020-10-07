// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

// EditModesEnum Enum with underlying type: string
type EditModesEnum string

// Set of constants representing the allowable values for EditModesEnum
const (
	EditModesReadOnly   EditModesEnum = "READ_ONLY"
	EditModesWritable   EditModesEnum = "WRITABLE"
	EditModesExtensible EditModesEnum = "EXTENSIBLE"
)

var mappingEditModes = map[string]EditModesEnum{
	"READ_ONLY":  EditModesReadOnly,
	"WRITABLE":   EditModesWritable,
	"EXTENSIBLE": EditModesExtensible,
}

// GetEditModesEnumValues Enumerates the set of values for EditModesEnum
func GetEditModesEnumValues() []EditModesEnum {
	values := make([]EditModesEnum, 0)
	for _, v := range mappingEditModes {
		values = append(values, v)
	}
	return values
}
