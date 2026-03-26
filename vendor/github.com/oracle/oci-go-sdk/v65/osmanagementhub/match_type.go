// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MatchTypeEnum Enum with underlying type: string
type MatchTypeEnum string

// Set of constants representing the allowable values for MatchTypeEnum
const (
	MatchTypeAny MatchTypeEnum = "ANY"
	MatchTypeAll MatchTypeEnum = "ALL"
)

var mappingMatchTypeEnum = map[string]MatchTypeEnum{
	"ANY": MatchTypeAny,
	"ALL": MatchTypeAll,
}

var mappingMatchTypeEnumLowerCase = map[string]MatchTypeEnum{
	"any": MatchTypeAny,
	"all": MatchTypeAll,
}

// GetMatchTypeEnumValues Enumerates the set of values for MatchTypeEnum
func GetMatchTypeEnumValues() []MatchTypeEnum {
	values := make([]MatchTypeEnum, 0)
	for _, v := range mappingMatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMatchTypeEnumStringValues Enumerates the set of values in String for MatchTypeEnum
func GetMatchTypeEnumStringValues() []string {
	return []string{
		"ANY",
		"ALL",
	}
}

// GetMappingMatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMatchTypeEnum(val string) (MatchTypeEnum, bool) {
	enum, ok := mappingMatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
