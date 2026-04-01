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

// ControlledUpdateTargetDbInstancesEnum Enum with underlying type: string
type ControlledUpdateTargetDbInstancesEnum string

// Set of constants representing the allowable values for ControlledUpdateTargetDbInstancesEnum
const (
	ControlledUpdateTargetDbInstancesAllButPrimary ControlledUpdateTargetDbInstancesEnum = "ALL_BUT_PRIMARY"
	ControlledUpdateTargetDbInstancesPrimaryOnly   ControlledUpdateTargetDbInstancesEnum = "PRIMARY_ONLY"
)

var mappingControlledUpdateTargetDbInstancesEnum = map[string]ControlledUpdateTargetDbInstancesEnum{
	"ALL_BUT_PRIMARY": ControlledUpdateTargetDbInstancesAllButPrimary,
	"PRIMARY_ONLY":    ControlledUpdateTargetDbInstancesPrimaryOnly,
}

var mappingControlledUpdateTargetDbInstancesEnumLowerCase = map[string]ControlledUpdateTargetDbInstancesEnum{
	"all_but_primary": ControlledUpdateTargetDbInstancesAllButPrimary,
	"primary_only":    ControlledUpdateTargetDbInstancesPrimaryOnly,
}

// GetControlledUpdateTargetDbInstancesEnumValues Enumerates the set of values for ControlledUpdateTargetDbInstancesEnum
func GetControlledUpdateTargetDbInstancesEnumValues() []ControlledUpdateTargetDbInstancesEnum {
	values := make([]ControlledUpdateTargetDbInstancesEnum, 0)
	for _, v := range mappingControlledUpdateTargetDbInstancesEnum {
		values = append(values, v)
	}
	return values
}

// GetControlledUpdateTargetDbInstancesEnumStringValues Enumerates the set of values in String for ControlledUpdateTargetDbInstancesEnum
func GetControlledUpdateTargetDbInstancesEnumStringValues() []string {
	return []string{
		"ALL_BUT_PRIMARY",
		"PRIMARY_ONLY",
	}
}

// GetMappingControlledUpdateTargetDbInstancesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingControlledUpdateTargetDbInstancesEnum(val string) (ControlledUpdateTargetDbInstancesEnum, bool) {
	enum, ok := mappingControlledUpdateTargetDbInstancesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
