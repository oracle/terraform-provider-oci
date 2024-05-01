// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"strings"
)

// DeletionScheduleEnum Enum with underlying type: string
type DeletionScheduleEnum string

// Set of constants representing the allowable values for DeletionScheduleEnum
const (
	DeletionScheduleDeleteAfterRetentionPeriod DeletionScheduleEnum = "DELETE_AFTER_RETENTION_PERIOD"
	DeletionScheduleDeleteAfter72Hours         DeletionScheduleEnum = "DELETE_AFTER_72_HOURS"
)

var mappingDeletionScheduleEnum = map[string]DeletionScheduleEnum{
	"DELETE_AFTER_RETENTION_PERIOD": DeletionScheduleDeleteAfterRetentionPeriod,
	"DELETE_AFTER_72_HOURS":         DeletionScheduleDeleteAfter72Hours,
}

var mappingDeletionScheduleEnumLowerCase = map[string]DeletionScheduleEnum{
	"delete_after_retention_period": DeletionScheduleDeleteAfterRetentionPeriod,
	"delete_after_72_hours":         DeletionScheduleDeleteAfter72Hours,
}

// GetDeletionScheduleEnumValues Enumerates the set of values for DeletionScheduleEnum
func GetDeletionScheduleEnumValues() []DeletionScheduleEnum {
	values := make([]DeletionScheduleEnum, 0)
	for _, v := range mappingDeletionScheduleEnum {
		values = append(values, v)
	}
	return values
}

// GetDeletionScheduleEnumStringValues Enumerates the set of values in String for DeletionScheduleEnum
func GetDeletionScheduleEnumStringValues() []string {
	return []string{
		"DELETE_AFTER_RETENTION_PERIOD",
		"DELETE_AFTER_72_HOURS",
	}
}

// GetMappingDeletionScheduleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeletionScheduleEnum(val string) (DeletionScheduleEnum, bool) {
	enum, ok := mappingDeletionScheduleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
