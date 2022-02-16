// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// ProblemLifecycleStateEnum Enum with underlying type: string
type ProblemLifecycleStateEnum string

// Set of constants representing the allowable values for ProblemLifecycleStateEnum
const (
	ProblemLifecycleStateActive   ProblemLifecycleStateEnum = "ACTIVE"
	ProblemLifecycleStateInactive ProblemLifecycleStateEnum = "INACTIVE"
)

var mappingProblemLifecycleStateEnum = map[string]ProblemLifecycleStateEnum{
	"ACTIVE":   ProblemLifecycleStateActive,
	"INACTIVE": ProblemLifecycleStateInactive,
}

// GetProblemLifecycleStateEnumValues Enumerates the set of values for ProblemLifecycleStateEnum
func GetProblemLifecycleStateEnumValues() []ProblemLifecycleStateEnum {
	values := make([]ProblemLifecycleStateEnum, 0)
	for _, v := range mappingProblemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetProblemLifecycleStateEnumStringValues Enumerates the set of values in String for ProblemLifecycleStateEnum
func GetProblemLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingProblemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProblemLifecycleStateEnum(val string) (ProblemLifecycleStateEnum, bool) {
	mappingProblemLifecycleStateEnumIgnoreCase := make(map[string]ProblemLifecycleStateEnum)
	for k, v := range mappingProblemLifecycleStateEnum {
		mappingProblemLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProblemLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
