// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MonitoringTemplateLifeCycleDetailsEnum Enum with underlying type: string
type MonitoringTemplateLifeCycleDetailsEnum string

// Set of constants representing the allowable values for MonitoringTemplateLifeCycleDetailsEnum
const (
	MonitoringTemplateLifeCycleDetailsNotApplied     MonitoringTemplateLifeCycleDetailsEnum = "NOT_APPLIED"
	MonitoringTemplateLifeCycleDetailsApplied        MonitoringTemplateLifeCycleDetailsEnum = "APPLIED"
	MonitoringTemplateLifeCycleDetailsPartialApplied MonitoringTemplateLifeCycleDetailsEnum = "PARTIAL_APPLIED"
)

var mappingMonitoringTemplateLifeCycleDetailsEnum = map[string]MonitoringTemplateLifeCycleDetailsEnum{
	"NOT_APPLIED":     MonitoringTemplateLifeCycleDetailsNotApplied,
	"APPLIED":         MonitoringTemplateLifeCycleDetailsApplied,
	"PARTIAL_APPLIED": MonitoringTemplateLifeCycleDetailsPartialApplied,
}

var mappingMonitoringTemplateLifeCycleDetailsEnumLowerCase = map[string]MonitoringTemplateLifeCycleDetailsEnum{
	"not_applied":     MonitoringTemplateLifeCycleDetailsNotApplied,
	"applied":         MonitoringTemplateLifeCycleDetailsApplied,
	"partial_applied": MonitoringTemplateLifeCycleDetailsPartialApplied,
}

// GetMonitoringTemplateLifeCycleDetailsEnumValues Enumerates the set of values for MonitoringTemplateLifeCycleDetailsEnum
func GetMonitoringTemplateLifeCycleDetailsEnumValues() []MonitoringTemplateLifeCycleDetailsEnum {
	values := make([]MonitoringTemplateLifeCycleDetailsEnum, 0)
	for _, v := range mappingMonitoringTemplateLifeCycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoringTemplateLifeCycleDetailsEnumStringValues Enumerates the set of values in String for MonitoringTemplateLifeCycleDetailsEnum
func GetMonitoringTemplateLifeCycleDetailsEnumStringValues() []string {
	return []string{
		"NOT_APPLIED",
		"APPLIED",
		"PARTIAL_APPLIED",
	}
}

// GetMappingMonitoringTemplateLifeCycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoringTemplateLifeCycleDetailsEnum(val string) (MonitoringTemplateLifeCycleDetailsEnum, bool) {
	enum, ok := mappingMonitoringTemplateLifeCycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
