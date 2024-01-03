// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// ObjectStatusEnum Enum with underlying type: string
type ObjectStatusEnum string

// Set of constants representing the allowable values for ObjectStatusEnum
const (
	ObjectStatusExclude ObjectStatusEnum = "EXCLUDE"
	ObjectStatusInclude ObjectStatusEnum = "INCLUDE"
)

var mappingObjectStatusEnum = map[string]ObjectStatusEnum{
	"EXCLUDE": ObjectStatusExclude,
	"INCLUDE": ObjectStatusInclude,
}

var mappingObjectStatusEnumLowerCase = map[string]ObjectStatusEnum{
	"exclude": ObjectStatusExclude,
	"include": ObjectStatusInclude,
}

// GetObjectStatusEnumValues Enumerates the set of values for ObjectStatusEnum
func GetObjectStatusEnumValues() []ObjectStatusEnum {
	values := make([]ObjectStatusEnum, 0)
	for _, v := range mappingObjectStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetObjectStatusEnumStringValues Enumerates the set of values in String for ObjectStatusEnum
func GetObjectStatusEnumStringValues() []string {
	return []string{
		"EXCLUDE",
		"INCLUDE",
	}
}

// GetMappingObjectStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingObjectStatusEnum(val string) (ObjectStatusEnum, bool) {
	enum, ok := mappingObjectStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
