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

// ApplicationInstallationSortByEnum Enum with underlying type: string
type ApplicationInstallationSortByEnum string

// Set of constants representing the allowable values for ApplicationInstallationSortByEnum
const (
	ApplicationInstallationSortByTimeFirstSeen                   ApplicationInstallationSortByEnum = "timeFirstSeen"
	ApplicationInstallationSortByTimeLastSeen                    ApplicationInstallationSortByEnum = "timeLastSeen"
	ApplicationInstallationSortByDisplayName                     ApplicationInstallationSortByEnum = "displayName"
	ApplicationInstallationSortByInstallationPath                ApplicationInstallationSortByEnum = "installationPath"
	ApplicationInstallationSortByOsName                          ApplicationInstallationSortByEnum = "osName"
	ApplicationInstallationSortByApproximateJreCount             ApplicationInstallationSortByEnum = "approximateJreCount"
	ApplicationInstallationSortByApproximateInstallationCount    ApplicationInstallationSortByEnum = "approximateInstallationCount"
	ApplicationInstallationSortByApproximateManagedInstanceCount ApplicationInstallationSortByEnum = "approximateManagedInstanceCount"
	ApplicationInstallationSortByApproximateLibraryCount         ApplicationInstallationSortByEnum = "approximateLibraryCount"
)

var mappingApplicationInstallationSortByEnum = map[string]ApplicationInstallationSortByEnum{
	"timeFirstSeen":                   ApplicationInstallationSortByTimeFirstSeen,
	"timeLastSeen":                    ApplicationInstallationSortByTimeLastSeen,
	"displayName":                     ApplicationInstallationSortByDisplayName,
	"installationPath":                ApplicationInstallationSortByInstallationPath,
	"osName":                          ApplicationInstallationSortByOsName,
	"approximateJreCount":             ApplicationInstallationSortByApproximateJreCount,
	"approximateInstallationCount":    ApplicationInstallationSortByApproximateInstallationCount,
	"approximateManagedInstanceCount": ApplicationInstallationSortByApproximateManagedInstanceCount,
	"approximateLibraryCount":         ApplicationInstallationSortByApproximateLibraryCount,
}

var mappingApplicationInstallationSortByEnumLowerCase = map[string]ApplicationInstallationSortByEnum{
	"timefirstseen":                   ApplicationInstallationSortByTimeFirstSeen,
	"timelastseen":                    ApplicationInstallationSortByTimeLastSeen,
	"displayname":                     ApplicationInstallationSortByDisplayName,
	"installationpath":                ApplicationInstallationSortByInstallationPath,
	"osname":                          ApplicationInstallationSortByOsName,
	"approximatejrecount":             ApplicationInstallationSortByApproximateJreCount,
	"approximateinstallationcount":    ApplicationInstallationSortByApproximateInstallationCount,
	"approximatemanagedinstancecount": ApplicationInstallationSortByApproximateManagedInstanceCount,
	"approximatelibrarycount":         ApplicationInstallationSortByApproximateLibraryCount,
}

// GetApplicationInstallationSortByEnumValues Enumerates the set of values for ApplicationInstallationSortByEnum
func GetApplicationInstallationSortByEnumValues() []ApplicationInstallationSortByEnum {
	values := make([]ApplicationInstallationSortByEnum, 0)
	for _, v := range mappingApplicationInstallationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationInstallationSortByEnumStringValues Enumerates the set of values in String for ApplicationInstallationSortByEnum
func GetApplicationInstallationSortByEnumStringValues() []string {
	return []string{
		"timeFirstSeen",
		"timeLastSeen",
		"displayName",
		"installationPath",
		"osName",
		"approximateJreCount",
		"approximateInstallationCount",
		"approximateManagedInstanceCount",
		"approximateLibraryCount",
	}
}

// GetMappingApplicationInstallationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationInstallationSortByEnum(val string) (ApplicationInstallationSortByEnum, bool) {
	enum, ok := mappingApplicationInstallationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
