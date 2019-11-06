// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// CapacityTypeEnum Enum with underlying type: string
type CapacityTypeEnum string

// Set of constants representing the allowable values for CapacityTypeEnum
const (
	CapacityTypeOlpuCount CapacityTypeEnum = "OLPU_COUNT"
	CapacityTypeUserCount CapacityTypeEnum = "USER_COUNT"
)

var mappingCapacityType = map[string]CapacityTypeEnum{
	"OLPU_COUNT": CapacityTypeOlpuCount,
	"USER_COUNT": CapacityTypeUserCount,
}

// GetCapacityTypeEnumValues Enumerates the set of values for CapacityTypeEnum
func GetCapacityTypeEnumValues() []CapacityTypeEnum {
	values := make([]CapacityTypeEnum, 0)
	for _, v := range mappingCapacityType {
		values = append(values, v)
	}
	return values
}
