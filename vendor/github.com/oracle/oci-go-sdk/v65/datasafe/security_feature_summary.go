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

// SecurityFeatureSummary The details of database security feature usage available on a given compartment.
type SecurityFeatureSummary struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID of the assessment that generates this security feature usage result.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// The usage of security feature - Unified Audit.
	UnifiedAudit SecurityFeatureSummaryUnifiedAuditEnum `mandatory:"true" json:"unifiedAudit"`

	// The usage of security feature - Fine Grained Audit.
	FineGrainedAudit SecurityFeatureSummaryFineGrainedAuditEnum `mandatory:"true" json:"fineGrainedAudit"`

	// The usage of security feature - Traditional Audit.
	TraditionalAudit SecurityFeatureSummaryTraditionalAuditEnum `mandatory:"true" json:"traditionalAudit"`

	// The usage of security feature - Database Vault.
	DatabaseVault SecurityFeatureSummaryDatabaseVaultEnum `mandatory:"true" json:"databaseVault"`

	// The usage of security feature - Privilege Analysis.
	PrivilegeAnalysis SecurityFeatureSummaryPrivilegeAnalysisEnum `mandatory:"true" json:"privilegeAnalysis"`

	// The usage of security feature - Tablespace Encryption.
	TablespaceEncryption SecurityFeatureSummaryTablespaceEncryptionEnum `mandatory:"true" json:"tablespaceEncryption"`

	// The usage of security feature - Column Encryption.
	ColumnEncryption SecurityFeatureSummaryColumnEncryptionEnum `mandatory:"true" json:"columnEncryption"`

	// The usage of security feature - Network Encryption.
	NetworkEncryption SecurityFeatureSummaryNetworkEncryptionEnum `mandatory:"true" json:"networkEncryption"`

	// The usage of security feature - Password Authentication.
	PasswordAuthentication SecurityFeatureSummaryPasswordAuthenticationEnum `mandatory:"true" json:"passwordAuthentication"`

	// The usage of security feature - Global Authentication.
	GlobalAuthentication SecurityFeatureSummaryGlobalAuthenticationEnum `mandatory:"true" json:"globalAuthentication"`

	// The usage of security feature - External Authentication.
	ExternalAuthentication SecurityFeatureSummaryExternalAuthenticationEnum `mandatory:"true" json:"externalAuthentication"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SecurityFeatureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityFeatureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityFeatureSummaryUnifiedAuditEnum(string(m.UnifiedAudit)); !ok && m.UnifiedAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnifiedAudit: %s. Supported values are: %s.", m.UnifiedAudit, strings.Join(GetSecurityFeatureSummaryUnifiedAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryFineGrainedAuditEnum(string(m.FineGrainedAudit)); !ok && m.FineGrainedAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FineGrainedAudit: %s. Supported values are: %s.", m.FineGrainedAudit, strings.Join(GetSecurityFeatureSummaryFineGrainedAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryTraditionalAuditEnum(string(m.TraditionalAudit)); !ok && m.TraditionalAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TraditionalAudit: %s. Supported values are: %s.", m.TraditionalAudit, strings.Join(GetSecurityFeatureSummaryTraditionalAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryDatabaseVaultEnum(string(m.DatabaseVault)); !ok && m.DatabaseVault != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseVault: %s. Supported values are: %s.", m.DatabaseVault, strings.Join(GetSecurityFeatureSummaryDatabaseVaultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryPrivilegeAnalysisEnum(string(m.PrivilegeAnalysis)); !ok && m.PrivilegeAnalysis != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrivilegeAnalysis: %s. Supported values are: %s.", m.PrivilegeAnalysis, strings.Join(GetSecurityFeatureSummaryPrivilegeAnalysisEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryTablespaceEncryptionEnum(string(m.TablespaceEncryption)); !ok && m.TablespaceEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TablespaceEncryption: %s. Supported values are: %s.", m.TablespaceEncryption, strings.Join(GetSecurityFeatureSummaryTablespaceEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryColumnEncryptionEnum(string(m.ColumnEncryption)); !ok && m.ColumnEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ColumnEncryption: %s. Supported values are: %s.", m.ColumnEncryption, strings.Join(GetSecurityFeatureSummaryColumnEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryNetworkEncryptionEnum(string(m.NetworkEncryption)); !ok && m.NetworkEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NetworkEncryption: %s. Supported values are: %s.", m.NetworkEncryption, strings.Join(GetSecurityFeatureSummaryNetworkEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryPasswordAuthenticationEnum(string(m.PasswordAuthentication)); !ok && m.PasswordAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PasswordAuthentication: %s. Supported values are: %s.", m.PasswordAuthentication, strings.Join(GetSecurityFeatureSummaryPasswordAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryGlobalAuthenticationEnum(string(m.GlobalAuthentication)); !ok && m.GlobalAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GlobalAuthentication: %s. Supported values are: %s.", m.GlobalAuthentication, strings.Join(GetSecurityFeatureSummaryGlobalAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityFeatureSummaryExternalAuthenticationEnum(string(m.ExternalAuthentication)); !ok && m.ExternalAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExternalAuthentication: %s. Supported values are: %s.", m.ExternalAuthentication, strings.Join(GetSecurityFeatureSummaryExternalAuthenticationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityFeatureSummaryUnifiedAuditEnum Enum with underlying type: string
type SecurityFeatureSummaryUnifiedAuditEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryUnifiedAuditEnum
const (
	SecurityFeatureSummaryUnifiedAuditEnabled  SecurityFeatureSummaryUnifiedAuditEnum = "ENABLED"
	SecurityFeatureSummaryUnifiedAuditDisabled SecurityFeatureSummaryUnifiedAuditEnum = "DISABLED"
	SecurityFeatureSummaryUnifiedAuditNone     SecurityFeatureSummaryUnifiedAuditEnum = "NONE"
)

var mappingSecurityFeatureSummaryUnifiedAuditEnum = map[string]SecurityFeatureSummaryUnifiedAuditEnum{
	"ENABLED":  SecurityFeatureSummaryUnifiedAuditEnabled,
	"DISABLED": SecurityFeatureSummaryUnifiedAuditDisabled,
	"NONE":     SecurityFeatureSummaryUnifiedAuditNone,
}

var mappingSecurityFeatureSummaryUnifiedAuditEnumLowerCase = map[string]SecurityFeatureSummaryUnifiedAuditEnum{
	"enabled":  SecurityFeatureSummaryUnifiedAuditEnabled,
	"disabled": SecurityFeatureSummaryUnifiedAuditDisabled,
	"none":     SecurityFeatureSummaryUnifiedAuditNone,
}

// GetSecurityFeatureSummaryUnifiedAuditEnumValues Enumerates the set of values for SecurityFeatureSummaryUnifiedAuditEnum
func GetSecurityFeatureSummaryUnifiedAuditEnumValues() []SecurityFeatureSummaryUnifiedAuditEnum {
	values := make([]SecurityFeatureSummaryUnifiedAuditEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryUnifiedAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryUnifiedAuditEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryUnifiedAuditEnum
func GetSecurityFeatureSummaryUnifiedAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryUnifiedAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryUnifiedAuditEnum(val string) (SecurityFeatureSummaryUnifiedAuditEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryUnifiedAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryFineGrainedAuditEnum Enum with underlying type: string
type SecurityFeatureSummaryFineGrainedAuditEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryFineGrainedAuditEnum
const (
	SecurityFeatureSummaryFineGrainedAuditEnabled  SecurityFeatureSummaryFineGrainedAuditEnum = "ENABLED"
	SecurityFeatureSummaryFineGrainedAuditDisabled SecurityFeatureSummaryFineGrainedAuditEnum = "DISABLED"
	SecurityFeatureSummaryFineGrainedAuditNone     SecurityFeatureSummaryFineGrainedAuditEnum = "NONE"
)

var mappingSecurityFeatureSummaryFineGrainedAuditEnum = map[string]SecurityFeatureSummaryFineGrainedAuditEnum{
	"ENABLED":  SecurityFeatureSummaryFineGrainedAuditEnabled,
	"DISABLED": SecurityFeatureSummaryFineGrainedAuditDisabled,
	"NONE":     SecurityFeatureSummaryFineGrainedAuditNone,
}

var mappingSecurityFeatureSummaryFineGrainedAuditEnumLowerCase = map[string]SecurityFeatureSummaryFineGrainedAuditEnum{
	"enabled":  SecurityFeatureSummaryFineGrainedAuditEnabled,
	"disabled": SecurityFeatureSummaryFineGrainedAuditDisabled,
	"none":     SecurityFeatureSummaryFineGrainedAuditNone,
}

// GetSecurityFeatureSummaryFineGrainedAuditEnumValues Enumerates the set of values for SecurityFeatureSummaryFineGrainedAuditEnum
func GetSecurityFeatureSummaryFineGrainedAuditEnumValues() []SecurityFeatureSummaryFineGrainedAuditEnum {
	values := make([]SecurityFeatureSummaryFineGrainedAuditEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryFineGrainedAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryFineGrainedAuditEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryFineGrainedAuditEnum
func GetSecurityFeatureSummaryFineGrainedAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryFineGrainedAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryFineGrainedAuditEnum(val string) (SecurityFeatureSummaryFineGrainedAuditEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryFineGrainedAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryTraditionalAuditEnum Enum with underlying type: string
type SecurityFeatureSummaryTraditionalAuditEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryTraditionalAuditEnum
const (
	SecurityFeatureSummaryTraditionalAuditEnabled  SecurityFeatureSummaryTraditionalAuditEnum = "ENABLED"
	SecurityFeatureSummaryTraditionalAuditDisabled SecurityFeatureSummaryTraditionalAuditEnum = "DISABLED"
	SecurityFeatureSummaryTraditionalAuditNone     SecurityFeatureSummaryTraditionalAuditEnum = "NONE"
)

var mappingSecurityFeatureSummaryTraditionalAuditEnum = map[string]SecurityFeatureSummaryTraditionalAuditEnum{
	"ENABLED":  SecurityFeatureSummaryTraditionalAuditEnabled,
	"DISABLED": SecurityFeatureSummaryTraditionalAuditDisabled,
	"NONE":     SecurityFeatureSummaryTraditionalAuditNone,
}

var mappingSecurityFeatureSummaryTraditionalAuditEnumLowerCase = map[string]SecurityFeatureSummaryTraditionalAuditEnum{
	"enabled":  SecurityFeatureSummaryTraditionalAuditEnabled,
	"disabled": SecurityFeatureSummaryTraditionalAuditDisabled,
	"none":     SecurityFeatureSummaryTraditionalAuditNone,
}

// GetSecurityFeatureSummaryTraditionalAuditEnumValues Enumerates the set of values for SecurityFeatureSummaryTraditionalAuditEnum
func GetSecurityFeatureSummaryTraditionalAuditEnumValues() []SecurityFeatureSummaryTraditionalAuditEnum {
	values := make([]SecurityFeatureSummaryTraditionalAuditEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryTraditionalAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryTraditionalAuditEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryTraditionalAuditEnum
func GetSecurityFeatureSummaryTraditionalAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryTraditionalAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryTraditionalAuditEnum(val string) (SecurityFeatureSummaryTraditionalAuditEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryTraditionalAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryDatabaseVaultEnum Enum with underlying type: string
type SecurityFeatureSummaryDatabaseVaultEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryDatabaseVaultEnum
const (
	SecurityFeatureSummaryDatabaseVaultEnabled  SecurityFeatureSummaryDatabaseVaultEnum = "ENABLED"
	SecurityFeatureSummaryDatabaseVaultDisabled SecurityFeatureSummaryDatabaseVaultEnum = "DISABLED"
	SecurityFeatureSummaryDatabaseVaultNone     SecurityFeatureSummaryDatabaseVaultEnum = "NONE"
)

var mappingSecurityFeatureSummaryDatabaseVaultEnum = map[string]SecurityFeatureSummaryDatabaseVaultEnum{
	"ENABLED":  SecurityFeatureSummaryDatabaseVaultEnabled,
	"DISABLED": SecurityFeatureSummaryDatabaseVaultDisabled,
	"NONE":     SecurityFeatureSummaryDatabaseVaultNone,
}

var mappingSecurityFeatureSummaryDatabaseVaultEnumLowerCase = map[string]SecurityFeatureSummaryDatabaseVaultEnum{
	"enabled":  SecurityFeatureSummaryDatabaseVaultEnabled,
	"disabled": SecurityFeatureSummaryDatabaseVaultDisabled,
	"none":     SecurityFeatureSummaryDatabaseVaultNone,
}

// GetSecurityFeatureSummaryDatabaseVaultEnumValues Enumerates the set of values for SecurityFeatureSummaryDatabaseVaultEnum
func GetSecurityFeatureSummaryDatabaseVaultEnumValues() []SecurityFeatureSummaryDatabaseVaultEnum {
	values := make([]SecurityFeatureSummaryDatabaseVaultEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryDatabaseVaultEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryDatabaseVaultEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryDatabaseVaultEnum
func GetSecurityFeatureSummaryDatabaseVaultEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryDatabaseVaultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryDatabaseVaultEnum(val string) (SecurityFeatureSummaryDatabaseVaultEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryDatabaseVaultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryPrivilegeAnalysisEnum Enum with underlying type: string
type SecurityFeatureSummaryPrivilegeAnalysisEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryPrivilegeAnalysisEnum
const (
	SecurityFeatureSummaryPrivilegeAnalysisEnabled  SecurityFeatureSummaryPrivilegeAnalysisEnum = "ENABLED"
	SecurityFeatureSummaryPrivilegeAnalysisDisabled SecurityFeatureSummaryPrivilegeAnalysisEnum = "DISABLED"
	SecurityFeatureSummaryPrivilegeAnalysisNone     SecurityFeatureSummaryPrivilegeAnalysisEnum = "NONE"
)

var mappingSecurityFeatureSummaryPrivilegeAnalysisEnum = map[string]SecurityFeatureSummaryPrivilegeAnalysisEnum{
	"ENABLED":  SecurityFeatureSummaryPrivilegeAnalysisEnabled,
	"DISABLED": SecurityFeatureSummaryPrivilegeAnalysisDisabled,
	"NONE":     SecurityFeatureSummaryPrivilegeAnalysisNone,
}

var mappingSecurityFeatureSummaryPrivilegeAnalysisEnumLowerCase = map[string]SecurityFeatureSummaryPrivilegeAnalysisEnum{
	"enabled":  SecurityFeatureSummaryPrivilegeAnalysisEnabled,
	"disabled": SecurityFeatureSummaryPrivilegeAnalysisDisabled,
	"none":     SecurityFeatureSummaryPrivilegeAnalysisNone,
}

// GetSecurityFeatureSummaryPrivilegeAnalysisEnumValues Enumerates the set of values for SecurityFeatureSummaryPrivilegeAnalysisEnum
func GetSecurityFeatureSummaryPrivilegeAnalysisEnumValues() []SecurityFeatureSummaryPrivilegeAnalysisEnum {
	values := make([]SecurityFeatureSummaryPrivilegeAnalysisEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryPrivilegeAnalysisEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryPrivilegeAnalysisEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryPrivilegeAnalysisEnum
func GetSecurityFeatureSummaryPrivilegeAnalysisEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryPrivilegeAnalysisEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryPrivilegeAnalysisEnum(val string) (SecurityFeatureSummaryPrivilegeAnalysisEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryPrivilegeAnalysisEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryTablespaceEncryptionEnum Enum with underlying type: string
type SecurityFeatureSummaryTablespaceEncryptionEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryTablespaceEncryptionEnum
const (
	SecurityFeatureSummaryTablespaceEncryptionEnabled  SecurityFeatureSummaryTablespaceEncryptionEnum = "ENABLED"
	SecurityFeatureSummaryTablespaceEncryptionDisabled SecurityFeatureSummaryTablespaceEncryptionEnum = "DISABLED"
	SecurityFeatureSummaryTablespaceEncryptionNone     SecurityFeatureSummaryTablespaceEncryptionEnum = "NONE"
)

var mappingSecurityFeatureSummaryTablespaceEncryptionEnum = map[string]SecurityFeatureSummaryTablespaceEncryptionEnum{
	"ENABLED":  SecurityFeatureSummaryTablespaceEncryptionEnabled,
	"DISABLED": SecurityFeatureSummaryTablespaceEncryptionDisabled,
	"NONE":     SecurityFeatureSummaryTablespaceEncryptionNone,
}

var mappingSecurityFeatureSummaryTablespaceEncryptionEnumLowerCase = map[string]SecurityFeatureSummaryTablespaceEncryptionEnum{
	"enabled":  SecurityFeatureSummaryTablespaceEncryptionEnabled,
	"disabled": SecurityFeatureSummaryTablespaceEncryptionDisabled,
	"none":     SecurityFeatureSummaryTablespaceEncryptionNone,
}

// GetSecurityFeatureSummaryTablespaceEncryptionEnumValues Enumerates the set of values for SecurityFeatureSummaryTablespaceEncryptionEnum
func GetSecurityFeatureSummaryTablespaceEncryptionEnumValues() []SecurityFeatureSummaryTablespaceEncryptionEnum {
	values := make([]SecurityFeatureSummaryTablespaceEncryptionEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryTablespaceEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryTablespaceEncryptionEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryTablespaceEncryptionEnum
func GetSecurityFeatureSummaryTablespaceEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryTablespaceEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryTablespaceEncryptionEnum(val string) (SecurityFeatureSummaryTablespaceEncryptionEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryTablespaceEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryColumnEncryptionEnum Enum with underlying type: string
type SecurityFeatureSummaryColumnEncryptionEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryColumnEncryptionEnum
const (
	SecurityFeatureSummaryColumnEncryptionEnabled  SecurityFeatureSummaryColumnEncryptionEnum = "ENABLED"
	SecurityFeatureSummaryColumnEncryptionDisabled SecurityFeatureSummaryColumnEncryptionEnum = "DISABLED"
	SecurityFeatureSummaryColumnEncryptionNone     SecurityFeatureSummaryColumnEncryptionEnum = "NONE"
)

var mappingSecurityFeatureSummaryColumnEncryptionEnum = map[string]SecurityFeatureSummaryColumnEncryptionEnum{
	"ENABLED":  SecurityFeatureSummaryColumnEncryptionEnabled,
	"DISABLED": SecurityFeatureSummaryColumnEncryptionDisabled,
	"NONE":     SecurityFeatureSummaryColumnEncryptionNone,
}

var mappingSecurityFeatureSummaryColumnEncryptionEnumLowerCase = map[string]SecurityFeatureSummaryColumnEncryptionEnum{
	"enabled":  SecurityFeatureSummaryColumnEncryptionEnabled,
	"disabled": SecurityFeatureSummaryColumnEncryptionDisabled,
	"none":     SecurityFeatureSummaryColumnEncryptionNone,
}

// GetSecurityFeatureSummaryColumnEncryptionEnumValues Enumerates the set of values for SecurityFeatureSummaryColumnEncryptionEnum
func GetSecurityFeatureSummaryColumnEncryptionEnumValues() []SecurityFeatureSummaryColumnEncryptionEnum {
	values := make([]SecurityFeatureSummaryColumnEncryptionEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryColumnEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryColumnEncryptionEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryColumnEncryptionEnum
func GetSecurityFeatureSummaryColumnEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryColumnEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryColumnEncryptionEnum(val string) (SecurityFeatureSummaryColumnEncryptionEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryColumnEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryNetworkEncryptionEnum Enum with underlying type: string
type SecurityFeatureSummaryNetworkEncryptionEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryNetworkEncryptionEnum
const (
	SecurityFeatureSummaryNetworkEncryptionEnabled  SecurityFeatureSummaryNetworkEncryptionEnum = "ENABLED"
	SecurityFeatureSummaryNetworkEncryptionDisabled SecurityFeatureSummaryNetworkEncryptionEnum = "DISABLED"
	SecurityFeatureSummaryNetworkEncryptionNone     SecurityFeatureSummaryNetworkEncryptionEnum = "NONE"
)

var mappingSecurityFeatureSummaryNetworkEncryptionEnum = map[string]SecurityFeatureSummaryNetworkEncryptionEnum{
	"ENABLED":  SecurityFeatureSummaryNetworkEncryptionEnabled,
	"DISABLED": SecurityFeatureSummaryNetworkEncryptionDisabled,
	"NONE":     SecurityFeatureSummaryNetworkEncryptionNone,
}

var mappingSecurityFeatureSummaryNetworkEncryptionEnumLowerCase = map[string]SecurityFeatureSummaryNetworkEncryptionEnum{
	"enabled":  SecurityFeatureSummaryNetworkEncryptionEnabled,
	"disabled": SecurityFeatureSummaryNetworkEncryptionDisabled,
	"none":     SecurityFeatureSummaryNetworkEncryptionNone,
}

// GetSecurityFeatureSummaryNetworkEncryptionEnumValues Enumerates the set of values for SecurityFeatureSummaryNetworkEncryptionEnum
func GetSecurityFeatureSummaryNetworkEncryptionEnumValues() []SecurityFeatureSummaryNetworkEncryptionEnum {
	values := make([]SecurityFeatureSummaryNetworkEncryptionEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryNetworkEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryNetworkEncryptionEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryNetworkEncryptionEnum
func GetSecurityFeatureSummaryNetworkEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryNetworkEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryNetworkEncryptionEnum(val string) (SecurityFeatureSummaryNetworkEncryptionEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryNetworkEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryPasswordAuthenticationEnum Enum with underlying type: string
type SecurityFeatureSummaryPasswordAuthenticationEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryPasswordAuthenticationEnum
const (
	SecurityFeatureSummaryPasswordAuthenticationEnabled  SecurityFeatureSummaryPasswordAuthenticationEnum = "ENABLED"
	SecurityFeatureSummaryPasswordAuthenticationDisabled SecurityFeatureSummaryPasswordAuthenticationEnum = "DISABLED"
	SecurityFeatureSummaryPasswordAuthenticationNone     SecurityFeatureSummaryPasswordAuthenticationEnum = "NONE"
)

var mappingSecurityFeatureSummaryPasswordAuthenticationEnum = map[string]SecurityFeatureSummaryPasswordAuthenticationEnum{
	"ENABLED":  SecurityFeatureSummaryPasswordAuthenticationEnabled,
	"DISABLED": SecurityFeatureSummaryPasswordAuthenticationDisabled,
	"NONE":     SecurityFeatureSummaryPasswordAuthenticationNone,
}

var mappingSecurityFeatureSummaryPasswordAuthenticationEnumLowerCase = map[string]SecurityFeatureSummaryPasswordAuthenticationEnum{
	"enabled":  SecurityFeatureSummaryPasswordAuthenticationEnabled,
	"disabled": SecurityFeatureSummaryPasswordAuthenticationDisabled,
	"none":     SecurityFeatureSummaryPasswordAuthenticationNone,
}

// GetSecurityFeatureSummaryPasswordAuthenticationEnumValues Enumerates the set of values for SecurityFeatureSummaryPasswordAuthenticationEnum
func GetSecurityFeatureSummaryPasswordAuthenticationEnumValues() []SecurityFeatureSummaryPasswordAuthenticationEnum {
	values := make([]SecurityFeatureSummaryPasswordAuthenticationEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryPasswordAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryPasswordAuthenticationEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryPasswordAuthenticationEnum
func GetSecurityFeatureSummaryPasswordAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryPasswordAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryPasswordAuthenticationEnum(val string) (SecurityFeatureSummaryPasswordAuthenticationEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryPasswordAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryGlobalAuthenticationEnum Enum with underlying type: string
type SecurityFeatureSummaryGlobalAuthenticationEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryGlobalAuthenticationEnum
const (
	SecurityFeatureSummaryGlobalAuthenticationEnabled  SecurityFeatureSummaryGlobalAuthenticationEnum = "ENABLED"
	SecurityFeatureSummaryGlobalAuthenticationDisabled SecurityFeatureSummaryGlobalAuthenticationEnum = "DISABLED"
	SecurityFeatureSummaryGlobalAuthenticationNone     SecurityFeatureSummaryGlobalAuthenticationEnum = "NONE"
)

var mappingSecurityFeatureSummaryGlobalAuthenticationEnum = map[string]SecurityFeatureSummaryGlobalAuthenticationEnum{
	"ENABLED":  SecurityFeatureSummaryGlobalAuthenticationEnabled,
	"DISABLED": SecurityFeatureSummaryGlobalAuthenticationDisabled,
	"NONE":     SecurityFeatureSummaryGlobalAuthenticationNone,
}

var mappingSecurityFeatureSummaryGlobalAuthenticationEnumLowerCase = map[string]SecurityFeatureSummaryGlobalAuthenticationEnum{
	"enabled":  SecurityFeatureSummaryGlobalAuthenticationEnabled,
	"disabled": SecurityFeatureSummaryGlobalAuthenticationDisabled,
	"none":     SecurityFeatureSummaryGlobalAuthenticationNone,
}

// GetSecurityFeatureSummaryGlobalAuthenticationEnumValues Enumerates the set of values for SecurityFeatureSummaryGlobalAuthenticationEnum
func GetSecurityFeatureSummaryGlobalAuthenticationEnumValues() []SecurityFeatureSummaryGlobalAuthenticationEnum {
	values := make([]SecurityFeatureSummaryGlobalAuthenticationEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryGlobalAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryGlobalAuthenticationEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryGlobalAuthenticationEnum
func GetSecurityFeatureSummaryGlobalAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryGlobalAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryGlobalAuthenticationEnum(val string) (SecurityFeatureSummaryGlobalAuthenticationEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryGlobalAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityFeatureSummaryExternalAuthenticationEnum Enum with underlying type: string
type SecurityFeatureSummaryExternalAuthenticationEnum string

// Set of constants representing the allowable values for SecurityFeatureSummaryExternalAuthenticationEnum
const (
	SecurityFeatureSummaryExternalAuthenticationEnabled  SecurityFeatureSummaryExternalAuthenticationEnum = "ENABLED"
	SecurityFeatureSummaryExternalAuthenticationDisabled SecurityFeatureSummaryExternalAuthenticationEnum = "DISABLED"
	SecurityFeatureSummaryExternalAuthenticationNone     SecurityFeatureSummaryExternalAuthenticationEnum = "NONE"
)

var mappingSecurityFeatureSummaryExternalAuthenticationEnum = map[string]SecurityFeatureSummaryExternalAuthenticationEnum{
	"ENABLED":  SecurityFeatureSummaryExternalAuthenticationEnabled,
	"DISABLED": SecurityFeatureSummaryExternalAuthenticationDisabled,
	"NONE":     SecurityFeatureSummaryExternalAuthenticationNone,
}

var mappingSecurityFeatureSummaryExternalAuthenticationEnumLowerCase = map[string]SecurityFeatureSummaryExternalAuthenticationEnum{
	"enabled":  SecurityFeatureSummaryExternalAuthenticationEnabled,
	"disabled": SecurityFeatureSummaryExternalAuthenticationDisabled,
	"none":     SecurityFeatureSummaryExternalAuthenticationNone,
}

// GetSecurityFeatureSummaryExternalAuthenticationEnumValues Enumerates the set of values for SecurityFeatureSummaryExternalAuthenticationEnum
func GetSecurityFeatureSummaryExternalAuthenticationEnumValues() []SecurityFeatureSummaryExternalAuthenticationEnum {
	values := make([]SecurityFeatureSummaryExternalAuthenticationEnum, 0)
	for _, v := range mappingSecurityFeatureSummaryExternalAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureSummaryExternalAuthenticationEnumStringValues Enumerates the set of values in String for SecurityFeatureSummaryExternalAuthenticationEnum
func GetSecurityFeatureSummaryExternalAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingSecurityFeatureSummaryExternalAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureSummaryExternalAuthenticationEnum(val string) (SecurityFeatureSummaryExternalAuthenticationEnum, bool) {
	enum, ok := mappingSecurityFeatureSummaryExternalAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
