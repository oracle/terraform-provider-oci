// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportRunbookDetails Request to Export Runbook.
type ExportRunbookDetails struct {
	ContentDestination TransferRunbookContentDetails `mandatory:"true" json:"contentDestination"`

	// Export Runbook As Type.
	ExportAs ExportRunbookDetailsExportAsEnum `mandatory:"false" json:"exportAs,omitempty"`
}

func (m ExportRunbookDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportRunbookDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExportRunbookDetailsExportAsEnum(string(m.ExportAs)); !ok && m.ExportAs != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExportAs: %s. Supported values are: %s.", m.ExportAs, strings.Join(GetExportRunbookDetailsExportAsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExportRunbookDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ExportAs           ExportRunbookDetailsExportAsEnum `json:"exportAs"`
		ContentDestination transferrunbookcontentdetails    `json:"contentDestination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ExportAs = model.ExportAs

	nn, e = model.ContentDestination.UnmarshalPolymorphicJSON(model.ContentDestination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ContentDestination = nn.(TransferRunbookContentDetails)
	} else {
		m.ContentDestination = nil
	}

	return
}

// ExportRunbookDetailsExportAsEnum Enum with underlying type: string
type ExportRunbookDetailsExportAsEnum string

// Set of constants representing the allowable values for ExportRunbookDetailsExportAsEnum
const (
	ExportRunbookDetailsExportAsTerraform    ExportRunbookDetailsExportAsEnum = "TERRAFORM"
	ExportRunbookDetailsExportAsNonTerraform ExportRunbookDetailsExportAsEnum = "NON_TERRAFORM"
)

var mappingExportRunbookDetailsExportAsEnum = map[string]ExportRunbookDetailsExportAsEnum{
	"TERRAFORM":     ExportRunbookDetailsExportAsTerraform,
	"NON_TERRAFORM": ExportRunbookDetailsExportAsNonTerraform,
}

var mappingExportRunbookDetailsExportAsEnumLowerCase = map[string]ExportRunbookDetailsExportAsEnum{
	"terraform":     ExportRunbookDetailsExportAsTerraform,
	"non_terraform": ExportRunbookDetailsExportAsNonTerraform,
}

// GetExportRunbookDetailsExportAsEnumValues Enumerates the set of values for ExportRunbookDetailsExportAsEnum
func GetExportRunbookDetailsExportAsEnumValues() []ExportRunbookDetailsExportAsEnum {
	values := make([]ExportRunbookDetailsExportAsEnum, 0)
	for _, v := range mappingExportRunbookDetailsExportAsEnum {
		values = append(values, v)
	}
	return values
}

// GetExportRunbookDetailsExportAsEnumStringValues Enumerates the set of values in String for ExportRunbookDetailsExportAsEnum
func GetExportRunbookDetailsExportAsEnumStringValues() []string {
	return []string{
		"TERRAFORM",
		"NON_TERRAFORM",
	}
}

// GetMappingExportRunbookDetailsExportAsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportRunbookDetailsExportAsEnum(val string) (ExportRunbookDetailsExportAsEnum, bool) {
	enum, ok := mappingExportRunbookDetailsExportAsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
