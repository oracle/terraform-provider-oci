// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TriggerSchedule Specifies a trigger schedule. Timing information for when to initiate automated syncs.
type TriggerSchedule struct {

	// Different types of trigger schedule:
	// None - No automated synchronization schedule.
	// Default - Trigger schedule is every 30 minutes.
	// Custom - Custom triggering schedule.
	ScheduleType TriggerScheduleScheduleTypeEnum `mandatory:"true" json:"scheduleType"`

	// Valid if type is CUSTOM. Following RFC 5545 recurrence rules, we can specify starting time, occurrence frequency, and interval size.
	// Example for frequency could be DAILY/WEEKLY/HOURLY or any RFC 5545 supported frequency, which is followed by start time of this window.
	// You can control the start time with BYHOUR, BYMINUTE and BYSECONDS. It is followed by the interval size.
	CustomSchedule *string `mandatory:"false" json:"customSchedule"`
}

func (m TriggerSchedule) String() string {
	return common.PointerString(m)
}

// TriggerScheduleScheduleTypeEnum Enum with underlying type: string
type TriggerScheduleScheduleTypeEnum string

// Set of constants representing the allowable values for TriggerScheduleScheduleTypeEnum
const (
	TriggerScheduleScheduleTypeNone    TriggerScheduleScheduleTypeEnum = "NONE"
	TriggerScheduleScheduleTypeDefault TriggerScheduleScheduleTypeEnum = "DEFAULT"
	TriggerScheduleScheduleTypeCustom  TriggerScheduleScheduleTypeEnum = "CUSTOM"
)

var mappingTriggerScheduleScheduleType = map[string]TriggerScheduleScheduleTypeEnum{
	"NONE":    TriggerScheduleScheduleTypeNone,
	"DEFAULT": TriggerScheduleScheduleTypeDefault,
	"CUSTOM":  TriggerScheduleScheduleTypeCustom,
}

// GetTriggerScheduleScheduleTypeEnumValues Enumerates the set of values for TriggerScheduleScheduleTypeEnum
func GetTriggerScheduleScheduleTypeEnumValues() []TriggerScheduleScheduleTypeEnum {
	values := make([]TriggerScheduleScheduleTypeEnum, 0)
	for _, v := range mappingTriggerScheduleScheduleType {
		values = append(values, v)
	}
	return values
}
