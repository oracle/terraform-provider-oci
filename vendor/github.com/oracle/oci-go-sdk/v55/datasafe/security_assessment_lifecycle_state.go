// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// SecurityAssessmentLifecycleStateEnum Enum with underlying type: string
type SecurityAssessmentLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityAssessmentLifecycleStateEnum
const (
	SecurityAssessmentLifecycleStateCreating  SecurityAssessmentLifecycleStateEnum = "CREATING"
	SecurityAssessmentLifecycleStateSucceeded SecurityAssessmentLifecycleStateEnum = "SUCCEEDED"
	SecurityAssessmentLifecycleStateUpdating  SecurityAssessmentLifecycleStateEnum = "UPDATING"
	SecurityAssessmentLifecycleStateDeleting  SecurityAssessmentLifecycleStateEnum = "DELETING"
	SecurityAssessmentLifecycleStateFailed    SecurityAssessmentLifecycleStateEnum = "FAILED"
)

var mappingSecurityAssessmentLifecycleState = map[string]SecurityAssessmentLifecycleStateEnum{
	"CREATING":  SecurityAssessmentLifecycleStateCreating,
	"SUCCEEDED": SecurityAssessmentLifecycleStateSucceeded,
	"UPDATING":  SecurityAssessmentLifecycleStateUpdating,
	"DELETING":  SecurityAssessmentLifecycleStateDeleting,
	"FAILED":    SecurityAssessmentLifecycleStateFailed,
}

// GetSecurityAssessmentLifecycleStateEnumValues Enumerates the set of values for SecurityAssessmentLifecycleStateEnum
func GetSecurityAssessmentLifecycleStateEnumValues() []SecurityAssessmentLifecycleStateEnum {
	values := make([]SecurityAssessmentLifecycleStateEnum, 0)
	for _, v := range mappingSecurityAssessmentLifecycleState {
		values = append(values, v)
	}
	return values
}
