// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// ScheduleTypesEnum Enum with underlying type: string
type ScheduleTypesEnum string

// Set of constants representing the allowable values for ScheduleTypesEnum
const (
	ScheduleTypesOnetime   ScheduleTypesEnum = "ONETIME"
	ScheduleTypesRecurring ScheduleTypesEnum = "RECURRING"
)

var mappingScheduleTypes = map[string]ScheduleTypesEnum{
	"ONETIME":   ScheduleTypesOnetime,
	"RECURRING": ScheduleTypesRecurring,
}

// GetScheduleTypesEnumValues Enumerates the set of values for ScheduleTypesEnum
func GetScheduleTypesEnumValues() []ScheduleTypesEnum {
	values := make([]ScheduleTypesEnum, 0)
	for _, v := range mappingScheduleTypes {
		values = append(values, v)
	}
	return values
}
