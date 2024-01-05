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

// JavaDownloadRecordSortByEnum Enum with underlying type: string
type JavaDownloadRecordSortByEnum string

// Set of constants representing the allowable values for JavaDownloadRecordSortByEnum
const (
	JavaDownloadRecordSortByTimeDownloaded   JavaDownloadRecordSortByEnum = "timeDownloaded"
	JavaDownloadRecordSortByDownloadSourceId JavaDownloadRecordSortByEnum = "downloadSourceId"
	JavaDownloadRecordSortByDownloadType     JavaDownloadRecordSortByEnum = "downloadType"
)

var mappingJavaDownloadRecordSortByEnum = map[string]JavaDownloadRecordSortByEnum{
	"timeDownloaded":   JavaDownloadRecordSortByTimeDownloaded,
	"downloadSourceId": JavaDownloadRecordSortByDownloadSourceId,
	"downloadType":     JavaDownloadRecordSortByDownloadType,
}

var mappingJavaDownloadRecordSortByEnumLowerCase = map[string]JavaDownloadRecordSortByEnum{
	"timedownloaded":   JavaDownloadRecordSortByTimeDownloaded,
	"downloadsourceid": JavaDownloadRecordSortByDownloadSourceId,
	"downloadtype":     JavaDownloadRecordSortByDownloadType,
}

// GetJavaDownloadRecordSortByEnumValues Enumerates the set of values for JavaDownloadRecordSortByEnum
func GetJavaDownloadRecordSortByEnumValues() []JavaDownloadRecordSortByEnum {
	values := make([]JavaDownloadRecordSortByEnum, 0)
	for _, v := range mappingJavaDownloadRecordSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaDownloadRecordSortByEnumStringValues Enumerates the set of values in String for JavaDownloadRecordSortByEnum
func GetJavaDownloadRecordSortByEnumStringValues() []string {
	return []string{
		"timeDownloaded",
		"downloadSourceId",
		"downloadType",
	}
}

// GetMappingJavaDownloadRecordSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaDownloadRecordSortByEnum(val string) (JavaDownloadRecordSortByEnum, bool) {
	enum, ok := mappingJavaDownloadRecordSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
