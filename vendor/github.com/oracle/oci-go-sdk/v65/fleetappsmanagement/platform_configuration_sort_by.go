// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// PlatformConfigurationSortByEnum Enum with underlying type: string
type PlatformConfigurationSortByEnum string

// Set of constants representing the allowable values for PlatformConfigurationSortByEnum
const (
	PlatformConfigurationSortByTimeCreated PlatformConfigurationSortByEnum = "timeCreated"
	PlatformConfigurationSortByDisplayName PlatformConfigurationSortByEnum = "displayName"
)

var mappingPlatformConfigurationSortByEnum = map[string]PlatformConfigurationSortByEnum{
	"timeCreated": PlatformConfigurationSortByTimeCreated,
	"displayName": PlatformConfigurationSortByDisplayName,
}

var mappingPlatformConfigurationSortByEnumLowerCase = map[string]PlatformConfigurationSortByEnum{
	"timecreated": PlatformConfigurationSortByTimeCreated,
	"displayname": PlatformConfigurationSortByDisplayName,
}

// GetPlatformConfigurationSortByEnumValues Enumerates the set of values for PlatformConfigurationSortByEnum
func GetPlatformConfigurationSortByEnumValues() []PlatformConfigurationSortByEnum {
	values := make([]PlatformConfigurationSortByEnum, 0)
	for _, v := range mappingPlatformConfigurationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformConfigurationSortByEnumStringValues Enumerates the set of values in String for PlatformConfigurationSortByEnum
func GetPlatformConfigurationSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingPlatformConfigurationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformConfigurationSortByEnum(val string) (PlatformConfigurationSortByEnum, bool) {
	enum, ok := mappingPlatformConfigurationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
