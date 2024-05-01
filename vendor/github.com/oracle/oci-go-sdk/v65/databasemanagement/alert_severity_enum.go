// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// AlertSeverityEnumEnum Enum with underlying type: string
type AlertSeverityEnumEnum string

// Set of constants representing the allowable values for AlertSeverityEnumEnum
const (
	AlertSeverityEnumClear    AlertSeverityEnumEnum = "CLEAR"
	AlertSeverityEnumInfo     AlertSeverityEnumEnum = "INFO"
	AlertSeverityEnumWarning  AlertSeverityEnumEnum = "WARNING"
	AlertSeverityEnumCritical AlertSeverityEnumEnum = "CRITICAL"
)

var mappingAlertSeverityEnumEnum = map[string]AlertSeverityEnumEnum{
	"CLEAR":    AlertSeverityEnumClear,
	"INFO":     AlertSeverityEnumInfo,
	"WARNING":  AlertSeverityEnumWarning,
	"CRITICAL": AlertSeverityEnumCritical,
}

var mappingAlertSeverityEnumEnumLowerCase = map[string]AlertSeverityEnumEnum{
	"clear":    AlertSeverityEnumClear,
	"info":     AlertSeverityEnumInfo,
	"warning":  AlertSeverityEnumWarning,
	"critical": AlertSeverityEnumCritical,
}

// GetAlertSeverityEnumEnumValues Enumerates the set of values for AlertSeverityEnumEnum
func GetAlertSeverityEnumEnumValues() []AlertSeverityEnumEnum {
	values := make([]AlertSeverityEnumEnum, 0)
	for _, v := range mappingAlertSeverityEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertSeverityEnumEnumStringValues Enumerates the set of values in String for AlertSeverityEnumEnum
func GetAlertSeverityEnumEnumStringValues() []string {
	return []string{
		"CLEAR",
		"INFO",
		"WARNING",
		"CRITICAL",
	}
}

// GetMappingAlertSeverityEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertSeverityEnumEnum(val string) (AlertSeverityEnumEnum, bool) {
	enum, ok := mappingAlertSeverityEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
