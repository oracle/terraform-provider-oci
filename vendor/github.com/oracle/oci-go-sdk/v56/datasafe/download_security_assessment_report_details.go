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

// DownloadSecurityAssessmentReportDetails The details used to download a security assessment report.
type DownloadSecurityAssessmentReportDetails struct {

	// Format of the report.
	Format DownloadSecurityAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m DownloadSecurityAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// DownloadSecurityAssessmentReportDetailsFormatEnum Enum with underlying type: string
type DownloadSecurityAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for DownloadSecurityAssessmentReportDetailsFormatEnum
const (
	DownloadSecurityAssessmentReportDetailsFormatPdf DownloadSecurityAssessmentReportDetailsFormatEnum = "PDF"
	DownloadSecurityAssessmentReportDetailsFormatXls DownloadSecurityAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingDownloadSecurityAssessmentReportDetailsFormat = map[string]DownloadSecurityAssessmentReportDetailsFormatEnum{
	"PDF": DownloadSecurityAssessmentReportDetailsFormatPdf,
	"XLS": DownloadSecurityAssessmentReportDetailsFormatXls,
}

// GetDownloadSecurityAssessmentReportDetailsFormatEnumValues Enumerates the set of values for DownloadSecurityAssessmentReportDetailsFormatEnum
func GetDownloadSecurityAssessmentReportDetailsFormatEnumValues() []DownloadSecurityAssessmentReportDetailsFormatEnum {
	values := make([]DownloadSecurityAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingDownloadSecurityAssessmentReportDetailsFormat {
		values = append(values, v)
	}
	return values
}
