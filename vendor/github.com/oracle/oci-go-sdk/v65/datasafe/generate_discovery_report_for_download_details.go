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

// GenerateDiscoveryReportForDownloadDetails Details to generate a downloadable discovery report.
type GenerateDiscoveryReportForDownloadDetails struct {

	// Format of the report.
	ReportFormat GenerateDiscoveryReportForDownloadDetailsReportFormatEnum `mandatory:"true" json:"reportFormat"`

	// The OCID of the discovery job.
	DiscoveryJobId *string `mandatory:"false" json:"discoveryJobId"`
}

func (m GenerateDiscoveryReportForDownloadDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateDiscoveryReportForDownloadDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetGenerateDiscoveryReportForDownloadDetailsReportFormatEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateDiscoveryReportForDownloadDetailsReportFormatEnum Enum with underlying type: string
type GenerateDiscoveryReportForDownloadDetailsReportFormatEnum string

// Set of constants representing the allowable values for GenerateDiscoveryReportForDownloadDetailsReportFormatEnum
const (
	GenerateDiscoveryReportForDownloadDetailsReportFormatPdf GenerateDiscoveryReportForDownloadDetailsReportFormatEnum = "PDF"
	GenerateDiscoveryReportForDownloadDetailsReportFormatXls GenerateDiscoveryReportForDownloadDetailsReportFormatEnum = "XLS"
)

var mappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnum = map[string]GenerateDiscoveryReportForDownloadDetailsReportFormatEnum{
	"PDF": GenerateDiscoveryReportForDownloadDetailsReportFormatPdf,
	"XLS": GenerateDiscoveryReportForDownloadDetailsReportFormatXls,
}

var mappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnumLowerCase = map[string]GenerateDiscoveryReportForDownloadDetailsReportFormatEnum{
	"pdf": GenerateDiscoveryReportForDownloadDetailsReportFormatPdf,
	"xls": GenerateDiscoveryReportForDownloadDetailsReportFormatXls,
}

// GetGenerateDiscoveryReportForDownloadDetailsReportFormatEnumValues Enumerates the set of values for GenerateDiscoveryReportForDownloadDetailsReportFormatEnum
func GetGenerateDiscoveryReportForDownloadDetailsReportFormatEnumValues() []GenerateDiscoveryReportForDownloadDetailsReportFormatEnum {
	values := make([]GenerateDiscoveryReportForDownloadDetailsReportFormatEnum, 0)
	for _, v := range mappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateDiscoveryReportForDownloadDetailsReportFormatEnumStringValues Enumerates the set of values in String for GenerateDiscoveryReportForDownloadDetailsReportFormatEnum
func GetGenerateDiscoveryReportForDownloadDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnum(val string) (GenerateDiscoveryReportForDownloadDetailsReportFormatEnum, bool) {
	enum, ok := mappingGenerateDiscoveryReportForDownloadDetailsReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
