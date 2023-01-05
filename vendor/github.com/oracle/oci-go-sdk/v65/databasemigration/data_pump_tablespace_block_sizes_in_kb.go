// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// DataPumpTablespaceBlockSizesInKbEnum Enum with underlying type: string
type DataPumpTablespaceBlockSizesInKbEnum string

// Set of constants representing the allowable values for DataPumpTablespaceBlockSizesInKbEnum
const (
	DataPumpTablespaceBlockSizesInKb8  DataPumpTablespaceBlockSizesInKbEnum = "8"
	DataPumpTablespaceBlockSizesInKb16 DataPumpTablespaceBlockSizesInKbEnum = "16"
)

var mappingDataPumpTablespaceBlockSizesInKbEnum = map[string]DataPumpTablespaceBlockSizesInKbEnum{
	"8":  DataPumpTablespaceBlockSizesInKb8,
	"16": DataPumpTablespaceBlockSizesInKb16,
}

var mappingDataPumpTablespaceBlockSizesInKbEnumLowerCase = map[string]DataPumpTablespaceBlockSizesInKbEnum{
	"8":  DataPumpTablespaceBlockSizesInKb8,
	"16": DataPumpTablespaceBlockSizesInKb16,
}

// GetDataPumpTablespaceBlockSizesInKbEnumValues Enumerates the set of values for DataPumpTablespaceBlockSizesInKbEnum
func GetDataPumpTablespaceBlockSizesInKbEnumValues() []DataPumpTablespaceBlockSizesInKbEnum {
	values := make([]DataPumpTablespaceBlockSizesInKbEnum, 0)
	for _, v := range mappingDataPumpTablespaceBlockSizesInKbEnum {
		values = append(values, v)
	}
	return values
}

// GetDataPumpTablespaceBlockSizesInKbEnumStringValues Enumerates the set of values in String for DataPumpTablespaceBlockSizesInKbEnum
func GetDataPumpTablespaceBlockSizesInKbEnumStringValues() []string {
	return []string{
		"8",
		"16",
	}
}

// GetMappingDataPumpTablespaceBlockSizesInKbEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataPumpTablespaceBlockSizesInKbEnum(val string) (DataPumpTablespaceBlockSizesInKbEnum, bool) {
	enum, ok := mappingDataPumpTablespaceBlockSizesInKbEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
