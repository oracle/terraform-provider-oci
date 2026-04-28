// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"strings"
)

// MaintenanceTypeEnumEnum Enum with underlying type: string
type MaintenanceTypeEnumEnum string

// Set of constants representing the allowable values for MaintenanceTypeEnumEnum
const (
	MaintenanceTypeEnumPlanned   MaintenanceTypeEnumEnum = "PLANNED"
	MaintenanceTypeEnumUnplanned MaintenanceTypeEnumEnum = "UNPLANNED"
)

var mappingMaintenanceTypeEnumEnum = map[string]MaintenanceTypeEnumEnum{
	"PLANNED":   MaintenanceTypeEnumPlanned,
	"UNPLANNED": MaintenanceTypeEnumUnplanned,
}

var mappingMaintenanceTypeEnumEnumLowerCase = map[string]MaintenanceTypeEnumEnum{
	"planned":   MaintenanceTypeEnumPlanned,
	"unplanned": MaintenanceTypeEnumUnplanned,
}

// GetMaintenanceTypeEnumEnumValues Enumerates the set of values for MaintenanceTypeEnumEnum
func GetMaintenanceTypeEnumEnumValues() []MaintenanceTypeEnumEnum {
	values := make([]MaintenanceTypeEnumEnum, 0)
	for _, v := range mappingMaintenanceTypeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceTypeEnumEnumStringValues Enumerates the set of values in String for MaintenanceTypeEnumEnum
func GetMaintenanceTypeEnumEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingMaintenanceTypeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceTypeEnumEnum(val string) (MaintenanceTypeEnumEnum, bool) {
	enum, ok := mappingMaintenanceTypeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
