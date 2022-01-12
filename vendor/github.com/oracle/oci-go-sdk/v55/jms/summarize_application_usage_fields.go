// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

// SummarizeApplicationUsageFieldsEnum Enum with underlying type: string
type SummarizeApplicationUsageFieldsEnum string

// Set of constants representing the allowable values for SummarizeApplicationUsageFieldsEnum
const (
	SummarizeApplicationUsageFieldsApproximateJreCount             SummarizeApplicationUsageFieldsEnum = "approximateJreCount"
	SummarizeApplicationUsageFieldsApproximateInstallationCount    SummarizeApplicationUsageFieldsEnum = "approximateInstallationCount"
	SummarizeApplicationUsageFieldsApproximateManagedInstanceCount SummarizeApplicationUsageFieldsEnum = "approximateManagedInstanceCount"
)

var mappingSummarizeApplicationUsageFields = map[string]SummarizeApplicationUsageFieldsEnum{
	"approximateJreCount":             SummarizeApplicationUsageFieldsApproximateJreCount,
	"approximateInstallationCount":    SummarizeApplicationUsageFieldsApproximateInstallationCount,
	"approximateManagedInstanceCount": SummarizeApplicationUsageFieldsApproximateManagedInstanceCount,
}

// GetSummarizeApplicationUsageFieldsEnumValues Enumerates the set of values for SummarizeApplicationUsageFieldsEnum
func GetSummarizeApplicationUsageFieldsEnumValues() []SummarizeApplicationUsageFieldsEnum {
	values := make([]SummarizeApplicationUsageFieldsEnum, 0)
	for _, v := range mappingSummarizeApplicationUsageFields {
		values = append(values, v)
	}
	return values
}
