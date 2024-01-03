// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"strings"
)

// SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
	SortByCapacityType   SortByEnum = "capacityType"
	SortByCapacityValue  SortByEnum = "capacityValue"
	SortByFeatureSet     SortByEnum = "featureSet"
	SortByLifecycleState SortByEnum = "lifecycleState"
	SortByName           SortByEnum = "name"
	SortByTimeCreated    SortByEnum = "timeCreated"
)

var mappingSortByEnum = map[string]SortByEnum{
	"capacityType":   SortByCapacityType,
	"capacityValue":  SortByCapacityValue,
	"featureSet":     SortByFeatureSet,
	"lifecycleState": SortByLifecycleState,
	"name":           SortByName,
	"timeCreated":    SortByTimeCreated,
}

var mappingSortByEnumLowerCase = map[string]SortByEnum{
	"capacitytype":   SortByCapacityType,
	"capacityvalue":  SortByCapacityValue,
	"featureset":     SortByFeatureSet,
	"lifecyclestate": SortByLifecycleState,
	"name":           SortByName,
	"timecreated":    SortByTimeCreated,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
	values := make([]SortByEnum, 0)
	for _, v := range mappingSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSortByEnumStringValues Enumerates the set of values in String for SortByEnum
func GetSortByEnumStringValues() []string {
	return []string{
		"capacityType",
		"capacityValue",
		"featureSet",
		"lifecycleState",
		"name",
		"timeCreated",
	}
}

// GetMappingSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortByEnum(val string) (SortByEnum, bool) {
	enum, ok := mappingSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
