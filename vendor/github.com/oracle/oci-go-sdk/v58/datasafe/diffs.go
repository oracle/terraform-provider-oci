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

// Diffs Results of the comparison of an item between two security assessments.
type Diffs struct {
	Current *Finding `mandatory:"false" json:"current"`

	Baseline *Finding `mandatory:"false" json:"baseline"`

	// This array identifies the items that are present in the baseline, but are missing from the current assessment.
	RemovedItems []string `mandatory:"false" json:"removedItems"`

	// This array identifies the items that are present in the current assessment, but are missing from the baseline.
	AddedItems []string `mandatory:"false" json:"addedItems"`

	// This array contains the items that are present in both the current assessment and the baseline, but are different in the two assessments.
	ModifiedItems []string `mandatory:"false" json:"modifiedItems"`

	// The severity of this diff.
	Severity DiffsSeverityEnum `mandatory:"false" json:"severity,omitempty"`
}

func (m Diffs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Diffs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiffsSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetDiffsSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiffsSeverityEnum Enum with underlying type: string
type DiffsSeverityEnum string

// Set of constants representing the allowable values for DiffsSeverityEnum
const (
	DiffsSeverityHigh     DiffsSeverityEnum = "HIGH"
	DiffsSeverityMedium   DiffsSeverityEnum = "MEDIUM"
	DiffsSeverityLow      DiffsSeverityEnum = "LOW"
	DiffsSeverityEvaluate DiffsSeverityEnum = "EVALUATE"
	DiffsSeverityAdvisory DiffsSeverityEnum = "ADVISORY"
	DiffsSeverityPass     DiffsSeverityEnum = "PASS"
)

var mappingDiffsSeverityEnum = map[string]DiffsSeverityEnum{
	"HIGH":     DiffsSeverityHigh,
	"MEDIUM":   DiffsSeverityMedium,
	"LOW":      DiffsSeverityLow,
	"EVALUATE": DiffsSeverityEvaluate,
	"ADVISORY": DiffsSeverityAdvisory,
	"PASS":     DiffsSeverityPass,
}

// GetDiffsSeverityEnumValues Enumerates the set of values for DiffsSeverityEnum
func GetDiffsSeverityEnumValues() []DiffsSeverityEnum {
	values := make([]DiffsSeverityEnum, 0)
	for _, v := range mappingDiffsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetDiffsSeverityEnumStringValues Enumerates the set of values in String for DiffsSeverityEnum
func GetDiffsSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
	}
}

// GetMappingDiffsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiffsSeverityEnum(val string) (DiffsSeverityEnum, bool) {
	mappingDiffsSeverityEnumIgnoreCase := make(map[string]DiffsSeverityEnum)
	for k, v := range mappingDiffsSeverityEnum {
		mappingDiffsSeverityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDiffsSeverityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
