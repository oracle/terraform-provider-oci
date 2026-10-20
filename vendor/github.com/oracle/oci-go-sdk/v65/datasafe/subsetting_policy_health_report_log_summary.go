// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsettingPolicyHealthReportLogSummary A log entry related to the pre-subsetting health check.
type SubsettingPolicyHealthReportLogSummary struct {

	// The log entry type.
	MessageType SubsettingPolicyHealthReportLogSummaryMessageTypeEnum `mandatory:"true" json:"messageType"`

	// The date and time the log entry was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// A human-readable log entry.
	Message *string `mandatory:"true" json:"message"`

	// A human-readable description for the log entry.
	Description *string `mandatory:"true" json:"description"`

	// A human-readable log entry to remedy any error or warnings in the subsetting policy.
	Remediation *string `mandatory:"false" json:"remediation"`

	// An enum type entry for each health check in the subsetting policy. Each enum describes a type of health check.
	// INVALID_OBJECT_CHECK checks if there exist any invalid objects in the subsetting tables.
	// PRIVILEGE_CHECK checks if the subsetting user has sufficient privilege to run subsetting.
	// TABLESPACE_CHECK checks if the user has sufficient default and TEMP tablespace. Also verifies that the specified tablespace by the user is valid, if user has provided one
	// DATABASE_OR_SYSTEM_TRIGGERS_CHECK checks if there exist any database/system triggers available.
	// UNDO_TABLESPACE_CHECK checks if for all the instances of undo tablespace the AUTOEXTEND feature is enabled.
	// If it's not enabled, it further checks if the undo tablespace has any space remaining.
	// STATE_STATS_CHECK checks if all the statistics of the subsetting table is upto date or not.
	// OLS_POLICY_CHECK , VPD_POLICY_CHECK and REDACTION_POLICY_CHECK checks if the subsetting tables has Oracle Label Security (OLS) or Virtual Private Database (VPD) or Redaction policies enabled.
	// DV_ENABLE_CHECK checks if database has Database Vault(DV) enabled
	// ACTIVE_JOB_CHECK checks if there is any active subsetting job running on the target database.
	// TABLE_EXIST_CHECK checks if the subsetting tables are available in the target database.
	// TIME_TRAVEL_CHECK checks if the subsetting tables have Time Travel enabled.
	// SYSTEM_OBJECTS_CHECK checks if the subsetting tables have dependent objects present in SYS schema.
	// INVALID_PACKAGE_CHECK checks if any of the required packages are in invalid state.
	// AUDIT_POLICY_CHECK checks if the subsetting tables have Audit policies enabled.
	// VALID_RULES_CHECK if the subsetting rules on the tables are valid.
	// CUSTOM_TABLESPACE_CHECK runs only when a custom tablespace is specified for the subsetting job. It checks if each policy schema can allocate space in that tablespace.
	HealthCheckType SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum `mandatory:"false" json:"healthCheckType,omitempty"`
}

func (m SubsettingPolicyHealthReportLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingPolicyHealthReportLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnum(string(m.MessageType)); !ok && m.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", m.MessageType, strings.Join(GetSubsettingPolicyHealthReportLogSummaryMessageTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum(string(m.HealthCheckType)); !ok && m.HealthCheckType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HealthCheckType: %s. Supported values are: %s.", m.HealthCheckType, strings.Join(GetSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingPolicyHealthReportLogSummaryMessageTypeEnum Enum with underlying type: string
type SubsettingPolicyHealthReportLogSummaryMessageTypeEnum string

// Set of constants representing the allowable values for SubsettingPolicyHealthReportLogSummaryMessageTypeEnum
const (
	SubsettingPolicyHealthReportLogSummaryMessageTypePass    SubsettingPolicyHealthReportLogSummaryMessageTypeEnum = "PASS"
	SubsettingPolicyHealthReportLogSummaryMessageTypeWarning SubsettingPolicyHealthReportLogSummaryMessageTypeEnum = "WARNING"
	SubsettingPolicyHealthReportLogSummaryMessageTypeError   SubsettingPolicyHealthReportLogSummaryMessageTypeEnum = "ERROR"
)

var mappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnum = map[string]SubsettingPolicyHealthReportLogSummaryMessageTypeEnum{
	"PASS":    SubsettingPolicyHealthReportLogSummaryMessageTypePass,
	"WARNING": SubsettingPolicyHealthReportLogSummaryMessageTypeWarning,
	"ERROR":   SubsettingPolicyHealthReportLogSummaryMessageTypeError,
}

var mappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase = map[string]SubsettingPolicyHealthReportLogSummaryMessageTypeEnum{
	"pass":    SubsettingPolicyHealthReportLogSummaryMessageTypePass,
	"warning": SubsettingPolicyHealthReportLogSummaryMessageTypeWarning,
	"error":   SubsettingPolicyHealthReportLogSummaryMessageTypeError,
}

// GetSubsettingPolicyHealthReportLogSummaryMessageTypeEnumValues Enumerates the set of values for SubsettingPolicyHealthReportLogSummaryMessageTypeEnum
func GetSubsettingPolicyHealthReportLogSummaryMessageTypeEnumValues() []SubsettingPolicyHealthReportLogSummaryMessageTypeEnum {
	values := make([]SubsettingPolicyHealthReportLogSummaryMessageTypeEnum, 0)
	for _, v := range mappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyHealthReportLogSummaryMessageTypeEnumStringValues Enumerates the set of values in String for SubsettingPolicyHealthReportLogSummaryMessageTypeEnum
func GetSubsettingPolicyHealthReportLogSummaryMessageTypeEnumStringValues() []string {
	return []string{
		"PASS",
		"WARNING",
		"ERROR",
	}
}

// GetMappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnum(val string) (SubsettingPolicyHealthReportLogSummaryMessageTypeEnum, bool) {
	enum, ok := mappingSubsettingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum Enum with underlying type: string
type SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum string

// Set of constants representing the allowable values for SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum
const (
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck            SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "INVALID_OBJECT_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "PRIVILEGE_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck               SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TABLESPACE_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "DATABASE_OR_SYSTEM_TRIGGERS_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeUndoTablespaceCheck           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "UNDO_TABLESPACE_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck               SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "STATE_STATS_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "OLS_POLICY_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "VPD_POLICY_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDvEnableCheck                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "DV_ENABLE_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeRedactionPolicyCheck          SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "REDACTION_POLICY_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeActiveJobCheck                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "ACTIVE_JOB_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck         SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TARGET_VALIDATION_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTableExistCheck               SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TABLE_EXIST_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTimeTravelCheck               SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TIME_TRAVEL_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeSystemObjectsCheck            SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "SYSTEM_OBJECTS_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidPackageCheck           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "INVALID_PACKAGE_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeAuditPolicyCheck              SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "AUDIT_POLICY_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeValidRulesCheck               SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "VALID_RULES_CHECK"
	SubsettingPolicyHealthReportLogSummaryHealthCheckTypeCustomTablespaceCheck         SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "CUSTOM_TABLESPACE_CHECK"
)

var mappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum = map[string]SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum{
	"INVALID_OBJECT_CHECK":              SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck,
	"PRIVILEGE_CHECK":                   SubsettingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck,
	"TABLESPACE_CHECK":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck,
	"DATABASE_OR_SYSTEM_TRIGGERS_CHECK": SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck,
	"UNDO_TABLESPACE_CHECK":             SubsettingPolicyHealthReportLogSummaryHealthCheckTypeUndoTablespaceCheck,
	"STATE_STATS_CHECK":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck,
	"OLS_POLICY_CHECK":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck,
	"VPD_POLICY_CHECK":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck,
	"DV_ENABLE_CHECK":                   SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDvEnableCheck,
	"REDACTION_POLICY_CHECK":            SubsettingPolicyHealthReportLogSummaryHealthCheckTypeRedactionPolicyCheck,
	"ACTIVE_JOB_CHECK":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeActiveJobCheck,
	"TARGET_VALIDATION_CHECK":           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck,
	"TABLE_EXIST_CHECK":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTableExistCheck,
	"TIME_TRAVEL_CHECK":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTimeTravelCheck,
	"SYSTEM_OBJECTS_CHECK":              SubsettingPolicyHealthReportLogSummaryHealthCheckTypeSystemObjectsCheck,
	"INVALID_PACKAGE_CHECK":             SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidPackageCheck,
	"AUDIT_POLICY_CHECK":                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeAuditPolicyCheck,
	"VALID_RULES_CHECK":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeValidRulesCheck,
	"CUSTOM_TABLESPACE_CHECK":           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeCustomTablespaceCheck,
}

var mappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumLowerCase = map[string]SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum{
	"invalid_object_check":              SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck,
	"privilege_check":                   SubsettingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck,
	"tablespace_check":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck,
	"database_or_system_triggers_check": SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck,
	"undo_tablespace_check":             SubsettingPolicyHealthReportLogSummaryHealthCheckTypeUndoTablespaceCheck,
	"state_stats_check":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck,
	"ols_policy_check":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck,
	"vpd_policy_check":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck,
	"dv_enable_check":                   SubsettingPolicyHealthReportLogSummaryHealthCheckTypeDvEnableCheck,
	"redaction_policy_check":            SubsettingPolicyHealthReportLogSummaryHealthCheckTypeRedactionPolicyCheck,
	"active_job_check":                  SubsettingPolicyHealthReportLogSummaryHealthCheckTypeActiveJobCheck,
	"target_validation_check":           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck,
	"table_exist_check":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTableExistCheck,
	"time_travel_check":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeTimeTravelCheck,
	"system_objects_check":              SubsettingPolicyHealthReportLogSummaryHealthCheckTypeSystemObjectsCheck,
	"invalid_package_check":             SubsettingPolicyHealthReportLogSummaryHealthCheckTypeInvalidPackageCheck,
	"audit_policy_check":                SubsettingPolicyHealthReportLogSummaryHealthCheckTypeAuditPolicyCheck,
	"valid_rules_check":                 SubsettingPolicyHealthReportLogSummaryHealthCheckTypeValidRulesCheck,
	"custom_tablespace_check":           SubsettingPolicyHealthReportLogSummaryHealthCheckTypeCustomTablespaceCheck,
}

// GetSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumValues Enumerates the set of values for SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum
func GetSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumValues() []SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum {
	values := make([]SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum, 0)
	for _, v := range mappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues Enumerates the set of values in String for SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum
func GetSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues() []string {
	return []string{
		"INVALID_OBJECT_CHECK",
		"PRIVILEGE_CHECK",
		"TABLESPACE_CHECK",
		"DATABASE_OR_SYSTEM_TRIGGERS_CHECK",
		"UNDO_TABLESPACE_CHECK",
		"STATE_STATS_CHECK",
		"OLS_POLICY_CHECK",
		"VPD_POLICY_CHECK",
		"DV_ENABLE_CHECK",
		"REDACTION_POLICY_CHECK",
		"ACTIVE_JOB_CHECK",
		"TARGET_VALIDATION_CHECK",
		"TABLE_EXIST_CHECK",
		"TIME_TRAVEL_CHECK",
		"SYSTEM_OBJECTS_CHECK",
		"INVALID_PACKAGE_CHECK",
		"AUDIT_POLICY_CHECK",
		"VALID_RULES_CHECK",
		"CUSTOM_TABLESPACE_CHECK",
	}
}

// GetMappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum(val string) (SubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnum, bool) {
	enum, ok := mappingSubsettingPolicyHealthReportLogSummaryHealthCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
