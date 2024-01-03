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

// IormPlanObjectiveEnumEnum Enum with underlying type: string
type IormPlanObjectiveEnumEnum string

// Set of constants representing the allowable values for IormPlanObjectiveEnumEnum
const (
	IormPlanObjectiveEnumAuto           IormPlanObjectiveEnumEnum = "AUTO"
	IormPlanObjectiveEnumHighThroughput IormPlanObjectiveEnumEnum = "HIGH_THROUGHPUT"
	IormPlanObjectiveEnumLowLatency     IormPlanObjectiveEnumEnum = "LOW_LATENCY"
	IormPlanObjectiveEnumBalanced       IormPlanObjectiveEnumEnum = "BALANCED"
	IormPlanObjectiveEnumBasic          IormPlanObjectiveEnumEnum = "BASIC"
	IormPlanObjectiveEnumOther          IormPlanObjectiveEnumEnum = "OTHER"
)

var mappingIormPlanObjectiveEnumEnum = map[string]IormPlanObjectiveEnumEnum{
	"AUTO":            IormPlanObjectiveEnumAuto,
	"HIGH_THROUGHPUT": IormPlanObjectiveEnumHighThroughput,
	"LOW_LATENCY":     IormPlanObjectiveEnumLowLatency,
	"BALANCED":        IormPlanObjectiveEnumBalanced,
	"BASIC":           IormPlanObjectiveEnumBasic,
	"OTHER":           IormPlanObjectiveEnumOther,
}

var mappingIormPlanObjectiveEnumEnumLowerCase = map[string]IormPlanObjectiveEnumEnum{
	"auto":            IormPlanObjectiveEnumAuto,
	"high_throughput": IormPlanObjectiveEnumHighThroughput,
	"low_latency":     IormPlanObjectiveEnumLowLatency,
	"balanced":        IormPlanObjectiveEnumBalanced,
	"basic":           IormPlanObjectiveEnumBasic,
	"other":           IormPlanObjectiveEnumOther,
}

// GetIormPlanObjectiveEnumEnumValues Enumerates the set of values for IormPlanObjectiveEnumEnum
func GetIormPlanObjectiveEnumEnumValues() []IormPlanObjectiveEnumEnum {
	values := make([]IormPlanObjectiveEnumEnum, 0)
	for _, v := range mappingIormPlanObjectiveEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetIormPlanObjectiveEnumEnumStringValues Enumerates the set of values in String for IormPlanObjectiveEnumEnum
func GetIormPlanObjectiveEnumEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH_THROUGHPUT",
		"LOW_LATENCY",
		"BALANCED",
		"BASIC",
		"OTHER",
	}
}

// GetMappingIormPlanObjectiveEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIormPlanObjectiveEnumEnum(val string) (IormPlanObjectiveEnumEnum, bool) {
	enum, ok := mappingIormPlanObjectiveEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
