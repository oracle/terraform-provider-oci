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

// OptimizerStatisticsGroupByTypesEnum Enum with underlying type: string
type OptimizerStatisticsGroupByTypesEnum string

// Set of constants representing the allowable values for OptimizerStatisticsGroupByTypesEnum
const (
	OptimizerStatisticsGroupByTypesTaskStatus        OptimizerStatisticsGroupByTypesEnum = "TASK_STATUS"
	OptimizerStatisticsGroupByTypesTaskObjectsStatus OptimizerStatisticsGroupByTypesEnum = "TASK_OBJECTS_STATUS"
)

var mappingOptimizerStatisticsGroupByTypesEnum = map[string]OptimizerStatisticsGroupByTypesEnum{
	"TASK_STATUS":         OptimizerStatisticsGroupByTypesTaskStatus,
	"TASK_OBJECTS_STATUS": OptimizerStatisticsGroupByTypesTaskObjectsStatus,
}

var mappingOptimizerStatisticsGroupByTypesEnumLowerCase = map[string]OptimizerStatisticsGroupByTypesEnum{
	"task_status":         OptimizerStatisticsGroupByTypesTaskStatus,
	"task_objects_status": OptimizerStatisticsGroupByTypesTaskObjectsStatus,
}

// GetOptimizerStatisticsGroupByTypesEnumValues Enumerates the set of values for OptimizerStatisticsGroupByTypesEnum
func GetOptimizerStatisticsGroupByTypesEnumValues() []OptimizerStatisticsGroupByTypesEnum {
	values := make([]OptimizerStatisticsGroupByTypesEnum, 0)
	for _, v := range mappingOptimizerStatisticsGroupByTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOptimizerStatisticsGroupByTypesEnumStringValues Enumerates the set of values in String for OptimizerStatisticsGroupByTypesEnum
func GetOptimizerStatisticsGroupByTypesEnumStringValues() []string {
	return []string{
		"TASK_STATUS",
		"TASK_OBJECTS_STATUS",
	}
}

// GetMappingOptimizerStatisticsGroupByTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOptimizerStatisticsGroupByTypesEnum(val string) (OptimizerStatisticsGroupByTypesEnum, bool) {
	enum, ok := mappingOptimizerStatisticsGroupByTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
