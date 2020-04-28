// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
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
