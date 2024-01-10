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

// DownloadUrlTypeEnum Enum with underlying type: string
type DownloadUrlTypeEnum string

// Set of constants representing the allowable values for DownloadUrlTypeEnum
const (
	DownloadUrlTypeOss DownloadUrlTypeEnum = "OSS"
	DownloadUrlTypeCdn DownloadUrlTypeEnum = "CDN"
)

var mappingDownloadUrlTypeEnum = map[string]DownloadUrlTypeEnum{
	"OSS": DownloadUrlTypeOss,
	"CDN": DownloadUrlTypeCdn,
}

var mappingDownloadUrlTypeEnumLowerCase = map[string]DownloadUrlTypeEnum{
	"oss": DownloadUrlTypeOss,
	"cdn": DownloadUrlTypeCdn,
}

// GetDownloadUrlTypeEnumValues Enumerates the set of values for DownloadUrlTypeEnum
func GetDownloadUrlTypeEnumValues() []DownloadUrlTypeEnum {
	values := make([]DownloadUrlTypeEnum, 0)
	for _, v := range mappingDownloadUrlTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDownloadUrlTypeEnumStringValues Enumerates the set of values in String for DownloadUrlTypeEnum
func GetDownloadUrlTypeEnumStringValues() []string {
	return []string{
		"OSS",
		"CDN",
	}
}

// GetMappingDownloadUrlTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDownloadUrlTypeEnum(val string) (DownloadUrlTypeEnum, bool) {
	enum, ok := mappingDownloadUrlTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
