// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

// SummarizeJreUsageFieldsEnum Enum with underlying type: string
type SummarizeJreUsageFieldsEnum string

// Set of constants representing the allowable values for SummarizeJreUsageFieldsEnum
const (
	SummarizeJreUsageFieldsApproximateInstallationCount    SummarizeJreUsageFieldsEnum = "approximateInstallationCount"
	SummarizeJreUsageFieldsApproximateApplicationCount     SummarizeJreUsageFieldsEnum = "approximateApplicationCount"
	SummarizeJreUsageFieldsApproximateManagedInstanceCount SummarizeJreUsageFieldsEnum = "approximateManagedInstanceCount"
)

var mappingSummarizeJreUsageFields = map[string]SummarizeJreUsageFieldsEnum{
	"approximateInstallationCount":    SummarizeJreUsageFieldsApproximateInstallationCount,
	"approximateApplicationCount":     SummarizeJreUsageFieldsApproximateApplicationCount,
	"approximateManagedInstanceCount": SummarizeJreUsageFieldsApproximateManagedInstanceCount,
}

// GetSummarizeJreUsageFieldsEnumValues Enumerates the set of values for SummarizeJreUsageFieldsEnum
func GetSummarizeJreUsageFieldsEnumValues() []SummarizeJreUsageFieldsEnum {
	values := make([]SummarizeJreUsageFieldsEnum, 0)
	for _, v := range mappingSummarizeJreUsageFields {
		values = append(values, v)
	}
	return values
}
