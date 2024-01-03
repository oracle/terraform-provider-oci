// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
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

var mappingDataMaskCategoryEnumLowerCase = map[string]DataMaskCategoryEnum{
	"actor":     DataMaskCategoryActor,
	"pii":       DataMaskCategoryPii,
	"phi":       DataMaskCategoryPhi,
	"financial": DataMaskCategoryFinancial,
	"location":  DataMaskCategoryLocation,
	"custom":    DataMaskCategoryCustom,
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
	enum, ok := mappingDataMaskCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
