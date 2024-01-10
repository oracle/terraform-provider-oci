// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"strings"
)

// DiscoveryScheduleLifecycleStateEnum Enum with underlying type: string
type DiscoveryScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for DiscoveryScheduleLifecycleStateEnum
const (
	DiscoveryScheduleLifecycleStateActive  DiscoveryScheduleLifecycleStateEnum = "ACTIVE"
	DiscoveryScheduleLifecycleStateDeleted DiscoveryScheduleLifecycleStateEnum = "DELETED"
)

var mappingDiscoveryScheduleLifecycleStateEnum = map[string]DiscoveryScheduleLifecycleStateEnum{
	"ACTIVE":  DiscoveryScheduleLifecycleStateActive,
	"DELETED": DiscoveryScheduleLifecycleStateDeleted,
}

var mappingDiscoveryScheduleLifecycleStateEnumLowerCase = map[string]DiscoveryScheduleLifecycleStateEnum{
	"active":  DiscoveryScheduleLifecycleStateActive,
	"deleted": DiscoveryScheduleLifecycleStateDeleted,
}

// GetDiscoveryScheduleLifecycleStateEnumValues Enumerates the set of values for DiscoveryScheduleLifecycleStateEnum
func GetDiscoveryScheduleLifecycleStateEnumValues() []DiscoveryScheduleLifecycleStateEnum {
	values := make([]DiscoveryScheduleLifecycleStateEnum, 0)
	for _, v := range mappingDiscoveryScheduleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryScheduleLifecycleStateEnumStringValues Enumerates the set of values in String for DiscoveryScheduleLifecycleStateEnum
func GetDiscoveryScheduleLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingDiscoveryScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryScheduleLifecycleStateEnum(val string) (DiscoveryScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingDiscoveryScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
