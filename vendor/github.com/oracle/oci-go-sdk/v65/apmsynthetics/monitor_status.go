// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
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

var mappingMonitorStatusEnumLowerCase = map[string]MonitorStatusEnum{
	"enabled":  MonitorStatusEnabled,
	"disabled": MonitorStatusDisabled,
	"invalid":  MonitorStatusInvalid,
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
	enum, ok := mappingMonitorStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
