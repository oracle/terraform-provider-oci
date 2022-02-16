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

// ProblemLifecycleDetailEnum Enum with underlying type: string
type ProblemLifecycleDetailEnum string

// Set of constants representing the allowable values for ProblemLifecycleDetailEnum
const (
	ProblemLifecycleDetailOpen      ProblemLifecycleDetailEnum = "OPEN"
	ProblemLifecycleDetailResolved  ProblemLifecycleDetailEnum = "RESOLVED"
	ProblemLifecycleDetailDismissed ProblemLifecycleDetailEnum = "DISMISSED"
	ProblemLifecycleDetailDeleted   ProblemLifecycleDetailEnum = "DELETED"
)

var mappingProblemLifecycleDetailEnum = map[string]ProblemLifecycleDetailEnum{
	"OPEN":      ProblemLifecycleDetailOpen,
	"RESOLVED":  ProblemLifecycleDetailResolved,
	"DISMISSED": ProblemLifecycleDetailDismissed,
	"DELETED":   ProblemLifecycleDetailDeleted,
}

// GetProblemLifecycleDetailEnumValues Enumerates the set of values for ProblemLifecycleDetailEnum
func GetProblemLifecycleDetailEnumValues() []ProblemLifecycleDetailEnum {
	values := make([]ProblemLifecycleDetailEnum, 0)
	for _, v := range mappingProblemLifecycleDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetProblemLifecycleDetailEnumStringValues Enumerates the set of values in String for ProblemLifecycleDetailEnum
func GetProblemLifecycleDetailEnumStringValues() []string {
	return []string{
		"OPEN",
		"RESOLVED",
		"DISMISSED",
		"DELETED",
	}
}

// GetMappingProblemLifecycleDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProblemLifecycleDetailEnum(val string) (ProblemLifecycleDetailEnum, bool) {
	mappingProblemLifecycleDetailEnumIgnoreCase := make(map[string]ProblemLifecycleDetailEnum)
	for k, v := range mappingProblemLifecycleDetailEnum {
		mappingProblemLifecycleDetailEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingProblemLifecycleDetailEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
