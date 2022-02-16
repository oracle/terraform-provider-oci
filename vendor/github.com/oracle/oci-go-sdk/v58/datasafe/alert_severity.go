// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// AlertSeverityEnum Enum with underlying type: string
type AlertSeverityEnum string

// Set of constants representing the allowable values for AlertSeverityEnum
const (
	AlertSeverityCritical AlertSeverityEnum = "CRITICAL"
	AlertSeverityHigh     AlertSeverityEnum = "HIGH"
	AlertSeverityMedium   AlertSeverityEnum = "MEDIUM"
	AlertSeverityLow      AlertSeverityEnum = "LOW"
	AlertSeverityEvaluate AlertSeverityEnum = "EVALUATE"
)

var mappingAlertSeverityEnum = map[string]AlertSeverityEnum{
	"CRITICAL": AlertSeverityCritical,
	"HIGH":     AlertSeverityHigh,
	"MEDIUM":   AlertSeverityMedium,
	"LOW":      AlertSeverityLow,
	"EVALUATE": AlertSeverityEvaluate,
}

// GetAlertSeverityEnumValues Enumerates the set of values for AlertSeverityEnum
func GetAlertSeverityEnumValues() []AlertSeverityEnum {
	values := make([]AlertSeverityEnum, 0)
	for _, v := range mappingAlertSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAlertSeverityEnumStringValues Enumerates the set of values in String for AlertSeverityEnum
func GetAlertSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
	}
}

// GetMappingAlertSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlertSeverityEnum(val string) (AlertSeverityEnum, bool) {
	mappingAlertSeverityEnumIgnoreCase := make(map[string]AlertSeverityEnum)
	for k, v := range mappingAlertSeverityEnum {
		mappingAlertSeverityEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAlertSeverityEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
