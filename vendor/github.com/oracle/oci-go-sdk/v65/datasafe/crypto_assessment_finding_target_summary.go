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

// CryptoAssessmentFindingTargetSummary Target-level occurrence details for a crypto finding.
type CryptoAssessmentFindingTargetSummary struct {

	// The OCID of the affected target.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Unique key of the finding affecting this target.
	FindingKey *string `mandatory:"true" json:"findingKey"`

	// The observed value for the selected finding on this target.
	ObservedValue *string `mandatory:"true" json:"observedValue"`

	// Database version of the affected target.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// The crypto assessment OCID associated with this finding occurrence.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	Priority *int `mandatory:"false" json:"priority"`

	// Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	Severity CryptoAssessmentFindingTargetSummarySeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Indicates whether this finding is part of quantum-readiness checks.
	IsQuantumReadinessCheck *bool `mandatory:"false" json:"isQuantumReadinessCheck"`
}

func (m CryptoAssessmentFindingTargetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentFindingTargetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoAssessmentFindingTargetSummarySeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetCryptoAssessmentFindingTargetSummarySeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentFindingTargetSummarySeverityEnum Enum with underlying type: string
type CryptoAssessmentFindingTargetSummarySeverityEnum string

// Set of constants representing the allowable values for CryptoAssessmentFindingTargetSummarySeverityEnum
const (
	CryptoAssessmentFindingTargetSummarySeverityCritical CryptoAssessmentFindingTargetSummarySeverityEnum = "CRITICAL"
	CryptoAssessmentFindingTargetSummarySeverityHigh     CryptoAssessmentFindingTargetSummarySeverityEnum = "HIGH"
	CryptoAssessmentFindingTargetSummarySeverityMedium   CryptoAssessmentFindingTargetSummarySeverityEnum = "MEDIUM"
	CryptoAssessmentFindingTargetSummarySeverityLow      CryptoAssessmentFindingTargetSummarySeverityEnum = "LOW"
)

var mappingCryptoAssessmentFindingTargetSummarySeverityEnum = map[string]CryptoAssessmentFindingTargetSummarySeverityEnum{
	"CRITICAL": CryptoAssessmentFindingTargetSummarySeverityCritical,
	"HIGH":     CryptoAssessmentFindingTargetSummarySeverityHigh,
	"MEDIUM":   CryptoAssessmentFindingTargetSummarySeverityMedium,
	"LOW":      CryptoAssessmentFindingTargetSummarySeverityLow,
}

var mappingCryptoAssessmentFindingTargetSummarySeverityEnumLowerCase = map[string]CryptoAssessmentFindingTargetSummarySeverityEnum{
	"critical": CryptoAssessmentFindingTargetSummarySeverityCritical,
	"high":     CryptoAssessmentFindingTargetSummarySeverityHigh,
	"medium":   CryptoAssessmentFindingTargetSummarySeverityMedium,
	"low":      CryptoAssessmentFindingTargetSummarySeverityLow,
}

// GetCryptoAssessmentFindingTargetSummarySeverityEnumValues Enumerates the set of values for CryptoAssessmentFindingTargetSummarySeverityEnum
func GetCryptoAssessmentFindingTargetSummarySeverityEnumValues() []CryptoAssessmentFindingTargetSummarySeverityEnum {
	values := make([]CryptoAssessmentFindingTargetSummarySeverityEnum, 0)
	for _, v := range mappingCryptoAssessmentFindingTargetSummarySeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentFindingTargetSummarySeverityEnumStringValues Enumerates the set of values in String for CryptoAssessmentFindingTargetSummarySeverityEnum
func GetCryptoAssessmentFindingTargetSummarySeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}

// GetMappingCryptoAssessmentFindingTargetSummarySeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentFindingTargetSummarySeverityEnum(val string) (CryptoAssessmentFindingTargetSummarySeverityEnum, bool) {
	enum, ok := mappingCryptoAssessmentFindingTargetSummarySeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
