// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"strings"
)

// ReportDefinitionDataSourceEnum Enum with underlying type: string
type ReportDefinitionDataSourceEnum string

// Set of constants representing the allowable values for ReportDefinitionDataSourceEnum
const (
	ReportDefinitionDataSourceEvents ReportDefinitionDataSourceEnum = "EVENTS"
	ReportDefinitionDataSourceAlerts ReportDefinitionDataSourceEnum = "ALERTS"
)

var mappingReportDefinitionDataSourceEnum = map[string]ReportDefinitionDataSourceEnum{
	"EVENTS": ReportDefinitionDataSourceEvents,
	"ALERTS": ReportDefinitionDataSourceAlerts,
}

// GetReportDefinitionDataSourceEnumValues Enumerates the set of values for ReportDefinitionDataSourceEnum
func GetReportDefinitionDataSourceEnumValues() []ReportDefinitionDataSourceEnum {
	values := make([]ReportDefinitionDataSourceEnum, 0)
	for _, v := range mappingReportDefinitionDataSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetReportDefinitionDataSourceEnumStringValues Enumerates the set of values in String for ReportDefinitionDataSourceEnum
func GetReportDefinitionDataSourceEnumStringValues() []string {
	return []string{
		"EVENTS",
		"ALERTS",
	}
}

// GetMappingReportDefinitionDataSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReportDefinitionDataSourceEnum(val string) (ReportDefinitionDataSourceEnum, bool) {
	mappingReportDefinitionDataSourceEnumIgnoreCase := make(map[string]ReportDefinitionDataSourceEnum)
	for k, v := range mappingReportDefinitionDataSourceEnum {
		mappingReportDefinitionDataSourceEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingReportDefinitionDataSourceEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
