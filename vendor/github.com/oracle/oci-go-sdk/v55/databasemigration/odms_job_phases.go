// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// OdmsJobPhasesEnum Enum with underlying type: string
type OdmsJobPhasesEnum string

// Set of constants representing the allowable values for OdmsJobPhasesEnum
const (
	OdmsJobPhasesOdmsValidateTgt                 OdmsJobPhasesEnum = "ODMS_VALIDATE_TGT"
	OdmsJobPhasesOdmsValidateSrc                 OdmsJobPhasesEnum = "ODMS_VALIDATE_SRC"
	OdmsJobPhasesOdmsValidatePremigrationAdvisor OdmsJobPhasesEnum = "ODMS_VALIDATE_PREMIGRATION_ADVISOR"
	OdmsJobPhasesOdmsValidateGgHub               OdmsJobPhasesEnum = "ODMS_VALIDATE_GG_HUB"
	OdmsJobPhasesOdmsValidateDatapumpSettings    OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS"
	OdmsJobPhasesOdmsValidateDatapumpSettingsSrc OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS_SRC"
	OdmsJobPhasesOdmsValidateDatapumpSettingsTgt OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SETTINGS_TGT"
	OdmsJobPhasesOdmsValidateDatapumpSrc         OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_SRC"
	OdmsJobPhasesOdmsValidateDatapumpEstimateSrc OdmsJobPhasesEnum = "ODMS_VALIDATE_DATAPUMP_ESTIMATE_SRC"
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

var mappingOdmsJobPhases = map[string]OdmsJobPhasesEnum{
	"ODMS_VALIDATE_TGT":                   OdmsJobPhasesOdmsValidateTgt,
	"ODMS_VALIDATE_SRC":                   OdmsJobPhasesOdmsValidateSrc,
	"ODMS_VALIDATE_PREMIGRATION_ADVISOR":  OdmsJobPhasesOdmsValidatePremigrationAdvisor,
	"ODMS_VALIDATE_GG_HUB":                OdmsJobPhasesOdmsValidateGgHub,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS":     OdmsJobPhasesOdmsValidateDatapumpSettings,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS_SRC": OdmsJobPhasesOdmsValidateDatapumpSettingsSrc,
	"ODMS_VALIDATE_DATAPUMP_SETTINGS_TGT": OdmsJobPhasesOdmsValidateDatapumpSettingsTgt,
	"ODMS_VALIDATE_DATAPUMP_SRC":          OdmsJobPhasesOdmsValidateDatapumpSrc,
	"ODMS_VALIDATE_DATAPUMP_ESTIMATE_SRC": OdmsJobPhasesOdmsValidateDatapumpEstimateSrc,
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

// GetOdmsJobPhasesEnumValues Enumerates the set of values for OdmsJobPhasesEnum
func GetOdmsJobPhasesEnumValues() []OdmsJobPhasesEnum {
	values := make([]OdmsJobPhasesEnum, 0)
	for _, v := range mappingOdmsJobPhases {
		values = append(values, v)
	}
	return values
}
