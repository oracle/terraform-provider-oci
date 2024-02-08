// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// ArchivalStateEnum Enum with underlying type: string
type ArchivalStateEnum string

// Set of constants representing the allowable values for ArchivalStateEnum
const (
	ArchivalStateArchived  ArchivalStateEnum = "ARCHIVED"
	ArchivalStateRestoring ArchivalStateEnum = "RESTORING"
	ArchivalStateRestored  ArchivalStateEnum = "RESTORED"
)

var mappingArchivalStateEnum = map[string]ArchivalStateEnum{
	"ARCHIVED":  ArchivalStateArchived,
	"RESTORING": ArchivalStateRestoring,
	"RESTORED":  ArchivalStateRestored,
}

var mappingArchivalStateEnumLowerCase = map[string]ArchivalStateEnum{
	"archived":  ArchivalStateArchived,
	"restoring": ArchivalStateRestoring,
	"restored":  ArchivalStateRestored,
}

// GetArchivalStateEnumValues Enumerates the set of values for ArchivalStateEnum
func GetArchivalStateEnumValues() []ArchivalStateEnum {
	values := make([]ArchivalStateEnum, 0)
	for _, v := range mappingArchivalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetArchivalStateEnumStringValues Enumerates the set of values in String for ArchivalStateEnum
func GetArchivalStateEnumStringValues() []string {
	return []string{
		"ARCHIVED",
		"RESTORING",
		"RESTORED",
	}
}

// GetMappingArchivalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchivalStateEnum(val string) (ArchivalStateEnum, bool) {
	enum, ok := mappingArchivalStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
