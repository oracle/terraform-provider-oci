// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateScheduleDetails Details for updating the custom table.
type UpdateScheduleDetails struct {

	// The description of the schedule.
	Description *string `mandatory:"false" json:"description"`

	// Specifies the supported output file format.
	OutputFileFormat UpdateScheduleDetailsOutputFileFormatEnum `mandatory:"false" json:"outputFileFormat,omitempty"`

	ResultLocation ResultLocation `mandatory:"false" json:"resultLocation"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
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

	if _, ok := GetMappingUpdateScheduleDetailsOutputFileFormatEnum(string(m.OutputFileFormat)); !ok && m.OutputFileFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputFileFormat: %s. Supported values are: %s.", m.OutputFileFormat, strings.Join(GetUpdateScheduleDetailsOutputFileFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateScheduleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                                   `json:"description"`
		OutputFileFormat UpdateScheduleDetailsOutputFileFormatEnum `json:"outputFileFormat"`
		ResultLocation   resultlocation                            `json:"resultLocation"`
		FreeformTags     map[string]string                         `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{}         `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.OutputFileFormat = model.OutputFileFormat

	nn, e = model.ResultLocation.UnmarshalPolymorphicJSON(model.ResultLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocation = nn.(ResultLocation)
	} else {
		m.ResultLocation = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}

// UpdateScheduleDetailsOutputFileFormatEnum Enum with underlying type: string
type UpdateScheduleDetailsOutputFileFormatEnum string

// Set of constants representing the allowable values for UpdateScheduleDetailsOutputFileFormatEnum
const (
	UpdateScheduleDetailsOutputFileFormatCsv UpdateScheduleDetailsOutputFileFormatEnum = "CSV"
	UpdateScheduleDetailsOutputFileFormatPdf UpdateScheduleDetailsOutputFileFormatEnum = "PDF"
)

var mappingUpdateScheduleDetailsOutputFileFormatEnum = map[string]UpdateScheduleDetailsOutputFileFormatEnum{
	"CSV": UpdateScheduleDetailsOutputFileFormatCsv,
	"PDF": UpdateScheduleDetailsOutputFileFormatPdf,
}

var mappingUpdateScheduleDetailsOutputFileFormatEnumLowerCase = map[string]UpdateScheduleDetailsOutputFileFormatEnum{
	"csv": UpdateScheduleDetailsOutputFileFormatCsv,
	"pdf": UpdateScheduleDetailsOutputFileFormatPdf,
}

// GetUpdateScheduleDetailsOutputFileFormatEnumValues Enumerates the set of values for UpdateScheduleDetailsOutputFileFormatEnum
func GetUpdateScheduleDetailsOutputFileFormatEnumValues() []UpdateScheduleDetailsOutputFileFormatEnum {
	values := make([]UpdateScheduleDetailsOutputFileFormatEnum, 0)
	for _, v := range mappingUpdateScheduleDetailsOutputFileFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateScheduleDetailsOutputFileFormatEnumStringValues Enumerates the set of values in String for UpdateScheduleDetailsOutputFileFormatEnum
func GetUpdateScheduleDetailsOutputFileFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"PDF",
	}
}

// GetMappingUpdateScheduleDetailsOutputFileFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateScheduleDetailsOutputFileFormatEnum(val string) (UpdateScheduleDetailsOutputFileFormatEnum, bool) {
	enum, ok := mappingUpdateScheduleDetailsOutputFileFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
