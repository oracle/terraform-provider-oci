// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// ServingTypeEnum Enum with underlying type: string
type ServingTypeEnum string

// Set of constants representing the allowable values for ServingTypeEnum
const (
	ServingTypeOnDemand ServingTypeEnum = "ON_DEMAND"
)

var mappingServingTypeEnum = map[string]ServingTypeEnum{
	"ON_DEMAND": ServingTypeOnDemand,
}

var mappingServingTypeEnumLowerCase = map[string]ServingTypeEnum{
	"on_demand": ServingTypeOnDemand,
}

// GetServingTypeEnumValues Enumerates the set of values for ServingTypeEnum
func GetServingTypeEnumValues() []ServingTypeEnum {
	values := make([]ServingTypeEnum, 0)
	for _, v := range mappingServingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServingTypeEnumStringValues Enumerates the set of values in String for ServingTypeEnum
func GetServingTypeEnumStringValues() []string {
	return []string{
		"ON_DEMAND",
	}
}

// GetMappingServingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServingTypeEnum(val string) (ServingTypeEnum, bool) {
	enum, ok := mappingServingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
