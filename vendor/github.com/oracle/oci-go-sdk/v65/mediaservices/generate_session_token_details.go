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

// GenerateSessionTokenDetails Information about the new session token.
type GenerateSessionTokenDetails struct {

	// Array of scopes the token can act upon.
	Scopes []GenerateSessionTokenDetailsScopesEnum `mandatory:"true" json:"scopes"`

	// The packaging config resource identifier used to limit the scope of the token.
	PackagingConfigId *string `mandatory:"true" json:"packagingConfigId"`

	// Token expiry time. An RFC3339 formatted datetime string.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Array of asset resource IDs used to limit the scope of the token.
	AssetIds []string `mandatory:"false" json:"assetIds"`
}

func (m GenerateSessionTokenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSessionTokenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Scopes {
		if _, ok := GetMappingGenerateSessionTokenDetailsScopesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scopes: %s. Supported values are: %s.", val, strings.Join(GetGenerateSessionTokenDetailsScopesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateSessionTokenDetailsScopesEnum Enum with underlying type: string
type GenerateSessionTokenDetailsScopesEnum string

// Set of constants representing the allowable values for GenerateSessionTokenDetailsScopesEnum
const (
	GenerateSessionTokenDetailsScopesPlaylist GenerateSessionTokenDetailsScopesEnum = "PLAYLIST"
	GenerateSessionTokenDetailsScopesEdge     GenerateSessionTokenDetailsScopesEnum = "EDGE"
)

var mappingGenerateSessionTokenDetailsScopesEnum = map[string]GenerateSessionTokenDetailsScopesEnum{
	"PLAYLIST": GenerateSessionTokenDetailsScopesPlaylist,
	"EDGE":     GenerateSessionTokenDetailsScopesEdge,
}

var mappingGenerateSessionTokenDetailsScopesEnumLowerCase = map[string]GenerateSessionTokenDetailsScopesEnum{
	"playlist": GenerateSessionTokenDetailsScopesPlaylist,
	"edge":     GenerateSessionTokenDetailsScopesEdge,
}

// GetGenerateSessionTokenDetailsScopesEnumValues Enumerates the set of values for GenerateSessionTokenDetailsScopesEnum
func GetGenerateSessionTokenDetailsScopesEnumValues() []GenerateSessionTokenDetailsScopesEnum {
	values := make([]GenerateSessionTokenDetailsScopesEnum, 0)
	for _, v := range mappingGenerateSessionTokenDetailsScopesEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateSessionTokenDetailsScopesEnumStringValues Enumerates the set of values in String for GenerateSessionTokenDetailsScopesEnum
func GetGenerateSessionTokenDetailsScopesEnumStringValues() []string {
	return []string{
		"PLAYLIST",
		"EDGE",
	}
}

// GetMappingGenerateSessionTokenDetailsScopesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateSessionTokenDetailsScopesEnum(val string) (GenerateSessionTokenDetailsScopesEnum, bool) {
	enum, ok := mappingGenerateSessionTokenDetailsScopesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
