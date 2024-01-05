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

// JavaDownloadCountAggregationTypeEnum Enum with underlying type: string
type JavaDownloadCountAggregationTypeEnum string

// Set of constants representing the allowable values for JavaDownloadCountAggregationTypeEnum
const (
	JavaDownloadCountAggregationTypeJavaFamily  JavaDownloadCountAggregationTypeEnum = "JAVA_FAMILY"
	JavaDownloadCountAggregationTypeJavaRelease JavaDownloadCountAggregationTypeEnum = "JAVA_RELEASE"
	JavaDownloadCountAggregationTypePlatform    JavaDownloadCountAggregationTypeEnum = "PLATFORM"
)

var mappingJavaDownloadCountAggregationTypeEnum = map[string]JavaDownloadCountAggregationTypeEnum{
	"JAVA_FAMILY":  JavaDownloadCountAggregationTypeJavaFamily,
	"JAVA_RELEASE": JavaDownloadCountAggregationTypeJavaRelease,
	"PLATFORM":     JavaDownloadCountAggregationTypePlatform,
}

var mappingJavaDownloadCountAggregationTypeEnumLowerCase = map[string]JavaDownloadCountAggregationTypeEnum{
	"java_family":  JavaDownloadCountAggregationTypeJavaFamily,
	"java_release": JavaDownloadCountAggregationTypeJavaRelease,
	"platform":     JavaDownloadCountAggregationTypePlatform,
}

// GetJavaDownloadCountAggregationTypeEnumValues Enumerates the set of values for JavaDownloadCountAggregationTypeEnum
func GetJavaDownloadCountAggregationTypeEnumValues() []JavaDownloadCountAggregationTypeEnum {
	values := make([]JavaDownloadCountAggregationTypeEnum, 0)
	for _, v := range mappingJavaDownloadCountAggregationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaDownloadCountAggregationTypeEnumStringValues Enumerates the set of values in String for JavaDownloadCountAggregationTypeEnum
func GetJavaDownloadCountAggregationTypeEnumStringValues() []string {
	return []string{
		"JAVA_FAMILY",
		"JAVA_RELEASE",
		"PLATFORM",
	}
}

// GetMappingJavaDownloadCountAggregationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaDownloadCountAggregationTypeEnum(val string) (JavaDownloadCountAggregationTypeEnum, bool) {
	enum, ok := mappingJavaDownloadCountAggregationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
