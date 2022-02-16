// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// JobTypesEnum Enum with underlying type: string
type JobTypesEnum string

// Set of constants representing the allowable values for JobTypesEnum
const (
	JobTypesEvaluation JobTypesEnum = "EVALUATION"
	JobTypesMigration  JobTypesEnum = "MIGRATION"
)

var mappingJobTypesEnum = map[string]JobTypesEnum{
	"EVALUATION": JobTypesEvaluation,
	"MIGRATION":  JobTypesMigration,
}

// GetJobTypesEnumValues Enumerates the set of values for JobTypesEnum
func GetJobTypesEnumValues() []JobTypesEnum {
	values := make([]JobTypesEnum, 0)
	for _, v := range mappingJobTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetJobTypesEnumStringValues Enumerates the set of values in String for JobTypesEnum
func GetJobTypesEnumStringValues() []string {
	return []string{
		"EVALUATION",
		"MIGRATION",
	}
}

// GetMappingJobTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobTypesEnum(val string) (JobTypesEnum, bool) {
	mappingJobTypesEnumIgnoreCase := make(map[string]JobTypesEnum)
	for k, v := range mappingJobTypesEnum {
		mappingJobTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJobTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
