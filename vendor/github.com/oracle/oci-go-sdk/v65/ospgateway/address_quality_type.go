// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"strings"
)

// AddressQualityTypeEnum Enum with underlying type: string
type AddressQualityTypeEnum string

// Set of constants representing the allowable values for AddressQualityTypeEnum
const (
	AddressQualityTypeExcellent AddressQualityTypeEnum = "EXCELLENT"
	AddressQualityTypeGood      AddressQualityTypeEnum = "GOOD"
	AddressQualityTypeAverage   AddressQualityTypeEnum = "AVERAGE"
	AddressQualityTypePoor      AddressQualityTypeEnum = "POOR"
	AddressQualityTypeBad       AddressQualityTypeEnum = "BAD"
)

var mappingAddressQualityTypeEnum = map[string]AddressQualityTypeEnum{
	"EXCELLENT": AddressQualityTypeExcellent,
	"GOOD":      AddressQualityTypeGood,
	"AVERAGE":   AddressQualityTypeAverage,
	"POOR":      AddressQualityTypePoor,
	"BAD":       AddressQualityTypeBad,
}

var mappingAddressQualityTypeEnumLowerCase = map[string]AddressQualityTypeEnum{
	"excellent": AddressQualityTypeExcellent,
	"good":      AddressQualityTypeGood,
	"average":   AddressQualityTypeAverage,
	"poor":      AddressQualityTypePoor,
	"bad":       AddressQualityTypeBad,
}

// GetAddressQualityTypeEnumValues Enumerates the set of values for AddressQualityTypeEnum
func GetAddressQualityTypeEnumValues() []AddressQualityTypeEnum {
	values := make([]AddressQualityTypeEnum, 0)
	for _, v := range mappingAddressQualityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddressQualityTypeEnumStringValues Enumerates the set of values in String for AddressQualityTypeEnum
func GetAddressQualityTypeEnumStringValues() []string {
	return []string{
		"EXCELLENT",
		"GOOD",
		"AVERAGE",
		"POOR",
		"BAD",
	}
}

// GetMappingAddressQualityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddressQualityTypeEnum(val string) (AddressQualityTypeEnum, bool) {
	enum, ok := mappingAddressQualityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
