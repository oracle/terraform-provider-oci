// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// JobTypesEnum Enum with underlying type: string
type JobTypesEnum string

// Set of constants representing the allowable values for JobTypesEnum
const (
	JobTypesSql JobTypesEnum = "SQL"
)

var mappingJobTypesEnum = map[string]JobTypesEnum{
	"SQL": JobTypesSql,
}

var mappingJobTypesEnumLowerCase = map[string]JobTypesEnum{
	"sql": JobTypesSql,
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
		"SQL",
	}
}

// GetMappingJobTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJobTypesEnum(val string) (JobTypesEnum, bool) {
	enum, ok := mappingJobTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
