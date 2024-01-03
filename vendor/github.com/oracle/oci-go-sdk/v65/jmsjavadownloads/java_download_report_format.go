// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"strings"
)

// JavaDownloadReportFormatEnum Enum with underlying type: string
type JavaDownloadReportFormatEnum string

// Set of constants representing the allowable values for JavaDownloadReportFormatEnum
const (
	JavaDownloadReportFormatCsv JavaDownloadReportFormatEnum = "CSV"
)

var mappingJavaDownloadReportFormatEnum = map[string]JavaDownloadReportFormatEnum{
	"CSV": JavaDownloadReportFormatCsv,
}

var mappingJavaDownloadReportFormatEnumLowerCase = map[string]JavaDownloadReportFormatEnum{
	"csv": JavaDownloadReportFormatCsv,
}

// GetJavaDownloadReportFormatEnumValues Enumerates the set of values for JavaDownloadReportFormatEnum
func GetJavaDownloadReportFormatEnumValues() []JavaDownloadReportFormatEnum {
	values := make([]JavaDownloadReportFormatEnum, 0)
	for _, v := range mappingJavaDownloadReportFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaDownloadReportFormatEnumStringValues Enumerates the set of values in String for JavaDownloadReportFormatEnum
func GetJavaDownloadReportFormatEnumStringValues() []string {
	return []string{
		"CSV",
	}
}

// GetMappingJavaDownloadReportFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaDownloadReportFormatEnum(val string) (JavaDownloadReportFormatEnum, bool) {
	enum, ok := mappingJavaDownloadReportFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
