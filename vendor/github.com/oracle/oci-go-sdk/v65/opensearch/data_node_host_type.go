// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DataNodeHostTypeEnum Enum with underlying type: string
type DataNodeHostTypeEnum string

// Set of constants representing the allowable values for DataNodeHostTypeEnum
const (
	DataNodeHostTypeFlex DataNodeHostTypeEnum = "FLEX"
	DataNodeHostTypeBm   DataNodeHostTypeEnum = "BM"
)

var mappingDataNodeHostTypeEnum = map[string]DataNodeHostTypeEnum{
	"FLEX": DataNodeHostTypeFlex,
	"BM":   DataNodeHostTypeBm,
}

var mappingDataNodeHostTypeEnumLowerCase = map[string]DataNodeHostTypeEnum{
	"flex": DataNodeHostTypeFlex,
	"bm":   DataNodeHostTypeBm,
}

// GetDataNodeHostTypeEnumValues Enumerates the set of values for DataNodeHostTypeEnum
func GetDataNodeHostTypeEnumValues() []DataNodeHostTypeEnum {
	values := make([]DataNodeHostTypeEnum, 0)
	for _, v := range mappingDataNodeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataNodeHostTypeEnumStringValues Enumerates the set of values in String for DataNodeHostTypeEnum
func GetDataNodeHostTypeEnumStringValues() []string {
	return []string{
		"FLEX",
		"BM",
	}
}

// GetMappingDataNodeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataNodeHostTypeEnum(val string) (DataNodeHostTypeEnum, bool) {
	enum, ok := mappingDataNodeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
