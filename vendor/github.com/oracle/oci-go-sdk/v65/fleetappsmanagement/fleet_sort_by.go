// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// FleetSortByEnum Enum with underlying type: string
type FleetSortByEnum string

// Set of constants representing the allowable values for FleetSortByEnum
const (
	FleetSortByTimeCreated FleetSortByEnum = "timeCreated"
	FleetSortByDisplayName FleetSortByEnum = "displayName"
)

var mappingFleetSortByEnum = map[string]FleetSortByEnum{
	"timeCreated": FleetSortByTimeCreated,
	"displayName": FleetSortByDisplayName,
}

var mappingFleetSortByEnumLowerCase = map[string]FleetSortByEnum{
	"timecreated": FleetSortByTimeCreated,
	"displayname": FleetSortByDisplayName,
}

// GetFleetSortByEnumValues Enumerates the set of values for FleetSortByEnum
func GetFleetSortByEnumValues() []FleetSortByEnum {
	values := make([]FleetSortByEnum, 0)
	for _, v := range mappingFleetSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetSortByEnumStringValues Enumerates the set of values in String for FleetSortByEnum
func GetFleetSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingFleetSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetSortByEnum(val string) (FleetSortByEnum, bool) {
	enum, ok := mappingFleetSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
