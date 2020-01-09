// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

// ApplicationLanguageEnum Enum with underlying type: string
type ApplicationLanguageEnum string

// Set of constants representing the allowable values for ApplicationLanguageEnum
const (
	ApplicationLanguageScala  ApplicationLanguageEnum = "SCALA"
	ApplicationLanguageJava   ApplicationLanguageEnum = "JAVA"
	ApplicationLanguagePython ApplicationLanguageEnum = "PYTHON"
	ApplicationLanguageSql    ApplicationLanguageEnum = "SQL"
)

var mappingApplicationLanguage = map[string]ApplicationLanguageEnum{
	"SCALA":  ApplicationLanguageScala,
	"JAVA":   ApplicationLanguageJava,
	"PYTHON": ApplicationLanguagePython,
	"SQL":    ApplicationLanguageSql,
}

// GetApplicationLanguageEnumValues Enumerates the set of values for ApplicationLanguageEnum
func GetApplicationLanguageEnumValues() []ApplicationLanguageEnum {
	values := make([]ApplicationLanguageEnum, 0)
	for _, v := range mappingApplicationLanguage {
		values = append(values, v)
	}
	return values
}
