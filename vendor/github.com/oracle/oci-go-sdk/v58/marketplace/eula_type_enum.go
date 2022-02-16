// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"strings"
)

// EulaTypeEnumEnum Enum with underlying type: string
type EulaTypeEnumEnum string

// Set of constants representing the allowable values for EulaTypeEnumEnum
const (
	EulaTypeEnumText EulaTypeEnumEnum = "TEXT"
)

var mappingEulaTypeEnumEnum = map[string]EulaTypeEnumEnum{
	"TEXT": EulaTypeEnumText,
}

// GetEulaTypeEnumEnumValues Enumerates the set of values for EulaTypeEnumEnum
func GetEulaTypeEnumEnumValues() []EulaTypeEnumEnum {
	values := make([]EulaTypeEnumEnum, 0)
	for _, v := range mappingEulaTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetEulaTypeEnumEnumStringValues Enumerates the set of values in String for EulaTypeEnumEnum
func GetEulaTypeEnumEnumStringValues() []string {
	return []string{
		"TEXT",
	}
}

// GetMappingEulaTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEulaTypeEnumEnum(val string) (EulaTypeEnumEnum, bool) {
	mappingEulaTypeEnumEnumIgnoreCase := make(map[string]EulaTypeEnumEnum)
	for k, v := range mappingEulaTypeEnumEnum {
		mappingEulaTypeEnumEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingEulaTypeEnumEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
