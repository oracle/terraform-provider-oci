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

// WorkItemTypeEnum Enum with underlying type: string
type WorkItemTypeEnum string

// Set of constants representing the allowable values for WorkItemTypeEnum
const (
	WorkItemTypeLcm                     WorkItemTypeEnum = "LCM"
	WorkItemTypeJfrCapture              WorkItemTypeEnum = "JFR_CAPTURE"
	WorkItemTypeJfrUpload               WorkItemTypeEnum = "JFR_UPLOAD"
	WorkItemTypeCryptoAnalysis          WorkItemTypeEnum = "CRYPTO_ANALYSIS"
	WorkItemTypeCryptoAnalysisMerge     WorkItemTypeEnum = "CRYPTO_ANALYSIS_MERGE"
	WorkItemTypeAdvancedUsageTracking   WorkItemTypeEnum = "ADVANCED_USAGE_TRACKING"
	WorkItemTypeAdvUsageServerMetadata  WorkItemTypeEnum = "ADV_USAGE_SERVER_METADATA"
	WorkItemTypeAdvUsageServerLibraries WorkItemTypeEnum = "ADV_USAGE_SERVER_LIBRARIES"
	WorkItemTypeAdvUsageJavaLibraries   WorkItemTypeEnum = "ADV_USAGE_JAVA_LIBRARIES"
	WorkItemTypePerformanceTuning       WorkItemTypeEnum = "PERFORMANCE_TUNING"
	WorkItemTypeJmigrateAnalysis        WorkItemTypeEnum = "JMIGRATE_ANALYSIS"
	WorkItemTypeJmigrateCreateReport    WorkItemTypeEnum = "JMIGRATE_CREATE_REPORT"
	WorkItemTypeDrs                     WorkItemTypeEnum = "DRS"
)

var mappingWorkItemTypeEnum = map[string]WorkItemTypeEnum{
	"LCM":                        WorkItemTypeLcm,
	"JFR_CAPTURE":                WorkItemTypeJfrCapture,
	"JFR_UPLOAD":                 WorkItemTypeJfrUpload,
	"CRYPTO_ANALYSIS":            WorkItemTypeCryptoAnalysis,
	"CRYPTO_ANALYSIS_MERGE":      WorkItemTypeCryptoAnalysisMerge,
	"ADVANCED_USAGE_TRACKING":    WorkItemTypeAdvancedUsageTracking,
	"ADV_USAGE_SERVER_METADATA":  WorkItemTypeAdvUsageServerMetadata,
	"ADV_USAGE_SERVER_LIBRARIES": WorkItemTypeAdvUsageServerLibraries,
	"ADV_USAGE_JAVA_LIBRARIES":   WorkItemTypeAdvUsageJavaLibraries,
	"PERFORMANCE_TUNING":         WorkItemTypePerformanceTuning,
	"JMIGRATE_ANALYSIS":          WorkItemTypeJmigrateAnalysis,
	"JMIGRATE_CREATE_REPORT":     WorkItemTypeJmigrateCreateReport,
	"DRS":                        WorkItemTypeDrs,
}

var mappingWorkItemTypeEnumLowerCase = map[string]WorkItemTypeEnum{
	"lcm":                        WorkItemTypeLcm,
	"jfr_capture":                WorkItemTypeJfrCapture,
	"jfr_upload":                 WorkItemTypeJfrUpload,
	"crypto_analysis":            WorkItemTypeCryptoAnalysis,
	"crypto_analysis_merge":      WorkItemTypeCryptoAnalysisMerge,
	"advanced_usage_tracking":    WorkItemTypeAdvancedUsageTracking,
	"adv_usage_server_metadata":  WorkItemTypeAdvUsageServerMetadata,
	"adv_usage_server_libraries": WorkItemTypeAdvUsageServerLibraries,
	"adv_usage_java_libraries":   WorkItemTypeAdvUsageJavaLibraries,
	"performance_tuning":         WorkItemTypePerformanceTuning,
	"jmigrate_analysis":          WorkItemTypeJmigrateAnalysis,
	"jmigrate_create_report":     WorkItemTypeJmigrateCreateReport,
	"drs":                        WorkItemTypeDrs,
}

// GetWorkItemTypeEnumValues Enumerates the set of values for WorkItemTypeEnum
func GetWorkItemTypeEnumValues() []WorkItemTypeEnum {
	values := make([]WorkItemTypeEnum, 0)
	for _, v := range mappingWorkItemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkItemTypeEnumStringValues Enumerates the set of values in String for WorkItemTypeEnum
func GetWorkItemTypeEnumStringValues() []string {
	return []string{
		"LCM",
		"JFR_CAPTURE",
		"JFR_UPLOAD",
		"CRYPTO_ANALYSIS",
		"CRYPTO_ANALYSIS_MERGE",
		"ADVANCED_USAGE_TRACKING",
		"ADV_USAGE_SERVER_METADATA",
		"ADV_USAGE_SERVER_LIBRARIES",
		"ADV_USAGE_JAVA_LIBRARIES",
		"PERFORMANCE_TUNING",
		"JMIGRATE_ANALYSIS",
		"JMIGRATE_CREATE_REPORT",
		"DRS",
	}
}

// GetMappingWorkItemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkItemTypeEnum(val string) (WorkItemTypeEnum, bool) {
	enum, ok := mappingWorkItemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
