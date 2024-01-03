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

// InstallationSortByEnum Enum with underlying type: string
type InstallationSortByEnum string

// Set of constants representing the allowable values for InstallationSortByEnum
const (
	InstallationSortByJreDistribution                 InstallationSortByEnum = "jreDistribution"
	InstallationSortByJreVendor                       InstallationSortByEnum = "jreVendor"
	InstallationSortByJreVersion                      InstallationSortByEnum = "jreVersion"
	InstallationSortByPath                            InstallationSortByEnum = "path"
	InstallationSortByTimeFirstSeen                   InstallationSortByEnum = "timeFirstSeen"
	InstallationSortByTimeLastSeen                    InstallationSortByEnum = "timeLastSeen"
	InstallationSortByApproximateApplicationCount     InstallationSortByEnum = "approximateApplicationCount"
	InstallationSortByApproximateManagedInstanceCount InstallationSortByEnum = "approximateManagedInstanceCount"
	InstallationSortByOsName                          InstallationSortByEnum = "osName"
)

var mappingInstallationSortByEnum = map[string]InstallationSortByEnum{
	"jreDistribution":                 InstallationSortByJreDistribution,
	"jreVendor":                       InstallationSortByJreVendor,
	"jreVersion":                      InstallationSortByJreVersion,
	"path":                            InstallationSortByPath,
	"timeFirstSeen":                   InstallationSortByTimeFirstSeen,
	"timeLastSeen":                    InstallationSortByTimeLastSeen,
	"approximateApplicationCount":     InstallationSortByApproximateApplicationCount,
	"approximateManagedInstanceCount": InstallationSortByApproximateManagedInstanceCount,
	"osName":                          InstallationSortByOsName,
}

var mappingInstallationSortByEnumLowerCase = map[string]InstallationSortByEnum{
	"jredistribution":                 InstallationSortByJreDistribution,
	"jrevendor":                       InstallationSortByJreVendor,
	"jreversion":                      InstallationSortByJreVersion,
	"path":                            InstallationSortByPath,
	"timefirstseen":                   InstallationSortByTimeFirstSeen,
	"timelastseen":                    InstallationSortByTimeLastSeen,
	"approximateapplicationcount":     InstallationSortByApproximateApplicationCount,
	"approximatemanagedinstancecount": InstallationSortByApproximateManagedInstanceCount,
	"osname":                          InstallationSortByOsName,
}

// GetInstallationSortByEnumValues Enumerates the set of values for InstallationSortByEnum
func GetInstallationSortByEnumValues() []InstallationSortByEnum {
	values := make([]InstallationSortByEnum, 0)
	for _, v := range mappingInstallationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetInstallationSortByEnumStringValues Enumerates the set of values in String for InstallationSortByEnum
func GetInstallationSortByEnumStringValues() []string {
	return []string{
		"jreDistribution",
		"jreVendor",
		"jreVersion",
		"path",
		"timeFirstSeen",
		"timeLastSeen",
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
		"osName",
	}
}

// GetMappingInstallationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstallationSortByEnum(val string) (InstallationSortByEnum, bool) {
	enum, ok := mappingInstallationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
