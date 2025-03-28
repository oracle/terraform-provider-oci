// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"strings"
)

// NewsSqlInsightsContentTypesResourceEnum Enum with underlying type: string
type NewsSqlInsightsContentTypesResourceEnum string

// Set of constants representing the allowable values for NewsSqlInsightsContentTypesResourceEnum
const (
	NewsSqlInsightsContentTypesResourceDatabase NewsSqlInsightsContentTypesResourceEnum = "DATABASE"
	NewsSqlInsightsContentTypesResourceExadata  NewsSqlInsightsContentTypesResourceEnum = "EXADATA"
)

var mappingNewsSqlInsightsContentTypesResourceEnum = map[string]NewsSqlInsightsContentTypesResourceEnum{
	"DATABASE": NewsSqlInsightsContentTypesResourceDatabase,
	"EXADATA":  NewsSqlInsightsContentTypesResourceExadata,
}

var mappingNewsSqlInsightsContentTypesResourceEnumLowerCase = map[string]NewsSqlInsightsContentTypesResourceEnum{
	"database": NewsSqlInsightsContentTypesResourceDatabase,
	"exadata":  NewsSqlInsightsContentTypesResourceExadata,
}

// GetNewsSqlInsightsContentTypesResourceEnumValues Enumerates the set of values for NewsSqlInsightsContentTypesResourceEnum
func GetNewsSqlInsightsContentTypesResourceEnumValues() []NewsSqlInsightsContentTypesResourceEnum {
	values := make([]NewsSqlInsightsContentTypesResourceEnum, 0)
	for _, v := range mappingNewsSqlInsightsContentTypesResourceEnum {
		values = append(values, v)
	}
	return values
}

// GetNewsSqlInsightsContentTypesResourceEnumStringValues Enumerates the set of values in String for NewsSqlInsightsContentTypesResourceEnum
func GetNewsSqlInsightsContentTypesResourceEnumStringValues() []string {
	return []string{
		"DATABASE",
		"EXADATA",
	}
}

// GetMappingNewsSqlInsightsContentTypesResourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNewsSqlInsightsContentTypesResourceEnum(val string) (NewsSqlInsightsContentTypesResourceEnum, bool) {
	enum, ok := mappingNewsSqlInsightsContentTypesResourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
