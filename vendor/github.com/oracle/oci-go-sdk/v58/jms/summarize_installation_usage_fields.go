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

// SummarizeInstallationUsageFieldsEnum Enum with underlying type: string
type SummarizeInstallationUsageFieldsEnum string

// Set of constants representing the allowable values for SummarizeInstallationUsageFieldsEnum
const (
	SummarizeInstallationUsageFieldsApproximateApplicationCount     SummarizeInstallationUsageFieldsEnum = "approximateApplicationCount"
	SummarizeInstallationUsageFieldsApproximateManagedInstanceCount SummarizeInstallationUsageFieldsEnum = "approximateManagedInstanceCount"
)

var mappingSummarizeInstallationUsageFieldsEnum = map[string]SummarizeInstallationUsageFieldsEnum{
	"approximateApplicationCount":     SummarizeInstallationUsageFieldsApproximateApplicationCount,
	"approximateManagedInstanceCount": SummarizeInstallationUsageFieldsApproximateManagedInstanceCount,
}

// GetSummarizeInstallationUsageFieldsEnumValues Enumerates the set of values for SummarizeInstallationUsageFieldsEnum
func GetSummarizeInstallationUsageFieldsEnumValues() []SummarizeInstallationUsageFieldsEnum {
	values := make([]SummarizeInstallationUsageFieldsEnum, 0)
	for _, v := range mappingSummarizeInstallationUsageFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeInstallationUsageFieldsEnumStringValues Enumerates the set of values in String for SummarizeInstallationUsageFieldsEnum
func GetSummarizeInstallationUsageFieldsEnumStringValues() []string {
	return []string{
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
	}
}

// GetMappingSummarizeInstallationUsageFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeInstallationUsageFieldsEnum(val string) (SummarizeInstallationUsageFieldsEnum, bool) {
	mappingSummarizeInstallationUsageFieldsEnumIgnoreCase := make(map[string]SummarizeInstallationUsageFieldsEnum)
	for k, v := range mappingSummarizeInstallationUsageFieldsEnum {
		mappingSummarizeInstallationUsageFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeInstallationUsageFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
