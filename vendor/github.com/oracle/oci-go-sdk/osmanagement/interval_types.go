// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// IntervalTypesEnum Enum with underlying type: string
type IntervalTypesEnum string

// Set of constants representing the allowable values for IntervalTypesEnum
const (
	IntervalTypesHour  IntervalTypesEnum = "HOUR"
	IntervalTypesDay   IntervalTypesEnum = "DAY"
	IntervalTypesWeek  IntervalTypesEnum = "WEEK"
	IntervalTypesMonth IntervalTypesEnum = "MONTH"
)

var mappingIntervalTypes = map[string]IntervalTypesEnum{
	"HOUR":  IntervalTypesHour,
	"DAY":   IntervalTypesDay,
	"WEEK":  IntervalTypesWeek,
	"MONTH": IntervalTypesMonth,
}

// GetIntervalTypesEnumValues Enumerates the set of values for IntervalTypesEnum
func GetIntervalTypesEnumValues() []IntervalTypesEnum {
	values := make([]IntervalTypesEnum, 0)
	for _, v := range mappingIntervalTypes {
		values = append(values, v)
	}
	return values
}
