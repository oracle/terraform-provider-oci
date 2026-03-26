// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// SoftwareSourceSubTypeEnum Enum with underlying type: string
type SoftwareSourceSubTypeEnum string

// Set of constants representing the allowable values for SoftwareSourceSubTypeEnum
const (
	SoftwareSourceSubTypeFilter   SoftwareSourceSubTypeEnum = "FILTER"
	SoftwareSourceSubTypeManifest SoftwareSourceSubTypeEnum = "MANIFEST"
	SoftwareSourceSubTypeSnapshot SoftwareSourceSubTypeEnum = "SNAPSHOT"
)

var mappingSoftwareSourceSubTypeEnum = map[string]SoftwareSourceSubTypeEnum{
	"FILTER":   SoftwareSourceSubTypeFilter,
	"MANIFEST": SoftwareSourceSubTypeManifest,
	"SNAPSHOT": SoftwareSourceSubTypeSnapshot,
}

var mappingSoftwareSourceSubTypeEnumLowerCase = map[string]SoftwareSourceSubTypeEnum{
	"filter":   SoftwareSourceSubTypeFilter,
	"manifest": SoftwareSourceSubTypeManifest,
	"snapshot": SoftwareSourceSubTypeSnapshot,
}

// GetSoftwareSourceSubTypeEnumValues Enumerates the set of values for SoftwareSourceSubTypeEnum
func GetSoftwareSourceSubTypeEnumValues() []SoftwareSourceSubTypeEnum {
	values := make([]SoftwareSourceSubTypeEnum, 0)
	for _, v := range mappingSoftwareSourceSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSoftwareSourceSubTypeEnumStringValues Enumerates the set of values in String for SoftwareSourceSubTypeEnum
func GetSoftwareSourceSubTypeEnumStringValues() []string {
	return []string{
		"FILTER",
		"MANIFEST",
		"SNAPSHOT",
	}
}

// GetMappingSoftwareSourceSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSoftwareSourceSubTypeEnum(val string) (SoftwareSourceSubTypeEnum, bool) {
	enum, ok := mappingSoftwareSourceSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
