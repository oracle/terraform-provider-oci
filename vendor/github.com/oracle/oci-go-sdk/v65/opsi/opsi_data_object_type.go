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

// OpsiDataObjectTypeEnum Enum with underlying type: string
type OpsiDataObjectTypeEnum string

// Set of constants representing the allowable values for OpsiDataObjectTypeEnum
const (
	OpsiDataObjectTypeDatabaseInsightsDataObject OpsiDataObjectTypeEnum = "DATABASE_INSIGHTS_DATA_OBJECT"
	OpsiDataObjectTypeHostInsightsDataObject     OpsiDataObjectTypeEnum = "HOST_INSIGHTS_DATA_OBJECT"
	OpsiDataObjectTypeExadataInsightsDataObject  OpsiDataObjectTypeEnum = "EXADATA_INSIGHTS_DATA_OBJECT"
)

var mappingOpsiDataObjectTypeEnum = map[string]OpsiDataObjectTypeEnum{
	"DATABASE_INSIGHTS_DATA_OBJECT": OpsiDataObjectTypeDatabaseInsightsDataObject,
	"HOST_INSIGHTS_DATA_OBJECT":     OpsiDataObjectTypeHostInsightsDataObject,
	"EXADATA_INSIGHTS_DATA_OBJECT":  OpsiDataObjectTypeExadataInsightsDataObject,
}

var mappingOpsiDataObjectTypeEnumLowerCase = map[string]OpsiDataObjectTypeEnum{
	"database_insights_data_object": OpsiDataObjectTypeDatabaseInsightsDataObject,
	"host_insights_data_object":     OpsiDataObjectTypeHostInsightsDataObject,
	"exadata_insights_data_object":  OpsiDataObjectTypeExadataInsightsDataObject,
}

// GetOpsiDataObjectTypeEnumValues Enumerates the set of values for OpsiDataObjectTypeEnum
func GetOpsiDataObjectTypeEnumValues() []OpsiDataObjectTypeEnum {
	values := make([]OpsiDataObjectTypeEnum, 0)
	for _, v := range mappingOpsiDataObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsiDataObjectTypeEnumStringValues Enumerates the set of values in String for OpsiDataObjectTypeEnum
func GetOpsiDataObjectTypeEnumStringValues() []string {
	return []string{
		"DATABASE_INSIGHTS_DATA_OBJECT",
		"HOST_INSIGHTS_DATA_OBJECT",
		"EXADATA_INSIGHTS_DATA_OBJECT",
	}
}

// GetMappingOpsiDataObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsiDataObjectTypeEnum(val string) (OpsiDataObjectTypeEnum, bool) {
	enum, ok := mappingOpsiDataObjectTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
