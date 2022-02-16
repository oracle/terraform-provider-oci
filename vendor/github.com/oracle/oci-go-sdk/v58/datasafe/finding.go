// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Finding) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFindingSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetFindingSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingFindingSeverityEnum = map[string]FindingSeverityEnum{
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
	for _, v := range mappingFindingSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetFindingSeverityEnumStringValues Enumerates the set of values in String for FindingSeverityEnum
func GetFindingSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
	}
}

// GetMappingFindingSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFindingSeverityEnum(val string) (FindingSeverityEnum, bool) {
	mappingFindingSeverityEnumIgnoreCase := make(map[string]FindingSeverityEnum)
	for k, v := range mappingFindingSeverityEnum {
		mappingFindingSeverityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFindingSeverityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
