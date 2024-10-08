// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MaintenanceWindowOperationStatusEnum Enum with underlying type: string
type MaintenanceWindowOperationStatusEnum string

// Set of constants representing the allowable values for MaintenanceWindowOperationStatusEnum
const (
	MaintenanceWindowOperationStatusInProgress MaintenanceWindowOperationStatusEnum = "IN_PROGRESS"
	MaintenanceWindowOperationStatusFailed     MaintenanceWindowOperationStatusEnum = "FAILED"
	MaintenanceWindowOperationStatusSucceeded  MaintenanceWindowOperationStatusEnum = "SUCCEEDED"
)

var mappingMaintenanceWindowOperationStatusEnum = map[string]MaintenanceWindowOperationStatusEnum{
	"IN_PROGRESS": MaintenanceWindowOperationStatusInProgress,
	"FAILED":      MaintenanceWindowOperationStatusFailed,
	"SUCCEEDED":   MaintenanceWindowOperationStatusSucceeded,
}

var mappingMaintenanceWindowOperationStatusEnumLowerCase = map[string]MaintenanceWindowOperationStatusEnum{
	"in_progress": MaintenanceWindowOperationStatusInProgress,
	"failed":      MaintenanceWindowOperationStatusFailed,
	"succeeded":   MaintenanceWindowOperationStatusSucceeded,
}

// GetMaintenanceWindowOperationStatusEnumValues Enumerates the set of values for MaintenanceWindowOperationStatusEnum
func GetMaintenanceWindowOperationStatusEnumValues() []MaintenanceWindowOperationStatusEnum {
	values := make([]MaintenanceWindowOperationStatusEnum, 0)
	for _, v := range mappingMaintenanceWindowOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceWindowOperationStatusEnumStringValues Enumerates the set of values in String for MaintenanceWindowOperationStatusEnum
func GetMaintenanceWindowOperationStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

// GetMappingMaintenanceWindowOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceWindowOperationStatusEnum(val string) (MaintenanceWindowOperationStatusEnum, bool) {
	enum, ok := mappingMaintenanceWindowOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
