// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"strings"
)

// TableStatisticsStatusCategoryEnum Enum with underlying type: string
type TableStatisticsStatusCategoryEnum string

// Set of constants representing the allowable values for TableStatisticsStatusCategoryEnum
const (
	TableStatisticsStatusCategoryNoStats  TableStatisticsStatusCategoryEnum = "NO_STATS"
	TableStatisticsStatusCategoryStale    TableStatisticsStatusCategoryEnum = "STALE"
	TableStatisticsStatusCategoryNotStale TableStatisticsStatusCategoryEnum = "NOT_STALE"
)

var mappingTableStatisticsStatusCategoryEnum = map[string]TableStatisticsStatusCategoryEnum{
	"NO_STATS":  TableStatisticsStatusCategoryNoStats,
	"STALE":     TableStatisticsStatusCategoryStale,
	"NOT_STALE": TableStatisticsStatusCategoryNotStale,
}

var mappingTableStatisticsStatusCategoryEnumLowerCase = map[string]TableStatisticsStatusCategoryEnum{
	"no_stats":  TableStatisticsStatusCategoryNoStats,
	"stale":     TableStatisticsStatusCategoryStale,
	"not_stale": TableStatisticsStatusCategoryNotStale,
}

// GetTableStatisticsStatusCategoryEnumValues Enumerates the set of values for TableStatisticsStatusCategoryEnum
func GetTableStatisticsStatusCategoryEnumValues() []TableStatisticsStatusCategoryEnum {
	values := make([]TableStatisticsStatusCategoryEnum, 0)
	for _, v := range mappingTableStatisticsStatusCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetTableStatisticsStatusCategoryEnumStringValues Enumerates the set of values in String for TableStatisticsStatusCategoryEnum
func GetTableStatisticsStatusCategoryEnumStringValues() []string {
	return []string{
		"NO_STATS",
		"STALE",
		"NOT_STALE",
	}
}

// GetMappingTableStatisticsStatusCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTableStatisticsStatusCategoryEnum(val string) (TableStatisticsStatusCategoryEnum, bool) {
	enum, ok := mappingTableStatisticsStatusCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
