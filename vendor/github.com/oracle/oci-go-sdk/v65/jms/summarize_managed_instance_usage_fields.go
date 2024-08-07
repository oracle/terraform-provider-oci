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

var mappingSummarizeManagedInstanceUsageFieldsEnumLowerCase = map[string]SummarizeManagedInstanceUsageFieldsEnum{
	"approximatejrecount":          SummarizeManagedInstanceUsageFieldsApproximateJreCount,
	"approximateinstallationcount": SummarizeManagedInstanceUsageFieldsApproximateInstallationCount,
	"approximateapplicationcount":  SummarizeManagedInstanceUsageFieldsApproximateApplicationCount,
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
	enum, ok := mappingSummarizeManagedInstanceUsageFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
