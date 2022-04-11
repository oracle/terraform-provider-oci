// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"strings"
)

// QueryStateEnum Enum with underlying type: string
type QueryStateEnum string

// Set of constants representing the allowable values for QueryStateEnum
const (
	QueryStateNotQueryable      QueryStateEnum = "NotQueryable"
	QueryStateMovingToQueryTier QueryStateEnum = "MovingToQueryTier"
	QueryStateQueryable         QueryStateEnum = "Queryable"
)

var mappingQueryStateEnum = map[string]QueryStateEnum{
	"NotQueryable":      QueryStateNotQueryable,
	"MovingToQueryTier": QueryStateMovingToQueryTier,
	"Queryable":         QueryStateQueryable,
}

var mappingQueryStateEnumLowerCase = map[string]QueryStateEnum{
	"notqueryable":      QueryStateNotQueryable,
	"movingtoquerytier": QueryStateMovingToQueryTier,
	"queryable":         QueryStateQueryable,
}

// GetQueryStateEnumValues Enumerates the set of values for QueryStateEnum
func GetQueryStateEnumValues() []QueryStateEnum {
	values := make([]QueryStateEnum, 0)
	for _, v := range mappingQueryStateEnum {
		values = append(values, v)
	}
	return values
}

// GetQueryStateEnumStringValues Enumerates the set of values in String for QueryStateEnum
func GetQueryStateEnumStringValues() []string {
	return []string{
		"NotQueryable",
		"MovingToQueryTier",
		"Queryable",
	}
}

// GetMappingQueryStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQueryStateEnum(val string) (QueryStateEnum, bool) {
	enum, ok := mappingQueryStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}