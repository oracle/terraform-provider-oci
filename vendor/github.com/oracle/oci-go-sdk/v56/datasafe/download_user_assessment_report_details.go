// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DownloadUserAssessmentReportDetails The details used to download a user assessment report.
type DownloadUserAssessmentReportDetails struct {

	// Format of the report.
	Format DownloadUserAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m DownloadUserAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// DownloadUserAssessmentReportDetailsFormatEnum Enum with underlying type: string
type DownloadUserAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for DownloadUserAssessmentReportDetailsFormatEnum
const (
	DownloadUserAssessmentReportDetailsFormatPdf DownloadUserAssessmentReportDetailsFormatEnum = "PDF"
	DownloadUserAssessmentReportDetailsFormatXls DownloadUserAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingDownloadUserAssessmentReportDetailsFormat = map[string]DownloadUserAssessmentReportDetailsFormatEnum{
	"PDF": DownloadUserAssessmentReportDetailsFormatPdf,
	"XLS": DownloadUserAssessmentReportDetailsFormatXls,
}

// GetDownloadUserAssessmentReportDetailsFormatEnumValues Enumerates the set of values for DownloadUserAssessmentReportDetailsFormatEnum
func GetDownloadUserAssessmentReportDetailsFormatEnumValues() []DownloadUserAssessmentReportDetailsFormatEnum {
	values := make([]DownloadUserAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingDownloadUserAssessmentReportDetailsFormat {
		values = append(values, v)
	}
	return values
}
