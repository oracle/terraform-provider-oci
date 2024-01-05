// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	DataPumpTablespaceBlockSizesInKbSize8K  DataPumpTablespaceBlockSizesInKbEnum = "SIZE_8K"
	DataPumpTablespaceBlockSizesInKbSize16K DataPumpTablespaceBlockSizesInKbEnum = "SIZE_16K"
)

var mappingDataPumpTablespaceBlockSizesInKbEnum = map[string]DataPumpTablespaceBlockSizesInKbEnum{
	"SIZE_8K":  DataPumpTablespaceBlockSizesInKbSize8K,
	"SIZE_16K": DataPumpTablespaceBlockSizesInKbSize16K,
}

var mappingDataPumpTablespaceBlockSizesInKbEnumLowerCase = map[string]DataPumpTablespaceBlockSizesInKbEnum{
	"size_8k":  DataPumpTablespaceBlockSizesInKbSize8K,
	"size_16k": DataPumpTablespaceBlockSizesInKbSize16K,
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
		"SIZE_8K",
		"SIZE_16K",
	}
}

// GetMappingDataPumpTablespaceBlockSizesInKbEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataPumpTablespaceBlockSizesInKbEnum(val string) (DataPumpTablespaceBlockSizesInKbEnum, bool) {
	enum, ok := mappingDataPumpTablespaceBlockSizesInKbEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
