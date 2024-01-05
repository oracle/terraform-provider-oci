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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MediaAssetTag Tags of the MediaAsset.
type MediaAssetTag struct {

	// Tag of the MediaAsset.
	Value *string `mandatory:"true" json:"value"`

	// Type of the tag.
	Type MediaAssetTagTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m MediaAssetTag) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MediaAssetTag) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMediaAssetTagTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetMediaAssetTagTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MediaAssetTagTypeEnum Enum with underlying type: string
type MediaAssetTagTypeEnum string

// Set of constants representing the allowable values for MediaAssetTagTypeEnum
const (
	MediaAssetTagTypeUser   MediaAssetTagTypeEnum = "USER"
	MediaAssetTagTypeSystem MediaAssetTagTypeEnum = "SYSTEM"
)

var mappingMediaAssetTagTypeEnum = map[string]MediaAssetTagTypeEnum{
	"USER":   MediaAssetTagTypeUser,
	"SYSTEM": MediaAssetTagTypeSystem,
}

var mappingMediaAssetTagTypeEnumLowerCase = map[string]MediaAssetTagTypeEnum{
	"user":   MediaAssetTagTypeUser,
	"system": MediaAssetTagTypeSystem,
}

// GetMediaAssetTagTypeEnumValues Enumerates the set of values for MediaAssetTagTypeEnum
func GetMediaAssetTagTypeEnumValues() []MediaAssetTagTypeEnum {
	values := make([]MediaAssetTagTypeEnum, 0)
	for _, v := range mappingMediaAssetTagTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaAssetTagTypeEnumStringValues Enumerates the set of values in String for MediaAssetTagTypeEnum
func GetMediaAssetTagTypeEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingMediaAssetTagTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaAssetTagTypeEnum(val string) (MediaAssetTagTypeEnum, bool) {
	enum, ok := mappingMediaAssetTagTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
