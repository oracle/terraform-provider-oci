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

// FleetPropertySortByEnum Enum with underlying type: string
type FleetPropertySortByEnum string

// Set of constants representing the allowable values for FleetPropertySortByEnum
const (
	FleetPropertySortByTimeCreated FleetPropertySortByEnum = "timeCreated"
	FleetPropertySortByDisplayName FleetPropertySortByEnum = "displayName"
)

var mappingFleetPropertySortByEnum = map[string]FleetPropertySortByEnum{
	"timeCreated": FleetPropertySortByTimeCreated,
	"displayName": FleetPropertySortByDisplayName,
}

var mappingFleetPropertySortByEnumLowerCase = map[string]FleetPropertySortByEnum{
	"timecreated": FleetPropertySortByTimeCreated,
	"displayname": FleetPropertySortByDisplayName,
}

// GetFleetPropertySortByEnumValues Enumerates the set of values for FleetPropertySortByEnum
func GetFleetPropertySortByEnumValues() []FleetPropertySortByEnum {
	values := make([]FleetPropertySortByEnum, 0)
	for _, v := range mappingFleetPropertySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetPropertySortByEnumStringValues Enumerates the set of values in String for FleetPropertySortByEnum
func GetFleetPropertySortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingFleetPropertySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetPropertySortByEnum(val string) (FleetPropertySortByEnum, bool) {
	enum, ok := mappingFleetPropertySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
