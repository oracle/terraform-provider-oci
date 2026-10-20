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

// CryptoAssessmentFindingSummary One crypto deviation finding for a specific assessment run.
type CryptoAssessmentFindingSummary struct {

	// Unique key identifier for the finding.
	FindingKey *string `mandatory:"true" json:"findingKey"`

	// Human-readable title for the finding.
	Title *string `mandatory:"true" json:"title"`

	// Status recorded for the finding.
	Status CryptoFindingStatusEnum `mandatory:"true" json:"status"`

	// Category value recorded for the finding.
	Category *string `mandatory:"false" json:"category"`

	// Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	Priority *int `mandatory:"false" json:"priority"`

	// Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	Severity CryptoAssessmentFindingSummarySeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Expected value recorded for the finding.
	ExpectedValue *string `mandatory:"false" json:"expectedValue"`

	// Observed values recorded for the finding.
	ObservedValue []string `mandatory:"false" json:"observedValue"`

	// Recommended value recorded for the finding.
	RecommendedValue *string `mandatory:"false" json:"recommendedValue"`

	// Indicates whether this finding is part of quantum-readiness checks.
	IsQuantumReadinessCheck *bool `mandatory:"false" json:"isQuantumReadinessCheck"`

	// Summary recorded for the finding.
	Summary *string `mandatory:"false" json:"summary"`

	// Short summary of the finding.
	ShortSummary *string `mandatory:"false" json:"shortSummary"`

	// URL recorded for the finding.
	Url *string `mandatory:"false" json:"url"`

	// Remediation text recorded for the finding.
	Remediation *string `mandatory:"false" json:"remediation"`

	// Short remediation for the finding.
	ShortRemediation *string `mandatory:"false" json:"shortRemediation"`

	// Compliance mapping recorded for the finding.
	Compliance *string `mandatory:"false" json:"compliance"`
}

func (m CryptoAssessmentFindingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentFindingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoFindingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoFindingStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCryptoAssessmentFindingSummarySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetCryptoAssessmentFindingSummarySeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentFindingSummarySeverityEnum Enum with underlying type: string
type CryptoAssessmentFindingSummarySeverityEnum string

// Set of constants representing the allowable values for CryptoAssessmentFindingSummarySeverityEnum
const (
	CryptoAssessmentFindingSummarySeverityCritical CryptoAssessmentFindingSummarySeverityEnum = "CRITICAL"
	CryptoAssessmentFindingSummarySeverityHigh     CryptoAssessmentFindingSummarySeverityEnum = "HIGH"
	CryptoAssessmentFindingSummarySeverityMedium   CryptoAssessmentFindingSummarySeverityEnum = "MEDIUM"
	CryptoAssessmentFindingSummarySeverityLow      CryptoAssessmentFindingSummarySeverityEnum = "LOW"
)

var mappingCryptoAssessmentFindingSummarySeverityEnum = map[string]CryptoAssessmentFindingSummarySeverityEnum{
	"CRITICAL": CryptoAssessmentFindingSummarySeverityCritical,
	"HIGH":     CryptoAssessmentFindingSummarySeverityHigh,
	"MEDIUM":   CryptoAssessmentFindingSummarySeverityMedium,
	"LOW":      CryptoAssessmentFindingSummarySeverityLow,
}

var mappingCryptoAssessmentFindingSummarySeverityEnumLowerCase = map[string]CryptoAssessmentFindingSummarySeverityEnum{
	"critical": CryptoAssessmentFindingSummarySeverityCritical,
	"high":     CryptoAssessmentFindingSummarySeverityHigh,
	"medium":   CryptoAssessmentFindingSummarySeverityMedium,
	"low":      CryptoAssessmentFindingSummarySeverityLow,
}

// GetCryptoAssessmentFindingSummarySeverityEnumValues Enumerates the set of values for CryptoAssessmentFindingSummarySeverityEnum
func GetCryptoAssessmentFindingSummarySeverityEnumValues() []CryptoAssessmentFindingSummarySeverityEnum {
	values := make([]CryptoAssessmentFindingSummarySeverityEnum, 0)
	for _, v := range mappingCryptoAssessmentFindingSummarySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentFindingSummarySeverityEnumStringValues Enumerates the set of values in String for CryptoAssessmentFindingSummarySeverityEnum
func GetCryptoAssessmentFindingSummarySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingCryptoAssessmentFindingSummarySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentFindingSummarySeverityEnum(val string) (CryptoAssessmentFindingSummarySeverityEnum, bool) {
	enum, ok := mappingCryptoAssessmentFindingSummarySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
