// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// MaintenanceScopeEnum Enum with underlying type: string
type MaintenanceScopeEnum string

// Set of constants representing the allowable values for MaintenanceScopeEnum
const (
	MaintenanceScopeAll           MaintenanceScopeEnum = "ALL"
	MaintenanceScopeAllButPrimary MaintenanceScopeEnum = "ALL_BUT_PRIMARY"
	MaintenanceScopePrimaryOnly   MaintenanceScopeEnum = "PRIMARY_ONLY"
)

var mappingMaintenanceScopeEnum = map[string]MaintenanceScopeEnum{
	"ALL":             MaintenanceScopeAll,
	"ALL_BUT_PRIMARY": MaintenanceScopeAllButPrimary,
	"PRIMARY_ONLY":    MaintenanceScopePrimaryOnly,
}

var mappingMaintenanceScopeEnumLowerCase = map[string]MaintenanceScopeEnum{
	"all":             MaintenanceScopeAll,
	"all_but_primary": MaintenanceScopeAllButPrimary,
	"primary_only":    MaintenanceScopePrimaryOnly,
}

// GetMaintenanceScopeEnumValues Enumerates the set of values for MaintenanceScopeEnum
func GetMaintenanceScopeEnumValues() []MaintenanceScopeEnum {
	values := make([]MaintenanceScopeEnum, 0)
	for _, v := range mappingMaintenanceScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceScopeEnumStringValues Enumerates the set of values in String for MaintenanceScopeEnum
func GetMaintenanceScopeEnumStringValues() []string {
	return []string{
		"ALL",
		"ALL_BUT_PRIMARY",
		"PRIMARY_ONLY",
	}
}

// GetMappingMaintenanceScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceScopeEnum(val string) (MaintenanceScopeEnum, bool) {
	enum, ok := mappingMaintenanceScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
