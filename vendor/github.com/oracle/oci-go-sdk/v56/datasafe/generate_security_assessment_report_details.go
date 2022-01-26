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

// GenerateSecurityAssessmentReportDetails The details used to generate a new security assessment report.
type GenerateSecurityAssessmentReportDetails struct {

	// Format of the report.
	Format GenerateSecurityAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m GenerateSecurityAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// GenerateSecurityAssessmentReportDetailsFormatEnum Enum with underlying type: string
type GenerateSecurityAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for GenerateSecurityAssessmentReportDetailsFormatEnum
const (
	GenerateSecurityAssessmentReportDetailsFormatPdf GenerateSecurityAssessmentReportDetailsFormatEnum = "PDF"
	GenerateSecurityAssessmentReportDetailsFormatXls GenerateSecurityAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingGenerateSecurityAssessmentReportDetailsFormat = map[string]GenerateSecurityAssessmentReportDetailsFormatEnum{
	"PDF": GenerateSecurityAssessmentReportDetailsFormatPdf,
	"XLS": GenerateSecurityAssessmentReportDetailsFormatXls,
}

// GetGenerateSecurityAssessmentReportDetailsFormatEnumValues Enumerates the set of values for GenerateSecurityAssessmentReportDetailsFormatEnum
func GetGenerateSecurityAssessmentReportDetailsFormatEnumValues() []GenerateSecurityAssessmentReportDetailsFormatEnum {
	values := make([]GenerateSecurityAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingGenerateSecurityAssessmentReportDetailsFormat {
		values = append(values, v)
	}
	return values
}
