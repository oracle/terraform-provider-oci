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

// DataPumpJobModeEnum Enum with underlying type: string
type DataPumpJobModeEnum string

// Set of constants representing the allowable values for DataPumpJobModeEnum
const (
	DataPumpJobModeFull          DataPumpJobModeEnum = "FULL"
	DataPumpJobModeSchema        DataPumpJobModeEnum = "SCHEMA"
	DataPumpJobModeTable         DataPumpJobModeEnum = "TABLE"
	DataPumpJobModeTablespace    DataPumpJobModeEnum = "TABLESPACE"
	DataPumpJobModeTransportable DataPumpJobModeEnum = "TRANSPORTABLE"
)

var mappingDataPumpJobModeEnum = map[string]DataPumpJobModeEnum{
	"FULL":          DataPumpJobModeFull,
	"SCHEMA":        DataPumpJobModeSchema,
	"TABLE":         DataPumpJobModeTable,
	"TABLESPACE":    DataPumpJobModeTablespace,
	"TRANSPORTABLE": DataPumpJobModeTransportable,
}

// GetDataPumpJobModeEnumValues Enumerates the set of values for DataPumpJobModeEnum
func GetDataPumpJobModeEnumValues() []DataPumpJobModeEnum {
	values := make([]DataPumpJobModeEnum, 0)
	for _, v := range mappingDataPumpJobModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataPumpJobModeEnumStringValues Enumerates the set of values in String for DataPumpJobModeEnum
func GetDataPumpJobModeEnumStringValues() []string {
	return []string{
		"FULL",
		"SCHEMA",
		"TABLE",
		"TABLESPACE",
		"TRANSPORTABLE",
	}
}

// GetMappingDataPumpJobModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataPumpJobModeEnum(val string) (DataPumpJobModeEnum, bool) {
	mappingDataPumpJobModeEnumIgnoreCase := make(map[string]DataPumpJobModeEnum)
	for k, v := range mappingDataPumpJobModeEnum {
		mappingDataPumpJobModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataPumpJobModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
