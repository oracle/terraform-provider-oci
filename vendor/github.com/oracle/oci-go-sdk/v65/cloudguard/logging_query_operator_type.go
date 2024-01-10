// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// LoggingQueryOperatorTypeEnum Enum with underlying type: string
type LoggingQueryOperatorTypeEnum string

// Set of constants representing the allowable values for LoggingQueryOperatorTypeEnum
const (
	LoggingQueryOperatorTypeEqual              LoggingQueryOperatorTypeEnum = "EQUAL"
	LoggingQueryOperatorTypeGreater            LoggingQueryOperatorTypeEnum = "GREATER"
	LoggingQueryOperatorTypeGreaterthanequalto LoggingQueryOperatorTypeEnum = "GREATERTHANEQUALTO"
	LoggingQueryOperatorTypeLess               LoggingQueryOperatorTypeEnum = "LESS"
	LoggingQueryOperatorTypeLessthanequalto    LoggingQueryOperatorTypeEnum = "LESSTHANEQUALTO"
)

var mappingLoggingQueryOperatorTypeEnum = map[string]LoggingQueryOperatorTypeEnum{
	"EQUAL":              LoggingQueryOperatorTypeEqual,
	"GREATER":            LoggingQueryOperatorTypeGreater,
	"GREATERTHANEQUALTO": LoggingQueryOperatorTypeGreaterthanequalto,
	"LESS":               LoggingQueryOperatorTypeLess,
	"LESSTHANEQUALTO":    LoggingQueryOperatorTypeLessthanequalto,
}

var mappingLoggingQueryOperatorTypeEnumLowerCase = map[string]LoggingQueryOperatorTypeEnum{
	"equal":              LoggingQueryOperatorTypeEqual,
	"greater":            LoggingQueryOperatorTypeGreater,
	"greaterthanequalto": LoggingQueryOperatorTypeGreaterthanequalto,
	"less":               LoggingQueryOperatorTypeLess,
	"lessthanequalto":    LoggingQueryOperatorTypeLessthanequalto,
}

// GetLoggingQueryOperatorTypeEnumValues Enumerates the set of values for LoggingQueryOperatorTypeEnum
func GetLoggingQueryOperatorTypeEnumValues() []LoggingQueryOperatorTypeEnum {
	values := make([]LoggingQueryOperatorTypeEnum, 0)
	for _, v := range mappingLoggingQueryOperatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoggingQueryOperatorTypeEnumStringValues Enumerates the set of values in String for LoggingQueryOperatorTypeEnum
func GetLoggingQueryOperatorTypeEnumStringValues() []string {
	return []string{
		"EQUAL",
		"GREATER",
		"GREATERTHANEQUALTO",
		"LESS",
		"LESSTHANEQUALTO",
	}
}

// GetMappingLoggingQueryOperatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoggingQueryOperatorTypeEnum(val string) (LoggingQueryOperatorTypeEnum, bool) {
	enum, ok := mappingLoggingQueryOperatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
