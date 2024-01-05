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

// DownloadUserAssessmentReportDetails The details used to download a user assessment report.
type DownloadUserAssessmentReportDetails struct {

	// Format of the report.
	Format DownloadUserAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m DownloadUserAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadUserAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadUserAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetDownloadUserAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadUserAssessmentReportDetailsFormatEnum Enum with underlying type: string
type DownloadUserAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for DownloadUserAssessmentReportDetailsFormatEnum
const (
	DownloadUserAssessmentReportDetailsFormatPdf DownloadUserAssessmentReportDetailsFormatEnum = "PDF"
	DownloadUserAssessmentReportDetailsFormatXls DownloadUserAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingDownloadUserAssessmentReportDetailsFormatEnum = map[string]DownloadUserAssessmentReportDetailsFormatEnum{
	"PDF": DownloadUserAssessmentReportDetailsFormatPdf,
	"XLS": DownloadUserAssessmentReportDetailsFormatXls,
}

var mappingDownloadUserAssessmentReportDetailsFormatEnumLowerCase = map[string]DownloadUserAssessmentReportDetailsFormatEnum{
	"pdf": DownloadUserAssessmentReportDetailsFormatPdf,
	"xls": DownloadUserAssessmentReportDetailsFormatXls,
}

// GetDownloadUserAssessmentReportDetailsFormatEnumValues Enumerates the set of values for DownloadUserAssessmentReportDetailsFormatEnum
func GetDownloadUserAssessmentReportDetailsFormatEnumValues() []DownloadUserAssessmentReportDetailsFormatEnum {
	values := make([]DownloadUserAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingDownloadUserAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadUserAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for DownloadUserAssessmentReportDetailsFormatEnum
func GetDownloadUserAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadUserAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadUserAssessmentReportDetailsFormatEnum(val string) (DownloadUserAssessmentReportDetailsFormatEnum, bool) {
	enum, ok := mappingDownloadUserAssessmentReportDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
