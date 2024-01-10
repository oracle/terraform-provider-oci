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

// MasterNodeHostTypeEnum Enum with underlying type: string
type MasterNodeHostTypeEnum string

// Set of constants representing the allowable values for MasterNodeHostTypeEnum
const (
	MasterNodeHostTypeFlex MasterNodeHostTypeEnum = "FLEX"
	MasterNodeHostTypeBm   MasterNodeHostTypeEnum = "BM"
)

var mappingMasterNodeHostTypeEnum = map[string]MasterNodeHostTypeEnum{
	"FLEX": MasterNodeHostTypeFlex,
	"BM":   MasterNodeHostTypeBm,
}

var mappingMasterNodeHostTypeEnumLowerCase = map[string]MasterNodeHostTypeEnum{
	"flex": MasterNodeHostTypeFlex,
	"bm":   MasterNodeHostTypeBm,
}

// GetMasterNodeHostTypeEnumValues Enumerates the set of values for MasterNodeHostTypeEnum
func GetMasterNodeHostTypeEnumValues() []MasterNodeHostTypeEnum {
	values := make([]MasterNodeHostTypeEnum, 0)
	for _, v := range mappingMasterNodeHostTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMasterNodeHostTypeEnumStringValues Enumerates the set of values in String for MasterNodeHostTypeEnum
func GetMasterNodeHostTypeEnumStringValues() []string {
	return []string{
		"FLEX",
		"BM",
	}
}

// GetMappingMasterNodeHostTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMasterNodeHostTypeEnum(val string) (MasterNodeHostTypeEnum, bool) {
	enum, ok := mappingMasterNodeHostTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
