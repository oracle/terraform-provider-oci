// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Utilities API
//
// The APIs for Analyze Applications and other utilities of Java Management Service.
//

package jmsutils

import (
	"strings"
)

// JavaMigrationAnalysisSortByEnum Enum with underlying type: string
type JavaMigrationAnalysisSortByEnum string

// Set of constants representing the allowable values for JavaMigrationAnalysisSortByEnum
const (
	JavaMigrationAnalysisSortByTimeCreated  JavaMigrationAnalysisSortByEnum = "TIME_CREATED"
	JavaMigrationAnalysisSortByTimeStarted  JavaMigrationAnalysisSortByEnum = "TIME_STARTED"
	JavaMigrationAnalysisSortByTimeFinished JavaMigrationAnalysisSortByEnum = "TIME_FINISHED"
)

var mappingJavaMigrationAnalysisSortByEnum = map[string]JavaMigrationAnalysisSortByEnum{
	"TIME_CREATED":  JavaMigrationAnalysisSortByTimeCreated,
	"TIME_STARTED":  JavaMigrationAnalysisSortByTimeStarted,
	"TIME_FINISHED": JavaMigrationAnalysisSortByTimeFinished,
}

var mappingJavaMigrationAnalysisSortByEnumLowerCase = map[string]JavaMigrationAnalysisSortByEnum{
	"time_created":  JavaMigrationAnalysisSortByTimeCreated,
	"time_started":  JavaMigrationAnalysisSortByTimeStarted,
	"time_finished": JavaMigrationAnalysisSortByTimeFinished,
}

// GetJavaMigrationAnalysisSortByEnumValues Enumerates the set of values for JavaMigrationAnalysisSortByEnum
func GetJavaMigrationAnalysisSortByEnumValues() []JavaMigrationAnalysisSortByEnum {
	values := make([]JavaMigrationAnalysisSortByEnum, 0)
	for _, v := range mappingJavaMigrationAnalysisSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaMigrationAnalysisSortByEnumStringValues Enumerates the set of values in String for JavaMigrationAnalysisSortByEnum
func GetJavaMigrationAnalysisSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"TIME_STARTED",
		"TIME_FINISHED",
	}
}

// GetMappingJavaMigrationAnalysisSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaMigrationAnalysisSortByEnum(val string) (JavaMigrationAnalysisSortByEnum, bool) {
	enum, ok := mappingJavaMigrationAnalysisSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
