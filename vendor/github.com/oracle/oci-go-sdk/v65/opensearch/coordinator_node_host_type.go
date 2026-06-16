// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// CoordinatorNodeHostTypeEnum Enum with underlying type: string
type CoordinatorNodeHostTypeEnum string

// Set of constants representing the allowable values for CoordinatorNodeHostTypeEnum
const (
	CoordinatorNodeHostTypeFlex CoordinatorNodeHostTypeEnum = "FLEX"
)

var mappingCoordinatorNodeHostTypeEnum = map[string]CoordinatorNodeHostTypeEnum{
	"FLEX": CoordinatorNodeHostTypeFlex,
}

var mappingCoordinatorNodeHostTypeEnumLowerCase = map[string]CoordinatorNodeHostTypeEnum{
	"flex": CoordinatorNodeHostTypeFlex,
}

// GetCoordinatorNodeHostTypeEnumValues Enumerates the set of values for CoordinatorNodeHostTypeEnum
func GetCoordinatorNodeHostTypeEnumValues() []CoordinatorNodeHostTypeEnum {
	values := make([]CoordinatorNodeHostTypeEnum, 0)
	for _, v := range mappingCoordinatorNodeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCoordinatorNodeHostTypeEnumStringValues Enumerates the set of values in String for CoordinatorNodeHostTypeEnum
func GetCoordinatorNodeHostTypeEnumStringValues() []string {
	return []string{
		"FLEX",
	}
}

// GetMappingCoordinatorNodeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCoordinatorNodeHostTypeEnum(val string) (CoordinatorNodeHostTypeEnum, bool) {
	enum, ok := mappingCoordinatorNodeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
