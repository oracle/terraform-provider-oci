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

// UserAssessmentComparison Provides a list of differences for user assessment when compared with the baseline value.
type UserAssessmentComparison struct {

	// The current state of the user assessment comparison.
	LifecycleState UserAssessmentComparisonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the user assessment comparison was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	Summary []interface{} `mandatory:"false" json:"summary"`
}

func (m UserAssessmentComparison) String() string {
	return common.PointerString(m)
}

// UserAssessmentComparisonLifecycleStateEnum Enum with underlying type: string
type UserAssessmentComparisonLifecycleStateEnum string

// Set of constants representing the allowable values for UserAssessmentComparisonLifecycleStateEnum
const (
	UserAssessmentComparisonLifecycleStateCreating  UserAssessmentComparisonLifecycleStateEnum = "CREATING"
	UserAssessmentComparisonLifecycleStateSucceeded UserAssessmentComparisonLifecycleStateEnum = "SUCCEEDED"
	UserAssessmentComparisonLifecycleStateFailed    UserAssessmentComparisonLifecycleStateEnum = "FAILED"
)

var mappingUserAssessmentComparisonLifecycleState = map[string]UserAssessmentComparisonLifecycleStateEnum{
	"CREATING":  UserAssessmentComparisonLifecycleStateCreating,
	"SUCCEEDED": UserAssessmentComparisonLifecycleStateSucceeded,
	"FAILED":    UserAssessmentComparisonLifecycleStateFailed,
}

// GetUserAssessmentComparisonLifecycleStateEnumValues Enumerates the set of values for UserAssessmentComparisonLifecycleStateEnum
func GetUserAssessmentComparisonLifecycleStateEnumValues() []UserAssessmentComparisonLifecycleStateEnum {
	values := make([]UserAssessmentComparisonLifecycleStateEnum, 0)
	for _, v := range mappingUserAssessmentComparisonLifecycleState {
		values = append(values, v)
	}
	return values
}
