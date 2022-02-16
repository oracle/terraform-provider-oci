// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// DataMaskCategoryEnum Enum with underlying type: string
type DataMaskCategoryEnum string

// Set of constants representing the allowable values for DataMaskCategoryEnum
const (
	DataMaskCategoryActor     DataMaskCategoryEnum = "ACTOR"
	DataMaskCategoryPii       DataMaskCategoryEnum = "PII"
	DataMaskCategoryPhi       DataMaskCategoryEnum = "PHI"
	DataMaskCategoryFinancial DataMaskCategoryEnum = "FINANCIAL"
	DataMaskCategoryLocation  DataMaskCategoryEnum = "LOCATION"
	DataMaskCategoryCustom    DataMaskCategoryEnum = "CUSTOM"
)

var mappingDataMaskCategoryEnum = map[string]DataMaskCategoryEnum{
	"ACTOR":     DataMaskCategoryActor,
	"PII":       DataMaskCategoryPii,
	"PHI":       DataMaskCategoryPhi,
	"FINANCIAL": DataMaskCategoryFinancial,
	"LOCATION":  DataMaskCategoryLocation,
	"CUSTOM":    DataMaskCategoryCustom,
}

// GetDataMaskCategoryEnumValues Enumerates the set of values for DataMaskCategoryEnum
func GetDataMaskCategoryEnumValues() []DataMaskCategoryEnum {
	values := make([]DataMaskCategoryEnum, 0)
	for _, v := range mappingDataMaskCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetDataMaskCategoryEnumStringValues Enumerates the set of values in String for DataMaskCategoryEnum
func GetDataMaskCategoryEnumStringValues() []string {
	return []string{
		"ACTOR",
		"PII",
		"PHI",
		"FINANCIAL",
		"LOCATION",
		"CUSTOM",
	}
}

// GetMappingDataMaskCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataMaskCategoryEnum(val string) (DataMaskCategoryEnum, bool) {
	mappingDataMaskCategoryEnumIgnoreCase := make(map[string]DataMaskCategoryEnum)
	for k, v := range mappingDataMaskCategoryEnum {
		mappingDataMaskCategoryEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDataMaskCategoryEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
