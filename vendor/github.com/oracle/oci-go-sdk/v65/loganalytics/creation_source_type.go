// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// CreationSourceTypeEnum Enum with underlying type: string
type CreationSourceTypeEnum string

// Set of constants representing the allowable values for CreationSourceTypeEnum
const (
	CreationSourceTypeEmBridge            CreationSourceTypeEnum = "EM_BRIDGE"
	CreationSourceTypeBulkDiscovery       CreationSourceTypeEnum = "BULK_DISCOVERY"
	CreationSourceTypeServiceConnectorHub CreationSourceTypeEnum = "SERVICE_CONNECTOR_HUB"
	CreationSourceTypeDiscovery           CreationSourceTypeEnum = "DISCOVERY"
	CreationSourceTypeLoggingAnalytics    CreationSourceTypeEnum = "LOGGING_ANALYTICS"
	CreationSourceTypeNone                CreationSourceTypeEnum = "NONE"
)

var mappingCreationSourceTypeEnum = map[string]CreationSourceTypeEnum{
	"EM_BRIDGE":             CreationSourceTypeEmBridge,
	"BULK_DISCOVERY":        CreationSourceTypeBulkDiscovery,
	"SERVICE_CONNECTOR_HUB": CreationSourceTypeServiceConnectorHub,
	"DISCOVERY":             CreationSourceTypeDiscovery,
	"LOGGING_ANALYTICS":     CreationSourceTypeLoggingAnalytics,
	"NONE":                  CreationSourceTypeNone,
}

var mappingCreationSourceTypeEnumLowerCase = map[string]CreationSourceTypeEnum{
	"em_bridge":             CreationSourceTypeEmBridge,
	"bulk_discovery":        CreationSourceTypeBulkDiscovery,
	"service_connector_hub": CreationSourceTypeServiceConnectorHub,
	"discovery":             CreationSourceTypeDiscovery,
	"logging_analytics":     CreationSourceTypeLoggingAnalytics,
	"none":                  CreationSourceTypeNone,
}

// GetCreationSourceTypeEnumValues Enumerates the set of values for CreationSourceTypeEnum
func GetCreationSourceTypeEnumValues() []CreationSourceTypeEnum {
	values := make([]CreationSourceTypeEnum, 0)
	for _, v := range mappingCreationSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreationSourceTypeEnumStringValues Enumerates the set of values in String for CreationSourceTypeEnum
func GetCreationSourceTypeEnumStringValues() []string {
	return []string{
		"EM_BRIDGE",
		"BULK_DISCOVERY",
		"SERVICE_CONNECTOR_HUB",
		"DISCOVERY",
		"LOGGING_ANALYTICS",
		"NONE",
	}
}

// GetMappingCreationSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreationSourceTypeEnum(val string) (CreationSourceTypeEnum, bool) {
	enum, ok := mappingCreationSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
