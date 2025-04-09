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

// RequestQueueActionModeEnum Enum with underlying type: string
type RequestQueueActionModeEnum string

// Set of constants representing the allowable values for RequestQueueActionModeEnum
const (
	RequestQueueActionModeRp RequestQueueActionModeEnum = "RP"
	RequestQueueActionModeSp RequestQueueActionModeEnum = "SP"
	RequestQueueActionModeUp RequestQueueActionModeEnum = "UP"
)

var mappingRequestQueueActionModeEnum = map[string]RequestQueueActionModeEnum{
	"RP": RequestQueueActionModeRp,
	"SP": RequestQueueActionModeSp,
	"UP": RequestQueueActionModeUp,
}

var mappingRequestQueueActionModeEnumLowerCase = map[string]RequestQueueActionModeEnum{
	"rp": RequestQueueActionModeRp,
	"sp": RequestQueueActionModeSp,
	"up": RequestQueueActionModeUp,
}

// GetRequestQueueActionModeEnumValues Enumerates the set of values for RequestQueueActionModeEnum
func GetRequestQueueActionModeEnumValues() []RequestQueueActionModeEnum {
	values := make([]RequestQueueActionModeEnum, 0)
	for _, v := range mappingRequestQueueActionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestQueueActionModeEnumStringValues Enumerates the set of values in String for RequestQueueActionModeEnum
func GetRequestQueueActionModeEnumStringValues() []string {
	return []string{
		"RP",
		"SP",
		"UP",
	}
}

// GetMappingRequestQueueActionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestQueueActionModeEnum(val string) (RequestQueueActionModeEnum, bool) {
	enum, ok := mappingRequestQueueActionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
