// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
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
