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

// MediaWorkflowTaskDeclarationSortByEnum Enum with underlying type: string
type MediaWorkflowTaskDeclarationSortByEnum string

// Set of constants representing the allowable values for MediaWorkflowTaskDeclarationSortByEnum
const (
	MediaWorkflowTaskDeclarationSortByName    MediaWorkflowTaskDeclarationSortByEnum = "name"
	MediaWorkflowTaskDeclarationSortByVersion MediaWorkflowTaskDeclarationSortByEnum = "version"
)

var mappingMediaWorkflowTaskDeclarationSortByEnum = map[string]MediaWorkflowTaskDeclarationSortByEnum{
	"name":    MediaWorkflowTaskDeclarationSortByName,
	"version": MediaWorkflowTaskDeclarationSortByVersion,
}

var mappingMediaWorkflowTaskDeclarationSortByEnumLowerCase = map[string]MediaWorkflowTaskDeclarationSortByEnum{
	"name":    MediaWorkflowTaskDeclarationSortByName,
	"version": MediaWorkflowTaskDeclarationSortByVersion,
}

// GetMediaWorkflowTaskDeclarationSortByEnumValues Enumerates the set of values for MediaWorkflowTaskDeclarationSortByEnum
func GetMediaWorkflowTaskDeclarationSortByEnumValues() []MediaWorkflowTaskDeclarationSortByEnum {
	values := make([]MediaWorkflowTaskDeclarationSortByEnum, 0)
	for _, v := range mappingMediaWorkflowTaskDeclarationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMediaWorkflowTaskDeclarationSortByEnumStringValues Enumerates the set of values in String for MediaWorkflowTaskDeclarationSortByEnum
func GetMediaWorkflowTaskDeclarationSortByEnumStringValues() []string {
	return []string{
		"name",
		"version",
	}
}

// GetMappingMediaWorkflowTaskDeclarationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMediaWorkflowTaskDeclarationSortByEnum(val string) (MediaWorkflowTaskDeclarationSortByEnum, bool) {
	enum, ok := mappingMediaWorkflowTaskDeclarationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
