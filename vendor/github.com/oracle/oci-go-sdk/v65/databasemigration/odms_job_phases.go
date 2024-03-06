// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// OdmsJobPhasesEnum Enum with underlying type: string
type OdmsJobPhasesEnum string

// Set of constants representing the allowable values for OdmsJobPhasesEnum
const (
	OdmsJobPhasesOdmsValidateTgt                 OdmsJobPhasesEnum = "ODMS_VALIDATE_TGT"
	OdmsJobPhasesOdmsValidateSrc                 OdmsJobPhasesEnum = "ODMS_VALIDATE_SRC"
	OdmsJobPhasesOdmsValidatePremigrationAdvisor OdmsJobPhasesEnum = "ODMS_VALIDATE_PREMIGRATION_ADVISOR"
	OdmsJobPhasesOdmsValidateGgHub               OdmsJobPhasesEnum = "ODMS_VALIDATE_GG_HUB"
	OdmsJobPhasesOdmsValidateGgService           OdmsJobPhasesEnum = "ODMS_VALIDATE_GG_SERVICE"
	OdmsJobPhasesOdmsValidateDatapumpSettings    OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS"
	OdmsJobPhasesOdmsValidateDatapumpSettingsSrc OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS_SRC"
	OdmsJobPhasesOdmsValidateDatapumpSettingsTgt OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS_TGT"
	OdmsJobPhasesOdmsValidateDatapumpSrc         OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SRC"
	OdmsJobPhasesOdmsValidateDatapumpEstimateSrc OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_ESTIMATE_SRC"
	OdmsJobPhasesOdmsInitializeGgs               OdmsJobPhasesEnum = "ODMS_INITIALIZE_GGS"
	OdmsJobPhasesOdmsValidate                    OdmsJobPhasesEnum = "ODMS_VALIDATE"
	OdmsJobPhasesOdmsPrepare                     OdmsJobPhasesEnum = "ODMS_PREPARE"
	OdmsJobPhasesOdmsInitialLoadExport           OdmsJobPhasesEnum = "ODMS_INITIAL_LOAD_EXPORT"
	OdmsJobPhasesOdmsDataUpload                  OdmsJobPhasesEnum = "ODMS_DATA_UPLOAD"
	OdmsJobPhasesOdmsInitialLoadImport           OdmsJobPhasesEnum = "ODMS_INITIAL_LOAD_IMPORT"
	OdmsJobPhasesOdmsPostInitialLoad             OdmsJobPhasesEnum = "ODMS_POST_INITIAL_LOAD"
	OdmsJobPhasesOdmsPrepareReplicationTarget    OdmsJobPhasesEnum = "ODMS_PREPARE_REPLICATION_TARGET"
	OdmsJobPhasesOdmsMonitorReplicationLag       OdmsJobPhasesEnum = "ODMS_MONITOR_REPLICATION_LAG"
	OdmsJobPhasesOdmsSwitchover                  OdmsJobPhasesEnum = "ODMS_SWITCHOVER"
	OdmsJobPhasesOdmsCleanup                     OdmsJobPhasesEnum = "ODMS_CLEANUP"
)

var mappingOdmsJobPhasesEnum = map[string]OdmsJobPhasesEnum{
	"ODMS_VALIDATE_TGT":                   OdmsJobPhasesOdmsValidateTgt,
	"ODMS_VALIDATE_SRC":                   OdmsJobPhasesOdmsValidateSrc,
	"ODMS_VALIDATE_PREMIGRATION_ADVISOR":  OdmsJobPhasesOdmsValidatePremigrationAdvisor,
	"ODMS_VALIDATE_GG_HUB":                OdmsJobPhasesOdmsValidateGgHub,
	"ODMS_VALIDATE_GG_SERVICE":            OdmsJobPhasesOdmsValidateGgService,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS":     OdmsJobPhasesOdmsValidateDatapumpSettings,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS_SRC": OdmsJobPhasesOdmsValidateDatapumpSettingsSrc,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS_TGT": OdmsJobPhasesOdmsValidateDatapumpSettingsTgt,
	"ODMS_VALIDATE_DATAPUMP_SRC":          OdmsJobPhasesOdmsValidateDatapumpSrc,
	"ODMS_VALIDATE_DATAPUMP_ESTIMATE_SRC": OdmsJobPhasesOdmsValidateDatapumpEstimateSrc,
	"ODMS_INITIALIZE_GGS":                 OdmsJobPhasesOdmsInitializeGgs,
	"ODMS_VALIDATE":                       OdmsJobPhasesOdmsValidate,
	"ODMS_PREPARE":                        OdmsJobPhasesOdmsPrepare,
	"ODMS_INITIAL_LOAD_EXPORT":            OdmsJobPhasesOdmsInitialLoadExport,
	"ODMS_DATA_UPLOAD":                    OdmsJobPhasesOdmsDataUpload,
	"ODMS_INITIAL_LOAD_IMPORT":            OdmsJobPhasesOdmsInitialLoadImport,
	"ODMS_POST_INITIAL_LOAD":              OdmsJobPhasesOdmsPostInitialLoad,
	"ODMS_PREPARE_REPLICATION_TARGET":     OdmsJobPhasesOdmsPrepareReplicationTarget,
	"ODMS_MONITOR_REPLICATION_LAG":        OdmsJobPhasesOdmsMonitorReplicationLag,
	"ODMS_SWITCHOVER":                     OdmsJobPhasesOdmsSwitchover,
	"ODMS_CLEANUP":                        OdmsJobPhasesOdmsCleanup,
}

