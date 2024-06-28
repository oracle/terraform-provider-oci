// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleSummary This is the summary information about a schedule.
type ScheduleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// This is the action that will be executed by the schedule.
	Action ScheduleSummaryActionEnum `mandatory:"true" json:"action"`

	// This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting
	// or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field.
	RecurrenceDetails *string `mandatory:"true" json:"recurrenceDetails"`

	// Type of recurrence of a schedule
	RecurrenceType ScheduleSummaryRecurrenceTypeEnum `mandatory:"true" json:"recurrenceType"`

	// This is the current state of the schedule.
	LifecycleState ScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// This is the description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// This is a list of resources filters.  The schedule will be applied to resources matching all of them.
	ResourceFilters []ResourceFilter `mandatory:"false" json:"resourceFilters"`

	// This is the list of resources to which the scheduled operation is applied.
	Resources []Resource `mandatory:"false" json:"resources"`

	// This is the date and time the schedule starts, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarts *common.SDKTime `mandatory:"false" json:"timeStarts"`

	// This is the date and time the schedule ends, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// Example: `2016-08-25T21:10:29.600Z`
	TimeEnds *common.SDKTime `mandatory:"false" json:"timeEnds"`

	// This is the date and time the schedule was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// This is the date and time the schedule was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This is the date and time the schedule runs last time, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeLastRun *common.SDKTime `mandatory:"false" json:"timeLastRun"`

	// This is the date and time the schedule run the next time, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeNextRun *common.SDKTime `mandatory:"false" json:"timeNextRun"`

	// This is the status of the last work request.
	LastRunStatus OperationStatusEnum `mandatory:"false" json:"lastRunStatus,omitempty"`

	// These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// These are defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// These are system tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ScheduleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleSummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetScheduleSummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduleSummaryRecurrenceTypeEnum(string(m.RecurrenceType)); !ok && m.RecurrenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecurrenceType: %s. Supported values are: %s.", m.RecurrenceType, strings.Join(GetScheduleSummaryRecurrenceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOperationStatusEnum(string(m.LastRunStatus)); !ok && m.LastRunStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastRunStatus: %s. Supported values are: %s.", m.LastRunStatus, strings.Join(GetOperationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                           `json:"description"`
		ResourceFilters   []resourcefilter                  `json:"resourceFilters"`
		Resources         []Resource                        `json:"resources"`
		TimeStarts        *common.SDKTime                   `json:"timeStarts"`
		TimeEnds          *common.SDKTime                   `json:"timeEnds"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		TimeLastRun       *common.SDKTime                   `json:"timeLastRun"`
		TimeNextRun       *common.SDKTime                   `json:"timeNextRun"`
		LastRunStatus     OperationStatusEnum               `json:"lastRunStatus"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		CompartmentId     *string                           `json:"compartmentId"`
		DisplayName       *string                           `json:"displayName"`
		Action            ScheduleSummaryActionEnum         `json:"action"`
		RecurrenceDetails *string                           `json:"recurrenceDetails"`
		RecurrenceType    ScheduleSummaryRecurrenceTypeEnum `json:"recurrenceType"`
		LifecycleState    ScheduleLifecycleStateEnum        `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.ResourceFilters = make([]ResourceFilter, len(model.ResourceFilters))
	for i, n := range model.ResourceFilters {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ResourceFilters[i] = nn.(ResourceFilter)
		} else {
			m.ResourceFilters[i] = nil
		}
	}
	m.Resources = make([]Resource, len(model.Resources))
	copy(m.Resources, model.Resources)
	m.TimeStarts = model.TimeStarts

	m.TimeEnds = model.TimeEnds

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.TimeLastRun = model.TimeLastRun

	m.TimeNextRun = model.TimeNextRun

	m.LastRunStatus = model.LastRunStatus

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Action = model.Action

	m.RecurrenceDetails = model.RecurrenceDetails

	m.RecurrenceType = model.RecurrenceType

	m.LifecycleState = model.LifecycleState

	return
}

// ScheduleSummaryActionEnum Enum with underlying type: string
type ScheduleSummaryActionEnum string

// Set of constants representing the allowable values for ScheduleSummaryActionEnum
const (
	ScheduleSummaryActionStartResource ScheduleSummaryActionEnum = "START_RESOURCE"
	ScheduleSummaryActionStopResource  ScheduleSummaryActionEnum = "STOP_RESOURCE"
)

var mappingScheduleSummaryActionEnum = map[string]ScheduleSummaryActionEnum{
	"START_RESOURCE": ScheduleSummaryActionStartResource,
	"STOP_RESOURCE":  ScheduleSummaryActionStopResource,
}

var mappingScheduleSummaryActionEnumLowerCase = map[string]ScheduleSummaryActionEnum{
	"start_resource": ScheduleSummaryActionStartResource,
	"stop_resource":  ScheduleSummaryActionStopResource,
}

// GetScheduleSummaryActionEnumValues Enumerates the set of values for ScheduleSummaryActionEnum
func GetScheduleSummaryActionEnumValues() []ScheduleSummaryActionEnum {
	values := make([]ScheduleSummaryActionEnum, 0)
	for _, v := range mappingScheduleSummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleSummaryActionEnumStringValues Enumerates the set of values in String for ScheduleSummaryActionEnum
func GetScheduleSummaryActionEnumStringValues() []string {
	return []string{
		"START_RESOURCE",
		"STOP_RESOURCE",
	}
}

// GetMappingScheduleSummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleSummaryActionEnum(val string) (ScheduleSummaryActionEnum, bool) {
	enum, ok := mappingScheduleSummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduleSummaryRecurrenceTypeEnum Enum with underlying type: string
type ScheduleSummaryRecurrenceTypeEnum string

// Set of constants representing the allowable values for ScheduleSummaryRecurrenceTypeEnum
const (
	ScheduleSummaryRecurrenceTypeCron ScheduleSummaryRecurrenceTypeEnum = "CRON"
	ScheduleSummaryRecurrenceTypeIcal ScheduleSummaryRecurrenceTypeEnum = "ICAL"
)

var mappingScheduleSummaryRecurrenceTypeEnum = map[string]ScheduleSummaryRecurrenceTypeEnum{
	"CRON": ScheduleSummaryRecurrenceTypeCron,
	"ICAL": ScheduleSummaryRecurrenceTypeIcal,
}

var mappingScheduleSummaryRecurrenceTypeEnumLowerCase = map[string]ScheduleSummaryRecurrenceTypeEnum{
	"cron": ScheduleSummaryRecurrenceTypeCron,
	"ical": ScheduleSummaryRecurrenceTypeIcal,
}

// GetScheduleSummaryRecurrenceTypeEnumValues Enumerates the set of values for ScheduleSummaryRecurrenceTypeEnum
func GetScheduleSummaryRecurrenceTypeEnumValues() []ScheduleSummaryRecurrenceTypeEnum {
	values := make([]ScheduleSummaryRecurrenceTypeEnum, 0)
	for _, v := range mappingScheduleSummaryRecurrenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleSummaryRecurrenceTypeEnumStringValues Enumerates the set of values in String for ScheduleSummaryRecurrenceTypeEnum
func GetScheduleSummaryRecurrenceTypeEnumStringValues() []string {
	return []string{
		"CRON",
		"ICAL",
	}
}

// GetMappingScheduleSummaryRecurrenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleSummaryRecurrenceTypeEnum(val string) (ScheduleSummaryRecurrenceTypeEnum, bool) {
	enum, ok := mappingScheduleSummaryRecurrenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
