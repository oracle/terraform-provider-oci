// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// DataPumpExcludeParametersEnum Enum with underlying type: string
type DataPumpExcludeParametersEnum string

// Set of constants representing the allowable values for DataPumpExcludeParametersEnum
const (
	DataPumpExcludeParametersIndex               DataPumpExcludeParametersEnum = "INDEX"
	DataPumpExcludeParametersMaterializedView    DataPumpExcludeParametersEnum = "MATERIALIZED_VIEW"
	DataPumpExcludeParametersMaterializedViewLog DataPumpExcludeParametersEnum = "MATERIALIZED_VIEW_LOG"
)

var mappingDataPumpExcludeParameters = map[string]DataPumpExcludeParametersEnum{
	"INDEX":                 DataPumpExcludeParametersIndex,
	"MATERIALIZED_VIEW":     DataPumpExcludeParametersMaterializedView,
	"MATERIALIZED_VIEW_LOG": DataPumpExcludeParametersMaterializedViewLog,
}

// GetDataPumpExcludeParametersEnumValues Enumerates the set of values for DataPumpExcludeParametersEnum
func GetDataPumpExcludeParametersEnumValues() []DataPumpExcludeParametersEnum {
	values := make([]DataPumpExcludeParametersEnum, 0)
	for _, v := range mappingDataPumpExcludeParameters {
		values = append(values, v)
	}
	return values
}
