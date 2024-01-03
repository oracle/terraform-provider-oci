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

// OptimizerStatisticsTaskFilterTypesEnum Enum with underlying type: string
type OptimizerStatisticsTaskFilterTypesEnum string

// Set of constants representing the allowable values for OptimizerStatisticsTaskFilterTypesEnum
const (
	OptimizerStatisticsTaskFilterTypesAll    OptimizerStatisticsTaskFilterTypesEnum = "ALL"
	OptimizerStatisticsTaskFilterTypesManual OptimizerStatisticsTaskFilterTypesEnum = "MANUAL"
	OptimizerStatisticsTaskFilterTypesAuto   OptimizerStatisticsTaskFilterTypesEnum = "AUTO"
)

var mappingOptimizerStatisticsTaskFilterTypesEnum = map[string]OptimizerStatisticsTaskFilterTypesEnum{
	"ALL":    OptimizerStatisticsTaskFilterTypesAll,
	"MANUAL": OptimizerStatisticsTaskFilterTypesManual,
	"AUTO":   OptimizerStatisticsTaskFilterTypesAuto,
}

var mappingOptimizerStatisticsTaskFilterTypesEnumLowerCase = map[string]OptimizerStatisticsTaskFilterTypesEnum{
	"all":    OptimizerStatisticsTaskFilterTypesAll,
	"manual": OptimizerStatisticsTaskFilterTypesManual,
	"auto":   OptimizerStatisticsTaskFilterTypesAuto,
}

// GetOptimizerStatisticsTaskFilterTypesEnumValues Enumerates the set of values for OptimizerStatisticsTaskFilterTypesEnum
func GetOptimizerStatisticsTaskFilterTypesEnumValues() []OptimizerStatisticsTaskFilterTypesEnum {
	values := make([]OptimizerStatisticsTaskFilterTypesEnum, 0)
	for _, v := range mappingOptimizerStatisticsTaskFilterTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsTaskFilterTypesEnumStringValues Enumerates the set of values in String for OptimizerStatisticsTaskFilterTypesEnum
func GetOptimizerStatisticsTaskFilterTypesEnumStringValues() []string {
	return []string{
		"ALL",
		"MANUAL",
		"AUTO",
	}
}

// GetMappingOptimizerStatisticsTaskFilterTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsTaskFilterTypesEnum(val string) (OptimizerStatisticsTaskFilterTypesEnum, bool) {
	enum, ok := mappingOptimizerStatisticsTaskFilterTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
