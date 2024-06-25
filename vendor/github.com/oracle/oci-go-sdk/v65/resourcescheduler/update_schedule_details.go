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

// UpdateScheduleDetails This is the data to update a schedule.
type UpdateScheduleDetails struct {

	// This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// This is the description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// This is the action that will be executed by the schedule.
	Action UpdateScheduleDetailsActionEnum `mandatory:"false" json:"action,omitempty"`

	// This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting
	// or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field.
	RecurrenceDetails *string `mandatory:"false" json:"recurrenceDetails"`

	// Type of recurrence of a schedule
	RecurrenceType UpdateScheduleDetailsRecurrenceTypeEnum `mandatory:"false" json:"recurrenceType,omitempty"`

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

	// These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// These are defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateScheduleDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUpdateScheduleDetailsActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateScheduleDetailsRecurrenceTypeEnum(string(m.RecurrenceType)); !ok && m.RecurrenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecurrenceType: %s. Supported values are: %s.", m.RecurrenceType, strings.Join(GetUpdateScheduleDetailsRecurrenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                                 `json:"displayName"`
		Description       *string                                 `json:"description"`
		Action            UpdateScheduleDetailsActionEnum         `json:"action"`
		RecurrenceDetails *string                                 `json:"recurrenceDetails"`
		RecurrenceType    UpdateScheduleDetailsRecurrenceTypeEnum `json:"recurrenceType"`
		ResourceFilters   []resourcefilter                        `json:"resourceFilters"`
		Resources         []Resource                              `json:"resources"`
		TimeStarts        *common.SDKTime                         `json:"timeStarts"`
		TimeEnds          *common.SDKTime                         `json:"timeEnds"`
		FreeformTags      map[string]string                       `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{}       `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.Action = model.Action

	m.RecurrenceDetails = model.RecurrenceDetails

	m.RecurrenceType = model.RecurrenceType

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// UpdateScheduleDetailsActionEnum Enum with underlying type: string
type UpdateScheduleDetailsActionEnum string

// Set of constants representing the allowable values for UpdateScheduleDetailsActionEnum
const (
	UpdateScheduleDetailsActionStartResource UpdateScheduleDetailsActionEnum = "START_RESOURCE"
	UpdateScheduleDetailsActionStopResource  UpdateScheduleDetailsActionEnum = "STOP_RESOURCE"
)

var mappingUpdateScheduleDetailsActionEnum = map[string]UpdateScheduleDetailsActionEnum{
	"START_RESOURCE": UpdateScheduleDetailsActionStartResource,
	"STOP_RESOURCE":  UpdateScheduleDetailsActionStopResource,
}

var mappingUpdateScheduleDetailsActionEnumLowerCase = map[string]UpdateScheduleDetailsActionEnum{
	"start_resource": UpdateScheduleDetailsActionStartResource,
	"stop_resource":  UpdateScheduleDetailsActionStopResource,
}

// GetUpdateScheduleDetailsActionEnumValues Enumerates the set of values for UpdateScheduleDetailsActionEnum
func GetUpdateScheduleDetailsActionEnumValues() []UpdateScheduleDetailsActionEnum {
	values := make([]UpdateScheduleDetailsActionEnum, 0)
	for _, v := range mappingUpdateScheduleDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateScheduleDetailsActionEnumStringValues Enumerates the set of values in String for UpdateScheduleDetailsActionEnum
func GetUpdateScheduleDetailsActionEnumStringValues() []string {
	return []string{
		"START_RESOURCE",
		"STOP_RESOURCE",
	}
}

// GetMappingUpdateScheduleDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateScheduleDetailsActionEnum(val string) (UpdateScheduleDetailsActionEnum, bool) {
	enum, ok := mappingUpdateScheduleDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateScheduleDetailsRecurrenceTypeEnum Enum with underlying type: string
type UpdateScheduleDetailsRecurrenceTypeEnum string

// Set of constants representing the allowable values for UpdateScheduleDetailsRecurrenceTypeEnum
const (
	UpdateScheduleDetailsRecurrenceTypeCron UpdateScheduleDetailsRecurrenceTypeEnum = "CRON"
	UpdateScheduleDetailsRecurrenceTypeIcal UpdateScheduleDetailsRecurrenceTypeEnum = "ICAL"
)

var mappingUpdateScheduleDetailsRecurrenceTypeEnum = map[string]UpdateScheduleDetailsRecurrenceTypeEnum{
	"CRON": UpdateScheduleDetailsRecurrenceTypeCron,
	"ICAL": UpdateScheduleDetailsRecurrenceTypeIcal,
}

var mappingUpdateScheduleDetailsRecurrenceTypeEnumLowerCase = map[string]UpdateScheduleDetailsRecurrenceTypeEnum{
	"cron": UpdateScheduleDetailsRecurrenceTypeCron,
	"ical": UpdateScheduleDetailsRecurrenceTypeIcal,
}

// GetUpdateScheduleDetailsRecurrenceTypeEnumValues Enumerates the set of values for UpdateScheduleDetailsRecurrenceTypeEnum
func GetUpdateScheduleDetailsRecurrenceTypeEnumValues() []UpdateScheduleDetailsRecurrenceTypeEnum {
	values := make([]UpdateScheduleDetailsRecurrenceTypeEnum, 0)
	for _, v := range mappingUpdateScheduleDetailsRecurrenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateScheduleDetailsRecurrenceTypeEnumStringValues Enumerates the set of values in String for UpdateScheduleDetailsRecurrenceTypeEnum
func GetUpdateScheduleDetailsRecurrenceTypeEnumStringValues() []string {
	return []string{
		"CRON",
		"ICAL",
	}
}

// GetMappingUpdateScheduleDetailsRecurrenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateScheduleDetailsRecurrenceTypeEnum(val string) (UpdateScheduleDetailsRecurrenceTypeEnum, bool) {
	enum, ok := mappingUpdateScheduleDetailsRecurrenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
