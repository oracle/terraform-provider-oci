// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"strings"
)

// ImageDetailsTypesEnum Enum with underlying type: string
type ImageDetailsTypesEnum string

// Set of constants representing the allowable values for ImageDetailsTypesEnum
const (
	ImageDetailsTypesObjectStorageBucket ImageDetailsTypesEnum = "OBJECT_STORAGE_BUCKET"
)

var mappingImageDetailsTypesEnum = map[string]ImageDetailsTypesEnum{
	"OBJECT_STORAGE_BUCKET": ImageDetailsTypesObjectStorageBucket,
}

var mappingImageDetailsTypesEnumLowerCase = map[string]ImageDetailsTypesEnum{
	"object_storage_bucket": ImageDetailsTypesObjectStorageBucket,
}

// GetImageDetailsTypesEnumValues Enumerates the set of values for ImageDetailsTypesEnum
func GetImageDetailsTypesEnumValues() []ImageDetailsTypesEnum {
	values := make([]ImageDetailsTypesEnum, 0)
	for _, v := range mappingImageDetailsTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetImageDetailsTypesEnumStringValues Enumerates the set of values in String for ImageDetailsTypesEnum
func GetImageDetailsTypesEnumStringValues() []string {
	return []string{
		"OBJECT_STORAGE_BUCKET",
	}
}

// GetMappingImageDetailsTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageDetailsTypesEnum(val string) (ImageDetailsTypesEnum, bool) {
	enum, ok := mappingImageDetailsTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
