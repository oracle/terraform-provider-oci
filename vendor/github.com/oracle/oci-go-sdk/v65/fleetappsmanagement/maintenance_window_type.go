// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// MaintenanceWindowTypeEnum Enum with underlying type: string
type MaintenanceWindowTypeEnum string

// Set of constants representing the allowable values for MaintenanceWindowTypeEnum
const (
	MaintenanceWindowTypeOpenEnded MaintenanceWindowTypeEnum = "OPEN_ENDED"
)

var mappingMaintenanceWindowTypeEnum = map[string]MaintenanceWindowTypeEnum{
	"OPEN_ENDED": MaintenanceWindowTypeOpenEnded,
}

var mappingMaintenanceWindowTypeEnumLowerCase = map[string]MaintenanceWindowTypeEnum{
	"open_ended": MaintenanceWindowTypeOpenEnded,
}

// GetMaintenanceWindowTypeEnumValues Enumerates the set of values for MaintenanceWindowTypeEnum
func GetMaintenanceWindowTypeEnumValues() []MaintenanceWindowTypeEnum {
	values := make([]MaintenanceWindowTypeEnum, 0)
	for _, v := range mappingMaintenanceWindowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowTypeEnumStringValues Enumerates the set of values in String for MaintenanceWindowTypeEnum
func GetMaintenanceWindowTypeEnumStringValues() []string {
	return []string{
		"OPEN_ENDED",
	}
}

// GetMappingMaintenanceWindowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowTypeEnum(val string) (MaintenanceWindowTypeEnum, bool) {
	enum, ok := mappingMaintenanceWindowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
