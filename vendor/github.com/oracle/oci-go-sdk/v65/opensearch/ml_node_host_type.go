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

// MlNodeHostTypeEnum Enum with underlying type: string
type MlNodeHostTypeEnum string

// Set of constants representing the allowable values for MlNodeHostTypeEnum
const (
	MlNodeHostTypeFlex MlNodeHostTypeEnum = "FLEX"
)

var mappingMlNodeHostTypeEnum = map[string]MlNodeHostTypeEnum{
	"FLEX": MlNodeHostTypeFlex,
}

var mappingMlNodeHostTypeEnumLowerCase = map[string]MlNodeHostTypeEnum{
	"flex": MlNodeHostTypeFlex,
}

// GetMlNodeHostTypeEnumValues Enumerates the set of values for MlNodeHostTypeEnum
func GetMlNodeHostTypeEnumValues() []MlNodeHostTypeEnum {
	values := make([]MlNodeHostTypeEnum, 0)
	for _, v := range mappingMlNodeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMlNodeHostTypeEnumStringValues Enumerates the set of values in String for MlNodeHostTypeEnum
func GetMlNodeHostTypeEnumStringValues() []string {
	return []string{
		"FLEX",
	}
}

// GetMappingMlNodeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMlNodeHostTypeEnum(val string) (MlNodeHostTypeEnum, bool) {
	enum, ok := mappingMlNodeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
