// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduleReportDetails The details of the report schedule.
type ScheduleReportDetails struct {

	// The schedule to generate the report periodically in the specified format:
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// 4. <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// No constraint introduced when it is '*'. When not, day of week must equal the given value
	// 5. <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"true" json:"schedule"`

	// Specifies if the report will be in .xls or .pdf or .json format
	MimeType ScheduleReportDetailsMimeTypeEnum `mandatory:"true" json:"mimeType"`

	// The OCID of the compartment
	// in which the resource should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ReportDetails ReportDetails `mandatory:"true" json:"reportDetails"`

	// The name of the report to be scheduled
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m ScheduleReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduleReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingScheduleReportDetailsMimeTypeEnum(string(m.MimeType)); !ok && m.MimeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MimeType: %s. Supported values are: %s.", m.MimeType, strings.Join(GetScheduleReportDetailsMimeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ScheduleReportDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName   *string                           `json:"displayName"`
		Schedule      *string                           `json:"schedule"`
		MimeType      ScheduleReportDetailsMimeTypeEnum `json:"mimeType"`
		CompartmentId *string                           `json:"compartmentId"`
		ReportDetails reportdetails                     `json:"reportDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Schedule = model.Schedule

	m.MimeType = model.MimeType

	m.CompartmentId = model.CompartmentId

	nn, e = model.ReportDetails.UnmarshalPolymorphicJSON(model.ReportDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ReportDetails = nn.(ReportDetails)
	} else {
		m.ReportDetails = nil
	}

	return
}

// ScheduleReportDetailsMimeTypeEnum Enum with underlying type: string
type ScheduleReportDetailsMimeTypeEnum string

// Set of constants representing the allowable values for ScheduleReportDetailsMimeTypeEnum
const (
	ScheduleReportDetailsMimeTypePdf  ScheduleReportDetailsMimeTypeEnum = "PDF"
	ScheduleReportDetailsMimeTypeXls  ScheduleReportDetailsMimeTypeEnum = "XLS"
	ScheduleReportDetailsMimeTypeJson ScheduleReportDetailsMimeTypeEnum = "JSON"
)

var mappingScheduleReportDetailsMimeTypeEnum = map[string]ScheduleReportDetailsMimeTypeEnum{
	"PDF":  ScheduleReportDetailsMimeTypePdf,
	"XLS":  ScheduleReportDetailsMimeTypeXls,
	"JSON": ScheduleReportDetailsMimeTypeJson,
}

var mappingScheduleReportDetailsMimeTypeEnumLowerCase = map[string]ScheduleReportDetailsMimeTypeEnum{
	"pdf":  ScheduleReportDetailsMimeTypePdf,
	"xls":  ScheduleReportDetailsMimeTypeXls,
	"json": ScheduleReportDetailsMimeTypeJson,
}

// GetScheduleReportDetailsMimeTypeEnumValues Enumerates the set of values for ScheduleReportDetailsMimeTypeEnum
func GetScheduleReportDetailsMimeTypeEnumValues() []ScheduleReportDetailsMimeTypeEnum {
	values := make([]ScheduleReportDetailsMimeTypeEnum, 0)
	for _, v := range mappingScheduleReportDetailsMimeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScheduleReportDetailsMimeTypeEnumStringValues Enumerates the set of values in String for ScheduleReportDetailsMimeTypeEnum
func GetScheduleReportDetailsMimeTypeEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
		"JSON",
	}
}

// GetMappingScheduleReportDetailsMimeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScheduleReportDetailsMimeTypeEnum(val string) (ScheduleReportDetailsMimeTypeEnum, bool) {
	enum, ok := mappingScheduleReportDetailsMimeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
