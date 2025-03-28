// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// MirrorTypeEnum Enum with underlying type: string
type MirrorTypeEnum string

// Set of constants representing the allowable values for MirrorTypeEnum
const (
	MirrorTypeCustom     MirrorTypeEnum = "CUSTOM"
	MirrorTypeVendor     MirrorTypeEnum = "VENDOR"
	MirrorTypeVersioned  MirrorTypeEnum = "VERSIONED"
	MirrorTypePrivate    MirrorTypeEnum = "PRIVATE"
	MirrorTypeThirdParty MirrorTypeEnum = "THIRD_PARTY"
)

var mappingMirrorTypeEnum = map[string]MirrorTypeEnum{
	"CUSTOM":      MirrorTypeCustom,
	"VENDOR":      MirrorTypeVendor,
	"VERSIONED":   MirrorTypeVersioned,
	"PRIVATE":     MirrorTypePrivate,
	"THIRD_PARTY": MirrorTypeThirdParty,
}

var mappingMirrorTypeEnumLowerCase = map[string]MirrorTypeEnum{
	"custom":      MirrorTypeCustom,
	"vendor":      MirrorTypeVendor,
	"versioned":   MirrorTypeVersioned,
	"private":     MirrorTypePrivate,
	"third_party": MirrorTypeThirdParty,
}

// GetMirrorTypeEnumValues Enumerates the set of values for MirrorTypeEnum
func GetMirrorTypeEnumValues() []MirrorTypeEnum {
	values := make([]MirrorTypeEnum, 0)
	for _, v := range mappingMirrorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMirrorTypeEnumStringValues Enumerates the set of values in String for MirrorTypeEnum
func GetMirrorTypeEnumStringValues() []string {
	return []string{
		"CUSTOM",
		"VENDOR",
		"VERSIONED",
		"PRIVATE",
		"THIRD_PARTY",
	}
}

// GetMappingMirrorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMirrorTypeEnum(val string) (MirrorTypeEnum, bool) {
	enum, ok := mappingMirrorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
