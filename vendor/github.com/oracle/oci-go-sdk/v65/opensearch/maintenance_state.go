// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"strings"
)

// MaintenanceStateEnum Enum with underlying type: string
type MaintenanceStateEnum string

// Set of constants representing the allowable values for MaintenanceStateEnum
const (
	MaintenanceStateScheduled     MaintenanceStateEnum = "SCHEDULED"
	MaintenanceStateRescheduled   MaintenanceStateEnum = "RESCHEDULED"
	MaintenanceStateToBeScheduled MaintenanceStateEnum = "TO_BE_SCHEDULED"
)

var mappingMaintenanceStateEnum = map[string]MaintenanceStateEnum{
	"SCHEDULED":       MaintenanceStateScheduled,
	"RESCHEDULED":     MaintenanceStateRescheduled,
	"TO_BE_SCHEDULED": MaintenanceStateToBeScheduled,
}

var mappingMaintenanceStateEnumLowerCase = map[string]MaintenanceStateEnum{
	"scheduled":       MaintenanceStateScheduled,
	"rescheduled":     MaintenanceStateRescheduled,
	"to_be_scheduled": MaintenanceStateToBeScheduled,
}

// GetMaintenanceStateEnumValues Enumerates the set of values for MaintenanceStateEnum
func GetMaintenanceStateEnumValues() []MaintenanceStateEnum {
	values := make([]MaintenanceStateEnum, 0)
	for _, v := range mappingMaintenanceStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMaintenanceStateEnumStringValues Enumerates the set of values in String for MaintenanceStateEnum
func GetMaintenanceStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RESCHEDULED",
		"TO_BE_SCHEDULED",
	}
}

// GetMappingMaintenanceStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaintenanceStateEnum(val string) (MaintenanceStateEnum, bool) {
	enum, ok := mappingMaintenanceStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
