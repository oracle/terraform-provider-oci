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

// MonitorTypesEnum Enum with underlying type: string
type MonitorTypesEnum string

// Set of constants representing the allowable values for MonitorTypesEnum
const (
	MonitorTypesScriptedBrowser MonitorTypesEnum = "SCRIPTED_BROWSER"
	MonitorTypesBrowser         MonitorTypesEnum = "BROWSER"
	MonitorTypesScriptedRest    MonitorTypesEnum = "SCRIPTED_REST"
	MonitorTypesRest            MonitorTypesEnum = "REST"
	MonitorTypesNetwork         MonitorTypesEnum = "NETWORK"
	MonitorTypesDns             MonitorTypesEnum = "DNS"
	MonitorTypesFtp             MonitorTypesEnum = "FTP"
	MonitorTypesSql             MonitorTypesEnum = "SQL"
)

var mappingMonitorTypesEnum = map[string]MonitorTypesEnum{
	"SCRIPTED_BROWSER": MonitorTypesScriptedBrowser,
	"BROWSER":          MonitorTypesBrowser,
	"SCRIPTED_REST":    MonitorTypesScriptedRest,
	"REST":             MonitorTypesRest,
	"NETWORK":          MonitorTypesNetwork,
	"DNS":              MonitorTypesDns,
	"FTP":              MonitorTypesFtp,
	"SQL":              MonitorTypesSql,
}

var mappingMonitorTypesEnumLowerCase = map[string]MonitorTypesEnum{
	"scripted_browser": MonitorTypesScriptedBrowser,
	"browser":          MonitorTypesBrowser,
	"scripted_rest":    MonitorTypesScriptedRest,
	"rest":             MonitorTypesRest,
	"network":          MonitorTypesNetwork,
	"dns":              MonitorTypesDns,
	"ftp":              MonitorTypesFtp,
	"sql":              MonitorTypesSql,
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
		"NETWORK",
		"DNS",
		"FTP",
		"SQL",
	}
}

// GetMappingMonitorTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitorTypesEnum(val string) (MonitorTypesEnum, bool) {
	enum, ok := mappingMonitorTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
