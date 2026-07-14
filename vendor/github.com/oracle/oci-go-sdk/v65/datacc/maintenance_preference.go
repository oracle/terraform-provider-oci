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

// MaintenancePreferenceEnum Enum with underlying type: string
type MaintenancePreferenceEnum string

// Set of constants representing the allowable values for MaintenancePreferenceEnum
const (
	MaintenancePreferenceNoPreference     MaintenancePreferenceEnum = "NO_PREFERENCE"
	MaintenancePreferenceCustomPreference MaintenancePreferenceEnum = "CUSTOM_PREFERENCE"
)

var mappingMaintenancePreferenceEnum = map[string]MaintenancePreferenceEnum{
	"NO_PREFERENCE":     MaintenancePreferenceNoPreference,
	"CUSTOM_PREFERENCE": MaintenancePreferenceCustomPreference,
}

var mappingMaintenancePreferenceEnumLowerCase = map[string]MaintenancePreferenceEnum{
	"no_preference":     MaintenancePreferenceNoPreference,
	"custom_preference": MaintenancePreferenceCustomPreference,
}

// GetMaintenancePreferenceEnumValues Enumerates the set of values for MaintenancePreferenceEnum
func GetMaintenancePreferenceEnumValues() []MaintenancePreferenceEnum {
	values := make([]MaintenancePreferenceEnum, 0)
	for _, v := range mappingMaintenancePreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenancePreferenceEnumStringValues Enumerates the set of values in String for MaintenancePreferenceEnum
func GetMaintenancePreferenceEnumStringValues() []string {
	return []string{
		"NO_PREFERENCE",
		"CUSTOM_PREFERENCE",
	}
}

// GetMappingMaintenancePreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenancePreferenceEnum(val string) (MaintenancePreferenceEnum, bool) {
	enum, ok := mappingMaintenancePreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
