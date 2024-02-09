// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by the Cost Analysis and Carbon Emissions Analysis tools in the Console. See Cost Analysis Overview (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm) and Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Schedule The schedule.
type Schedule struct {

	// The OCID representing a unique shedule.
	Id *string `mandatory:"true" json:"id"`

	// The unique name of the schedule created by the user.
	Name *string `mandatory:"true" json:"name"`

	// The customer tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ResultLocation ResultLocation `mandatory:"true" json:"resultLocation"`

	// Specifies the frequency according to when the schedule will be run,
	// in the x-obmcs-recurring-time format described in RFC 5545 section 3.3.10 (https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10).
	// Supported values are : ONE_TIME, DAILY, WEEKLY and MONTHLY.
	ScheduleRecurrences *string `mandatory:"true" json:"scheduleRecurrences"`

	// The date and time of the first time job execution.
	TimeScheduled *common.SDKTime `mandatory:"true" json:"timeScheduled"`

	// The date and time the schedule was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The schedule lifecycle state.
	LifecycleState ScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// The date and time of the next job execution.
	TimeNextRun *common.SDKTime `mandatory:"false" json:"timeNextRun"`

	// Specifies the supported output file format.
	OutputFileFormat ScheduleOutputFileFormatEnum `mandatory:"false" json:"outputFileFormat,omitempty"`

	// The saved report ID which can also be used to generate a query.
	SavedReportId *string `mandatory:"false" json:"savedReportId"`

	QueryProperties *QueryProperties `mandatory:"false" json:"queryProperties"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
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
	if _, ok := GetMappingScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScheduleOutputFileFormatEnum(string(m.OutputFileFormat)); !ok && m.OutputFileFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputFileFormat: %s. Supported values are: %s.", m.OutputFileFormat, strings.Join(GetScheduleOutputFileFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Schedule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description         *string                           `json:"description"`
		TimeNextRun         *common.SDKTime                   `json:"timeNextRun"`
		OutputFileFormat    ScheduleOutputFileFormatEnum      `json:"outputFileFormat"`
		SavedReportId       *string                           `json:"savedReportId"`
		QueryProperties     *QueryProperties                  `json:"queryProperties"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		Id                  *string                           `json:"id"`
		Name                *string                           `json:"name"`
		CompartmentId       *string                           `json:"compartmentId"`
		ResultLocation      resultlocation                    `json:"resultLocation"`
		ScheduleRecurrences *string                           `json:"scheduleRecurrences"`
		TimeScheduled       *common.SDKTime                   `json:"timeScheduled"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		LifecycleState      ScheduleLifecycleStateEnum        `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeNextRun = model.TimeNextRun

	m.OutputFileFormat = model.OutputFileFormat

	m.SavedReportId = model.SavedReportId

	m.QueryProperties = model.QueryProperties

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.CompartmentId = model.CompartmentId

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(ResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.ScheduleRecurrences = model.ScheduleRecurrences

	m.TimeScheduled = model.TimeScheduled

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	return
}

// ScheduleOutputFileFormatEnum Enum with underlying type: string
type ScheduleOutputFileFormatEnum string

// Set of constants representing the allowable values for ScheduleOutputFileFormatEnum
const (
	ScheduleOutputFileFormatCsv ScheduleOutputFileFormatEnum = "CSV"
	ScheduleOutputFileFormatPdf ScheduleOutputFileFormatEnum = "PDF"
)

var mappingScheduleOutputFileFormatEnum = map[string]ScheduleOutputFileFormatEnum{
	"CSV": ScheduleOutputFileFormatCsv,
	"PDF": ScheduleOutputFileFormatPdf,
}

var mappingScheduleOutputFileFormatEnumLowerCase = map[string]ScheduleOutputFileFormatEnum{
	"csv": ScheduleOutputFileFormatCsv,
	"pdf": ScheduleOutputFileFormatPdf,
}

// GetScheduleOutputFileFormatEnumValues Enumerates the set of values for ScheduleOutputFileFormatEnum
func GetScheduleOutputFileFormatEnumValues() []ScheduleOutputFileFormatEnum {
	values := make([]ScheduleOutputFileFormatEnum, 0)
	for _, v := range mappingScheduleOutputFileFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleOutputFileFormatEnumStringValues Enumerates the set of values in String for ScheduleOutputFileFormatEnum
func GetScheduleOutputFileFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"PDF",
	}
}

// GetMappingScheduleOutputFileFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleOutputFileFormatEnum(val string) (ScheduleOutputFileFormatEnum, bool) {
	enum, ok := mappingScheduleOutputFileFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ScheduleLifecycleStateEnum Enum with underlying type: string
type ScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for ScheduleLifecycleStateEnum
const (
	ScheduleLifecycleStateActive   ScheduleLifecycleStateEnum = "ACTIVE"
	ScheduleLifecycleStateInactive ScheduleLifecycleStateEnum = "INACTIVE"
)

var mappingScheduleLifecycleStateEnum = map[string]ScheduleLifecycleStateEnum{
	"ACTIVE":   ScheduleLifecycleStateActive,
	"INACTIVE": ScheduleLifecycleStateInactive,
}

var mappingScheduleLifecycleStateEnumLowerCase = map[string]ScheduleLifecycleStateEnum{
	"active":   ScheduleLifecycleStateActive,
	"inactive": ScheduleLifecycleStateInactive,
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
	}
}

// GetMappingScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleLifecycleStateEnum(val string) (ScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
