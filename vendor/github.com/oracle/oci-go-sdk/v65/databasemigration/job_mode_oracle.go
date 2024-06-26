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

// JobModeOracleEnum Enum with underlying type: string
type JobModeOracleEnum string

// Set of constants representing the allowable values for JobModeOracleEnum
const (
	JobModeOracleFull          JobModeOracleEnum = "FULL"
	JobModeOracleSchema        JobModeOracleEnum = "SCHEMA"
	JobModeOracleTable         JobModeOracleEnum = "TABLE"
	JobModeOracleTablespace    JobModeOracleEnum = "TABLESPACE"
	JobModeOracleTransportable JobModeOracleEnum = "TRANSPORTABLE"
)

var mappingJobModeOracleEnum = map[string]JobModeOracleEnum{
	"FULL":          JobModeOracleFull,
	"SCHEMA":        JobModeOracleSchema,
	"TABLE":         JobModeOracleTable,
	"TABLESPACE":    JobModeOracleTablespace,
	"TRANSPORTABLE": JobModeOracleTransportable,
}

var mappingJobModeOracleEnumLowerCase = map[string]JobModeOracleEnum{
	"full":          JobModeOracleFull,
	"schema":        JobModeOracleSchema,
	"table":         JobModeOracleTable,
	"tablespace":    JobModeOracleTablespace,
	"transportable": JobModeOracleTransportable,
}

// GetJobModeOracleEnumValues Enumerates the set of values for JobModeOracleEnum
func GetJobModeOracleEnumValues() []JobModeOracleEnum {
	values := make([]JobModeOracleEnum, 0)
	for _, v := range mappingJobModeOracleEnum {
		values = append(values, v)
	}
	return values
}

// GetJobModeOracleEnumStringValues Enumerates the set of values in String for JobModeOracleEnum
func GetJobModeOracleEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
		"TABLE",
		"TABLESPACE",
		"TRANSPORTABLE",
	}
}

// GetMappingJobModeOracleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobModeOracleEnum(val string) (JobModeOracleEnum, bool) {
	enum, ok := mappingJobModeOracleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
