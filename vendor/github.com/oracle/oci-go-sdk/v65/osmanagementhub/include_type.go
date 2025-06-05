// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// IncludeTypeEnum Enum with underlying type: string
type IncludeTypeEnum string

// Set of constants representing the allowable values for IncludeTypeEnum
const (
	IncludeTypeAny IncludeTypeEnum = "ANY"
	IncludeTypeAll IncludeTypeEnum = "ALL"
)

var mappingIncludeTypeEnum = map[string]IncludeTypeEnum{
	"ANY": IncludeTypeAny,
	"ALL": IncludeTypeAll,
}

var mappingIncludeTypeEnumLowerCase = map[string]IncludeTypeEnum{
	"any": IncludeTypeAny,
	"all": IncludeTypeAll,
}

// GetIncludeTypeEnumValues Enumerates the set of values for IncludeTypeEnum
func GetIncludeTypeEnumValues() []IncludeTypeEnum {
	values := make([]IncludeTypeEnum, 0)
	for _, v := range mappingIncludeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIncludeTypeEnumStringValues Enumerates the set of values in String for IncludeTypeEnum
func GetIncludeTypeEnumStringValues() []string {
	return []string{
		"ANY",
		"ALL",
	}
}

// GetMappingIncludeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIncludeTypeEnum(val string) (IncludeTypeEnum, bool) {
	enum, ok := mappingIncludeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
