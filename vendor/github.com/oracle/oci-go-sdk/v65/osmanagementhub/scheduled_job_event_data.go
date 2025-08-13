// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledJobEventData Provides information about a scheduled job event.
type ScheduledJobEventData struct {

	// Type of management station operation.
	OperationType ScheduledJobEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the scheduled job operation.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	// The time this scheduled job is set to be paused (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeStartPause *common.SDKTime `mandatory:"false" json:"timeStartPause"`

	// The time this scheduled job is set to be resumed (in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) format).
	TimeStartResume *common.SDKTime `mandatory:"false" json:"timeStartResume"`
}

func (m ScheduledJobEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledJobEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduledJobEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetScheduledJobEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScheduledJobEventDataOperationTypeEnum Enum with underlying type: string
type ScheduledJobEventDataOperationTypeEnum string

// Set of constants representing the allowable values for ScheduledJobEventDataOperationTypeEnum
const (
	ScheduledJobEventDataOperationTypeImmediatelyPause  ScheduledJobEventDataOperationTypeEnum = "IMMEDIATELY_PAUSE"
	ScheduledJobEventDataOperationTypeSchedulePause     ScheduledJobEventDataOperationTypeEnum = "SCHEDULE_PAUSE"
	ScheduledJobEventDataOperationTypeCancelPause       ScheduledJobEventDataOperationTypeEnum = "CANCEL_PAUSE"
	ScheduledJobEventDataOperationTypeImmediatelyResume ScheduledJobEventDataOperationTypeEnum = "IMMEDIATELY_RESUME"
	ScheduledJobEventDataOperationTypeScheduleResume    ScheduledJobEventDataOperationTypeEnum = "SCHEDULE_RESUME"
	ScheduledJobEventDataOperationTypeCancelResume      ScheduledJobEventDataOperationTypeEnum = "CANCEL_RESUME"
)

var mappingScheduledJobEventDataOperationTypeEnum = map[string]ScheduledJobEventDataOperationTypeEnum{
	"IMMEDIATELY_PAUSE":  ScheduledJobEventDataOperationTypeImmediatelyPause,
	"SCHEDULE_PAUSE":     ScheduledJobEventDataOperationTypeSchedulePause,
	"CANCEL_PAUSE":       ScheduledJobEventDataOperationTypeCancelPause,
	"IMMEDIATELY_RESUME": ScheduledJobEventDataOperationTypeImmediatelyResume,
	"SCHEDULE_RESUME":    ScheduledJobEventDataOperationTypeScheduleResume,
	"CANCEL_RESUME":      ScheduledJobEventDataOperationTypeCancelResume,
}

var mappingScheduledJobEventDataOperationTypeEnumLowerCase = map[string]ScheduledJobEventDataOperationTypeEnum{
	"immediately_pause":  ScheduledJobEventDataOperationTypeImmediatelyPause,
	"schedule_pause":     ScheduledJobEventDataOperationTypeSchedulePause,
	"cancel_pause":       ScheduledJobEventDataOperationTypeCancelPause,
	"immediately_resume": ScheduledJobEventDataOperationTypeImmediatelyResume,
	"schedule_resume":    ScheduledJobEventDataOperationTypeScheduleResume,
	"cancel_resume":      ScheduledJobEventDataOperationTypeCancelResume,
}

// GetScheduledJobEventDataOperationTypeEnumValues Enumerates the set of values for ScheduledJobEventDataOperationTypeEnum
func GetScheduledJobEventDataOperationTypeEnumValues() []ScheduledJobEventDataOperationTypeEnum {
	values := make([]ScheduledJobEventDataOperationTypeEnum, 0)
	for _, v := range mappingScheduledJobEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduledJobEventDataOperationTypeEnumStringValues Enumerates the set of values in String for ScheduledJobEventDataOperationTypeEnum
func GetScheduledJobEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"IMMEDIATELY_PAUSE",
		"SCHEDULE_PAUSE",
		"CANCEL_PAUSE",
		"IMMEDIATELY_RESUME",
		"SCHEDULE_RESUME",
		"CANCEL_RESUME",
	}
}

// GetMappingScheduledJobEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduledJobEventDataOperationTypeEnum(val string) (ScheduledJobEventDataOperationTypeEnum, bool) {
	enum, ok := mappingScheduledJobEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
