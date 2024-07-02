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

// CreateScheduleDetails This is the data to create a schedule.
type CreateScheduleDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the schedule is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// This is the action that will be executed by the schedule.
	Action CreateScheduleDetailsActionEnum `mandatory:"true" json:"action"`

	// This is the frequency of recurrence of a schedule. The frequency field can either conform to RFC-5545 formatting
	// or UNIX cron formatting for recurrences, based on the value specified by the recurrenceType field.
	RecurrenceDetails *string `mandatory:"true" json:"recurrenceDetails"`

	// Type of recurrence of a schedule
	RecurrenceType CreateScheduleDetailsRecurrenceTypeEnum `mandatory:"true" json:"recurrenceType"`

	// This is a user-friendly name for the schedule. It does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

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

	// These are free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// These are defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateScheduleDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetCreateScheduleDetailsActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateScheduleDetailsRecurrenceTypeEnum(string(m.RecurrenceType)); !ok && m.RecurrenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecurrenceType: %s. Supported values are: %s.", m.RecurrenceType, strings.Join(GetCreateScheduleDetailsRecurrenceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName       *string                                 `json:"displayName"`
		Description       *string                                 `json:"description"`
		ResourceFilters   []resourcefilter                        `json:"resourceFilters"`
		Resources         []Resource                              `json:"resources"`
		TimeStarts        *common.SDKTime                         `json:"timeStarts"`
		TimeEnds          *common.SDKTime                         `json:"timeEnds"`
		FreeformTags      map[string]string                       `json:"freeformTags"`
		DefinedTags       map[string]map[string]interface{}       `json:"definedTags"`
		CompartmentId     *string                                 `json:"compartmentId"`
		Action            CreateScheduleDetailsActionEnum         `json:"action"`
		RecurrenceDetails *string                                 `json:"recurrenceDetails"`
		RecurrenceType    CreateScheduleDetailsRecurrenceTypeEnum `json:"recurrenceType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

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

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.Action = model.Action

	m.RecurrenceDetails = model.RecurrenceDetails

	m.RecurrenceType = model.RecurrenceType

	return
}

// CreateScheduleDetailsActionEnum Enum with underlying type: string
type CreateScheduleDetailsActionEnum string

// Set of constants representing the allowable values for CreateScheduleDetailsActionEnum
const (
	CreateScheduleDetailsActionStartResource CreateScheduleDetailsActionEnum = "START_RESOURCE"
	CreateScheduleDetailsActionStopResource  CreateScheduleDetailsActionEnum = "STOP_RESOURCE"
)

var mappingCreateScheduleDetailsActionEnum = map[string]CreateScheduleDetailsActionEnum{
	"START_RESOURCE": CreateScheduleDetailsActionStartResource,
	"STOP_RESOURCE":  CreateScheduleDetailsActionStopResource,
}

var mappingCreateScheduleDetailsActionEnumLowerCase = map[string]CreateScheduleDetailsActionEnum{
	"start_resource": CreateScheduleDetailsActionStartResource,
	"stop_resource":  CreateScheduleDetailsActionStopResource,
}

// GetCreateScheduleDetailsActionEnumValues Enumerates the set of values for CreateScheduleDetailsActionEnum
func GetCreateScheduleDetailsActionEnumValues() []CreateScheduleDetailsActionEnum {
	values := make([]CreateScheduleDetailsActionEnum, 0)
	for _, v := range mappingCreateScheduleDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduleDetailsActionEnumStringValues Enumerates the set of values in String for CreateScheduleDetailsActionEnum
func GetCreateScheduleDetailsActionEnumStringValues() []string {
	return []string{
		"START_RESOURCE",
		"STOP_RESOURCE",
	}
}

// GetMappingCreateScheduleDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduleDetailsActionEnum(val string) (CreateScheduleDetailsActionEnum, bool) {
	enum, ok := mappingCreateScheduleDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateScheduleDetailsRecurrenceTypeEnum Enum with underlying type: string
type CreateScheduleDetailsRecurrenceTypeEnum string

// Set of constants representing the allowable values for CreateScheduleDetailsRecurrenceTypeEnum
const (
	CreateScheduleDetailsRecurrenceTypeCron CreateScheduleDetailsRecurrenceTypeEnum = "CRON"
	CreateScheduleDetailsRecurrenceTypeIcal CreateScheduleDetailsRecurrenceTypeEnum = "ICAL"
)

var mappingCreateScheduleDetailsRecurrenceTypeEnum = map[string]CreateScheduleDetailsRecurrenceTypeEnum{
	"CRON": CreateScheduleDetailsRecurrenceTypeCron,
	"ICAL": CreateScheduleDetailsRecurrenceTypeIcal,
}

var mappingCreateScheduleDetailsRecurrenceTypeEnumLowerCase = map[string]CreateScheduleDetailsRecurrenceTypeEnum{
	"cron": CreateScheduleDetailsRecurrenceTypeCron,
	"ical": CreateScheduleDetailsRecurrenceTypeIcal,
}

// GetCreateScheduleDetailsRecurrenceTypeEnumValues Enumerates the set of values for CreateScheduleDetailsRecurrenceTypeEnum
func GetCreateScheduleDetailsRecurrenceTypeEnumValues() []CreateScheduleDetailsRecurrenceTypeEnum {
	values := make([]CreateScheduleDetailsRecurrenceTypeEnum, 0)
	for _, v := range mappingCreateScheduleDetailsRecurrenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduleDetailsRecurrenceTypeEnumStringValues Enumerates the set of values in String for CreateScheduleDetailsRecurrenceTypeEnum
func GetCreateScheduleDetailsRecurrenceTypeEnumStringValues() []string {
	return []string{
		"CRON",
		"ICAL",
	}
}

// GetMappingCreateScheduleDetailsRecurrenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduleDetailsRecurrenceTypeEnum(val string) (CreateScheduleDetailsRecurrenceTypeEnum, bool) {
	enum, ok := mappingCreateScheduleDetailsRecurrenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
