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

// PluginErrorSortByEnum Enum with underlying type: string
type PluginErrorSortByEnum string

// Set of constants representing the allowable values for PluginErrorSortByEnum
const (
	PluginErrorSortByTimeFirstSeen PluginErrorSortByEnum = "TIME_FIRST_SEEN"
	PluginErrorSortByTimeLastSeen  PluginErrorSortByEnum = "TIME_LAST_SEEN"
)

var mappingPluginErrorSortByEnum = map[string]PluginErrorSortByEnum{
	"TIME_FIRST_SEEN": PluginErrorSortByTimeFirstSeen,
	"TIME_LAST_SEEN":  PluginErrorSortByTimeLastSeen,
}

var mappingPluginErrorSortByEnumLowerCase = map[string]PluginErrorSortByEnum{
	"time_first_seen": PluginErrorSortByTimeFirstSeen,
	"time_last_seen":  PluginErrorSortByTimeLastSeen,
}

// GetPluginErrorSortByEnumValues Enumerates the set of values for PluginErrorSortByEnum
func GetPluginErrorSortByEnumValues() []PluginErrorSortByEnum {
	values := make([]PluginErrorSortByEnum, 0)
	for _, v := range mappingPluginErrorSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetPluginErrorSortByEnumStringValues Enumerates the set of values in String for PluginErrorSortByEnum
func GetPluginErrorSortByEnumStringValues() []string {
	return []string{
		"TIME_FIRST_SEEN",
		"TIME_LAST_SEEN",
	}
}

// GetMappingPluginErrorSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluginErrorSortByEnum(val string) (PluginErrorSortByEnum, bool) {
	enum, ok := mappingPluginErrorSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
