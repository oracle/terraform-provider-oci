// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// SecurityAssessmentLifecycleStateEnum Enum with underlying type: string
type SecurityAssessmentLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityAssessmentLifecycleStateEnum
const (
	SecurityAssessmentLifecycleStateCreating  SecurityAssessmentLifecycleStateEnum = "CREATING"
	SecurityAssessmentLifecycleStateSucceeded SecurityAssessmentLifecycleStateEnum = "SUCCEEDED"
	SecurityAssessmentLifecycleStateUpdating  SecurityAssessmentLifecycleStateEnum = "UPDATING"
	SecurityAssessmentLifecycleStateDeleting  SecurityAssessmentLifecycleStateEnum = "DELETING"
	SecurityAssessmentLifecycleStateDeleted   SecurityAssessmentLifecycleStateEnum = "DELETED"
	SecurityAssessmentLifecycleStateFailed    SecurityAssessmentLifecycleStateEnum = "FAILED"
)

var mappingSecurityAssessmentLifecycleStateEnum = map[string]SecurityAssessmentLifecycleStateEnum{
	"CREATING":  SecurityAssessmentLifecycleStateCreating,
	"SUCCEEDED": SecurityAssessmentLifecycleStateSucceeded,
	"UPDATING":  SecurityAssessmentLifecycleStateUpdating,
	"DELETING":  SecurityAssessmentLifecycleStateDeleting,
	"DELETED":   SecurityAssessmentLifecycleStateDeleted,
	"FAILED":    SecurityAssessmentLifecycleStateFailed,
}

var mappingSecurityAssessmentLifecycleStateEnumLowerCase = map[string]SecurityAssessmentLifecycleStateEnum{
	"creating":  SecurityAssessmentLifecycleStateCreating,
	"succeeded": SecurityAssessmentLifecycleStateSucceeded,
	"updating":  SecurityAssessmentLifecycleStateUpdating,
	"deleting":  SecurityAssessmentLifecycleStateDeleting,
	"deleted":   SecurityAssessmentLifecycleStateDeleted,
	"failed":    SecurityAssessmentLifecycleStateFailed,
}

// GetSecurityAssessmentLifecycleStateEnumValues Enumerates the set of values for SecurityAssessmentLifecycleStateEnum
func GetSecurityAssessmentLifecycleStateEnumValues() []SecurityAssessmentLifecycleStateEnum {
	values := make([]SecurityAssessmentLifecycleStateEnum, 0)
	for _, v := range mappingSecurityAssessmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityAssessmentLifecycleStateEnum
func GetSecurityAssessmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SUCCEEDED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingSecurityAssessmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentLifecycleStateEnum(val string) (SecurityAssessmentLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityAssessmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
