// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateScheduleDetails The saved schedule.
type CreateScheduleDetails struct {

	// The unique name of the user-created schedule.
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

	// The description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the supported output file format.
	OutputFileFormat CreateScheduleDetailsOutputFileFormatEnum `mandatory:"false" json:"outputFileFormat,omitempty"`

	// The saved report ID which can also be used to generate a query.
	SavedReportId *string `mandatory:"false" json:"savedReportId"`

	QueryProperties *QueryProperties `mandatory:"false" json:"queryProperties"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
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

	if _, ok := GetMappingCreateScheduleDetailsOutputFileFormatEnum(string(m.OutputFileFormat)); !ok && m.OutputFileFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputFileFormat: %s. Supported values are: %s.", m.OutputFileFormat, strings.Join(GetCreateScheduleDetailsOutputFileFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description         *string                                   `json:"description"`
		OutputFileFormat    CreateScheduleDetailsOutputFileFormatEnum `json:"outputFileFormat"`
		SavedReportId       *string                                   `json:"savedReportId"`
		QueryProperties     *QueryProperties                          `json:"queryProperties"`
		FreeformTags        map[string]string                         `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}         `json:"definedTags"`
		Name                *string                                   `json:"name"`
		CompartmentId       *string                                   `json:"compartmentId"`
		ResultLocation      resultlocation                            `json:"resultLocation"`
		ScheduleRecurrences *string                                   `json:"scheduleRecurrences"`
		TimeScheduled       *common.SDKTime                           `json:"timeScheduled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.OutputFileFormat = model.OutputFileFormat

	m.SavedReportId = model.SavedReportId

	m.QueryProperties = model.QueryProperties

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

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

	return
}

// CreateScheduleDetailsOutputFileFormatEnum Enum with underlying type: string
type CreateScheduleDetailsOutputFileFormatEnum string

// Set of constants representing the allowable values for CreateScheduleDetailsOutputFileFormatEnum
const (
	CreateScheduleDetailsOutputFileFormatCsv CreateScheduleDetailsOutputFileFormatEnum = "CSV"
	CreateScheduleDetailsOutputFileFormatPdf CreateScheduleDetailsOutputFileFormatEnum = "PDF"
)

var mappingCreateScheduleDetailsOutputFileFormatEnum = map[string]CreateScheduleDetailsOutputFileFormatEnum{
	"CSV": CreateScheduleDetailsOutputFileFormatCsv,
	"PDF": CreateScheduleDetailsOutputFileFormatPdf,
}

var mappingCreateScheduleDetailsOutputFileFormatEnumLowerCase = map[string]CreateScheduleDetailsOutputFileFormatEnum{
	"csv": CreateScheduleDetailsOutputFileFormatCsv,
	"pdf": CreateScheduleDetailsOutputFileFormatPdf,
}

// GetCreateScheduleDetailsOutputFileFormatEnumValues Enumerates the set of values for CreateScheduleDetailsOutputFileFormatEnum
func GetCreateScheduleDetailsOutputFileFormatEnumValues() []CreateScheduleDetailsOutputFileFormatEnum {
	values := make([]CreateScheduleDetailsOutputFileFormatEnum, 0)
	for _, v := range mappingCreateScheduleDetailsOutputFileFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateScheduleDetailsOutputFileFormatEnumStringValues Enumerates the set of values in String for CreateScheduleDetailsOutputFileFormatEnum
func GetCreateScheduleDetailsOutputFileFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"PDF",
	}
}

// GetMappingCreateScheduleDetailsOutputFileFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateScheduleDetailsOutputFileFormatEnum(val string) (CreateScheduleDetailsOutputFileFormatEnum, bool) {
	enum, ok := mappingCreateScheduleDetailsOutputFileFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
