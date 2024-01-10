// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ReplicatPerformanceProfileEnum Enum with underlying type: string
type ReplicatPerformanceProfileEnum string

// Set of constants representing the allowable values for ReplicatPerformanceProfileEnum
const (
	ReplicatPerformanceProfileLow  ReplicatPerformanceProfileEnum = "LOW"
	ReplicatPerformanceProfileHigh ReplicatPerformanceProfileEnum = "HIGH"
)

var mappingReplicatPerformanceProfileEnum = map[string]ReplicatPerformanceProfileEnum{
	"LOW":  ReplicatPerformanceProfileLow,
	"HIGH": ReplicatPerformanceProfileHigh,
}

var mappingReplicatPerformanceProfileEnumLowerCase = map[string]ReplicatPerformanceProfileEnum{
	"low":  ReplicatPerformanceProfileLow,
	"high": ReplicatPerformanceProfileHigh,
}

// GetReplicatPerformanceProfileEnumValues Enumerates the set of values for ReplicatPerformanceProfileEnum
func GetReplicatPerformanceProfileEnumValues() []ReplicatPerformanceProfileEnum {
	values := make([]ReplicatPerformanceProfileEnum, 0)
	for _, v := range mappingReplicatPerformanceProfileEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicatPerformanceProfileEnumStringValues Enumerates the set of values in String for ReplicatPerformanceProfileEnum
func GetReplicatPerformanceProfileEnumStringValues() []string {
	return []string{
		"LOW",
		"HIGH",
	}
}

// GetMappingReplicatPerformanceProfileEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicatPerformanceProfileEnum(val string) (ReplicatPerformanceProfileEnum, bool) {
	enum, ok := mappingReplicatPerformanceProfileEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
