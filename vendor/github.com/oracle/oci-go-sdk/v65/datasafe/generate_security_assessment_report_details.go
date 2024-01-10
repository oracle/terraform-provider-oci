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

// GenerateSecurityAssessmentReportDetails The details used to generate a new security assessment report.
type GenerateSecurityAssessmentReportDetails struct {

	// Format of the report.
	Format GenerateSecurityAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m GenerateSecurityAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSecurityAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateSecurityAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetGenerateSecurityAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateSecurityAssessmentReportDetailsFormatEnum Enum with underlying type: string
type GenerateSecurityAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for GenerateSecurityAssessmentReportDetailsFormatEnum
const (
	GenerateSecurityAssessmentReportDetailsFormatPdf GenerateSecurityAssessmentReportDetailsFormatEnum = "PDF"
	GenerateSecurityAssessmentReportDetailsFormatXls GenerateSecurityAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingGenerateSecurityAssessmentReportDetailsFormatEnum = map[string]GenerateSecurityAssessmentReportDetailsFormatEnum{
	"PDF": GenerateSecurityAssessmentReportDetailsFormatPdf,
	"XLS": GenerateSecurityAssessmentReportDetailsFormatXls,
}

var mappingGenerateSecurityAssessmentReportDetailsFormatEnumLowerCase = map[string]GenerateSecurityAssessmentReportDetailsFormatEnum{
	"pdf": GenerateSecurityAssessmentReportDetailsFormatPdf,
	"xls": GenerateSecurityAssessmentReportDetailsFormatXls,
}

// GetGenerateSecurityAssessmentReportDetailsFormatEnumValues Enumerates the set of values for GenerateSecurityAssessmentReportDetailsFormatEnum
func GetGenerateSecurityAssessmentReportDetailsFormatEnumValues() []GenerateSecurityAssessmentReportDetailsFormatEnum {
	values := make([]GenerateSecurityAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingGenerateSecurityAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateSecurityAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for GenerateSecurityAssessmentReportDetailsFormatEnum
func GetGenerateSecurityAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateSecurityAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateSecurityAssessmentReportDetailsFormatEnum(val string) (GenerateSecurityAssessmentReportDetailsFormatEnum, bool) {
	enum, ok := mappingGenerateSecurityAssessmentReportDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
