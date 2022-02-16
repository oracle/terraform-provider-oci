// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Recurrence) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRecurrenceIntervalTypeEnum(string(m.IntervalType)); !ok && m.IntervalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IntervalType: %s. Supported values are: %s.", m.IntervalType, strings.Join(GetRecurrenceIntervalTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingRecurrenceIntervalTypeEnum = map[string]RecurrenceIntervalTypeEnum{
	"MINUTES": RecurrenceIntervalTypeMinutes,
	"HOURS":   RecurrenceIntervalTypeHours,
	"DAYS":    RecurrenceIntervalTypeDays,
	"WEEKS":   RecurrenceIntervalTypeWeeks,
}

// GetRecurrenceIntervalTypeEnumValues Enumerates the set of values for RecurrenceIntervalTypeEnum
func GetRecurrenceIntervalTypeEnumValues() []RecurrenceIntervalTypeEnum {
	values := make([]RecurrenceIntervalTypeEnum, 0)
	for _, v := range mappingRecurrenceIntervalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecurrenceIntervalTypeEnumStringValues Enumerates the set of values in String for RecurrenceIntervalTypeEnum
func GetRecurrenceIntervalTypeEnumStringValues() []string {
	return []string{
		"MINUTES",
		"HOURS",
		"DAYS",
		"WEEKS",
	}
}

// GetMappingRecurrenceIntervalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecurrenceIntervalTypeEnum(val string) (RecurrenceIntervalTypeEnum, bool) {
	mappingRecurrenceIntervalTypeEnumIgnoreCase := make(map[string]RecurrenceIntervalTypeEnum)
	for k, v := range mappingRecurrenceIntervalTypeEnum {
		mappingRecurrenceIntervalTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRecurrenceIntervalTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
