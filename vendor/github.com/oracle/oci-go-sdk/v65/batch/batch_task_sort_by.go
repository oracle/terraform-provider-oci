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

// BatchTaskSortByEnum Enum with underlying type: string
type BatchTaskSortByEnum string

// Set of constants representing the allowable values for BatchTaskSortByEnum
const (
	BatchTaskSortByName BatchTaskSortByEnum = "name"
)

var mappingBatchTaskSortByEnum = map[string]BatchTaskSortByEnum{
	"name": BatchTaskSortByName,
}

var mappingBatchTaskSortByEnumLowerCase = map[string]BatchTaskSortByEnum{
	"name": BatchTaskSortByName,
}

// GetBatchTaskSortByEnumValues Enumerates the set of values for BatchTaskSortByEnum
func GetBatchTaskSortByEnumValues() []BatchTaskSortByEnum {
	values := make([]BatchTaskSortByEnum, 0)
	for _, v := range mappingBatchTaskSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskSortByEnumStringValues Enumerates the set of values in String for BatchTaskSortByEnum
func GetBatchTaskSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingBatchTaskSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskSortByEnum(val string) (BatchTaskSortByEnum, bool) {
	enum, ok := mappingBatchTaskSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
