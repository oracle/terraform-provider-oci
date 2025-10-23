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

// MaintenanceActionEnum Enum with underlying type: string
type MaintenanceActionEnum string

// Set of constants representing the allowable values for MaintenanceActionEnum
const (
	MaintenanceActionDatabase MaintenanceActionEnum = "DATABASE"
	MaintenanceActionOsUpdate MaintenanceActionEnum = "OS_UPDATE"
)

var mappingMaintenanceActionEnum = map[string]MaintenanceActionEnum{
	"DATABASE":  MaintenanceActionDatabase,
	"OS_UPDATE": MaintenanceActionOsUpdate,
}

var mappingMaintenanceActionEnumLowerCase = map[string]MaintenanceActionEnum{
	"database":  MaintenanceActionDatabase,
	"os_update": MaintenanceActionOsUpdate,
}

// GetMaintenanceActionEnumValues Enumerates the set of values for MaintenanceActionEnum
func GetMaintenanceActionEnumValues() []MaintenanceActionEnum {
	values := make([]MaintenanceActionEnum, 0)
	for _, v := range mappingMaintenanceActionEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceActionEnumStringValues Enumerates the set of values in String for MaintenanceActionEnum
func GetMaintenanceActionEnumStringValues() []string {
	return []string{
		"DATABASE",
		"OS_UPDATE",
	}
}

// GetMappingMaintenanceActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceActionEnum(val string) (MaintenanceActionEnum, bool) {
	enum, ok := mappingMaintenanceActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
