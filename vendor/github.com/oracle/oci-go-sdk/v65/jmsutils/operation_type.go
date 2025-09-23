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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeRequestJavaMigrationSaAnalysis     OperationTypeEnum = "REQUEST_JAVA_MIGRATION_SA_ANALYSIS"
	OperationTypeRequestPerformanceTuningSaAnalysis OperationTypeEnum = "REQUEST_PERFORMANCE_TUNING_SA_ANALYSIS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"REQUEST_JAVA_MIGRATION_SA_ANALYSIS":     OperationTypeRequestJavaMigrationSaAnalysis,
	"REQUEST_PERFORMANCE_TUNING_SA_ANALYSIS": OperationTypeRequestPerformanceTuningSaAnalysis,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"request_java_migration_sa_analysis":     OperationTypeRequestJavaMigrationSaAnalysis,
	"request_performance_tuning_sa_analysis": OperationTypeRequestPerformanceTuningSaAnalysis,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"REQUEST_JAVA_MIGRATION_SA_ANALYSIS",
		"REQUEST_PERFORMANCE_TUNING_SA_ANALYSIS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
