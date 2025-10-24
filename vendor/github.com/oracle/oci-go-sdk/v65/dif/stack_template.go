// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"strings"
)

// StackTemplateEnum Enum with underlying type: string
type StackTemplateEnum string

// Set of constants representing the allowable values for StackTemplateEnum
const (
	StackTemplateDatalake           StackTemplateEnum = "DATALAKE"
	StackTemplateDatapipeline       StackTemplateEnum = "DATAPIPELINE"
	StackTemplateAiservices         StackTemplateEnum = "AISERVICES"
	StackTemplateDatatransformation StackTemplateEnum = "DATATRANSFORMATION"
)

var mappingStackTemplateEnum = map[string]StackTemplateEnum{
	"DATALAKE":           StackTemplateDatalake,
	"DATAPIPELINE":       StackTemplateDatapipeline,
	"AISERVICES":         StackTemplateAiservices,
	"DATATRANSFORMATION": StackTemplateDatatransformation,
}

var mappingStackTemplateEnumLowerCase = map[string]StackTemplateEnum{
	"datalake":           StackTemplateDatalake,
	"datapipeline":       StackTemplateDatapipeline,
	"aiservices":         StackTemplateAiservices,
	"datatransformation": StackTemplateDatatransformation,
}

// GetStackTemplateEnumValues Enumerates the set of values for StackTemplateEnum
func GetStackTemplateEnumValues() []StackTemplateEnum {
	values := make([]StackTemplateEnum, 0)
	for _, v := range mappingStackTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetStackTemplateEnumStringValues Enumerates the set of values in String for StackTemplateEnum
func GetStackTemplateEnumStringValues() []string {
	return []string{
		"DATALAKE",
		"DATAPIPELINE",
		"AISERVICES",
		"DATATRANSFORMATION",
	}
}

// GetMappingStackTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackTemplateEnum(val string) (StackTemplateEnum, bool) {
	enum, ok := mappingStackTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
