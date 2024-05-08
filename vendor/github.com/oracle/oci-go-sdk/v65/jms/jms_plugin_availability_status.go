// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// JmsPluginAvailabilityStatusEnum Enum with underlying type: string
type JmsPluginAvailabilityStatusEnum string

// Set of constants representing the allowable values for JmsPluginAvailabilityStatusEnum
const (
	JmsPluginAvailabilityStatusActive       JmsPluginAvailabilityStatusEnum = "ACTIVE"
	JmsPluginAvailabilityStatusSilent       JmsPluginAvailabilityStatusEnum = "SILENT"
	JmsPluginAvailabilityStatusNotAvailable JmsPluginAvailabilityStatusEnum = "NOT_AVAILABLE"
)

var mappingJmsPluginAvailabilityStatusEnum = map[string]JmsPluginAvailabilityStatusEnum{
	"ACTIVE":        JmsPluginAvailabilityStatusActive,
	"SILENT":        JmsPluginAvailabilityStatusSilent,
	"NOT_AVAILABLE": JmsPluginAvailabilityStatusNotAvailable,
}

var mappingJmsPluginAvailabilityStatusEnumLowerCase = map[string]JmsPluginAvailabilityStatusEnum{
	"active":        JmsPluginAvailabilityStatusActive,
	"silent":        JmsPluginAvailabilityStatusSilent,
	"not_available": JmsPluginAvailabilityStatusNotAvailable,
}

// GetJmsPluginAvailabilityStatusEnumValues Enumerates the set of values for JmsPluginAvailabilityStatusEnum
func GetJmsPluginAvailabilityStatusEnumValues() []JmsPluginAvailabilityStatusEnum {
	values := make([]JmsPluginAvailabilityStatusEnum, 0)
	for _, v := range mappingJmsPluginAvailabilityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetJmsPluginAvailabilityStatusEnumStringValues Enumerates the set of values in String for JmsPluginAvailabilityStatusEnum
func GetJmsPluginAvailabilityStatusEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"SILENT",
		"NOT_AVAILABLE",
	}
}

// GetMappingJmsPluginAvailabilityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJmsPluginAvailabilityStatusEnum(val string) (JmsPluginAvailabilityStatusEnum, bool) {
	enum, ok := mappingJmsPluginAvailabilityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
