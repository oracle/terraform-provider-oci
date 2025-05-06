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

// FleetResourceSortByEnum Enum with underlying type: string
type FleetResourceSortByEnum string

// Set of constants representing the allowable values for FleetResourceSortByEnum
const (
	FleetResourceSortByTimeCreated FleetResourceSortByEnum = "timeCreated"
	FleetResourceSortByDisplayName FleetResourceSortByEnum = "displayName"
)

var mappingFleetResourceSortByEnum = map[string]FleetResourceSortByEnum{
	"timeCreated": FleetResourceSortByTimeCreated,
	"displayName": FleetResourceSortByDisplayName,
}

var mappingFleetResourceSortByEnumLowerCase = map[string]FleetResourceSortByEnum{
	"timecreated": FleetResourceSortByTimeCreated,
	"displayname": FleetResourceSortByDisplayName,
}

// GetFleetResourceSortByEnumValues Enumerates the set of values for FleetResourceSortByEnum
func GetFleetResourceSortByEnumValues() []FleetResourceSortByEnum {
	values := make([]FleetResourceSortByEnum, 0)
	for _, v := range mappingFleetResourceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetResourceSortByEnumStringValues Enumerates the set of values in String for FleetResourceSortByEnum
func GetFleetResourceSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingFleetResourceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetResourceSortByEnum(val string) (FleetResourceSortByEnum, bool) {
	enum, ok := mappingFleetResourceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
