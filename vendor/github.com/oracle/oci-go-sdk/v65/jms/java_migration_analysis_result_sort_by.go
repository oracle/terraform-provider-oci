// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// JavaMigrationAnalysisResultSortByEnum Enum with underlying type: string
type JavaMigrationAnalysisResultSortByEnum string

// Set of constants representing the allowable values for JavaMigrationAnalysisResultSortByEnum
const (
	JavaMigrationAnalysisResultSortByTimeCreated       JavaMigrationAnalysisResultSortByEnum = "timeCreated"
	JavaMigrationAnalysisResultSortByManagedInstanceId JavaMigrationAnalysisResultSortByEnum = "managedInstanceId"
	JavaMigrationAnalysisResultSortByWorkRequestId     JavaMigrationAnalysisResultSortByEnum = "workRequestId"
)

var mappingJavaMigrationAnalysisResultSortByEnum = map[string]JavaMigrationAnalysisResultSortByEnum{
	"timeCreated":       JavaMigrationAnalysisResultSortByTimeCreated,
	"managedInstanceId": JavaMigrationAnalysisResultSortByManagedInstanceId,
	"workRequestId":     JavaMigrationAnalysisResultSortByWorkRequestId,
}

var mappingJavaMigrationAnalysisResultSortByEnumLowerCase = map[string]JavaMigrationAnalysisResultSortByEnum{
	"timecreated":       JavaMigrationAnalysisResultSortByTimeCreated,
	"managedinstanceid": JavaMigrationAnalysisResultSortByManagedInstanceId,
	"workrequestid":     JavaMigrationAnalysisResultSortByWorkRequestId,
}

// GetJavaMigrationAnalysisResultSortByEnumValues Enumerates the set of values for JavaMigrationAnalysisResultSortByEnum
func GetJavaMigrationAnalysisResultSortByEnumValues() []JavaMigrationAnalysisResultSortByEnum {
	values := make([]JavaMigrationAnalysisResultSortByEnum, 0)
	for _, v := range mappingJavaMigrationAnalysisResultSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaMigrationAnalysisResultSortByEnumStringValues Enumerates the set of values in String for JavaMigrationAnalysisResultSortByEnum
func GetJavaMigrationAnalysisResultSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"managedInstanceId",
		"workRequestId",
	}
}

// GetMappingJavaMigrationAnalysisResultSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaMigrationAnalysisResultSortByEnum(val string) (JavaMigrationAnalysisResultSortByEnum, bool) {
	enum, ok := mappingJavaMigrationAnalysisResultSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
