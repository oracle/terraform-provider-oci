// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AnnouncementSortByEnum Enum with underlying type: string
type AnnouncementSortByEnum string

// Set of constants representing the allowable values for AnnouncementSortByEnum
const (
	AnnouncementSortByTimeReleased AnnouncementSortByEnum = "timeReleased"
	AnnouncementSortBySummary      AnnouncementSortByEnum = "summary"
)

var mappingAnnouncementSortByEnum = map[string]AnnouncementSortByEnum{
	"timeReleased": AnnouncementSortByTimeReleased,
	"summary":      AnnouncementSortBySummary,
}

var mappingAnnouncementSortByEnumLowerCase = map[string]AnnouncementSortByEnum{
	"timereleased": AnnouncementSortByTimeReleased,
	"summary":      AnnouncementSortBySummary,
}

// GetAnnouncementSortByEnumValues Enumerates the set of values for AnnouncementSortByEnum
func GetAnnouncementSortByEnumValues() []AnnouncementSortByEnum {
	values := make([]AnnouncementSortByEnum, 0)
	for _, v := range mappingAnnouncementSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAnnouncementSortByEnumStringValues Enumerates the set of values in String for AnnouncementSortByEnum
func GetAnnouncementSortByEnumStringValues() []string {
	return []string{
		"timeReleased",
		"summary",
	}
}

// GetMappingAnnouncementSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnnouncementSortByEnum(val string) (AnnouncementSortByEnum, bool) {
	enum, ok := mappingAnnouncementSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
