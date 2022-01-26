// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Finding The particular finding reported by the security assessment.
type Finding struct {

	// A unique identifier for the finding. This is common for the finding across targets.
	Key *string `mandatory:"false" json:"key"`

	// The severity of the finding.
	Severity FindingSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The short title for the finding.
	Title *string `mandatory:"false" json:"title"`

	// The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
	Remarks *string `mandatory:"false" json:"remarks"`

	// The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
	Details *interface{} `mandatory:"false" json:"details"`

	// The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
	Summary *string `mandatory:"false" json:"summary"`

	// Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, STIG rule, or related to a GDPR Article/Recital.
	References *References `mandatory:"false" json:"references"`
}

func (m Finding) String() string {
	return common.PointerString(m)
}

// FindingSeverityEnum Enum with underlying type: string
type FindingSeverityEnum string

// Set of constants representing the allowable values for FindingSeverityEnum
const (
	FindingSeverityHigh     FindingSeverityEnum = "HIGH"
	FindingSeverityMedium   FindingSeverityEnum = "MEDIUM"
	FindingSeverityLow      FindingSeverityEnum = "LOW"
	FindingSeverityEvaluate FindingSeverityEnum = "EVALUATE"
	FindingSeverityAdvisory FindingSeverityEnum = "ADVISORY"
	FindingSeverityPass     FindingSeverityEnum = "PASS"
)

var mappingFindingSeverity = map[string]FindingSeverityEnum{
	"HIGH":     FindingSeverityHigh,
	"MEDIUM":   FindingSeverityMedium,
	"LOW":      FindingSeverityLow,
	"EVALUATE": FindingSeverityEvaluate,
	"ADVISORY": FindingSeverityAdvisory,
	"PASS":     FindingSeverityPass,
}

// GetFindingSeverityEnumValues Enumerates the set of values for FindingSeverityEnum
func GetFindingSeverityEnumValues() []FindingSeverityEnum {
	values := make([]FindingSeverityEnum, 0)
	for _, v := range mappingFindingSeverity {
		values = append(values, v)
	}
	return values
}
