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

// GenerateMaskingReportForDownloadDetails Details to generate a downloadable masking report.
type GenerateMaskingReportForDownloadDetails struct {

	// The OCID of the masking report for which a downloadable file is to be generated.
	ReportId *string `mandatory:"true" json:"reportId"`

	// Format of the report.
	ReportFormat GenerateMaskingReportForDownloadDetailsReportFormatEnum `mandatory:"true" json:"reportFormat"`
}

func (m GenerateMaskingReportForDownloadDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateMaskingReportForDownloadDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateMaskingReportForDownloadDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetGenerateMaskingReportForDownloadDetailsReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateMaskingReportForDownloadDetailsReportFormatEnum Enum with underlying type: string
type GenerateMaskingReportForDownloadDetailsReportFormatEnum string

// Set of constants representing the allowable values for GenerateMaskingReportForDownloadDetailsReportFormatEnum
const (
	GenerateMaskingReportForDownloadDetailsReportFormatPdf GenerateMaskingReportForDownloadDetailsReportFormatEnum = "PDF"
	GenerateMaskingReportForDownloadDetailsReportFormatXls GenerateMaskingReportForDownloadDetailsReportFormatEnum = "XLS"
)

var mappingGenerateMaskingReportForDownloadDetailsReportFormatEnum = map[string]GenerateMaskingReportForDownloadDetailsReportFormatEnum{
	"PDF": GenerateMaskingReportForDownloadDetailsReportFormatPdf,
	"XLS": GenerateMaskingReportForDownloadDetailsReportFormatXls,
}

var mappingGenerateMaskingReportForDownloadDetailsReportFormatEnumLowerCase = map[string]GenerateMaskingReportForDownloadDetailsReportFormatEnum{
	"pdf": GenerateMaskingReportForDownloadDetailsReportFormatPdf,
	"xls": GenerateMaskingReportForDownloadDetailsReportFormatXls,
}

// GetGenerateMaskingReportForDownloadDetailsReportFormatEnumValues Enumerates the set of values for GenerateMaskingReportForDownloadDetailsReportFormatEnum
func GetGenerateMaskingReportForDownloadDetailsReportFormatEnumValues() []GenerateMaskingReportForDownloadDetailsReportFormatEnum {
	values := make([]GenerateMaskingReportForDownloadDetailsReportFormatEnum, 0)
	for _, v := range mappingGenerateMaskingReportForDownloadDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateMaskingReportForDownloadDetailsReportFormatEnumStringValues Enumerates the set of values in String for GenerateMaskingReportForDownloadDetailsReportFormatEnum
func GetGenerateMaskingReportForDownloadDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateMaskingReportForDownloadDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateMaskingReportForDownloadDetailsReportFormatEnum(val string) (GenerateMaskingReportForDownloadDetailsReportFormatEnum, bool) {
	enum, ok := mappingGenerateMaskingReportForDownloadDetailsReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
