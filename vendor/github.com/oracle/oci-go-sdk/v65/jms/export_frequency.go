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

// ExportFrequencyEnum Enum with underlying type: string
type ExportFrequencyEnum string

// Set of constants representing the allowable values for ExportFrequencyEnum
const (
	ExportFrequencyDaily   ExportFrequencyEnum = "DAILY"
	ExportFrequencyWeekly  ExportFrequencyEnum = "WEEKLY"
	ExportFrequencyMonthly ExportFrequencyEnum = "MONTHLY"
)

var mappingExportFrequencyEnum = map[string]ExportFrequencyEnum{
	"DAILY":   ExportFrequencyDaily,
	"WEEKLY":  ExportFrequencyWeekly,
	"MONTHLY": ExportFrequencyMonthly,
}

var mappingExportFrequencyEnumLowerCase = map[string]ExportFrequencyEnum{
	"daily":   ExportFrequencyDaily,
	"weekly":  ExportFrequencyWeekly,
	"monthly": ExportFrequencyMonthly,
}

// GetExportFrequencyEnumValues Enumerates the set of values for ExportFrequencyEnum
func GetExportFrequencyEnumValues() []ExportFrequencyEnum {
	values := make([]ExportFrequencyEnum, 0)
	for _, v := range mappingExportFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetExportFrequencyEnumStringValues Enumerates the set of values in String for ExportFrequencyEnum
func GetExportFrequencyEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingExportFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportFrequencyEnum(val string) (ExportFrequencyEnum, bool) {
	enum, ok := mappingExportFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
