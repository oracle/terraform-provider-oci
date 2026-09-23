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

// CryptoAssessmentFindingAnalyticsSummary Aggregated finding details with number of affected targets.
type CryptoAssessmentFindingAnalyticsSummary struct {

	// Unique key of the finding.
	FindingKey *string `mandatory:"true" json:"findingKey"`

	// Display title of the finding.
	Title *string `mandatory:"true" json:"title"`

	// Category key to which the finding belongs.
	Category CryptoAssessmentFindingAnalyticsSummaryCategoryEnum `mandatory:"true" json:"category"`

	// Number of targets impacted by this finding in the queried scope.
	TargetCount *int `mandatory:"true" json:"targetCount"`

	// Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	Priority *int `mandatory:"false" json:"priority"`

	// Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	Severity CryptoAssessmentFindingAnalyticsSummarySeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Short summary of the finding.
	ShortSummary *string `mandatory:"false" json:"shortSummary"`

	// Short remediation for the finding.
	ShortRemediation *string `mandatory:"false" json:"shortRemediation"`
}

func (m CryptoAssessmentFindingAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentFindingAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetCryptoAssessmentFindingAnalyticsSummaryCategoryEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCryptoAssessmentFindingAnalyticsSummarySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetCryptoAssessmentFindingAnalyticsSummarySeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentFindingAnalyticsSummaryCategoryEnum Enum with underlying type: string
type CryptoAssessmentFindingAnalyticsSummaryCategoryEnum string

// Set of constants representing the allowable values for CryptoAssessmentFindingAnalyticsSummaryCategoryEnum
const (
	CryptoAssessmentFindingAnalyticsSummaryCategoryNetworkEncryption            CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "NETWORK_ENCRYPTION"
	CryptoAssessmentFindingAnalyticsSummaryCategoryDataEncryption               CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "DATA_ENCRYPTION"
	CryptoAssessmentFindingAnalyticsSummaryCategoryCertificatesAndKeyManagement CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "CERTIFICATES_AND_KEY_MANAGEMENT"
	CryptoAssessmentFindingAnalyticsSummaryCategoryBackupAndExportEncryption    CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "BACKUP_AND_EXPORT_ENCRYPTION"
	CryptoAssessmentFindingAnalyticsSummaryCategoryPostQuantumReadiness         CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "POST_QUANTUM_READINESS"
	CryptoAssessmentFindingAnalyticsSummaryCategoryNotSupported                 CryptoAssessmentFindingAnalyticsSummaryCategoryEnum = "NOT_SUPPORTED"
)

var mappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnum = map[string]CryptoAssessmentFindingAnalyticsSummaryCategoryEnum{
	"NETWORK_ENCRYPTION":              CryptoAssessmentFindingAnalyticsSummaryCategoryNetworkEncryption,
	"DATA_ENCRYPTION":                 CryptoAssessmentFindingAnalyticsSummaryCategoryDataEncryption,
	"CERTIFICATES_AND_KEY_MANAGEMENT": CryptoAssessmentFindingAnalyticsSummaryCategoryCertificatesAndKeyManagement,
	"BACKUP_AND_EXPORT_ENCRYPTION":    CryptoAssessmentFindingAnalyticsSummaryCategoryBackupAndExportEncryption,
	"POST_QUANTUM_READINESS":          CryptoAssessmentFindingAnalyticsSummaryCategoryPostQuantumReadiness,
	"NOT_SUPPORTED":                   CryptoAssessmentFindingAnalyticsSummaryCategoryNotSupported,
}

var mappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnumLowerCase = map[string]CryptoAssessmentFindingAnalyticsSummaryCategoryEnum{
	"network_encryption":              CryptoAssessmentFindingAnalyticsSummaryCategoryNetworkEncryption,
	"data_encryption":                 CryptoAssessmentFindingAnalyticsSummaryCategoryDataEncryption,
	"certificates_and_key_management": CryptoAssessmentFindingAnalyticsSummaryCategoryCertificatesAndKeyManagement,
	"backup_and_export_encryption":    CryptoAssessmentFindingAnalyticsSummaryCategoryBackupAndExportEncryption,
	"post_quantum_readiness":          CryptoAssessmentFindingAnalyticsSummaryCategoryPostQuantumReadiness,
	"not_supported":                   CryptoAssessmentFindingAnalyticsSummaryCategoryNotSupported,
}

// GetCryptoAssessmentFindingAnalyticsSummaryCategoryEnumValues Enumerates the set of values for CryptoAssessmentFindingAnalyticsSummaryCategoryEnum
func GetCryptoAssessmentFindingAnalyticsSummaryCategoryEnumValues() []CryptoAssessmentFindingAnalyticsSummaryCategoryEnum {
	values := make([]CryptoAssessmentFindingAnalyticsSummaryCategoryEnum, 0)
	for _, v := range mappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentFindingAnalyticsSummaryCategoryEnumStringValues Enumerates the set of values in String for CryptoAssessmentFindingAnalyticsSummaryCategoryEnum
func GetCryptoAssessmentFindingAnalyticsSummaryCategoryEnumStringValues() []string {
	return []string{
		"NETWORK_ENCRYPTION",
		"DATA_ENCRYPTION",
		"CERTIFICATES_AND_KEY_MANAGEMENT",
		"BACKUP_AND_EXPORT_ENCRYPTION",
		"POST_QUANTUM_READINESS",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnum(val string) (CryptoAssessmentFindingAnalyticsSummaryCategoryEnum, bool) {
	enum, ok := mappingCryptoAssessmentFindingAnalyticsSummaryCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentFindingAnalyticsSummarySeverityEnum Enum with underlying type: string
type CryptoAssessmentFindingAnalyticsSummarySeverityEnum string

// Set of constants representing the allowable values for CryptoAssessmentFindingAnalyticsSummarySeverityEnum
const (
	CryptoAssessmentFindingAnalyticsSummarySeverityCritical CryptoAssessmentFindingAnalyticsSummarySeverityEnum = "CRITICAL"
	CryptoAssessmentFindingAnalyticsSummarySeverityHigh     CryptoAssessmentFindingAnalyticsSummarySeverityEnum = "HIGH"
	CryptoAssessmentFindingAnalyticsSummarySeverityMedium   CryptoAssessmentFindingAnalyticsSummarySeverityEnum = "MEDIUM"
	CryptoAssessmentFindingAnalyticsSummarySeverityLow      CryptoAssessmentFindingAnalyticsSummarySeverityEnum = "LOW"
)

var mappingCryptoAssessmentFindingAnalyticsSummarySeverityEnum = map[string]CryptoAssessmentFindingAnalyticsSummarySeverityEnum{
	"CRITICAL": CryptoAssessmentFindingAnalyticsSummarySeverityCritical,
	"HIGH":     CryptoAssessmentFindingAnalyticsSummarySeverityHigh,
	"MEDIUM":   CryptoAssessmentFindingAnalyticsSummarySeverityMedium,
	"LOW":      CryptoAssessmentFindingAnalyticsSummarySeverityLow,
}

var mappingCryptoAssessmentFindingAnalyticsSummarySeverityEnumLowerCase = map[string]CryptoAssessmentFindingAnalyticsSummarySeverityEnum{
	"critical": CryptoAssessmentFindingAnalyticsSummarySeverityCritical,
	"high":     CryptoAssessmentFindingAnalyticsSummarySeverityHigh,
	"medium":   CryptoAssessmentFindingAnalyticsSummarySeverityMedium,
	"low":      CryptoAssessmentFindingAnalyticsSummarySeverityLow,
}

// GetCryptoAssessmentFindingAnalyticsSummarySeverityEnumValues Enumerates the set of values for CryptoAssessmentFindingAnalyticsSummarySeverityEnum
func GetCryptoAssessmentFindingAnalyticsSummarySeverityEnumValues() []CryptoAssessmentFindingAnalyticsSummarySeverityEnum {
	values := make([]CryptoAssessmentFindingAnalyticsSummarySeverityEnum, 0)
	for _, v := range mappingCryptoAssessmentFindingAnalyticsSummarySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentFindingAnalyticsSummarySeverityEnumStringValues Enumerates the set of values in String for CryptoAssessmentFindingAnalyticsSummarySeverityEnum
func GetCryptoAssessmentFindingAnalyticsSummarySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingCryptoAssessmentFindingAnalyticsSummarySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentFindingAnalyticsSummarySeverityEnum(val string) (CryptoAssessmentFindingAnalyticsSummarySeverityEnum, bool) {
	enum, ok := mappingCryptoAssessmentFindingAnalyticsSummarySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
