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

// FilterTypeEnum Enum with underlying type: string
type FilterTypeEnum string

// Set of constants representing the allowable values for FilterTypeEnum
const (
	FilterTypeInclude FilterTypeEnum = "INCLUDE"
	FilterTypeExclude FilterTypeEnum = "EXCLUDE"
)

var mappingFilterTypeEnum = map[string]FilterTypeEnum{
	"INCLUDE": FilterTypeInclude,
	"EXCLUDE": FilterTypeExclude,
}

var mappingFilterTypeEnumLowerCase = map[string]FilterTypeEnum{
	"include": FilterTypeInclude,
	"exclude": FilterTypeExclude,
}

// GetFilterTypeEnumValues Enumerates the set of values for FilterTypeEnum
func GetFilterTypeEnumValues() []FilterTypeEnum {
	values := make([]FilterTypeEnum, 0)
	for _, v := range mappingFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFilterTypeEnumStringValues Enumerates the set of values in String for FilterTypeEnum
func GetFilterTypeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilterTypeEnum(val string) (FilterTypeEnum, bool) {
	enum, ok := mappingFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
