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

// GenerateCryptoAssessmentReportDetails The details used to generate a new crypto assessment report.
type GenerateCryptoAssessmentReportDetails struct {

	// Format of the crypto assessment report.
	Format GenerateCryptoAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m GenerateCryptoAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateCryptoAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateCryptoAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetGenerateCryptoAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateCryptoAssessmentReportDetailsFormatEnum Enum with underlying type: string
type GenerateCryptoAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for GenerateCryptoAssessmentReportDetailsFormatEnum
const (
	GenerateCryptoAssessmentReportDetailsFormatPdf GenerateCryptoAssessmentReportDetailsFormatEnum = "PDF"
	GenerateCryptoAssessmentReportDetailsFormatXls GenerateCryptoAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingGenerateCryptoAssessmentReportDetailsFormatEnum = map[string]GenerateCryptoAssessmentReportDetailsFormatEnum{
	"PDF": GenerateCryptoAssessmentReportDetailsFormatPdf,
	"XLS": GenerateCryptoAssessmentReportDetailsFormatXls,
}

var mappingGenerateCryptoAssessmentReportDetailsFormatEnumLowerCase = map[string]GenerateCryptoAssessmentReportDetailsFormatEnum{
	"pdf": GenerateCryptoAssessmentReportDetailsFormatPdf,
	"xls": GenerateCryptoAssessmentReportDetailsFormatXls,
}

// GetGenerateCryptoAssessmentReportDetailsFormatEnumValues Enumerates the set of values for GenerateCryptoAssessmentReportDetailsFormatEnum
func GetGenerateCryptoAssessmentReportDetailsFormatEnumValues() []GenerateCryptoAssessmentReportDetailsFormatEnum {
	values := make([]GenerateCryptoAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingGenerateCryptoAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateCryptoAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for GenerateCryptoAssessmentReportDetailsFormatEnum
func GetGenerateCryptoAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateCryptoAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateCryptoAssessmentReportDetailsFormatEnum(val string) (GenerateCryptoAssessmentReportDetailsFormatEnum, bool) {
	enum, ok := mappingGenerateCryptoAssessmentReportDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
