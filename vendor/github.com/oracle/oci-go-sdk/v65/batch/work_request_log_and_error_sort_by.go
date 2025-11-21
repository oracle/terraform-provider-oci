// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"strings"
)

// WorkRequestLogAndErrorSortByEnum Enum with underlying type: string
type WorkRequestLogAndErrorSortByEnum string

// Set of constants representing the allowable values for WorkRequestLogAndErrorSortByEnum
const (
	WorkRequestLogAndErrorSortByTimeCreated WorkRequestLogAndErrorSortByEnum = "timeCreated"
)

var mappingWorkRequestLogAndErrorSortByEnum = map[string]WorkRequestLogAndErrorSortByEnum{
	"timeCreated": WorkRequestLogAndErrorSortByTimeCreated,
}

var mappingWorkRequestLogAndErrorSortByEnumLowerCase = map[string]WorkRequestLogAndErrorSortByEnum{
	"timecreated": WorkRequestLogAndErrorSortByTimeCreated,
}

// GetWorkRequestLogAndErrorSortByEnumValues Enumerates the set of values for WorkRequestLogAndErrorSortByEnum
func GetWorkRequestLogAndErrorSortByEnumValues() []WorkRequestLogAndErrorSortByEnum {
	values := make([]WorkRequestLogAndErrorSortByEnum, 0)
	for _, v := range mappingWorkRequestLogAndErrorSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestLogAndErrorSortByEnumStringValues Enumerates the set of values in String for WorkRequestLogAndErrorSortByEnum
func GetWorkRequestLogAndErrorSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingWorkRequestLogAndErrorSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestLogAndErrorSortByEnum(val string) (WorkRequestLogAndErrorSortByEnum, bool) {
	enum, ok := mappingWorkRequestLogAndErrorSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
