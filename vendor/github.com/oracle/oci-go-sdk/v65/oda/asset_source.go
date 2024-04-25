// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AssetSourceEnum Enum with underlying type: string
type AssetSourceEnum string

// Set of constants representing the allowable values for AssetSourceEnum
const (
	AssetSourceLocalUpload      AssetSourceEnum = "LOCAL_UPLOAD"
	AssetSourceObjectStorageUrl AssetSourceEnum = "OBJECT_STORAGE_URL"
)

var mappingAssetSourceEnum = map[string]AssetSourceEnum{
	"LOCAL_UPLOAD":       AssetSourceLocalUpload,
	"OBJECT_STORAGE_URL": AssetSourceObjectStorageUrl,
}

var mappingAssetSourceEnumLowerCase = map[string]AssetSourceEnum{
	"local_upload":       AssetSourceLocalUpload,
	"object_storage_url": AssetSourceObjectStorageUrl,
}

// GetAssetSourceEnumValues Enumerates the set of values for AssetSourceEnum
func GetAssetSourceEnumValues() []AssetSourceEnum {
	values := make([]AssetSourceEnum, 0)
	for _, v := range mappingAssetSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetSourceEnumStringValues Enumerates the set of values in String for AssetSourceEnum
func GetAssetSourceEnumStringValues() []string {
	return []string{
		"LOCAL_UPLOAD",
		"OBJECT_STORAGE_URL",
	}
}

// GetMappingAssetSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetSourceEnum(val string) (AssetSourceEnum, bool) {
	enum, ok := mappingAssetSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
