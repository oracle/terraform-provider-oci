// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SearchNodeHostTypeEnum Enum with underlying type: string
type SearchNodeHostTypeEnum string

// Set of constants representing the allowable values for SearchNodeHostTypeEnum
const (
	SearchNodeHostTypeFlex SearchNodeHostTypeEnum = "FLEX"
)

var mappingSearchNodeHostTypeEnum = map[string]SearchNodeHostTypeEnum{
	"FLEX": SearchNodeHostTypeFlex,
}

var mappingSearchNodeHostTypeEnumLowerCase = map[string]SearchNodeHostTypeEnum{
	"flex": SearchNodeHostTypeFlex,
}

// GetSearchNodeHostTypeEnumValues Enumerates the set of values for SearchNodeHostTypeEnum
func GetSearchNodeHostTypeEnumValues() []SearchNodeHostTypeEnum {
	values := make([]SearchNodeHostTypeEnum, 0)
	for _, v := range mappingSearchNodeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchNodeHostTypeEnumStringValues Enumerates the set of values in String for SearchNodeHostTypeEnum
func GetSearchNodeHostTypeEnumStringValues() []string {
	return []string{
		"FLEX",
	}
}

// GetMappingSearchNodeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchNodeHostTypeEnum(val string) (SearchNodeHostTypeEnum, bool) {
	enum, ok := mappingSearchNodeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
