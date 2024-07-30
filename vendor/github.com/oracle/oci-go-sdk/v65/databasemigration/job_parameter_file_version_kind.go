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

// JobParameterFileVersionKindEnum Enum with underlying type: string
type JobParameterFileVersionKindEnum string

// Set of constants representing the allowable values for JobParameterFileVersionKindEnum
const (
	JobParameterFileVersionKindExtract  JobParameterFileVersionKindEnum = "EXTRACT"
	JobParameterFileVersionKindReplicat JobParameterFileVersionKindEnum = "REPLICAT"
)

var mappingJobParameterFileVersionKindEnum = map[string]JobParameterFileVersionKindEnum{
	"EXTRACT":  JobParameterFileVersionKindExtract,
	"REPLICAT": JobParameterFileVersionKindReplicat,
}

var mappingJobParameterFileVersionKindEnumLowerCase = map[string]JobParameterFileVersionKindEnum{
	"extract":  JobParameterFileVersionKindExtract,
	"replicat": JobParameterFileVersionKindReplicat,
}

// GetJobParameterFileVersionKindEnumValues Enumerates the set of values for JobParameterFileVersionKindEnum
func GetJobParameterFileVersionKindEnumValues() []JobParameterFileVersionKindEnum {
	values := make([]JobParameterFileVersionKindEnum, 0)
	for _, v := range mappingJobParameterFileVersionKindEnum {
		values = append(values, v)
	}
	return values
}

// GetJobParameterFileVersionKindEnumStringValues Enumerates the set of values in String for JobParameterFileVersionKindEnum
func GetJobParameterFileVersionKindEnumStringValues() []string {
	return []string{
		"EXTRACT",
		"REPLICAT",
	}
}

// GetMappingJobParameterFileVersionKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobParameterFileVersionKindEnum(val string) (JobParameterFileVersionKindEnum, bool) {
	enum, ok := mappingJobParameterFileVersionKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