var mappingOdmsJobPhasesEnumLowerCase = map[string]OdmsJobPhasesEnum{
	"odms_validate_tgt":                   OdmsJobPhasesOdmsValidateTgt,
	"odms_validate_src":                   OdmsJobPhasesOdmsValidateSrc,
	"odms_validate_premigration_advisor":  OdmsJobPhasesOdmsValidatePremigrationAdvisor,
	"odms_validate_gg_hub":                OdmsJobPhasesOdmsValidateGgHub,
	"odms_validate_gg_service":            OdmsJobPhasesOdmsValidateGgService,
	"odms_validate_datapump_settings":     OdmsJobPhasesOdmsValidateDatapumpSettings,
	"odms_validate_datapump_settings_src": OdmsJobPhasesOdmsValidateDatapumpSettingsSrc,
	"odms_validate_datapump_settings_tgt": OdmsJobPhasesOdmsValidateDatapumpSettingsTgt,
	"odms_validate_datapump_src":          OdmsJobPhasesOdmsValidateDatapumpSrc,
	"odms_validate_datapump_estimate_src": OdmsJobPhasesOdmsValidateDatapumpEstimateSrc,
	"odms_initialize_ggs":                 OdmsJobPhasesOdmsInitializeGgs,
	"odms_validate":                       OdmsJobPhasesOdmsValidate,
	"odms_prepare":                        OdmsJobPhasesOdmsPrepare,
	"odms_initial_load_export":            OdmsJobPhasesOdmsInitialLoadExport,
	"odms_data_upload":                    OdmsJobPhasesOdmsDataUpload,
	"odms_initial_load_import":            OdmsJobPhasesOdmsInitialLoadImport,
	"odms_post_initial_load":              OdmsJobPhasesOdmsPostInitialLoad,
	"odms_prepare_replication_target":     OdmsJobPhasesOdmsPrepareReplicationTarget,
	"odms_monitor_replication_lag":        OdmsJobPhasesOdmsMonitorReplicationLag,
	"odms_switchover":                     OdmsJobPhasesOdmsSwitchover,
	"odms_cleanup":                        OdmsJobPhasesOdmsCleanup,
}

// GetOdmsJobPhasesEnumValues Enumerates the set of values for OdmsJobPhasesEnum
func GetOdmsJobPhasesEnumValues() []OdmsJobPhasesEnum {
	values := make([]OdmsJobPhasesEnum, 0)
	for _, v := range mappingOdmsJobPhasesEnum {
		values = append(values, v)
	}
	return values
}

// GetOdmsJobPhasesEnumStringValues Enumerates the set of values in String for OdmsJobPhasesEnum
func GetOdmsJobPhasesEnumStringValues() []string {
	return []string{
		"ODMS_VALIDATE_TGT",
		"ODMS_VALIDATE_SRC",
		"ODMS_VALIDATE_PREMIGRATION_ADVISOR",
		"ODMS_VALIDATE_GG_HUB",
		"ODMS_VALIDATE_GG_SERVICE",
		"ODMS_VALIDATE_DATAPUMP_SETTINGS",
		"ODMS_VALIDATE_DATAPUMP_SETTINGS_SRC",
		"ODMS_VALIDATE_DATAPUMP_SETTINGS_TGT",
		"ODMS_VALIDATE_DATAPUMP_SRC",
		"ODMS_VALIDATE_DATAPUMP_ESTIMATE_SRC",
		"ODMS_INITIALIZE_GGS",
		"ODMS_VALIDATE",
		"ODMS_PREPARE",
		"ODMS_INITIAL_LOAD_EXPORT",
		"ODMS_DATA_UPLOAD",
		"ODMS_INITIAL_LOAD_IMPORT",
		"ODMS_POST_INITIAL_LOAD",
		"ODMS_PREPARE_REPLICATION_TARGET",
		"ODMS_MONITOR_REPLICATION_LAG",
		"ODMS_SWITCHOVER",
		"ODMS_CLEANUP",
	}
}

// GetMappingOdmsJobPhasesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdmsJobPhasesEnum(val string) (OdmsJobPhasesEnum, bool) {
	enum, ok := mappingOdmsJobPhasesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
