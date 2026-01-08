// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"strings"
)

// AssetFileContentTypeEnum Enum with underlying type: string
type AssetFileContentTypeEnum string

// Set of constants representing the allowable values for AssetFileContentTypeEnum
const (
	AssetFileContentTypeText  AssetFileContentTypeEnum = "TEXT"
	AssetFileContentTypeImage AssetFileContentTypeEnum = "IMAGE"
	AssetFileContentTypeAudio AssetFileContentTypeEnum = "AUDIO"
	AssetFileContentTypeVideo AssetFileContentTypeEnum = "VIDEO"
)

var mappingAssetFileContentTypeEnum = map[string]AssetFileContentTypeEnum{
	"TEXT":  AssetFileContentTypeText,
	"IMAGE": AssetFileContentTypeImage,
	"AUDIO": AssetFileContentTypeAudio,
	"VIDEO": AssetFileContentTypeVideo,
}

var mappingAssetFileContentTypeEnumLowerCase = map[string]AssetFileContentTypeEnum{
	"text":  AssetFileContentTypeText,
	"image": AssetFileContentTypeImage,
	"audio": AssetFileContentTypeAudio,
	"video": AssetFileContentTypeVideo,
}

// GetAssetFileContentTypeEnumValues Enumerates the set of values for AssetFileContentTypeEnum
func GetAssetFileContentTypeEnumValues() []AssetFileContentTypeEnum {
	values := make([]AssetFileContentTypeEnum, 0)
	for _, v := range mappingAssetFileContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetFileContentTypeEnumStringValues Enumerates the set of values in String for AssetFileContentTypeEnum
func GetAssetFileContentTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"IMAGE",
		"AUDIO",
		"VIDEO",
	}
}

// GetMappingAssetFileContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetFileContentTypeEnum(val string) (AssetFileContentTypeEnum, bool) {
	enum, ok := mappingAssetFileContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
