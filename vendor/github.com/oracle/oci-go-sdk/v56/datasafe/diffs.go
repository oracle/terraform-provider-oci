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

var mappingDiffsSeverity = map[string]DiffsSeverityEnum{
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
	for _, v := range mappingDiffsSeverity {
		values = append(values, v)
	}
	return values
}
