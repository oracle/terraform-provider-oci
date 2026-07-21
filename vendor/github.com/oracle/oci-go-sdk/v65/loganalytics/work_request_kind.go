// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// WorkRequestKindEnum Enum with underlying type: string
type WorkRequestKindEnum string

// Set of constants representing the allowable values for WorkRequestKindEnum
const (
	WorkRequestKindGenerateAiContent WorkRequestKindEnum = "GENERATE_AI_CONTENT"
)

var mappingWorkRequestKindEnum = map[string]WorkRequestKindEnum{
	"GENERATE_AI_CONTENT": WorkRequestKindGenerateAiContent,
}

var mappingWorkRequestKindEnumLowerCase = map[string]WorkRequestKindEnum{
	"generate_ai_content": WorkRequestKindGenerateAiContent,
}

// GetWorkRequestKindEnumValues Enumerates the set of values for WorkRequestKindEnum
func GetWorkRequestKindEnumValues() []WorkRequestKindEnum {
	values := make([]WorkRequestKindEnum, 0)
	for _, v := range mappingWorkRequestKindEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestKindEnumStringValues Enumerates the set of values in String for WorkRequestKindEnum
func GetWorkRequestKindEnumStringValues() []string {
	return []string{
		"GENERATE_AI_CONTENT",
	}
}

// GetMappingWorkRequestKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestKindEnum(val string) (WorkRequestKindEnum, bool) {
	enum, ok := mappingWorkRequestKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
