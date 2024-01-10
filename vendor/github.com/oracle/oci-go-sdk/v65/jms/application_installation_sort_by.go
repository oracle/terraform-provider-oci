// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
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
	}
}

// GetMappingApplicationInstallationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationInstallationSortByEnum(val string) (ApplicationInstallationSortByEnum, bool) {
	enum, ok := mappingApplicationInstallationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
