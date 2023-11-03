// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet                      OperationTypeEnum = "CREATE_FLEET"
	OperationTypeDeleteFleet                      OperationTypeEnum = "DELETE_FLEET"
	OperationTypeMoveFleet                        OperationTypeEnum = "MOVE_FLEET"
	OperationTypeUpdateFleet                      OperationTypeEnum = "UPDATE_FLEET"
	OperationTypeUpdateFleetAgentConfiguration    OperationTypeEnum = "UPDATE_FLEET_AGENT_CONFIGURATION"
	OperationTypeDeleteJavaInstallation           OperationTypeEnum = "DELETE_JAVA_INSTALLATION"
	OperationTypeCreateJavaInstallation           OperationTypeEnum = "CREATE_JAVA_INSTALLATION"
	OperationTypeCollectJfr                       OperationTypeEnum = "COLLECT_JFR"
	OperationTypeRequestCryptoEventAnalysis       OperationTypeEnum = "REQUEST_CRYPTO_EVENT_ANALYSIS"
	OperationTypeRequestPerformanceTuningAnalysis OperationTypeEnum = "REQUEST_PERFORMANCE_TUNING_ANALYSIS"
	OperationTypeRequestJavaMigrationAnalysis     OperationTypeEnum = "REQUEST_JAVA_MIGRATION_ANALYSIS"
	OperationTypeDeleteJmsReport                  OperationTypeEnum = "DELETE_JMS_REPORT"
	OperationTypeScanJavaServerUsage              OperationTypeEnum = "SCAN_JAVA_SERVER_USAGE"
	OperationTypeScanLibraryUsage                 OperationTypeEnum = "SCAN_LIBRARY_USAGE"
	OperationTypeExportDataCsv                    OperationTypeEnum = "EXPORT_DATA_CSV"
	OperationTypeCreateDrsFile                    OperationTypeEnum = "CREATE_DRS_FILE"
	OperationTypeUpdateDrsFile                    OperationTypeEnum = "UPDATE_DRS_FILE"
	OperationTypeDeleteDrsFile                    OperationTypeEnum = "DELETE_DRS_FILE"
	OperationTypeEnableDrs                        OperationTypeEnum = "ENABLE_DRS"
	OperationTypeDisableDrs                       OperationTypeEnum = "DISABLE_DRS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_FLEET":                        OperationTypeCreateFleet,
	"DELETE_FLEET":                        OperationTypeDeleteFleet,
	"MOVE_FLEET":                          OperationTypeMoveFleet,
	"UPDATE_FLEET":                        OperationTypeUpdateFleet,
	"UPDATE_FLEET_AGENT_CONFIGURATION":    OperationTypeUpdateFleetAgentConfiguration,
	"DELETE_JAVA_INSTALLATION":            OperationTypeDeleteJavaInstallation,
	"CREATE_JAVA_INSTALLATION":            OperationTypeCreateJavaInstallation,
	"COLLECT_JFR":                         OperationTypeCollectJfr,
	"REQUEST_CRYPTO_EVENT_ANALYSIS":       OperationTypeRequestCryptoEventAnalysis,
	"REQUEST_PERFORMANCE_TUNING_ANALYSIS": OperationTypeRequestPerformanceTuningAnalysis,
	"REQUEST_JAVA_MIGRATION_ANALYSIS":     OperationTypeRequestJavaMigrationAnalysis,
	"DELETE_JMS_REPORT":                   OperationTypeDeleteJmsReport,
	"SCAN_JAVA_SERVER_USAGE":              OperationTypeScanJavaServerUsage,
	"SCAN_LIBRARY_USAGE":                  OperationTypeScanLibraryUsage,
	"EXPORT_DATA_CSV":                     OperationTypeExportDataCsv,
	"CREATE_DRS_FILE":                     OperationTypeCreateDrsFile,
	"UPDATE_DRS_FILE":                     OperationTypeUpdateDrsFile,
	"DELETE_DRS_FILE":                     OperationTypeDeleteDrsFile,
	"ENABLE_DRS":                          OperationTypeEnableDrs,
	"DISABLE_DRS":                         OperationTypeDisableDrs,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_fleet":                        OperationTypeCreateFleet,
	"delete_fleet":                        OperationTypeDeleteFleet,
	"move_fleet":                          OperationTypeMoveFleet,
	"update_fleet":                        OperationTypeUpdateFleet,
	"update_fleet_agent_configuration":    OperationTypeUpdateFleetAgentConfiguration,
	"delete_java_installation":            OperationTypeDeleteJavaInstallation,
	"create_java_installation":            OperationTypeCreateJavaInstallation,
	"collect_jfr":                         OperationTypeCollectJfr,
	"request_crypto_event_analysis":       OperationTypeRequestCryptoEventAnalysis,
	"request_performance_tuning_analysis": OperationTypeRequestPerformanceTuningAnalysis,
	"request_java_migration_analysis":     OperationTypeRequestJavaMigrationAnalysis,
	"delete_jms_report":                   OperationTypeDeleteJmsReport,
	"scan_java_server_usage":              OperationTypeScanJavaServerUsage,
	"scan_library_usage":                  OperationTypeScanLibraryUsage,
	"export_data_csv":                     OperationTypeExportDataCsv,
	"create_drs_file":                     OperationTypeCreateDrsFile,
	"update_drs_file":                     OperationTypeUpdateDrsFile,
	"delete_drs_file":                     OperationTypeDeleteDrsFile,
	"enable_drs":                          OperationTypeEnableDrs,
	"disable_drs":                         OperationTypeDisableDrs,
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
		"CREATE_FLEET",
		"DELETE_FLEET",
		"MOVE_FLEET",
		"UPDATE_FLEET",
		"UPDATE_FLEET_AGENT_CONFIGURATION",
		"DELETE_JAVA_INSTALLATION",
		"CREATE_JAVA_INSTALLATION",
		"COLLECT_JFR",
		"REQUEST_CRYPTO_EVENT_ANALYSIS",
		"REQUEST_PERFORMANCE_TUNING_ANALYSIS",
		"REQUEST_JAVA_MIGRATION_ANALYSIS",
		"DELETE_JMS_REPORT",
		"SCAN_JAVA_SERVER_USAGE",
		"SCAN_LIBRARY_USAGE",
		"EXPORT_DATA_CSV",
		"CREATE_DRS_FILE",
		"UPDATE_DRS_FILE",
		"DELETE_DRS_FILE",
		"ENABLE_DRS",
		"DISABLE_DRS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
