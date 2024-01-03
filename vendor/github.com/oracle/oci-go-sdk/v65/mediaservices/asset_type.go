// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Media Services API
//
// Media Services (includes Media Flow and Media Streams) is a fully managed service for processing media (video) source content. Use Media Flow and Media Streams to transcode and package digital video using configurable workflows and stream video outputs.
// Use the Media Services API to configure media workflows and run Media Flow jobs, create distribution channels, ingest assets, create Preview URLs and play assets. For more information, see Media Flow (https://docs.cloud.oracle.com/iaas/Content/dms-mediaflow/home.htm) and Media Streams (https://docs.cloud.oracle.com/iaas/Content/dms-mediastream/home.htm).
//

package mediaservices

import (
	"strings"
)

// AssetTypeEnum Enum with underlying type: string
type AssetTypeEnum string

// Set of constants representing the allowable values for AssetTypeEnum
const (
	AssetTypeAudio       AssetTypeEnum = "AUDIO"
	AssetTypeVideo       AssetTypeEnum = "VIDEO"
	AssetTypePlaylist    AssetTypeEnum = "PLAYLIST"
	AssetTypeImage       AssetTypeEnum = "IMAGE"
	AssetTypeCaptionFile AssetTypeEnum = "CAPTION_FILE"
	AssetTypeUnknown     AssetTypeEnum = "UNKNOWN"
)

var mappingAssetTypeEnum = map[string]AssetTypeEnum{
	"AUDIO":        AssetTypeAudio,
	"VIDEO":        AssetTypeVideo,
	"PLAYLIST":     AssetTypePlaylist,
	"IMAGE":        AssetTypeImage,
	"CAPTION_FILE": AssetTypeCaptionFile,
	"UNKNOWN":      AssetTypeUnknown,
}

var mappingAssetTypeEnumLowerCase = map[string]AssetTypeEnum{
	"audio":        AssetTypeAudio,
	"video":        AssetTypeVideo,
	"playlist":     AssetTypePlaylist,
	"image":        AssetTypeImage,
	"caption_file": AssetTypeCaptionFile,
	"unknown":      AssetTypeUnknown,
}

// GetAssetTypeEnumValues Enumerates the set of values for AssetTypeEnum
func GetAssetTypeEnumValues() []AssetTypeEnum {
	values := make([]AssetTypeEnum, 0)
	for _, v := range mappingAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssetTypeEnumStringValues Enumerates the set of values in String for AssetTypeEnum
func GetAssetTypeEnumStringValues() []string {
	return []string{
		"AUDIO",
		"VIDEO",
		"PLAYLIST",
		"IMAGE",
		"CAPTION_FILE",
		"UNKNOWN",
	}
}

// GetMappingAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssetTypeEnum(val string) (AssetTypeEnum, bool) {
	enum, ok := mappingAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
