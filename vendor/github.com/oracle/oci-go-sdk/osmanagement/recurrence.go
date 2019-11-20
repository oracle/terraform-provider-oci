// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Recurrence An object for representing a recurrence time interval
type Recurrence struct {

	// the interval period for the recurrence
	IntervalType RecurrenceIntervalTypeEnum `mandatory:"true" json:"intervalType"`

	// the value for the interval period for the recurrence
	IntervalValue *string `mandatory:"true" json:"intervalValue"`
}

func (m Recurrence) String() string {
	return common.PointerString(m)
}

// RecurrenceIntervalTypeEnum Enum with underlying type: string
type RecurrenceIntervalTypeEnum string

// Set of constants representing the allowable values for RecurrenceIntervalTypeEnum
const (
	RecurrenceIntervalTypeMinutes RecurrenceIntervalTypeEnum = "MINUTES"
	RecurrenceIntervalTypeHours   RecurrenceIntervalTypeEnum = "HOURS"
	RecurrenceIntervalTypeDays    RecurrenceIntervalTypeEnum = "DAYS"
	RecurrenceIntervalTypeWeeks   RecurrenceIntervalTypeEnum = "WEEKS"
)

var mappingRecurrenceIntervalType = map[string]RecurrenceIntervalTypeEnum{
	"MINUTES": RecurrenceIntervalTypeMinutes,
	"HOURS":   RecurrenceIntervalTypeHours,
	"DAYS":    RecurrenceIntervalTypeDays,
	"WEEKS":   RecurrenceIntervalTypeWeeks,
}

// GetRecurrenceIntervalTypeEnumValues Enumerates the set of values for RecurrenceIntervalTypeEnum
func GetRecurrenceIntervalTypeEnumValues() []RecurrenceIntervalTypeEnum {
	values := make([]RecurrenceIntervalTypeEnum, 0)
	for _, v := range mappingRecurrenceIntervalType {
		values = append(values, v)
	}
	return values
}
