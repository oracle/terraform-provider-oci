// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// AwrReportFormatTypeEnum Enum with underlying type: string
type AwrReportFormatTypeEnum string

// Set of constants representing the allowable values for AwrReportFormatTypeEnum
const (
	AwrReportFormatTypeHtml AwrReportFormatTypeEnum = "HTML"
	AwrReportFormatTypeText AwrReportFormatTypeEnum = "TEXT"
)

var mappingAwrReportFormatTypeEnum = map[string]AwrReportFormatTypeEnum{
	"HTML": AwrReportFormatTypeHtml,
	"TEXT": AwrReportFormatTypeText,
}

var mappingAwrReportFormatTypeEnumLowerCase = map[string]AwrReportFormatTypeEnum{
	"html": AwrReportFormatTypeHtml,
	"text": AwrReportFormatTypeText,
}

// GetAwrReportFormatTypeEnumValues Enumerates the set of values for AwrReportFormatTypeEnum
func GetAwrReportFormatTypeEnumValues() []AwrReportFormatTypeEnum {
	values := make([]AwrReportFormatTypeEnum, 0)
	for _, v := range mappingAwrReportFormatTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAwrReportFormatTypeEnumStringValues Enumerates the set of values in String for AwrReportFormatTypeEnum
func GetAwrReportFormatTypeEnumStringValues() []string {
	return []string{
		"HTML",
		"TEXT",
	}
}

// GetMappingAwrReportFormatTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAwrReportFormatTypeEnum(val string) (AwrReportFormatTypeEnum, bool) {
	enum, ok := mappingAwrReportFormatTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
