// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MaintenanceTypeEnum Enum with underlying type: string
type MaintenanceTypeEnum string

// Set of constants representing the allowable values for MaintenanceTypeEnum
const (
	MaintenanceTypeAutomatic MaintenanceTypeEnum = "AUTOMATIC"
	MaintenanceTypeManual    MaintenanceTypeEnum = "MANUAL"
	MaintenanceTypeShape     MaintenanceTypeEnum = "SHAPE"
)

var mappingMaintenanceTypeEnum = map[string]MaintenanceTypeEnum{
	"AUTOMATIC": MaintenanceTypeAutomatic,
	"MANUAL":    MaintenanceTypeManual,
	"SHAPE":     MaintenanceTypeShape,
}

var mappingMaintenanceTypeEnumLowerCase = map[string]MaintenanceTypeEnum{
	"automatic": MaintenanceTypeAutomatic,
	"manual":    MaintenanceTypeManual,
	"shape":     MaintenanceTypeShape,
}

// GetMaintenanceTypeEnumValues Enumerates the set of values for MaintenanceTypeEnum
func GetMaintenanceTypeEnumValues() []MaintenanceTypeEnum {
	values := make([]MaintenanceTypeEnum, 0)
	for _, v := range mappingMaintenanceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceTypeEnumStringValues Enumerates the set of values in String for MaintenanceTypeEnum
func GetMaintenanceTypeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
		"SHAPE",
	}
}

// GetMappingMaintenanceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceTypeEnum(val string) (MaintenanceTypeEnum, bool) {
	enum, ok := mappingMaintenanceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
