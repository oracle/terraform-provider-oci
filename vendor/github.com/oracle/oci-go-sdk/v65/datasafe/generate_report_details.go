// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateReportDetails Details for the report generation.
type GenerateReportDetails struct {

	// The name of the report to be generated
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// into which the resource should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies the format of report to be .xls or .pdf or .json
	MimeType GenerateReportDetailsMimeTypeEnum `mandatory:"true" json:"mimeType"`

	// Array of database target OCIDs.
	TargetIds []string `mandatory:"false" json:"targetIds"`

	// The description of the report to be generated
	Description *string `mandatory:"false" json:"description"`

	// Specifies the time until which the data needs to be reported.
	TimeLessThan *common.SDKTime `mandatory:"false" json:"timeLessThan"`

	// Specifies the time after which the data needs to be reported.
	TimeGreaterThan *common.SDKTime `mandatory:"false" json:"timeGreaterThan"`

	// Specifies the limit on the number of rows in the report.
	RowLimit *int `mandatory:"false" json:"rowLimit"`
}

func (m GenerateReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateReportDetailsMimeTypeEnum(string(m.MimeType)); !ok && m.MimeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MimeType: %s. Supported values are: %s.", m.MimeType, strings.Join(GetGenerateReportDetailsMimeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateReportDetailsMimeTypeEnum Enum with underlying type: string
type GenerateReportDetailsMimeTypeEnum string

// Set of constants representing the allowable values for GenerateReportDetailsMimeTypeEnum
const (
	GenerateReportDetailsMimeTypePdf  GenerateReportDetailsMimeTypeEnum = "PDF"
	GenerateReportDetailsMimeTypeXls  GenerateReportDetailsMimeTypeEnum = "XLS"
	GenerateReportDetailsMimeTypeJson GenerateReportDetailsMimeTypeEnum = "JSON"
)

var mappingGenerateReportDetailsMimeTypeEnum = map[string]GenerateReportDetailsMimeTypeEnum{
	"PDF":  GenerateReportDetailsMimeTypePdf,
	"XLS":  GenerateReportDetailsMimeTypeXls,
	"JSON": GenerateReportDetailsMimeTypeJson,
}

var mappingGenerateReportDetailsMimeTypeEnumLowerCase = map[string]GenerateReportDetailsMimeTypeEnum{
	"pdf":  GenerateReportDetailsMimeTypePdf,
	"xls":  GenerateReportDetailsMimeTypeXls,
	"json": GenerateReportDetailsMimeTypeJson,
}

// GetGenerateReportDetailsMimeTypeEnumValues Enumerates the set of values for GenerateReportDetailsMimeTypeEnum
func GetGenerateReportDetailsMimeTypeEnumValues() []GenerateReportDetailsMimeTypeEnum {
	values := make([]GenerateReportDetailsMimeTypeEnum, 0)
	for _, v := range mappingGenerateReportDetailsMimeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateReportDetailsMimeTypeEnumStringValues Enumerates the set of values in String for GenerateReportDetailsMimeTypeEnum
func GetGenerateReportDetailsMimeTypeEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
		"JSON",
	}
}

// GetMappingGenerateReportDetailsMimeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateReportDetailsMimeTypeEnum(val string) (GenerateReportDetailsMimeTypeEnum, bool) {
	enum, ok := mappingGenerateReportDetailsMimeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
