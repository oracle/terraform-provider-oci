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

// SecurityAssessmentComparison Provides a list of the differences in a comparison of the security assessment with the baseline value.
type SecurityAssessmentComparison struct {

	// The current state of the security assessment comparison.
	LifecycleState SecurityAssessmentComparisonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the security assessment comparison was created. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the security assessment that is being compared with a baseline security assessment.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the security assessment that is set as a baseline.
	BaselineId *string `mandatory:"false" json:"baselineId"`

	// A target-based comparison between two security assessments.
	Targets []SecurityAssessmentComparisonPerTarget `mandatory:"false" json:"targets"`
}

func (m SecurityAssessmentComparison) String() string {
	return common.PointerString(m)
}

// SecurityAssessmentComparisonLifecycleStateEnum Enum with underlying type: string
type SecurityAssessmentComparisonLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityAssessmentComparisonLifecycleStateEnum
const (
	SecurityAssessmentComparisonLifecycleStateCreating  SecurityAssessmentComparisonLifecycleStateEnum = "CREATING"
	SecurityAssessmentComparisonLifecycleStateSucceeded SecurityAssessmentComparisonLifecycleStateEnum = "SUCCEEDED"
	SecurityAssessmentComparisonLifecycleStateFailed    SecurityAssessmentComparisonLifecycleStateEnum = "FAILED"
)

var mappingSecurityAssessmentComparisonLifecycleState = map[string]SecurityAssessmentComparisonLifecycleStateEnum{
	"CREATING":  SecurityAssessmentComparisonLifecycleStateCreating,
	"SUCCEEDED": SecurityAssessmentComparisonLifecycleStateSucceeded,
	"FAILED":    SecurityAssessmentComparisonLifecycleStateFailed,
}

// GetSecurityAssessmentComparisonLifecycleStateEnumValues Enumerates the set of values for SecurityAssessmentComparisonLifecycleStateEnum
func GetSecurityAssessmentComparisonLifecycleStateEnumValues() []SecurityAssessmentComparisonLifecycleStateEnum {
	values := make([]SecurityAssessmentComparisonLifecycleStateEnum, 0)
	for _, v := range mappingSecurityAssessmentComparisonLifecycleState {
		values = append(values, v)
	}
	return values
}
