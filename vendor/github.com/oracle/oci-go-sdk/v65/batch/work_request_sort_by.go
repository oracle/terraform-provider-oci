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

// WorkRequestSortByEnum Enum with underlying type: string
type WorkRequestSortByEnum string

// Set of constants representing the allowable values for WorkRequestSortByEnum
const (
	WorkRequestSortByTimeAccepted WorkRequestSortByEnum = "timeAccepted"
)

var mappingWorkRequestSortByEnum = map[string]WorkRequestSortByEnum{
	"timeAccepted": WorkRequestSortByTimeAccepted,
}

var mappingWorkRequestSortByEnumLowerCase = map[string]WorkRequestSortByEnum{
	"timeaccepted": WorkRequestSortByTimeAccepted,
}

// GetWorkRequestSortByEnumValues Enumerates the set of values for WorkRequestSortByEnum
func GetWorkRequestSortByEnumValues() []WorkRequestSortByEnum {
	values := make([]WorkRequestSortByEnum, 0)
	for _, v := range mappingWorkRequestSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSortByEnumStringValues Enumerates the set of values in String for WorkRequestSortByEnum
func GetWorkRequestSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingWorkRequestSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSortByEnum(val string) (WorkRequestSortByEnum, bool) {
	enum, ok := mappingWorkRequestSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
