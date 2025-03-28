// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ExportDurationEnum Enum with underlying type: string
type ExportDurationEnum string

// Set of constants representing the allowable values for ExportDurationEnum
const (
	ExportDurationLast30Days ExportDurationEnum = "LAST_30_DAYS"
	ExportDurationLast60Days ExportDurationEnum = "LAST_60_DAYS"
	ExportDurationLast90Days ExportDurationEnum = "LAST_90_DAYS"
)

var mappingExportDurationEnum = map[string]ExportDurationEnum{
	"LAST_30_DAYS": ExportDurationLast30Days,
	"LAST_60_DAYS": ExportDurationLast60Days,
	"LAST_90_DAYS": ExportDurationLast90Days,
}

var mappingExportDurationEnumLowerCase = map[string]ExportDurationEnum{
	"last_30_days": ExportDurationLast30Days,
	"last_60_days": ExportDurationLast60Days,
	"last_90_days": ExportDurationLast90Days,
}

// GetExportDurationEnumValues Enumerates the set of values for ExportDurationEnum
func GetExportDurationEnumValues() []ExportDurationEnum {
	values := make([]ExportDurationEnum, 0)
	for _, v := range mappingExportDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetExportDurationEnumStringValues Enumerates the set of values in String for ExportDurationEnum
func GetExportDurationEnumStringValues() []string {
	return []string{
		"LAST_30_DAYS",
		"LAST_60_DAYS",
		"LAST_90_DAYS",
	}
}

// GetMappingExportDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportDurationEnum(val string) (ExportDurationEnum, bool) {
	enum, ok := mappingExportDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
