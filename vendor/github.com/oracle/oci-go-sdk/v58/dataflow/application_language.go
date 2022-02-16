// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// ApplicationLanguageEnum Enum with underlying type: string
type ApplicationLanguageEnum string

// Set of constants representing the allowable values for ApplicationLanguageEnum
const (
	ApplicationLanguageScala  ApplicationLanguageEnum = "SCALA"
	ApplicationLanguageJava   ApplicationLanguageEnum = "JAVA"
	ApplicationLanguagePython ApplicationLanguageEnum = "PYTHON"
	ApplicationLanguageSql    ApplicationLanguageEnum = "SQL"
)

var mappingApplicationLanguageEnum = map[string]ApplicationLanguageEnum{
	"SCALA":  ApplicationLanguageScala,
	"JAVA":   ApplicationLanguageJava,
	"PYTHON": ApplicationLanguagePython,
	"SQL":    ApplicationLanguageSql,
}

// GetApplicationLanguageEnumValues Enumerates the set of values for ApplicationLanguageEnum
func GetApplicationLanguageEnumValues() []ApplicationLanguageEnum {
	values := make([]ApplicationLanguageEnum, 0)
	for _, v := range mappingApplicationLanguageEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationLanguageEnumStringValues Enumerates the set of values in String for ApplicationLanguageEnum
func GetApplicationLanguageEnumStringValues() []string {
	return []string{
		"SCALA",
		"JAVA",
		"PYTHON",
		"SQL",
	}
}

// GetMappingApplicationLanguageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationLanguageEnum(val string) (ApplicationLanguageEnum, bool) {
	mappingApplicationLanguageEnumIgnoreCase := make(map[string]ApplicationLanguageEnum)
	for k, v := range mappingApplicationLanguageEnum {
		mappingApplicationLanguageEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingApplicationLanguageEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
