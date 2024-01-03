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

// DownloadMaskingReportDetails Details to download a masking report.
type DownloadMaskingReportDetails struct {

	// The OCID of the masking report to be downloaded.
	ReportId *string `mandatory:"true" json:"reportId"`

	// Format of the report.
	ReportFormat DownloadMaskingReportDetailsReportFormatEnum `mandatory:"true" json:"reportFormat"`
}

func (m DownloadMaskingReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadMaskingReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadMaskingReportDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetDownloadMaskingReportDetailsReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadMaskingReportDetailsReportFormatEnum Enum with underlying type: string
type DownloadMaskingReportDetailsReportFormatEnum string

// Set of constants representing the allowable values for DownloadMaskingReportDetailsReportFormatEnum
const (
	DownloadMaskingReportDetailsReportFormatPdf DownloadMaskingReportDetailsReportFormatEnum = "PDF"
	DownloadMaskingReportDetailsReportFormatXls DownloadMaskingReportDetailsReportFormatEnum = "XLS"
)

var mappingDownloadMaskingReportDetailsReportFormatEnum = map[string]DownloadMaskingReportDetailsReportFormatEnum{
	"PDF": DownloadMaskingReportDetailsReportFormatPdf,
	"XLS": DownloadMaskingReportDetailsReportFormatXls,
}

var mappingDownloadMaskingReportDetailsReportFormatEnumLowerCase = map[string]DownloadMaskingReportDetailsReportFormatEnum{
	"pdf": DownloadMaskingReportDetailsReportFormatPdf,
	"xls": DownloadMaskingReportDetailsReportFormatXls,
}

// GetDownloadMaskingReportDetailsReportFormatEnumValues Enumerates the set of values for DownloadMaskingReportDetailsReportFormatEnum
func GetDownloadMaskingReportDetailsReportFormatEnumValues() []DownloadMaskingReportDetailsReportFormatEnum {
	values := make([]DownloadMaskingReportDetailsReportFormatEnum, 0)
	for _, v := range mappingDownloadMaskingReportDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadMaskingReportDetailsReportFormatEnumStringValues Enumerates the set of values in String for DownloadMaskingReportDetailsReportFormatEnum
func GetDownloadMaskingReportDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadMaskingReportDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadMaskingReportDetailsReportFormatEnum(val string) (DownloadMaskingReportDetailsReportFormatEnum, bool) {
	enum, ok := mappingDownloadMaskingReportDetailsReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
