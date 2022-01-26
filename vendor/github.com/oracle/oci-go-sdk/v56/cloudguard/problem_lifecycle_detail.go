// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

// ProblemLifecycleDetailEnum Enum with underlying type: string
type ProblemLifecycleDetailEnum string

// Set of constants representing the allowable values for ProblemLifecycleDetailEnum
const (
	ProblemLifecycleDetailOpen      ProblemLifecycleDetailEnum = "OPEN"
	ProblemLifecycleDetailResolved  ProblemLifecycleDetailEnum = "RESOLVED"
	ProblemLifecycleDetailDismissed ProblemLifecycleDetailEnum = "DISMISSED"
	ProblemLifecycleDetailDeleted   ProblemLifecycleDetailEnum = "DELETED"
)

var mappingProblemLifecycleDetail = map[string]ProblemLifecycleDetailEnum{
	"OPEN":      ProblemLifecycleDetailOpen,
	"RESOLVED":  ProblemLifecycleDetailResolved,
	"DISMISSED": ProblemLifecycleDetailDismissed,
	"DELETED":   ProblemLifecycleDetailDeleted,
}

// GetProblemLifecycleDetailEnumValues Enumerates the set of values for ProblemLifecycleDetailEnum
func GetProblemLifecycleDetailEnumValues() []ProblemLifecycleDetailEnum {
	values := make([]ProblemLifecycleDetailEnum, 0)
	for _, v := range mappingProblemLifecycleDetail {
		values = append(values, v)
	}
	return values
}
