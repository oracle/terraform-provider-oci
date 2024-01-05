// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidateLabelConditionResult The result of the label condition validation
type ValidateLabelConditionResult struct {

	// String representation of the validated label condition.
	ConditionString *string `mandatory:"true" json:"conditionString"`

	ConditionBlock *ConditionBlock `mandatory:"true" json:"conditionBlock"`

	// The validation status.
	Status *string `mandatory:"true" json:"status"`

	// Field values against which the label condition was evaluated.
	FieldValues []LogAnalyticsProperty `mandatory:"false" json:"fieldValues"`

	// The validation status description.
	StatusDescription *string `mandatory:"false" json:"statusDescription"`

	// The result of evaluating the condition blocks against the specified field values. Either true or false.
	EvaluationResult *bool `mandatory:"false" json:"evaluationResult"`
}

func (m ValidateLabelConditionResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidateLabelConditionResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
