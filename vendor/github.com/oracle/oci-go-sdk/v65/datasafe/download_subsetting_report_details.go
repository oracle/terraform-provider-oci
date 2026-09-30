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

// DownloadSubsettingReportDetails Details to download a subsetting report
type DownloadSubsettingReportDetails struct {

	// The OCID of the subsetting report to be downloaded
	ReportId *string `mandatory:"true" json:"reportId"`

	// Format of the report.
	ReportFormat DownloadSubsettingReportDetailsReportFormatEnum `mandatory:"true" json:"reportFormat"`
}

func (m DownloadSubsettingReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadSubsettingReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadSubsettingReportDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetDownloadSubsettingReportDetailsReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadSubsettingReportDetailsReportFormatEnum Enum with underlying type: string
type DownloadSubsettingReportDetailsReportFormatEnum string

// Set of constants representing the allowable values for DownloadSubsettingReportDetailsReportFormatEnum
const (
	DownloadSubsettingReportDetailsReportFormatPdf DownloadSubsettingReportDetailsReportFormatEnum = "PDF"
	DownloadSubsettingReportDetailsReportFormatXls DownloadSubsettingReportDetailsReportFormatEnum = "XLS"
)

var mappingDownloadSubsettingReportDetailsReportFormatEnum = map[string]DownloadSubsettingReportDetailsReportFormatEnum{
	"PDF": DownloadSubsettingReportDetailsReportFormatPdf,
	"XLS": DownloadSubsettingReportDetailsReportFormatXls,
}

var mappingDownloadSubsettingReportDetailsReportFormatEnumLowerCase = map[string]DownloadSubsettingReportDetailsReportFormatEnum{
	"pdf": DownloadSubsettingReportDetailsReportFormatPdf,
	"xls": DownloadSubsettingReportDetailsReportFormatXls,
}

// GetDownloadSubsettingReportDetailsReportFormatEnumValues Enumerates the set of values for DownloadSubsettingReportDetailsReportFormatEnum
func GetDownloadSubsettingReportDetailsReportFormatEnumValues() []DownloadSubsettingReportDetailsReportFormatEnum {
	values := make([]DownloadSubsettingReportDetailsReportFormatEnum, 0)
	for _, v := range mappingDownloadSubsettingReportDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadSubsettingReportDetailsReportFormatEnumStringValues Enumerates the set of values in String for DownloadSubsettingReportDetailsReportFormatEnum
func GetDownloadSubsettingReportDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadSubsettingReportDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadSubsettingReportDetailsReportFormatEnum(val string) (DownloadSubsettingReportDetailsReportFormatEnum, bool) {
	enum, ok := mappingDownloadSubsettingReportDetailsReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
