// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// SummarizeJreUsageFieldsEnum Enum with underlying type: string
type SummarizeJreUsageFieldsEnum string

// Set of constants representing the allowable values for SummarizeJreUsageFieldsEnum
const (
	SummarizeJreUsageFieldsApproximateInstallationCount    SummarizeJreUsageFieldsEnum = "approximateInstallationCount"
	SummarizeJreUsageFieldsApproximateApplicationCount     SummarizeJreUsageFieldsEnum = "approximateApplicationCount"
	SummarizeJreUsageFieldsApproximateManagedInstanceCount SummarizeJreUsageFieldsEnum = "approximateManagedInstanceCount"
)

var mappingSummarizeJreUsageFieldsEnum = map[string]SummarizeJreUsageFieldsEnum{
	"approximateInstallationCount":    SummarizeJreUsageFieldsApproximateInstallationCount,
	"approximateApplicationCount":     SummarizeJreUsageFieldsApproximateApplicationCount,
	"approximateManagedInstanceCount": SummarizeJreUsageFieldsApproximateManagedInstanceCount,
}

var mappingSummarizeJreUsageFieldsEnumLowerCase = map[string]SummarizeJreUsageFieldsEnum{
	"approximateinstallationcount":    SummarizeJreUsageFieldsApproximateInstallationCount,
	"approximateapplicationcount":     SummarizeJreUsageFieldsApproximateApplicationCount,
	"approximatemanagedinstancecount": SummarizeJreUsageFieldsApproximateManagedInstanceCount,
}

// GetSummarizeJreUsageFieldsEnumValues Enumerates the set of values for SummarizeJreUsageFieldsEnum
func GetSummarizeJreUsageFieldsEnumValues() []SummarizeJreUsageFieldsEnum {
	values := make([]SummarizeJreUsageFieldsEnum, 0)
	for _, v := range mappingSummarizeJreUsageFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeJreUsageFieldsEnumStringValues Enumerates the set of values in String for SummarizeJreUsageFieldsEnum
func GetSummarizeJreUsageFieldsEnumStringValues() []string {
	return []string{
		"approximateInstallationCount",
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
	}
}

// GetMappingSummarizeJreUsageFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeJreUsageFieldsEnum(val string) (SummarizeJreUsageFieldsEnum, bool) {
	enum, ok := mappingSummarizeJreUsageFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
