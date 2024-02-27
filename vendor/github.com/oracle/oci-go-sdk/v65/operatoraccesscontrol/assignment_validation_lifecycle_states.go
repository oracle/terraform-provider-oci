// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"strings"
)

// AssignmentValidationLifecycleStatesEnum Enum with underlying type: string
type AssignmentValidationLifecycleStatesEnum string

// Set of constants representing the allowable values for AssignmentValidationLifecycleStatesEnum
const (
	AssignmentValidationLifecycleStatesProgress AssignmentValidationLifecycleStatesEnum = "PROGRESS"
	AssignmentValidationLifecycleStatesSuccess  AssignmentValidationLifecycleStatesEnum = "SUCCESS"
	AssignmentValidationLifecycleStatesFailed   AssignmentValidationLifecycleStatesEnum = "FAILED"
)

var mappingAssignmentValidationLifecycleStatesEnum = map[string]AssignmentValidationLifecycleStatesEnum{
	"PROGRESS": AssignmentValidationLifecycleStatesProgress,
	"SUCCESS":  AssignmentValidationLifecycleStatesSuccess,
	"FAILED":   AssignmentValidationLifecycleStatesFailed,
}

var mappingAssignmentValidationLifecycleStatesEnumLowerCase = map[string]AssignmentValidationLifecycleStatesEnum{
	"progress": AssignmentValidationLifecycleStatesProgress,
	"success":  AssignmentValidationLifecycleStatesSuccess,
	"failed":   AssignmentValidationLifecycleStatesFailed,
}

// GetAssignmentValidationLifecycleStatesEnumValues Enumerates the set of values for AssignmentValidationLifecycleStatesEnum
func GetAssignmentValidationLifecycleStatesEnumValues() []AssignmentValidationLifecycleStatesEnum {
	values := make([]AssignmentValidationLifecycleStatesEnum, 0)
	for _, v := range mappingAssignmentValidationLifecycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAssignmentValidationLifecycleStatesEnumStringValues Enumerates the set of values in String for AssignmentValidationLifecycleStatesEnum
func GetAssignmentValidationLifecycleStatesEnumStringValues() []string {
	return []string{
		"PROGRESS",
		"SUCCESS",
		"FAILED",
	}
}

// GetMappingAssignmentValidationLifecycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssignmentValidationLifecycleStatesEnum(val string) (AssignmentValidationLifecycleStatesEnum, bool) {
	enum, ok := mappingAssignmentValidationLifecycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
