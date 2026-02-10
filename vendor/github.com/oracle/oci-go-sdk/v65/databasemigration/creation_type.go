// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreationTypeEnum Enum with underlying type: string
type CreationTypeEnum string

// Set of constants representing the allowable values for CreationTypeEnum
const (
	CreationTypeCreateOnly            CreationTypeEnum = "CREATE_ONLY"
	CreationTypeCreateAndRunAssessors CreationTypeEnum = "CREATE_AND_RUN_ASSESSORS"
)

var mappingCreationTypeEnum = map[string]CreationTypeEnum{
	"CREATE_ONLY":              CreationTypeCreateOnly,
	"CREATE_AND_RUN_ASSESSORS": CreationTypeCreateAndRunAssessors,
}

var mappingCreationTypeEnumLowerCase = map[string]CreationTypeEnum{
	"create_only":              CreationTypeCreateOnly,
	"create_and_run_assessors": CreationTypeCreateAndRunAssessors,
}

// GetCreationTypeEnumValues Enumerates the set of values for CreationTypeEnum
func GetCreationTypeEnumValues() []CreationTypeEnum {
	values := make([]CreationTypeEnum, 0)
	for _, v := range mappingCreationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreationTypeEnumStringValues Enumerates the set of values in String for CreationTypeEnum
func GetCreationTypeEnumStringValues() []string {
	return []string{
		"CREATE_ONLY",
		"CREATE_AND_RUN_ASSESSORS",
	}
}

// GetMappingCreationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreationTypeEnum(val string) (CreationTypeEnum, bool) {
	enum, ok := mappingCreationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
