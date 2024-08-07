// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ManagedInstanceSortByEnum Enum with underlying type: string
type ManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ManagedInstanceSortByEnum
const (
	ManagedInstanceSortByTimeFirstSeen                ManagedInstanceSortByEnum = "timeFirstSeen"
	ManagedInstanceSortByTimeLastSeen                 ManagedInstanceSortByEnum = "timeLastSeen"
	ManagedInstanceSortByApproximateJreCount          ManagedInstanceSortByEnum = "approximateJreCount"
	ManagedInstanceSortByApproximateInstallationCount ManagedInstanceSortByEnum = "approximateInstallationCount"
	ManagedInstanceSortByApproximateApplicationCount  ManagedInstanceSortByEnum = "approximateApplicationCount"
	ManagedInstanceSortByOsName                       ManagedInstanceSortByEnum = "osName"
)

var mappingManagedInstanceSortByEnum = map[string]ManagedInstanceSortByEnum{
	"timeFirstSeen":                ManagedInstanceSortByTimeFirstSeen,
	"timeLastSeen":                 ManagedInstanceSortByTimeLastSeen,
	"approximateJreCount":          ManagedInstanceSortByApproximateJreCount,
	"approximateInstallationCount": ManagedInstanceSortByApproximateInstallationCount,
	"approximateApplicationCount":  ManagedInstanceSortByApproximateApplicationCount,
	"osName":                       ManagedInstanceSortByOsName,
}

var mappingManagedInstanceSortByEnumLowerCase = map[string]ManagedInstanceSortByEnum{
	"timefirstseen":                ManagedInstanceSortByTimeFirstSeen,
	"timelastseen":                 ManagedInstanceSortByTimeLastSeen,
	"approximatejrecount":          ManagedInstanceSortByApproximateJreCount,
	"approximateinstallationcount": ManagedInstanceSortByApproximateInstallationCount,
	"approximateapplicationcount":  ManagedInstanceSortByApproximateApplicationCount,
	"osname":                       ManagedInstanceSortByOsName,
}

// GetManagedInstanceSortByEnumValues Enumerates the set of values for ManagedInstanceSortByEnum
func GetManagedInstanceSortByEnumValues() []ManagedInstanceSortByEnum {
	values := make([]ManagedInstanceSortByEnum, 0)
	for _, v := range mappingManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedInstanceSortByEnumStringValues Enumerates the set of values in String for ManagedInstanceSortByEnum
func GetManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"timeFirstSeen",
		"timeLastSeen",
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateApplicationCount",
		"osName",
	}
}

// GetMappingManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedInstanceSortByEnum(val string) (ManagedInstanceSortByEnum, bool) {
	enum, ok := mappingManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
