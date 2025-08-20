// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CheckSummary The summary of the security rule to be evaluated by security assessment to create finding.
type CheckSummary struct {

	// A unique identifier for the check.
	Key *string `mandatory:"true" json:"key"`

	// The short title for the check.
	Title *string `mandatory:"true" json:"title"`

	// The explanation of the issue in this check. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
	Remarks *string `mandatory:"false" json:"remarks"`

	// Provides information on whether the check is related to a CIS Oracle Database Benchmark recommendation, STIG rule, GDPR Article/Recital or related to the Oracle Recommended Practice.
	References *References `mandatory:"false" json:"references"`

	// The category to which the check belongs to.
	Category *string `mandatory:"false" json:"category"`

	// Provides a recommended approach to take to remediate the check reported.
	Oneline *string `mandatory:"false" json:"oneline"`

	// The severity of the check as suggested by Data Safe security assessment. This will be the default severity in the template baseline security assessment.
	SuggestedSeverity CheckSummarySuggestedSeverityEnum `mandatory:"false" json:"suggestedSeverity,omitempty"`
}

func (m CheckSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CheckSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCheckSummarySuggestedSeverityEnum(string(m.SuggestedSeverity)); !ok && m.SuggestedSeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SuggestedSeverity: %s. Supported values are: %s.", m.SuggestedSeverity, strings.Join(GetCheckSummarySuggestedSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CheckSummarySuggestedSeverityEnum Enum with underlying type: string
type CheckSummarySuggestedSeverityEnum string

// Set of constants representing the allowable values for CheckSummarySuggestedSeverityEnum
const (
	CheckSummarySuggestedSeverityHigh     CheckSummarySuggestedSeverityEnum = "HIGH"
	CheckSummarySuggestedSeverityMedium   CheckSummarySuggestedSeverityEnum = "MEDIUM"
	CheckSummarySuggestedSeverityLow      CheckSummarySuggestedSeverityEnum = "LOW"
	CheckSummarySuggestedSeverityEvaluate CheckSummarySuggestedSeverityEnum = "EVALUATE"
	CheckSummarySuggestedSeverityAdvisory CheckSummarySuggestedSeverityEnum = "ADVISORY"
	CheckSummarySuggestedSeverityPass     CheckSummarySuggestedSeverityEnum = "PASS"
	CheckSummarySuggestedSeverityDeferred CheckSummarySuggestedSeverityEnum = "DEFERRED"
)

var mappingCheckSummarySuggestedSeverityEnum = map[string]CheckSummarySuggestedSeverityEnum{
	"HIGH":     CheckSummarySuggestedSeverityHigh,
	"MEDIUM":   CheckSummarySuggestedSeverityMedium,
	"LOW":      CheckSummarySuggestedSeverityLow,
	"EVALUATE": CheckSummarySuggestedSeverityEvaluate,
	"ADVISORY": CheckSummarySuggestedSeverityAdvisory,
	"PASS":     CheckSummarySuggestedSeverityPass,
	"DEFERRED": CheckSummarySuggestedSeverityDeferred,
}

var mappingCheckSummarySuggestedSeverityEnumLowerCase = map[string]CheckSummarySuggestedSeverityEnum{
	"high":     CheckSummarySuggestedSeverityHigh,
	"medium":   CheckSummarySuggestedSeverityMedium,
	"low":      CheckSummarySuggestedSeverityLow,
	"evaluate": CheckSummarySuggestedSeverityEvaluate,
	"advisory": CheckSummarySuggestedSeverityAdvisory,
	"pass":     CheckSummarySuggestedSeverityPass,
	"deferred": CheckSummarySuggestedSeverityDeferred,
}

// GetCheckSummarySuggestedSeverityEnumValues Enumerates the set of values for CheckSummarySuggestedSeverityEnum
func GetCheckSummarySuggestedSeverityEnumValues() []CheckSummarySuggestedSeverityEnum {
	values := make([]CheckSummarySuggestedSeverityEnum, 0)
	for _, v := range mappingCheckSummarySuggestedSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetCheckSummarySuggestedSeverityEnumStringValues Enumerates the set of values in String for CheckSummarySuggestedSeverityEnum
func GetCheckSummarySuggestedSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingCheckSummarySuggestedSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCheckSummarySuggestedSeverityEnum(val string) (CheckSummarySuggestedSeverityEnum, bool) {
	enum, ok := mappingCheckSummarySuggestedSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
