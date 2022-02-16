// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// SqlTuningTaskCredentialTypesEnum Enum with underlying type: string
type SqlTuningTaskCredentialTypesEnum string

// Set of constants representing the allowable values for SqlTuningTaskCredentialTypesEnum
const (
	SqlTuningTaskCredentialTypesSecret   SqlTuningTaskCredentialTypesEnum = "SECRET"
	SqlTuningTaskCredentialTypesPassword SqlTuningTaskCredentialTypesEnum = "PASSWORD"
)

var mappingSqlTuningTaskCredentialTypesEnum = map[string]SqlTuningTaskCredentialTypesEnum{
	"SECRET":   SqlTuningTaskCredentialTypesSecret,
	"PASSWORD": SqlTuningTaskCredentialTypesPassword,
}

// GetSqlTuningTaskCredentialTypesEnumValues Enumerates the set of values for SqlTuningTaskCredentialTypesEnum
func GetSqlTuningTaskCredentialTypesEnumValues() []SqlTuningTaskCredentialTypesEnum {
	values := make([]SqlTuningTaskCredentialTypesEnum, 0)
	for _, v := range mappingSqlTuningTaskCredentialTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningTaskCredentialTypesEnumStringValues Enumerates the set of values in String for SqlTuningTaskCredentialTypesEnum
func GetSqlTuningTaskCredentialTypesEnumStringValues() []string {
	return []string{
		"SECRET",
		"PASSWORD",
	}
}

// GetMappingSqlTuningTaskCredentialTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningTaskCredentialTypesEnum(val string) (SqlTuningTaskCredentialTypesEnum, bool) {
	mappingSqlTuningTaskCredentialTypesEnumIgnoreCase := make(map[string]SqlTuningTaskCredentialTypesEnum)
	for k, v := range mappingSqlTuningTaskCredentialTypesEnum {
		mappingSqlTuningTaskCredentialTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSqlTuningTaskCredentialTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
