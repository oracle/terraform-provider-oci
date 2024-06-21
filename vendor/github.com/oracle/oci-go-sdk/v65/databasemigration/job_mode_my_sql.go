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

// JobModeMySqlEnum Enum with underlying type: string
type JobModeMySqlEnum string

// Set of constants representing the allowable values for JobModeMySqlEnum
const (
	JobModeMySqlFull   JobModeMySqlEnum = "FULL"
	JobModeMySqlSchema JobModeMySqlEnum = "SCHEMA"
)

var mappingJobModeMySqlEnum = map[string]JobModeMySqlEnum{
	"FULL":   JobModeMySqlFull,
	"SCHEMA": JobModeMySqlSchema,
}

var mappingJobModeMySqlEnumLowerCase = map[string]JobModeMySqlEnum{
	"full":   JobModeMySqlFull,
	"schema": JobModeMySqlSchema,
}

// GetJobModeMySqlEnumValues Enumerates the set of values for JobModeMySqlEnum
func GetJobModeMySqlEnumValues() []JobModeMySqlEnum {
	values := make([]JobModeMySqlEnum, 0)
	for _, v := range mappingJobModeMySqlEnum {
		values = append(values, v)
	}
	return values
}

// GetJobModeMySqlEnumStringValues Enumerates the set of values in String for JobModeMySqlEnum
func GetJobModeMySqlEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
	}
}

// GetMappingJobModeMySqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobModeMySqlEnum(val string) (JobModeMySqlEnum, bool) {
	enum, ok := mappingJobModeMySqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
