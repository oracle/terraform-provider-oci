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

// DeleteModeEnum Enum with underlying type: string
type DeleteModeEnum string

// Set of constants representing the allowable values for DeleteModeEnum
const (
	DeleteModeDeleteChildren    DeleteModeEnum = "DELETE_CHILDREN"
	DeleteModeDeleteDerivations DeleteModeEnum = "DELETE_DERIVATIONS"
)

var mappingDeleteModeEnum = map[string]DeleteModeEnum{
	"DELETE_CHILDREN":    DeleteModeDeleteChildren,
	"DELETE_DERIVATIONS": DeleteModeDeleteDerivations,
}

var mappingDeleteModeEnumLowerCase = map[string]DeleteModeEnum{
	"delete_children":    DeleteModeDeleteChildren,
	"delete_derivations": DeleteModeDeleteDerivations,
}

// GetDeleteModeEnumValues Enumerates the set of values for DeleteModeEnum
func GetDeleteModeEnumValues() []DeleteModeEnum {
	values := make([]DeleteModeEnum, 0)
	for _, v := range mappingDeleteModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteModeEnumStringValues Enumerates the set of values in String for DeleteModeEnum
func GetDeleteModeEnumStringValues() []string {
	return []string{
		"DELETE_CHILDREN",
		"DELETE_DERIVATIONS",
	}
}

// GetMappingDeleteModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteModeEnum(val string) (DeleteModeEnum, bool) {
	enum, ok := mappingDeleteModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
