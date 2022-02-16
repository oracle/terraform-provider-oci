// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"strings"
)

// ManagementAgentGroupByEnum Enum with underlying type: string
type ManagementAgentGroupByEnum string

// Set of constants representing the allowable values for ManagementAgentGroupByEnum
const (
	ManagementAgentGroupByAvailabilityStatus ManagementAgentGroupByEnum = "availabilityStatus"
	ManagementAgentGroupByPlatformType       ManagementAgentGroupByEnum = "platformType"
	ManagementAgentGroupByVersion            ManagementAgentGroupByEnum = "version"
)

var mappingManagementAgentGroupByEnum = map[string]ManagementAgentGroupByEnum{
	"availabilityStatus": ManagementAgentGroupByAvailabilityStatus,
	"platformType":       ManagementAgentGroupByPlatformType,
	"version":            ManagementAgentGroupByVersion,
}

// GetManagementAgentGroupByEnumValues Enumerates the set of values for ManagementAgentGroupByEnum
func GetManagementAgentGroupByEnumValues() []ManagementAgentGroupByEnum {
	values := make([]ManagementAgentGroupByEnum, 0)
	for _, v := range mappingManagementAgentGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetManagementAgentGroupByEnumStringValues Enumerates the set of values in String for ManagementAgentGroupByEnum
func GetManagementAgentGroupByEnumStringValues() []string {
	return []string{
		"availabilityStatus",
		"platformType",
		"version",
	}
}

// GetMappingManagementAgentGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagementAgentGroupByEnum(val string) (ManagementAgentGroupByEnum, bool) {
	mappingManagementAgentGroupByEnumIgnoreCase := make(map[string]ManagementAgentGroupByEnum)
	for k, v := range mappingManagementAgentGroupByEnum {
		mappingManagementAgentGroupByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingManagementAgentGroupByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
