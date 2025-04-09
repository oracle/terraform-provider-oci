// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// RequestQueueEntryTypeEnum Enum with underlying type: string
type RequestQueueEntryTypeEnum string

// Set of constants representing the allowable values for RequestQueueEntryTypeEnum
const (
	RequestQueueEntryTypeManual       RequestQueueEntryTypeEnum = "MANUAL"
	RequestQueueEntryTypeOnDemandAuto RequestQueueEntryTypeEnum = "ON_DEMAND_AUTO"
	RequestQueueEntryTypeSystemAuto   RequestQueueEntryTypeEnum = "SYSTEM_AUTO"
)

var mappingRequestQueueEntryTypeEnum = map[string]RequestQueueEntryTypeEnum{
	"MANUAL":         RequestQueueEntryTypeManual,
	"ON_DEMAND_AUTO": RequestQueueEntryTypeOnDemandAuto,
	"SYSTEM_AUTO":    RequestQueueEntryTypeSystemAuto,
}

var mappingRequestQueueEntryTypeEnumLowerCase = map[string]RequestQueueEntryTypeEnum{
	"manual":         RequestQueueEntryTypeManual,
	"on_demand_auto": RequestQueueEntryTypeOnDemandAuto,
	"system_auto":    RequestQueueEntryTypeSystemAuto,
}

// GetRequestQueueEntryTypeEnumValues Enumerates the set of values for RequestQueueEntryTypeEnum
func GetRequestQueueEntryTypeEnumValues() []RequestQueueEntryTypeEnum {
	values := make([]RequestQueueEntryTypeEnum, 0)
	for _, v := range mappingRequestQueueEntryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestQueueEntryTypeEnumStringValues Enumerates the set of values in String for RequestQueueEntryTypeEnum
func GetRequestQueueEntryTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"ON_DEMAND_AUTO",
		"SYSTEM_AUTO",
	}
}

// GetMappingRequestQueueEntryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestQueueEntryTypeEnum(val string) (RequestQueueEntryTypeEnum, bool) {
	enum, ok := mappingRequestQueueEntryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
