// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// GenerateSubsettingReportForDownloadDetails Details to generate a downloadable subsetting report
type GenerateSubsettingReportForDownloadDetails struct {

	// The OCID of the subsetting report for which a downloadable file is to be generated
	ReportId *string `mandatory:"true" json:"reportId"`

	// Format of the report.
	ReportFormat GenerateSubsettingReportForDownloadDetailsReportFormatEnum `mandatory:"true" json:"reportFormat"`
}

func (m GenerateSubsettingReportForDownloadDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSubsettingReportForDownloadDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateSubsettingReportForDownloadDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetGenerateSubsettingReportForDownloadDetailsReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateSubsettingReportForDownloadDetailsReportFormatEnum Enum with underlying type: string
type GenerateSubsettingReportForDownloadDetailsReportFormatEnum string

// Set of constants representing the allowable values for GenerateSubsettingReportForDownloadDetailsReportFormatEnum
const (
	GenerateSubsettingReportForDownloadDetailsReportFormatPdf GenerateSubsettingReportForDownloadDetailsReportFormatEnum = "PDF"
	GenerateSubsettingReportForDownloadDetailsReportFormatXls GenerateSubsettingReportForDownloadDetailsReportFormatEnum = "XLS"
)

var mappingGenerateSubsettingReportForDownloadDetailsReportFormatEnum = map[string]GenerateSubsettingReportForDownloadDetailsReportFormatEnum{
	"PDF": GenerateSubsettingReportForDownloadDetailsReportFormatPdf,
	"XLS": GenerateSubsettingReportForDownloadDetailsReportFormatXls,
}

var mappingGenerateSubsettingReportForDownloadDetailsReportFormatEnumLowerCase = map[string]GenerateSubsettingReportForDownloadDetailsReportFormatEnum{
	"pdf": GenerateSubsettingReportForDownloadDetailsReportFormatPdf,
	"xls": GenerateSubsettingReportForDownloadDetailsReportFormatXls,
}

// GetGenerateSubsettingReportForDownloadDetailsReportFormatEnumValues Enumerates the set of values for GenerateSubsettingReportForDownloadDetailsReportFormatEnum
func GetGenerateSubsettingReportForDownloadDetailsReportFormatEnumValues() []GenerateSubsettingReportForDownloadDetailsReportFormatEnum {
	values := make([]GenerateSubsettingReportForDownloadDetailsReportFormatEnum, 0)
	for _, v := range mappingGenerateSubsettingReportForDownloadDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateSubsettingReportForDownloadDetailsReportFormatEnumStringValues Enumerates the set of values in String for GenerateSubsettingReportForDownloadDetailsReportFormatEnum
func GetGenerateSubsettingReportForDownloadDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateSubsettingReportForDownloadDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateSubsettingReportForDownloadDetailsReportFormatEnum(val string) (GenerateSubsettingReportForDownloadDetailsReportFormatEnum, bool) {
	enum, ok := mappingGenerateSubsettingReportForDownloadDetailsReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
