// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Discovery and Monitoring Control API
//
// Use the Resource Discovery and Monitoring Control API to get details about monitored instances and perform actions. For more information, see Resource Discovery and Monitoring (https://docs.oracle.com/iaas/os-management/osms/osms-resource-discovery-monitoring.htm).
//

package appmgmtcontrol

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeActivateResourceMonitoringPlugin OperationTypeEnum = "ACTIVATE_RESOURCE_MONITORING_PLUGIN"
	OperationTypePublishTopProcessesMetrics       OperationTypeEnum = "PUBLISH_TOP_PROCESSES_METRICS"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"ACTIVATE_RESOURCE_MONITORING_PLUGIN": OperationTypeActivateResourceMonitoringPlugin,
	"PUBLISH_TOP_PROCESSES_METRICS":       OperationTypePublishTopProcessesMetrics,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"activate_resource_monitoring_plugin": OperationTypeActivateResourceMonitoringPlugin,
	"publish_top_processes_metrics":       OperationTypePublishTopProcessesMetrics,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"ACTIVATE_RESOURCE_MONITORING_PLUGIN",
		"PUBLISH_TOP_PROCESSES_METRICS",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
