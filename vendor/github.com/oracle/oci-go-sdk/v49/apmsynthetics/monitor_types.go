// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

// MonitorTypesEnum Enum with underlying type: string
type MonitorTypesEnum string

// Set of constants representing the allowable values for MonitorTypesEnum
const (
	MonitorTypesScriptedBrowser MonitorTypesEnum = "SCRIPTED_BROWSER"
	MonitorTypesBrowser         MonitorTypesEnum = "BROWSER"
	MonitorTypesScriptedRest    MonitorTypesEnum = "SCRIPTED_REST"
	MonitorTypesRest            MonitorTypesEnum = "REST"
)

var mappingMonitorTypes = map[string]MonitorTypesEnum{
	"SCRIPTED_BROWSER": MonitorTypesScriptedBrowser,
	"BROWSER":          MonitorTypesBrowser,
	"SCRIPTED_REST":    MonitorTypesScriptedRest,
	"REST":             MonitorTypesRest,
}

// GetMonitorTypesEnumValues Enumerates the set of values for MonitorTypesEnum
func GetMonitorTypesEnumValues() []MonitorTypesEnum {
	values := make([]MonitorTypesEnum, 0)
	for _, v := range mappingMonitorTypes {
		values = append(values, v)
	}
	return values
}
