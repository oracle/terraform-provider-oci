// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// SummarizeManagedInstanceUsageFieldsEnum Enum with underlying type: string
type SummarizeManagedInstanceUsageFieldsEnum string

// Set of constants representing the allowable values for SummarizeManagedInstanceUsageFieldsEnum
const (
	SummarizeManagedInstanceUsageFieldsApproximateJreCount          SummarizeManagedInstanceUsageFieldsEnum = "approximateJreCount"
	SummarizeManagedInstanceUsageFieldsApproximateInstallationCount SummarizeManagedInstanceUsageFieldsEnum = "approximateInstallationCount"
	SummarizeManagedInstanceUsageFieldsApproximateApplicationCount  SummarizeManagedInstanceUsageFieldsEnum = "approximateApplicationCount"
)

var mappingSummarizeManagedInstanceUsageFieldsEnum = map[string]SummarizeManagedInstanceUsageFieldsEnum{
	"approximateJreCount":          SummarizeManagedInstanceUsageFieldsApproximateJreCount,
	"approximateInstallationCount": SummarizeManagedInstanceUsageFieldsApproximateInstallationCount,
	"approximateApplicationCount":  SummarizeManagedInstanceUsageFieldsApproximateApplicationCount,
}

// GetSummarizeManagedInstanceUsageFieldsEnumValues Enumerates the set of values for SummarizeManagedInstanceUsageFieldsEnum
func GetSummarizeManagedInstanceUsageFieldsEnumValues() []SummarizeManagedInstanceUsageFieldsEnum {
	values := make([]SummarizeManagedInstanceUsageFieldsEnum, 0)
	for _, v := range mappingSummarizeManagedInstanceUsageFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeManagedInstanceUsageFieldsEnumStringValues Enumerates the set of values in String for SummarizeManagedInstanceUsageFieldsEnum
func GetSummarizeManagedInstanceUsageFieldsEnumStringValues() []string {
	return []string{
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateApplicationCount",
	}
}

// GetMappingSummarizeManagedInstanceUsageFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeManagedInstanceUsageFieldsEnum(val string) (SummarizeManagedInstanceUsageFieldsEnum, bool) {
	mappingSummarizeManagedInstanceUsageFieldsEnumIgnoreCase := make(map[string]SummarizeManagedInstanceUsageFieldsEnum)
	for k, v := range mappingSummarizeManagedInstanceUsageFieldsEnum {
		mappingSummarizeManagedInstanceUsageFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeManagedInstanceUsageFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
