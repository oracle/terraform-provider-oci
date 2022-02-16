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

// DownloadDiscoveryReportDetails Details to download a discovery report.
type DownloadDiscoveryReportDetails struct {

	// The OCID of the discovery job.
	DiscoveryJobId *string `mandatory:"false" json:"discoveryJobId"`

	// Format of the report.
	ReportFormat DownloadDiscoveryReportDetailsReportFormatEnum `mandatory:"false" json:"reportFormat,omitempty"`
}

func (m DownloadDiscoveryReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DownloadDiscoveryReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDownloadDiscoveryReportDetailsReportFormatEnum(string(m.ReportFormat)); !ok && m.ReportFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReportFormat: %s. Supported values are: %s.", m.ReportFormat, strings.Join(GetDownloadDiscoveryReportDetailsReportFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DownloadDiscoveryReportDetailsReportFormatEnum Enum with underlying type: string
type DownloadDiscoveryReportDetailsReportFormatEnum string

// Set of constants representing the allowable values for DownloadDiscoveryReportDetailsReportFormatEnum
const (
	DownloadDiscoveryReportDetailsReportFormatPdf DownloadDiscoveryReportDetailsReportFormatEnum = "PDF"
	DownloadDiscoveryReportDetailsReportFormatXls DownloadDiscoveryReportDetailsReportFormatEnum = "XLS"
)

var mappingDownloadDiscoveryReportDetailsReportFormatEnum = map[string]DownloadDiscoveryReportDetailsReportFormatEnum{
	"PDF": DownloadDiscoveryReportDetailsReportFormatPdf,
	"XLS": DownloadDiscoveryReportDetailsReportFormatXls,
}

// GetDownloadDiscoveryReportDetailsReportFormatEnumValues Enumerates the set of values for DownloadDiscoveryReportDetailsReportFormatEnum
func GetDownloadDiscoveryReportDetailsReportFormatEnumValues() []DownloadDiscoveryReportDetailsReportFormatEnum {
	values := make([]DownloadDiscoveryReportDetailsReportFormatEnum, 0)
	for _, v := range mappingDownloadDiscoveryReportDetailsReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadDiscoveryReportDetailsReportFormatEnumStringValues Enumerates the set of values in String for DownloadDiscoveryReportDetailsReportFormatEnum
func GetDownloadDiscoveryReportDetailsReportFormatEnumStringValues() []string {
	return []string{
		"PDF",
		"XLS",
	}
}

// GetMappingDownloadDiscoveryReportDetailsReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadDiscoveryReportDetailsReportFormatEnum(val string) (DownloadDiscoveryReportDetailsReportFormatEnum, bool) {
	mappingDownloadDiscoveryReportDetailsReportFormatEnumIgnoreCase := make(map[string]DownloadDiscoveryReportDetailsReportFormatEnum)
	for k, v := range mappingDownloadDiscoveryReportDetailsReportFormatEnum {
		mappingDownloadDiscoveryReportDetailsReportFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDownloadDiscoveryReportDetailsReportFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
