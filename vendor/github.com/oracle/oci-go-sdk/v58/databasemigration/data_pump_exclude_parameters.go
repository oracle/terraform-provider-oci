// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// DataPumpExcludeParametersEnum Enum with underlying type: string
type DataPumpExcludeParametersEnum string

// Set of constants representing the allowable values for DataPumpExcludeParametersEnum
const (
	DataPumpExcludeParametersIndex               DataPumpExcludeParametersEnum = "INDEX"
	DataPumpExcludeParametersMaterializedView    DataPumpExcludeParametersEnum = "MATERIALIZED_VIEW"
	DataPumpExcludeParametersMaterializedViewLog DataPumpExcludeParametersEnum = "MATERIALIZED_VIEW_LOG"
)

var mappingDataPumpExcludeParametersEnum = map[string]DataPumpExcludeParametersEnum{
	"INDEX":                 DataPumpExcludeParametersIndex,
	"MATERIALIZED_VIEW":     DataPumpExcludeParametersMaterializedView,
	"MATERIALIZED_VIEW_LOG": DataPumpExcludeParametersMaterializedViewLog,
}

// GetDataPumpExcludeParametersEnumValues Enumerates the set of values for DataPumpExcludeParametersEnum
func GetDataPumpExcludeParametersEnumValues() []DataPumpExcludeParametersEnum {
	values := make([]DataPumpExcludeParametersEnum, 0)
	for _, v := range mappingDataPumpExcludeParametersEnum {
		values = append(values, v)
	}
	return values
}

// GetDataPumpExcludeParametersEnumStringValues Enumerates the set of values in String for DataPumpExcludeParametersEnum
func GetDataPumpExcludeParametersEnumStringValues() []string {
	return []string{
		"INDEX",
		"MATERIALIZED_VIEW",
		"MATERIALIZED_VIEW_LOG",
	}
}

// GetMappingDataPumpExcludeParametersEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataPumpExcludeParametersEnum(val string) (DataPumpExcludeParametersEnum, bool) {
	mappingDataPumpExcludeParametersEnumIgnoreCase := make(map[string]DataPumpExcludeParametersEnum)
	for k, v := range mappingDataPumpExcludeParametersEnum {
		mappingDataPumpExcludeParametersEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataPumpExcludeParametersEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
