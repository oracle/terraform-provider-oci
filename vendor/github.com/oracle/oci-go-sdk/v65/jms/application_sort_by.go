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

// ApplicationSortByEnum Enum with underlying type: string
type ApplicationSortByEnum string

// Set of constants representing the allowable values for ApplicationSortByEnum
const (
	ApplicationSortByTimeFirstSeen                   ApplicationSortByEnum = "timeFirstSeen"
	ApplicationSortByTimeLastSeen                    ApplicationSortByEnum = "timeLastSeen"
	ApplicationSortByDisplayName                     ApplicationSortByEnum = "displayName"
	ApplicationSortByApproximateJreCount             ApplicationSortByEnum = "approximateJreCount"
	ApplicationSortByApproximateInstallationCount    ApplicationSortByEnum = "approximateInstallationCount"
	ApplicationSortByApproximateManagedInstanceCount ApplicationSortByEnum = "approximateManagedInstanceCount"
	ApplicationSortByApproximateLibraryCount         ApplicationSortByEnum = "approximateLibraryCount"
	ApplicationSortByOsName                          ApplicationSortByEnum = "osName"
)

var mappingApplicationSortByEnum = map[string]ApplicationSortByEnum{
	"timeFirstSeen":                   ApplicationSortByTimeFirstSeen,
	"timeLastSeen":                    ApplicationSortByTimeLastSeen,
	"displayName":                     ApplicationSortByDisplayName,
	"approximateJreCount":             ApplicationSortByApproximateJreCount,
	"approximateInstallationCount":    ApplicationSortByApproximateInstallationCount,
	"approximateManagedInstanceCount": ApplicationSortByApproximateManagedInstanceCount,
	"approximateLibraryCount":         ApplicationSortByApproximateLibraryCount,
	"osName":                          ApplicationSortByOsName,
}

var mappingApplicationSortByEnumLowerCase = map[string]ApplicationSortByEnum{
	"timefirstseen":                   ApplicationSortByTimeFirstSeen,
	"timelastseen":                    ApplicationSortByTimeLastSeen,
	"displayname":                     ApplicationSortByDisplayName,
	"approximatejrecount":             ApplicationSortByApproximateJreCount,
	"approximateinstallationcount":    ApplicationSortByApproximateInstallationCount,
	"approximatemanagedinstancecount": ApplicationSortByApproximateManagedInstanceCount,
	"approximatelibrarycount":         ApplicationSortByApproximateLibraryCount,
	"osname":                          ApplicationSortByOsName,
}

// GetApplicationSortByEnumValues Enumerates the set of values for ApplicationSortByEnum
func GetApplicationSortByEnumValues() []ApplicationSortByEnum {
	values := make([]ApplicationSortByEnum, 0)
	for _, v := range mappingApplicationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationSortByEnumStringValues Enumerates the set of values in String for ApplicationSortByEnum
func GetApplicationSortByEnumStringValues() []string {
	return []string{
		"timeFirstSeen",
		"timeLastSeen",
		"displayName",
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateManagedInstanceCount",
		"approximateLibraryCount",
		"osName",
	}
}

// GetMappingApplicationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationSortByEnum(val string) (ApplicationSortByEnum, bool) {
	enum, ok := mappingApplicationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
