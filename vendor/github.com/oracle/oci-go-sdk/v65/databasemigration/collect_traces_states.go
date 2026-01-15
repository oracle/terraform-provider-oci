// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// CollectTracesStatesEnum Enum with underlying type: string
type CollectTracesStatesEnum string

// Set of constants representing the allowable values for CollectTracesStatesEnum
const (
	CollectTracesStatesInProgress CollectTracesStatesEnum = "IN_PROGRESS"
	CollectTracesStatesSucceeded  CollectTracesStatesEnum = "SUCCEEDED"
	CollectTracesStatesFailed     CollectTracesStatesEnum = "FAILED"
)

var mappingCollectTracesStatesEnum = map[string]CollectTracesStatesEnum{
	"IN_PROGRESS": CollectTracesStatesInProgress,
	"SUCCEEDED":   CollectTracesStatesSucceeded,
	"FAILED":      CollectTracesStatesFailed,
}

var mappingCollectTracesStatesEnumLowerCase = map[string]CollectTracesStatesEnum{
	"in_progress": CollectTracesStatesInProgress,
	"succeeded":   CollectTracesStatesSucceeded,
	"failed":      CollectTracesStatesFailed,
}

// GetCollectTracesStatesEnumValues Enumerates the set of values for CollectTracesStatesEnum
func GetCollectTracesStatesEnumValues() []CollectTracesStatesEnum {
	values := make([]CollectTracesStatesEnum, 0)
	for _, v := range mappingCollectTracesStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetCollectTracesStatesEnumStringValues Enumerates the set of values in String for CollectTracesStatesEnum
func GetCollectTracesStatesEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingCollectTracesStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCollectTracesStatesEnum(val string) (CollectTracesStatesEnum, bool) {
	enum, ok := mappingCollectTracesStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
