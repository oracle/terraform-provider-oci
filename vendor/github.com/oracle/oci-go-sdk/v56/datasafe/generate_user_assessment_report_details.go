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

// GenerateUserAssessmentReportDetails The details used to generate a new user assessment report.
type GenerateUserAssessmentReportDetails struct {

	// Format of the report.
	Format GenerateUserAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m GenerateUserAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// GenerateUserAssessmentReportDetailsFormatEnum Enum with underlying type: string
type GenerateUserAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for GenerateUserAssessmentReportDetailsFormatEnum
const (
	GenerateUserAssessmentReportDetailsFormatPdf GenerateUserAssessmentReportDetailsFormatEnum = "PDF"
	GenerateUserAssessmentReportDetailsFormatXls GenerateUserAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingGenerateUserAssessmentReportDetailsFormat = map[string]GenerateUserAssessmentReportDetailsFormatEnum{
	"PDF": GenerateUserAssessmentReportDetailsFormatPdf,
	"XLS": GenerateUserAssessmentReportDetailsFormatXls,
}

// GetGenerateUserAssessmentReportDetailsFormatEnumValues Enumerates the set of values for GenerateUserAssessmentReportDetailsFormatEnum
func GetGenerateUserAssessmentReportDetailsFormatEnumValues() []GenerateUserAssessmentReportDetailsFormatEnum {
	values := make([]GenerateUserAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingGenerateUserAssessmentReportDetailsFormat {
		values = append(values, v)
	}
	return values
}
