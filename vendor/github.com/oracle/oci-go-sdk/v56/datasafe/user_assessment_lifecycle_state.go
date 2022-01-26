// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// UserAssessmentLifecycleStateEnum Enum with underlying type: string
type UserAssessmentLifecycleStateEnum string

// Set of constants representing the allowable values for UserAssessmentLifecycleStateEnum
const (
	UserAssessmentLifecycleStateCreating  UserAssessmentLifecycleStateEnum = "CREATING"
	UserAssessmentLifecycleStateSucceeded UserAssessmentLifecycleStateEnum = "SUCCEEDED"
	UserAssessmentLifecycleStateUpdating  UserAssessmentLifecycleStateEnum = "UPDATING"
	UserAssessmentLifecycleStateDeleting  UserAssessmentLifecycleStateEnum = "DELETING"
	UserAssessmentLifecycleStateFailed    UserAssessmentLifecycleStateEnum = "FAILED"
)

var mappingUserAssessmentLifecycleState = map[string]UserAssessmentLifecycleStateEnum{
	"CREATING":  UserAssessmentLifecycleStateCreating,
	"SUCCEEDED": UserAssessmentLifecycleStateSucceeded,
	"UPDATING":  UserAssessmentLifecycleStateUpdating,
	"DELETING":  UserAssessmentLifecycleStateDeleting,
	"FAILED":    UserAssessmentLifecycleStateFailed,
}

// GetUserAssessmentLifecycleStateEnumValues Enumerates the set of values for UserAssessmentLifecycleStateEnum
func GetUserAssessmentLifecycleStateEnumValues() []UserAssessmentLifecycleStateEnum {
	values := make([]UserAssessmentLifecycleStateEnum, 0)
	for _, v := range mappingUserAssessmentLifecycleState {
		values = append(values, v)
	}
	return values
}
