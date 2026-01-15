// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"strings"
)

// FleetErrorSortByEnum Enum with underlying type: string
type FleetErrorSortByEnum string

// Set of constants representing the allowable values for FleetErrorSortByEnum
const (
	FleetErrorSortByTimeFirstSeen FleetErrorSortByEnum = "TIME_FIRST_SEEN"
	FleetErrorSortByTimeLastSeen  FleetErrorSortByEnum = "TIME_LAST_SEEN"
)

var mappingFleetErrorSortByEnum = map[string]FleetErrorSortByEnum{
	"TIME_FIRST_SEEN": FleetErrorSortByTimeFirstSeen,
	"TIME_LAST_SEEN":  FleetErrorSortByTimeLastSeen,
}

var mappingFleetErrorSortByEnumLowerCase = map[string]FleetErrorSortByEnum{
	"time_first_seen": FleetErrorSortByTimeFirstSeen,
	"time_last_seen":  FleetErrorSortByTimeLastSeen,
}

// GetFleetErrorSortByEnumValues Enumerates the set of values for FleetErrorSortByEnum
func GetFleetErrorSortByEnumValues() []FleetErrorSortByEnum {
	values := make([]FleetErrorSortByEnum, 0)
	for _, v := range mappingFleetErrorSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetErrorSortByEnumStringValues Enumerates the set of values in String for FleetErrorSortByEnum
func GetFleetErrorSortByEnumStringValues() []string {
	return []string{
		"TIME_FIRST_SEEN",
		"TIME_LAST_SEEN",
	}
}

// GetMappingFleetErrorSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetErrorSortByEnum(val string) (FleetErrorSortByEnum, bool) {
	enum, ok := mappingFleetErrorSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
