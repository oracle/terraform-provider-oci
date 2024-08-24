// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"strings"
)

// PlatformTypeEnum Enum with underlying type: string
type PlatformTypeEnum string

// Set of constants representing the allowable values for PlatformTypeEnum
const (
	PlatformTypeIaas PlatformTypeEnum = "IAAS"
	PlatformTypeSaas PlatformTypeEnum = "SAAS"
	PlatformTypePaas PlatformTypeEnum = "PAAS"
)

var mappingPlatformTypeEnum = map[string]PlatformTypeEnum{
	"IAAS": PlatformTypeIaas,
	"SAAS": PlatformTypeSaas,
	"PAAS": PlatformTypePaas,
}

var mappingPlatformTypeEnumLowerCase = map[string]PlatformTypeEnum{
	"iaas": PlatformTypeIaas,
	"saas": PlatformTypeSaas,
	"paas": PlatformTypePaas,
}

// GetPlatformTypeEnumValues Enumerates the set of values for PlatformTypeEnum
func GetPlatformTypeEnumValues() []PlatformTypeEnum {
	values := make([]PlatformTypeEnum, 0)
	for _, v := range mappingPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformTypeEnumStringValues Enumerates the set of values in String for PlatformTypeEnum
func GetPlatformTypeEnumStringValues() []string {
	return []string{
		"IAAS",
		"SAAS",
		"PAAS",
	}
}

// GetMappingPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformTypeEnum(val string) (PlatformTypeEnum, bool) {
	enum, ok := mappingPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
