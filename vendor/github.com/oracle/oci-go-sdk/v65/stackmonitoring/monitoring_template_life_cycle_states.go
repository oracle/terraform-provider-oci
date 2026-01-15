// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MonitoringTemplateLifeCycleStatesEnum Enum with underlying type: string
type MonitoringTemplateLifeCycleStatesEnum string

// Set of constants representing the allowable values for MonitoringTemplateLifeCycleStatesEnum
const (
	MonitoringTemplateLifeCycleStatesCreating MonitoringTemplateLifeCycleStatesEnum = "CREATING"
	MonitoringTemplateLifeCycleStatesActive   MonitoringTemplateLifeCycleStatesEnum = "ACTIVE"
	MonitoringTemplateLifeCycleStatesInactive MonitoringTemplateLifeCycleStatesEnum = "INACTIVE"
	MonitoringTemplateLifeCycleStatesUpdating MonitoringTemplateLifeCycleStatesEnum = "UPDATING"
	MonitoringTemplateLifeCycleStatesDeleted  MonitoringTemplateLifeCycleStatesEnum = "DELETED"
)

var mappingMonitoringTemplateLifeCycleStatesEnum = map[string]MonitoringTemplateLifeCycleStatesEnum{
	"CREATING": MonitoringTemplateLifeCycleStatesCreating,
	"ACTIVE":   MonitoringTemplateLifeCycleStatesActive,
	"INACTIVE": MonitoringTemplateLifeCycleStatesInactive,
	"UPDATING": MonitoringTemplateLifeCycleStatesUpdating,
	"DELETED":  MonitoringTemplateLifeCycleStatesDeleted,
}

var mappingMonitoringTemplateLifeCycleStatesEnumLowerCase = map[string]MonitoringTemplateLifeCycleStatesEnum{
	"creating": MonitoringTemplateLifeCycleStatesCreating,
	"active":   MonitoringTemplateLifeCycleStatesActive,
	"inactive": MonitoringTemplateLifeCycleStatesInactive,
	"updating": MonitoringTemplateLifeCycleStatesUpdating,
	"deleted":  MonitoringTemplateLifeCycleStatesDeleted,
}

// GetMonitoringTemplateLifeCycleStatesEnumValues Enumerates the set of values for MonitoringTemplateLifeCycleStatesEnum
func GetMonitoringTemplateLifeCycleStatesEnumValues() []MonitoringTemplateLifeCycleStatesEnum {
	values := make([]MonitoringTemplateLifeCycleStatesEnum, 0)
	for _, v := range mappingMonitoringTemplateLifeCycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoringTemplateLifeCycleStatesEnumStringValues Enumerates the set of values in String for MonitoringTemplateLifeCycleStatesEnum
func GetMonitoringTemplateLifeCycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingMonitoringTemplateLifeCycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoringTemplateLifeCycleStatesEnum(val string) (MonitoringTemplateLifeCycleStatesEnum, bool) {
	enum, ok := mappingMonitoringTemplateLifeCycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
