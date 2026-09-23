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

// DownloadCryptoAssessmentReportDetails The details used to download a crypto assessment report.
type DownloadCryptoAssessmentReportDetails struct {

	// Format of the crypto assessment report.
	Format DownloadCryptoAssessmentReportDetailsFormatEnum `mandatory:"true" json:"format"`
}

func (m DownloadCryptoAssessmentReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadCryptoAssessmentReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDownloadCryptoAssessmentReportDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetDownloadCryptoAssessmentReportDetailsFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadCryptoAssessmentReportDetailsFormatEnum Enum with underlying type: string
type DownloadCryptoAssessmentReportDetailsFormatEnum string

// Set of constants representing the allowable values for DownloadCryptoAssessmentReportDetailsFormatEnum
const (
	DownloadCryptoAssessmentReportDetailsFormatPdf DownloadCryptoAssessmentReportDetailsFormatEnum = "PDF"
	DownloadCryptoAssessmentReportDetailsFormatXls DownloadCryptoAssessmentReportDetailsFormatEnum = "XLS"
)

var mappingDownloadCryptoAssessmentReportDetailsFormatEnum = map[string]DownloadCryptoAssessmentReportDetailsFormatEnum{
	"PDF": DownloadCryptoAssessmentReportDetailsFormatPdf,
	"XLS": DownloadCryptoAssessmentReportDetailsFormatXls,
}

var mappingDownloadCryptoAssessmentReportDetailsFormatEnumLowerCase = map[string]DownloadCryptoAssessmentReportDetailsFormatEnum{
	"pdf": DownloadCryptoAssessmentReportDetailsFormatPdf,
	"xls": DownloadCryptoAssessmentReportDetailsFormatXls,
}

// GetDownloadCryptoAssessmentReportDetailsFormatEnumValues Enumerates the set of values for DownloadCryptoAssessmentReportDetailsFormatEnum
func GetDownloadCryptoAssessmentReportDetailsFormatEnumValues() []DownloadCryptoAssessmentReportDetailsFormatEnum {
	values := make([]DownloadCryptoAssessmentReportDetailsFormatEnum, 0)
	for _, v := range mappingDownloadCryptoAssessmentReportDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadCryptoAssessmentReportDetailsFormatEnumStringValues Enumerates the set of values in String for DownloadCryptoAssessmentReportDetailsFormatEnum
func GetDownloadCryptoAssessmentReportDetailsFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadCryptoAssessmentReportDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadCryptoAssessmentReportDetailsFormatEnum(val string) (DownloadCryptoAssessmentReportDetailsFormatEnum, bool) {
	enum, ok := mappingDownloadCryptoAssessmentReportDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
