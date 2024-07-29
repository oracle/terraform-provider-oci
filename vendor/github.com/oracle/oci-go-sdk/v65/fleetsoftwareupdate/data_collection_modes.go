// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// DataCollectionModesEnum Enum with underlying type: string
type DataCollectionModesEnum string

// Set of constants representing the allowable values for DataCollectionModesEnum
const (
	DataCollectionModesEnable           DataCollectionModesEnum = "ENABLE"
	DataCollectionModesEnableAndRestore DataCollectionModesEnum = "ENABLE_AND_RESTORE"
	DataCollectionModesNoChange         DataCollectionModesEnum = "NO_CHANGE"
)

var mappingDataCollectionModesEnum = map[string]DataCollectionModesEnum{
	"ENABLE":             DataCollectionModesEnable,
	"ENABLE_AND_RESTORE": DataCollectionModesEnableAndRestore,
	"NO_CHANGE":          DataCollectionModesNoChange,
}

var mappingDataCollectionModesEnumLowerCase = map[string]DataCollectionModesEnum{
	"enable":             DataCollectionModesEnable,
	"enable_and_restore": DataCollectionModesEnableAndRestore,
	"no_change":          DataCollectionModesNoChange,
}

// GetDataCollectionModesEnumValues Enumerates the set of values for DataCollectionModesEnum
func GetDataCollectionModesEnumValues() []DataCollectionModesEnum {
	values := make([]DataCollectionModesEnum, 0)
	for _, v := range mappingDataCollectionModesEnum {
		values = append(values, v)
	}
	return values
}

// GetDataCollectionModesEnumStringValues Enumerates the set of values in String for DataCollectionModesEnum
func GetDataCollectionModesEnumStringValues() []string {
	return []string{
		"ENABLE",
		"ENABLE_AND_RESTORE",
		"NO_CHANGE",
	}
}

// GetMappingDataCollectionModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataCollectionModesEnum(val string) (DataCollectionModesEnum, bool) {
	enum, ok := mappingDataCollectionModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
