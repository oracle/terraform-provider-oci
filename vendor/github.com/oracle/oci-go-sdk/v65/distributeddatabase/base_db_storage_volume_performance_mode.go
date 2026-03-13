// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"strings"
)

// BaseDbStorageVolumePerformanceModeEnum Enum with underlying type: string
type BaseDbStorageVolumePerformanceModeEnum string

// Set of constants representing the allowable values for BaseDbStorageVolumePerformanceModeEnum
const (
	BaseDbStorageVolumePerformanceModeBalanced        BaseDbStorageVolumePerformanceModeEnum = "BALANCED"
	BaseDbStorageVolumePerformanceModeHighPerformance BaseDbStorageVolumePerformanceModeEnum = "HIGH_PERFORMANCE"
)

var mappingBaseDbStorageVolumePerformanceModeEnum = map[string]BaseDbStorageVolumePerformanceModeEnum{
	"BALANCED":         BaseDbStorageVolumePerformanceModeBalanced,
	"HIGH_PERFORMANCE": BaseDbStorageVolumePerformanceModeHighPerformance,
}

var mappingBaseDbStorageVolumePerformanceModeEnumLowerCase = map[string]BaseDbStorageVolumePerformanceModeEnum{
	"balanced":         BaseDbStorageVolumePerformanceModeBalanced,
	"high_performance": BaseDbStorageVolumePerformanceModeHighPerformance,
}

// GetBaseDbStorageVolumePerformanceModeEnumValues Enumerates the set of values for BaseDbStorageVolumePerformanceModeEnum
func GetBaseDbStorageVolumePerformanceModeEnumValues() []BaseDbStorageVolumePerformanceModeEnum {
	values := make([]BaseDbStorageVolumePerformanceModeEnum, 0)
	for _, v := range mappingBaseDbStorageVolumePerformanceModeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseDbStorageVolumePerformanceModeEnumStringValues Enumerates the set of values in String for BaseDbStorageVolumePerformanceModeEnum
func GetBaseDbStorageVolumePerformanceModeEnumStringValues() []string {
	return []string{
		"BALANCED",
		"HIGH_PERFORMANCE",
	}
}

// GetMappingBaseDbStorageVolumePerformanceModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseDbStorageVolumePerformanceModeEnum(val string) (BaseDbStorageVolumePerformanceModeEnum, bool) {
	enum, ok := mappingBaseDbStorageVolumePerformanceModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
