// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MaskingPolicyHealthReportLogSummary A log entry related to the pre-masking health check.
type MaskingPolicyHealthReportLogSummary struct {

	// The log entry type.
	MessageType MaskingPolicyHealthReportLogSummaryMessageTypeEnum `mandatory:"true" json:"messageType"`

	// The date and time the log entry was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// A human-readable log entry.
	Message *string `mandatory:"true" json:"message"`

	// A human-readable description for the log entry.
	Description *string `mandatory:"true" json:"description"`

	// A human-readable log entry to remedy any error or warnings in the masking policy.
	Remediation *string `mandatory:"false" json:"remediation"`

	// An enum type entry for each health check in the masking policy. Each enum describes a type of health check.
	// INVALID_OBJECT_CHECK checks if there exist any invalid objects in the masking tables.
	// PRIVILEGE_CHECK checks if the masking user has sufficient privilege to run masking.
	// TABLESPACE_CHECK checks if the user has sufficient default and TEMP tablespace.
	// DATABASE_OR_SYSTEM_TRIGGERS_CHECK checks if there exist any database/system triggers available.
	// STATE_STATS_CHECK checks if all the statistics of the masking table is upto date or not.
	// OLS_POLICY_CHECK and VPD_POLICY_CHECK checks if the masking tables has Oracle Label Security (OLS) or Virtual Private Database (VPD) policies enabled.
	// ACTIVE_MASK_JOB_CHECK checks if there is any active masking job running on the target database.
	// DETERMINISTIC_ENCRYPTION_FORMAT_CHECK checks if any masking column has deterministic encryption masking format.
	// COLUMN_EXIST_CHECK checks if the masking columns are available in the target database.
	HealthCheckType MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum `mandatory:"false" json:"healthCheckType,omitempty"`
}

func (m MaskingPolicyHealthReportLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingPolicyHealthReportLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum(string(m.MessageType)); !ok && m.MessageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageType: %s. Supported values are: %s.", m.MessageType, strings.Join(GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum(string(m.HealthCheckType)); !ok && m.HealthCheckType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HealthCheckType: %s. Supported values are: %s.", m.HealthCheckType, strings.Join(GetMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingPolicyHealthReportLogSummaryMessageTypeEnum Enum with underlying type: string
type MaskingPolicyHealthReportLogSummaryMessageTypeEnum string

// Set of constants representing the allowable values for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
const (
	MaskingPolicyHealthReportLogSummaryMessageTypePass    MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "PASS"
	MaskingPolicyHealthReportLogSummaryMessageTypeWarning MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "WARNING"
	MaskingPolicyHealthReportLogSummaryMessageTypeError   MaskingPolicyHealthReportLogSummaryMessageTypeEnum = "ERROR"
)

var mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum = map[string]MaskingPolicyHealthReportLogSummaryMessageTypeEnum{
	"PASS":    MaskingPolicyHealthReportLogSummaryMessageTypePass,
	"WARNING": MaskingPolicyHealthReportLogSummaryMessageTypeWarning,
	"ERROR":   MaskingPolicyHealthReportLogSummaryMessageTypeError,
}

var mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase = map[string]MaskingPolicyHealthReportLogSummaryMessageTypeEnum{
	"pass":    MaskingPolicyHealthReportLogSummaryMessageTypePass,
	"warning": MaskingPolicyHealthReportLogSummaryMessageTypeWarning,
	"error":   MaskingPolicyHealthReportLogSummaryMessageTypeError,
}

// GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumValues Enumerates the set of values for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
func GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumValues() []MaskingPolicyHealthReportLogSummaryMessageTypeEnum {
	values := make([]MaskingPolicyHealthReportLogSummaryMessageTypeEnum, 0)
	for _, v := range mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues Enumerates the set of values in String for MaskingPolicyHealthReportLogSummaryMessageTypeEnum
func GetMaskingPolicyHealthReportLogSummaryMessageTypeEnumStringValues() []string {
	return []string{
		"PASS",
		"WARNING",
		"ERROR",
	}
}

// GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyHealthReportLogSummaryMessageTypeEnum(val string) (MaskingPolicyHealthReportLogSummaryMessageTypeEnum, bool) {
	enum, ok := mappingMaskingPolicyHealthReportLogSummaryMessageTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum Enum with underlying type: string
type MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum string

// Set of constants representing the allowable values for MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum
const (
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck                 MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "INVALID_OBJECT_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck                     MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "PRIVILEGE_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck                    MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TABLESPACE_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck      MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "DATABASE_OR_SYSTEM_TRIGGERS_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck                    MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "STATE_STATS_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck                     MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "OLS_POLICY_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck                     MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "VPD_POLICY_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeActiveMaskJobCheck                 MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "ACTIVE_MASK_JOB_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck              MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "TARGET_VALIDATION_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeDeterministicEncryptionFormatCheck MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "DETERMINISTIC_ENCRYPTION_FORMAT_CHECK"
	MaskingPolicyHealthReportLogSummaryHealthCheckTypeColumnExistCheck                   MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = "COLUMN_EXIST_CHECK"
)

var mappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum = map[string]MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum{
	"INVALID_OBJECT_CHECK":                  MaskingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck,
	"PRIVILEGE_CHECK":                       MaskingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck,
	"TABLESPACE_CHECK":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck,
	"DATABASE_OR_SYSTEM_TRIGGERS_CHECK":     MaskingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck,
	"STATE_STATS_CHECK":                     MaskingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck,
	"OLS_POLICY_CHECK":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck,
	"VPD_POLICY_CHECK":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck,
	"ACTIVE_MASK_JOB_CHECK":                 MaskingPolicyHealthReportLogSummaryHealthCheckTypeActiveMaskJobCheck,
	"TARGET_VALIDATION_CHECK":               MaskingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck,
	"DETERMINISTIC_ENCRYPTION_FORMAT_CHECK": MaskingPolicyHealthReportLogSummaryHealthCheckTypeDeterministicEncryptionFormatCheck,
	"COLUMN_EXIST_CHECK":                    MaskingPolicyHealthReportLogSummaryHealthCheckTypeColumnExistCheck,
}

var mappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumLowerCase = map[string]MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum{
	"invalid_object_check":                  MaskingPolicyHealthReportLogSummaryHealthCheckTypeInvalidObjectCheck,
	"privilege_check":                       MaskingPolicyHealthReportLogSummaryHealthCheckTypePrivilegeCheck,
	"tablespace_check":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeTablespaceCheck,
	"database_or_system_triggers_check":     MaskingPolicyHealthReportLogSummaryHealthCheckTypeDatabaseOrSystemTriggersCheck,
	"state_stats_check":                     MaskingPolicyHealthReportLogSummaryHealthCheckTypeStateStatsCheck,
	"ols_policy_check":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeOlsPolicyCheck,
	"vpd_policy_check":                      MaskingPolicyHealthReportLogSummaryHealthCheckTypeVpdPolicyCheck,
	"active_mask_job_check":                 MaskingPolicyHealthReportLogSummaryHealthCheckTypeActiveMaskJobCheck,
	"target_validation_check":               MaskingPolicyHealthReportLogSummaryHealthCheckTypeTargetValidationCheck,
	"deterministic_encryption_format_check": MaskingPolicyHealthReportLogSummaryHealthCheckTypeDeterministicEncryptionFormatCheck,
	"column_exist_check":                    MaskingPolicyHealthReportLogSummaryHealthCheckTypeColumnExistCheck,
}

// GetMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumValues Enumerates the set of values for MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum
func GetMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumValues() []MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum {
	values := make([]MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum, 0)
	for _, v := range mappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues Enumerates the set of values in String for MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum
func GetMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumStringValues() []string {
	return []string{
		"INVALID_OBJECT_CHECK",
		"PRIVILEGE_CHECK",
		"TABLESPACE_CHECK",
		"DATABASE_OR_SYSTEM_TRIGGERS_CHECK",
		"STATE_STATS_CHECK",
		"OLS_POLICY_CHECK",
		"VPD_POLICY_CHECK",
		"ACTIVE_MASK_JOB_CHECK",
		"TARGET_VALIDATION_CHECK",
		"DETERMINISTIC_ENCRYPTION_FORMAT_CHECK",
		"COLUMN_EXIST_CHECK",
	}
}

// GetMappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum(val string) (MaskingPolicyHealthReportLogSummaryHealthCheckTypeEnum, bool) {
	enum, ok := mappingMaskingPolicyHealthReportLogSummaryHealthCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
