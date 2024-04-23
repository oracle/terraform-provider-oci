// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// EventTypeEnum Enum with underlying type: string
type EventTypeEnum string

// Set of constants representing the allowable values for EventTypeEnum
const (
	EventTypeKernelOops        EventTypeEnum = "KERNEL_OOPS"
	EventTypeKernelCrash       EventTypeEnum = "KERNEL_CRASH"
	EventTypeExploitAttempt    EventTypeEnum = "EXPLOIT_ATTEMPT"
	EventTypeSoftwareUpdate    EventTypeEnum = "SOFTWARE_UPDATE"
	EventTypeKspliceUpdate     EventTypeEnum = "KSPLICE_UPDATE"
	EventTypeSoftwareSource    EventTypeEnum = "SOFTWARE_SOURCE"
	EventTypeAgent             EventTypeEnum = "AGENT"
	EventTypeManagementStation EventTypeEnum = "MANAGEMENT_STATION"
)

var mappingEventTypeEnum = map[string]EventTypeEnum{
	"KERNEL_OOPS":        EventTypeKernelOops,
	"KERNEL_CRASH":       EventTypeKernelCrash,
	"EXPLOIT_ATTEMPT":    EventTypeExploitAttempt,
	"SOFTWARE_UPDATE":    EventTypeSoftwareUpdate,
	"KSPLICE_UPDATE":     EventTypeKspliceUpdate,
	"SOFTWARE_SOURCE":    EventTypeSoftwareSource,
	"AGENT":              EventTypeAgent,
	"MANAGEMENT_STATION": EventTypeManagementStation,
}

var mappingEventTypeEnumLowerCase = map[string]EventTypeEnum{
	"kernel_oops":        EventTypeKernelOops,
	"kernel_crash":       EventTypeKernelCrash,
	"exploit_attempt":    EventTypeExploitAttempt,
	"software_update":    EventTypeSoftwareUpdate,
	"ksplice_update":     EventTypeKspliceUpdate,
	"software_source":    EventTypeSoftwareSource,
	"agent":              EventTypeAgent,
	"management_station": EventTypeManagementStation,
}

// GetEventTypeEnumValues Enumerates the set of values for EventTypeEnum
func GetEventTypeEnumValues() []EventTypeEnum {
	values := make([]EventTypeEnum, 0)
	for _, v := range mappingEventTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetEventTypeEnumStringValues Enumerates the set of values in String for EventTypeEnum
func GetEventTypeEnumStringValues() []string {
	return []string{
		"KERNEL_OOPS",
		"KERNEL_CRASH",
		"EXPLOIT_ATTEMPT",
		"SOFTWARE_UPDATE",
		"KSPLICE_UPDATE",
		"SOFTWARE_SOURCE",
		"AGENT",
		"MANAGEMENT_STATION",
	}
}

// GetMappingEventTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEventTypeEnum(val string) (EventTypeEnum, bool) {
	enum, ok := mappingEventTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
