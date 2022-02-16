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

// DownloadSecurityAssessmentReportDetails The details used to download a security assessment report.
type DownloadSecurityAssessmentReportDetails struct {

	// Format of the report.
	Format DownloadSecurityAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m DownloadSecurityAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadSecurityAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadSecurityAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetDownloadSecurityAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadSecurityAssessmentReportDetailsFormatEnum Enum with underlying type: string
type DownloadSecurityAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for DownloadSecurityAssessmentReportDetailsFormatEnum
const (
	DownloadSecurityAssessmentReportDetailsFormatPdf DownloadSecurityAssessmentReportDetailsFormatEnum = "PDF"
	DownloadSecurityAssessmentReportDetailsFormatXls DownloadSecurityAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingDownloadSecurityAssessmentReportDetailsFormatEnum = map[string]DownloadSecurityAssessmentReportDetailsFormatEnum{
	"PDF": DownloadSecurityAssessmentReportDetailsFormatPdf,
	"XLS": DownloadSecurityAssessmentReportDetailsFormatXls,
}

// GetDownloadSecurityAssessmentReportDetailsFormatEnumValues Enumerates the set of values for DownloadSecurityAssessmentReportDetailsFormatEnum
func GetDownloadSecurityAssessmentReportDetailsFormatEnumValues() []DownloadSecurityAssessmentReportDetailsFormatEnum {
	values := make([]DownloadSecurityAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingDownloadSecurityAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadSecurityAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for DownloadSecurityAssessmentReportDetailsFormatEnum
func GetDownloadSecurityAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadSecurityAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadSecurityAssessmentReportDetailsFormatEnum(val string) (DownloadSecurityAssessmentReportDetailsFormatEnum, bool) {
	mappingDownloadSecurityAssessmentReportDetailsFormatEnumIgnoreCase := make(map[string]DownloadSecurityAssessmentReportDetailsFormatEnum)
	for k, v := range mappingDownloadSecurityAssessmentReportDetailsFormatEnum {
		mappingDownloadSecurityAssessmentReportDetailsFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDownloadSecurityAssessmentReportDetailsFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
