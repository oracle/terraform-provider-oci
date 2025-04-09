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

// RequestQueueActionTargetEnum Enum with underlying type: string
type RequestQueueActionTargetEnum string

// Set of constants representing the allowable values for RequestQueueActionTargetEnum
const (
	RequestQueueActionTargetDbaas  RequestQueueActionTargetEnum = "DBAAS"
	RequestQueueActionTargetDbmReg RequestQueueActionTargetEnum = "DBM_REG"
)

var mappingRequestQueueActionTargetEnum = map[string]RequestQueueActionTargetEnum{
	"DBAAS":   RequestQueueActionTargetDbaas,
	"DBM_REG": RequestQueueActionTargetDbmReg,
}

var mappingRequestQueueActionTargetEnumLowerCase = map[string]RequestQueueActionTargetEnum{
	"dbaas":   RequestQueueActionTargetDbaas,
	"dbm_reg": RequestQueueActionTargetDbmReg,
}

// GetRequestQueueActionTargetEnumValues Enumerates the set of values for RequestQueueActionTargetEnum
func GetRequestQueueActionTargetEnumValues() []RequestQueueActionTargetEnum {
	values := make([]RequestQueueActionTargetEnum, 0)
	for _, v := range mappingRequestQueueActionTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestQueueActionTargetEnumStringValues Enumerates the set of values in String for RequestQueueActionTargetEnum
func GetRequestQueueActionTargetEnumStringValues() []string {
	return []string{
		"DBAAS",
		"DBM_REG",
	}
}

// GetMappingRequestQueueActionTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestQueueActionTargetEnum(val string) (RequestQueueActionTargetEnum, bool) {
	enum, ok := mappingRequestQueueActionTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
