// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// ExtractPerformanceProfileEnum Enum with underlying type: string
type ExtractPerformanceProfileEnum string

// Set of constants representing the allowable values for ExtractPerformanceProfileEnum
const (
	ExtractPerformanceProfileLow    ExtractPerformanceProfileEnum = "LOW"
	ExtractPerformanceProfileMedium ExtractPerformanceProfileEnum = "MEDIUM"
	ExtractPerformanceProfileHigh   ExtractPerformanceProfileEnum = "HIGH"
)

var mappingExtractPerformanceProfile = map[string]ExtractPerformanceProfileEnum{
	"LOW":    ExtractPerformanceProfileLow,
	"MEDIUM": ExtractPerformanceProfileMedium,
	"HIGH":   ExtractPerformanceProfileHigh,
}

// GetExtractPerformanceProfileEnumValues Enumerates the set of values for ExtractPerformanceProfileEnum
func GetExtractPerformanceProfileEnumValues() []ExtractPerformanceProfileEnum {
	values := make([]ExtractPerformanceProfileEnum, 0)
	for _, v := range mappingExtractPerformanceProfile {
		values = append(values, v)
	}
	return values
}
