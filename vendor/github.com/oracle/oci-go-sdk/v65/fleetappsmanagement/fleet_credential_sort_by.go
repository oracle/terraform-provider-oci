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

// FleetCredentialSortByEnum Enum with underlying type: string
type FleetCredentialSortByEnum string

// Set of constants representing the allowable values for FleetCredentialSortByEnum
const (
	FleetCredentialSortByTimeCreated FleetCredentialSortByEnum = "timeCreated"
	FleetCredentialSortByDisplayName FleetCredentialSortByEnum = "displayName"
)

var mappingFleetCredentialSortByEnum = map[string]FleetCredentialSortByEnum{
	"timeCreated": FleetCredentialSortByTimeCreated,
	"displayName": FleetCredentialSortByDisplayName,
}

var mappingFleetCredentialSortByEnumLowerCase = map[string]FleetCredentialSortByEnum{
	"timecreated": FleetCredentialSortByTimeCreated,
	"displayname": FleetCredentialSortByDisplayName,
}

// GetFleetCredentialSortByEnumValues Enumerates the set of values for FleetCredentialSortByEnum
func GetFleetCredentialSortByEnumValues() []FleetCredentialSortByEnum {
	values := make([]FleetCredentialSortByEnum, 0)
	for _, v := range mappingFleetCredentialSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetCredentialSortByEnumStringValues Enumerates the set of values in String for FleetCredentialSortByEnum
func GetFleetCredentialSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingFleetCredentialSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetCredentialSortByEnum(val string) (FleetCredentialSortByEnum, bool) {
	enum, ok := mappingFleetCredentialSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
