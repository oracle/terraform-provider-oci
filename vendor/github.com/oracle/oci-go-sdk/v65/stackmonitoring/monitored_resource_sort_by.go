// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// MonitoredResourceSortByEnum Enum with underlying type: string
type MonitoredResourceSortByEnum string

// Set of constants representing the allowable values for MonitoredResourceSortByEnum
const (
	MonitoredResourceSortByName        MonitoredResourceSortByEnum = "NAME"
	MonitoredResourceSortByTimeCreated MonitoredResourceSortByEnum = "TIME_CREATED"
)

var mappingMonitoredResourceSortByEnum = map[string]MonitoredResourceSortByEnum{
	"NAME":         MonitoredResourceSortByName,
	"TIME_CREATED": MonitoredResourceSortByTimeCreated,
}

var mappingMonitoredResourceSortByEnumLowerCase = map[string]MonitoredResourceSortByEnum{
	"name":         MonitoredResourceSortByName,
	"time_created": MonitoredResourceSortByTimeCreated,
}

// GetMonitoredResourceSortByEnumValues Enumerates the set of values for MonitoredResourceSortByEnum
func GetMonitoredResourceSortByEnumValues() []MonitoredResourceSortByEnum {
	values := make([]MonitoredResourceSortByEnum, 0)
	for _, v := range mappingMonitoredResourceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredResourceSortByEnumStringValues Enumerates the set of values in String for MonitoredResourceSortByEnum
func GetMonitoredResourceSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"TIME_CREATED",
	}
}

// GetMappingMonitoredResourceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredResourceSortByEnum(val string) (MonitoredResourceSortByEnum, bool) {
	enum, ok := mappingMonitoredResourceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
