// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"strings"
)

// RankingMeasureEnum Enum with underlying type: string
type RankingMeasureEnum string

// Set of constants representing the allowable values for RankingMeasureEnum
const (
	RankingMeasureElapsedTime   RankingMeasureEnum = "ELAPSED_TIME"
	RankingMeasureCpuTime       RankingMeasureEnum = "CPU_TIME"
	RankingMeasureOptimizerCost RankingMeasureEnum = "OPTIMIZER_COST"
	RankingMeasureBufferGets    RankingMeasureEnum = "BUFFER_GETS"
	RankingMeasureDiskReads     RankingMeasureEnum = "DISK_READS"
	RankingMeasureDirectWrites  RankingMeasureEnum = "DIRECT_WRITES"
)

var mappingRankingMeasureEnum = map[string]RankingMeasureEnum{
	"ELAPSED_TIME":   RankingMeasureElapsedTime,
	"CPU_TIME":       RankingMeasureCpuTime,
	"OPTIMIZER_COST": RankingMeasureOptimizerCost,
	"BUFFER_GETS":    RankingMeasureBufferGets,
	"DISK_READS":     RankingMeasureDiskReads,
	"DIRECT_WRITES":  RankingMeasureDirectWrites,
}

var mappingRankingMeasureEnumLowerCase = map[string]RankingMeasureEnum{
	"elapsed_time":   RankingMeasureElapsedTime,
	"cpu_time":       RankingMeasureCpuTime,
	"optimizer_cost": RankingMeasureOptimizerCost,
	"buffer_gets":    RankingMeasureBufferGets,
	"disk_reads":     RankingMeasureDiskReads,
	"direct_writes":  RankingMeasureDirectWrites,
}

// GetRankingMeasureEnumValues Enumerates the set of values for RankingMeasureEnum
func GetRankingMeasureEnumValues() []RankingMeasureEnum {
	values := make([]RankingMeasureEnum, 0)
	for _, v := range mappingRankingMeasureEnum {
		values = append(values, v)
	}
	return values
}

// GetRankingMeasureEnumStringValues Enumerates the set of values in String for RankingMeasureEnum
func GetRankingMeasureEnumStringValues() []string {
	return []string{
		"ELAPSED_TIME",
		"CPU_TIME",
		"OPTIMIZER_COST",
		"BUFFER_GETS",
		"DISK_READS",
		"DIRECT_WRITES",
	}
}

// GetMappingRankingMeasureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRankingMeasureEnum(val string) (RankingMeasureEnum, bool) {
	enum, ok := mappingRankingMeasureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
