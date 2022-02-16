// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"strings"
)

// MonitorStatusEnum Enum with underlying type: string
type MonitorStatusEnum string

// Set of constants representing the allowable values for MonitorStatusEnum
const (
	MonitorStatusEnabled  MonitorStatusEnum = "ENABLED"
	MonitorStatusDisabled MonitorStatusEnum = "DISABLED"
	MonitorStatusInvalid  MonitorStatusEnum = "INVALID"
)

var mappingMonitorStatusEnum = map[string]MonitorStatusEnum{
	"ENABLED":  MonitorStatusEnabled,
	"DISABLED": MonitorStatusDisabled,
	"INVALID":  MonitorStatusInvalid,
}

// GetMonitorStatusEnumValues Enumerates the set of values for MonitorStatusEnum
func GetMonitorStatusEnumValues() []MonitorStatusEnum {
	values := make([]MonitorStatusEnum, 0)
	for _, v := range mappingMonitorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitorStatusEnumStringValues Enumerates the set of values in String for MonitorStatusEnum
func GetMonitorStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"INVALID",
	}
}

// GetMappingMonitorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitorStatusEnum(val string) (MonitorStatusEnum, bool) {
	mappingMonitorStatusEnumIgnoreCase := make(map[string]MonitorStatusEnum)
	for k, v := range mappingMonitorStatusEnum {
		mappingMonitorStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMonitorStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
