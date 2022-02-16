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

// MonitorTypesEnum Enum with underlying type: string
type MonitorTypesEnum string

// Set of constants representing the allowable values for MonitorTypesEnum
const (
	MonitorTypesScriptedBrowser MonitorTypesEnum = "SCRIPTED_BROWSER"
	MonitorTypesBrowser         MonitorTypesEnum = "BROWSER"
	MonitorTypesScriptedRest    MonitorTypesEnum = "SCRIPTED_REST"
	MonitorTypesRest            MonitorTypesEnum = "REST"
)

var mappingMonitorTypesEnum = map[string]MonitorTypesEnum{
	"SCRIPTED_BROWSER": MonitorTypesScriptedBrowser,
	"BROWSER":          MonitorTypesBrowser,
	"SCRIPTED_REST":    MonitorTypesScriptedRest,
	"REST":             MonitorTypesRest,
}

// GetMonitorTypesEnumValues Enumerates the set of values for MonitorTypesEnum
func GetMonitorTypesEnumValues() []MonitorTypesEnum {
	values := make([]MonitorTypesEnum, 0)
	for _, v := range mappingMonitorTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitorTypesEnumStringValues Enumerates the set of values in String for MonitorTypesEnum
func GetMonitorTypesEnumStringValues() []string {
	return []string{
		"SCRIPTED_BROWSER",
		"BROWSER",
		"SCRIPTED_REST",
		"REST",
	}
}

// GetMappingMonitorTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitorTypesEnum(val string) (MonitorTypesEnum, bool) {
	mappingMonitorTypesEnumIgnoreCase := make(map[string]MonitorTypesEnum)
	for k, v := range mappingMonitorTypesEnum {
		mappingMonitorTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMonitorTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
