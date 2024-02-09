// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LanguageEnum Enum with underlying type: string
type LanguageEnum string

// Set of constants representing the allowable values for LanguageEnum
const (
	LanguageJava   LanguageEnum = "JAVA"
	LanguageScala  LanguageEnum = "SCALA"
	LanguagePython LanguageEnum = "PYTHON"
	LanguageSql    LanguageEnum = "SQL"
)

var mappingLanguageEnum = map[string]LanguageEnum{
	"JAVA":   LanguageJava,
	"SCALA":  LanguageScala,
	"PYTHON": LanguagePython,
	"SQL":    LanguageSql,
}

var mappingLanguageEnumLowerCase = map[string]LanguageEnum{
	"java":   LanguageJava,
	"scala":  LanguageScala,
	"python": LanguagePython,
	"sql":    LanguageSql,
}

// GetLanguageEnumValues Enumerates the set of values for LanguageEnum
func GetLanguageEnumValues() []LanguageEnum {
	values := make([]LanguageEnum, 0)
	for _, v := range mappingLanguageEnum {
		values = append(values, v)
	}
	return values
}

// GetLanguageEnumStringValues Enumerates the set of values in String for LanguageEnum
func GetLanguageEnumStringValues() []string {
	return []string{
		"JAVA",
		"SCALA",
		"PYTHON",
		"SQL",
	}
}

// GetMappingLanguageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLanguageEnum(val string) (LanguageEnum, bool) {
	enum, ok := mappingLanguageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
