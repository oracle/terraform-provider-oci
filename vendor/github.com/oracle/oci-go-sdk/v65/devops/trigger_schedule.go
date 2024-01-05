// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TriggerSchedule Specifies a trigger schedule. Timing information for when to initiate automated syncs.
type TriggerSchedule struct {

	// Different types of trigger schedule:
	// NONE - No automated synchronization schedule.
	// DEFAULT - Trigger schedule is every 30 minutes.
	// CUSTOM - Custom triggering schedule.
	ScheduleType TriggerScheduleScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// Valid if type is CUSTOM. Following RFC 5545 recurrence rules, we can specify starting time, occurrence frequency, and interval size.
	// Example for frequency could be DAILY/WEEKLY/HOURLY or any RFC 5545 supported frequency, which is followed by start time of this window.
	// You can control the start time with BYHOUR, BYMINUTE and BYSECONDS. It is followed by the interval size.
	CustomSchedule *string `mandatory:"false" json:"customSchedule"`
}

func (m TriggerSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TriggerSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTriggerScheduleScheduleTypeEnum(string(m.ScheduleType)); !ok && m.ScheduleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduleType: %s. Supported values are: %s.", m.ScheduleType, strings.Join(GetTriggerScheduleScheduleTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TriggerScheduleScheduleTypeEnum Enum with underlying type: string
type TriggerScheduleScheduleTypeEnum string

// Set of constants representing the allowable values for TriggerScheduleScheduleTypeEnum
const (
	TriggerScheduleScheduleTypeNone    TriggerScheduleScheduleTypeEnum = "NONE"
	TriggerScheduleScheduleTypeDefault TriggerScheduleScheduleTypeEnum = "DEFAULT"
	TriggerScheduleScheduleTypeCustom  TriggerScheduleScheduleTypeEnum = "CUSTOM"
)

var mappingTriggerScheduleScheduleTypeEnum = map[string]TriggerScheduleScheduleTypeEnum{
	"NONE":    TriggerScheduleScheduleTypeNone,
	"DEFAULT": TriggerScheduleScheduleTypeDefault,
	"CUSTOM":  TriggerScheduleScheduleTypeCustom,
}

var mappingTriggerScheduleScheduleTypeEnumLowerCase = map[string]TriggerScheduleScheduleTypeEnum{
	"none":    TriggerScheduleScheduleTypeNone,
	"default": TriggerScheduleScheduleTypeDefault,
	"custom":  TriggerScheduleScheduleTypeCustom,
}

// GetTriggerScheduleScheduleTypeEnumValues Enumerates the set of values for TriggerScheduleScheduleTypeEnum
func GetTriggerScheduleScheduleTypeEnumValues() []TriggerScheduleScheduleTypeEnum {
	values := make([]TriggerScheduleScheduleTypeEnum, 0)
	for _, v := range mappingTriggerScheduleScheduleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTriggerScheduleScheduleTypeEnumStringValues Enumerates the set of values in String for TriggerScheduleScheduleTypeEnum
func GetTriggerScheduleScheduleTypeEnumStringValues() []string {
	return []string{
		"NONE",
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingTriggerScheduleScheduleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTriggerScheduleScheduleTypeEnum(val string) (TriggerScheduleScheduleTypeEnum, bool) {
	enum, ok := mappingTriggerScheduleScheduleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
