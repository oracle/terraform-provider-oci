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

// JavaDownloadReportSortByEnum Enum with underlying type: string
type JavaDownloadReportSortByEnum string

// Set of constants representing the allowable values for JavaDownloadReportSortByEnum
const (
	JavaDownloadReportSortByTimeCreated JavaDownloadReportSortByEnum = "timeCreated"
	JavaDownloadReportSortByDisplayName JavaDownloadReportSortByEnum = "displayName"
)

var mappingJavaDownloadReportSortByEnum = map[string]JavaDownloadReportSortByEnum{
	"timeCreated": JavaDownloadReportSortByTimeCreated,
	"displayName": JavaDownloadReportSortByDisplayName,
}

var mappingJavaDownloadReportSortByEnumLowerCase = map[string]JavaDownloadReportSortByEnum{
	"timecreated": JavaDownloadReportSortByTimeCreated,
	"displayname": JavaDownloadReportSortByDisplayName,
}

// GetJavaDownloadReportSortByEnumValues Enumerates the set of values for JavaDownloadReportSortByEnum
func GetJavaDownloadReportSortByEnumValues() []JavaDownloadReportSortByEnum {
	values := make([]JavaDownloadReportSortByEnum, 0)
	for _, v := range mappingJavaDownloadReportSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaDownloadReportSortByEnumStringValues Enumerates the set of values in String for JavaDownloadReportSortByEnum
func GetJavaDownloadReportSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingJavaDownloadReportSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaDownloadReportSortByEnum(val string) (JavaDownloadReportSortByEnum, bool) {
	enum, ok := mappingJavaDownloadReportSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
