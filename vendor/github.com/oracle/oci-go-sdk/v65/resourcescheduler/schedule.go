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

// Schedule A Schedule describes the date and time when an operation will be or has been applied to a set of resources. You must specify either
// the resources directly or provide a set of resource filters to select the resources.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, contact your
// administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Schedule struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the schedule
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// This is the action that will be executed by the schedule.
	Action ScheduleActionEnum `mandatory:"true" json:"action"`

	// This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting
	// or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field.
	RecurrenceDetails *string `mandatory:"true" json:"recurrenceDetails"`

	// Type of recurrence of a schedule
	RecurrenceType ScheduleRecurrenceTypeEnum `mandatory:"true" json:"recurrenceType"`

	// This is the date and time the schedule was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// This is the current state of a schedule.
	LifecycleState ScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// These are defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

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

	// This is the date and time the schedule was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// This is the date and time the schedule runs last time, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeLastRun *common.SDKTime `mandatory:"false" json:"timeLastRun"`

	// This is the date and time the schedule run the next time, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeNextRun *common.SDKTime `mandatory:"false" json:"timeNextRun"`

	// These are system tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Schedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Schedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetScheduleActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduleRecurrenceTypeEnum(string(m.RecurrenceType)); !ok && m.RecurrenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecurrenceType: %s. Supported values are: %s.", m.RecurrenceType, strings.Join(GetScheduleRecurrenceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Schedule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description       *string                           `json:"description"`
		ResourceFilters   []resourcefilter                  `json:"resourceFilters"`
		Resources         []Resource                        `json:"resources"`
		TimeStarts        *common.SDKTime                   `json:"timeStarts"`
		TimeEnds          *common.SDKTime                   `json:"timeEnds"`
		TimeUpdated       *common.SDKTime                   `json:"timeUpdated"`
		TimeLastRun       *common.SDKTime                   `json:"timeLastRun"`
		TimeNextRun       *common.SDKTime                   `json:"timeNextRun"`
		SystemTags        map[string]map[string]interface{} `json:"systemTags"`
		Id                *string                           `json:"id"`
		CompartmentId     *string                           `json:"compartmentId"`
		DisplayName       *string                           `json:"displayName"`
		Action            ScheduleActionEnum                `json:"action"`
		RecurrenceDetails *string                           `json:"recurrenceDetails"`
		RecurrenceType    ScheduleRecurrenceTypeEnum        `json:"recurrenceType"`
		TimeCreated       *common.SDKTime                   `json:"timeCreated"`
		LifecycleState    ScheduleLifecycleStateEnum        `json:"lifecycleState"`
		FreeformTags      map[string]string                 `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{} `json:"definedTags"`
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

	m.TimeUpdated = model.TimeUpdated

	m.TimeLastRun = model.TimeLastRun

	m.TimeNextRun = model.TimeNextRun

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.Action = model.Action

	m.RecurrenceDetails = model.RecurrenceDetails

	m.RecurrenceType = model.RecurrenceType

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// ScheduleActionEnum Enum with underlying type: string
type ScheduleActionEnum string

// Set of constants representing the allowable values for ScheduleActionEnum
const (
	ScheduleActionStartResource ScheduleActionEnum = "START_RESOURCE"
	ScheduleActionStopResource  ScheduleActionEnum = "STOP_RESOURCE"
)

var mappingScheduleActionEnum = map[string]ScheduleActionEnum{
	"START_RESOURCE": ScheduleActionStartResource,
	"STOP_RESOURCE":  ScheduleActionStopResource,
}

var mappingScheduleActionEnumLowerCase = map[string]ScheduleActionEnum{
	"start_resource": ScheduleActionStartResource,
	"stop_resource":  ScheduleActionStopResource,
}

// GetScheduleActionEnumValues Enumerates the set of values for ScheduleActionEnum
func GetScheduleActionEnumValues() []ScheduleActionEnum {
	values := make([]ScheduleActionEnum, 0)
	for _, v := range mappingScheduleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleActionEnumStringValues Enumerates the set of values in String for ScheduleActionEnum
func GetScheduleActionEnumStringValues() []string {
	return []string{
		"START_RESOURCE",
		"STOP_RESOURCE",
	}
}

// GetMappingScheduleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleActionEnum(val string) (ScheduleActionEnum, bool) {
	enum, ok := mappingScheduleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduleRecurrenceTypeEnum Enum with underlying type: string
type ScheduleRecurrenceTypeEnum string

// Set of constants representing the allowable values for ScheduleRecurrenceTypeEnum
const (
	ScheduleRecurrenceTypeCron ScheduleRecurrenceTypeEnum = "CRON"
	ScheduleRecurrenceTypeIcal ScheduleRecurrenceTypeEnum = "ICAL"
)

var mappingScheduleRecurrenceTypeEnum = map[string]ScheduleRecurrenceTypeEnum{
	"CRON": ScheduleRecurrenceTypeCron,
	"ICAL": ScheduleRecurrenceTypeIcal,
}

var mappingScheduleRecurrenceTypeEnumLowerCase = map[string]ScheduleRecurrenceTypeEnum{
	"cron": ScheduleRecurrenceTypeCron,
	"ical": ScheduleRecurrenceTypeIcal,
}

// GetScheduleRecurrenceTypeEnumValues Enumerates the set of values for ScheduleRecurrenceTypeEnum
func GetScheduleRecurrenceTypeEnumValues() []ScheduleRecurrenceTypeEnum {
	values := make([]ScheduleRecurrenceTypeEnum, 0)
	for _, v := range mappingScheduleRecurrenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleRecurrenceTypeEnumStringValues Enumerates the set of values in String for ScheduleRecurrenceTypeEnum
func GetScheduleRecurrenceTypeEnumStringValues() []string {
	return []string{
		"CRON",
		"ICAL",
	}
}

// GetMappingScheduleRecurrenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleRecurrenceTypeEnum(val string) (ScheduleRecurrenceTypeEnum, bool) {
	enum, ok := mappingScheduleRecurrenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduleLifecycleStateEnum Enum with underlying type: string
type ScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduleLifecycleStateEnum
const (
	ScheduleLifecycleStateActive   ScheduleLifecycleStateEnum = "ACTIVE"
	ScheduleLifecycleStateInactive ScheduleLifecycleStateEnum = "INACTIVE"
	ScheduleLifecycleStateCreating ScheduleLifecycleStateEnum = "CREATING"
	ScheduleLifecycleStateUpdating ScheduleLifecycleStateEnum = "UPDATING"
	ScheduleLifecycleStateDeleting ScheduleLifecycleStateEnum = "DELETING"
	ScheduleLifecycleStateDeleted  ScheduleLifecycleStateEnum = "DELETED"
	ScheduleLifecycleStateFailed   ScheduleLifecycleStateEnum = "FAILED"
)

var mappingScheduleLifecycleStateEnum = map[string]ScheduleLifecycleStateEnum{
	"ACTIVE":   ScheduleLifecycleStateActive,
	"INACTIVE": ScheduleLifecycleStateInactive,
	"CREATING": ScheduleLifecycleStateCreating,
	"UPDATING": ScheduleLifecycleStateUpdating,
	"DELETING": ScheduleLifecycleStateDeleting,
	"DELETED":  ScheduleLifecycleStateDeleted,
	"FAILED":   ScheduleLifecycleStateFailed,
}

var mappingScheduleLifecycleStateEnumLowerCase = map[string]ScheduleLifecycleStateEnum{
	"active":   ScheduleLifecycleStateActive,
	"inactive": ScheduleLifecycleStateInactive,
	"creating": ScheduleLifecycleStateCreating,
	"updating": ScheduleLifecycleStateUpdating,
	"deleting": ScheduleLifecycleStateDeleting,
	"deleted":  ScheduleLifecycleStateDeleted,
	"failed":   ScheduleLifecycleStateFailed,
}

// GetScheduleLifecycleStateEnumValues Enumerates the set of values for ScheduleLifecycleStateEnum
func GetScheduleLifecycleStateEnumValues() []ScheduleLifecycleStateEnum {
	values := make([]ScheduleLifecycleStateEnum, 0)
	for _, v := range mappingScheduleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleLifecycleStateEnumStringValues Enumerates the set of values in String for ScheduleLifecycleStateEnum
func GetScheduleLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleLifecycleStateEnum(val string) (ScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
