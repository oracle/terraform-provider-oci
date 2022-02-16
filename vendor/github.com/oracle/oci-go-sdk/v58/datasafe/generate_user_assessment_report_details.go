// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// GenerateUserAssessmentReportDetails The details used to generate a new user assessment report.
type GenerateUserAssessmentReportDetails struct {

	// Format of the report.
	Format GenerateUserAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m GenerateUserAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateUserAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateUserAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetGenerateUserAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateUserAssessmentReportDetailsFormatEnum Enum with underlying type: string
type GenerateUserAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for GenerateUserAssessmentReportDetailsFormatEnum
const (
	GenerateUserAssessmentReportDetailsFormatPdf GenerateUserAssessmentReportDetailsFormatEnum = "PDF"
	GenerateUserAssessmentReportDetailsFormatXls GenerateUserAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingGenerateUserAssessmentReportDetailsFormatEnum = map[string]GenerateUserAssessmentReportDetailsFormatEnum{
	"PDF": GenerateUserAssessmentReportDetailsFormatPdf,
	"XLS": GenerateUserAssessmentReportDetailsFormatXls,
}

// GetGenerateUserAssessmentReportDetailsFormatEnumValues Enumerates the set of values for GenerateUserAssessmentReportDetailsFormatEnum
func GetGenerateUserAssessmentReportDetailsFormatEnumValues() []GenerateUserAssessmentReportDetailsFormatEnum {
	values := make([]GenerateUserAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingGenerateUserAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateUserAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for GenerateUserAssessmentReportDetailsFormatEnum
func GetGenerateUserAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateUserAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateUserAssessmentReportDetailsFormatEnum(val string) (GenerateUserAssessmentReportDetailsFormatEnum, bool) {
	mappingGenerateUserAssessmentReportDetailsFormatEnumIgnoreCase := make(map[string]GenerateUserAssessmentReportDetailsFormatEnum)
	for k, v := range mappingGenerateUserAssessmentReportDetailsFormatEnum {
		mappingGenerateUserAssessmentReportDetailsFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGenerateUserAssessmentReportDetailsFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
