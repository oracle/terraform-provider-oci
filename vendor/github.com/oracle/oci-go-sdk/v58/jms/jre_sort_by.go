// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// JreSortByEnum Enum with underlying type: string
type JreSortByEnum string

// Set of constants representing the allowable values for JreSortByEnum
const (
	JreSortByDistribution                    JreSortByEnum = "distribution"
	JreSortByTimeFirstSeen                   JreSortByEnum = "timeFirstSeen"
	JreSortByTimeLastSeen                    JreSortByEnum = "timeLastSeen"
	JreSortByVendor                          JreSortByEnum = "vendor"
	JreSortByVersion                         JreSortByEnum = "version"
	JreSortByApproximateInstallationCount    JreSortByEnum = "approximateInstallationCount"
	JreSortByApproximateApplicationCount     JreSortByEnum = "approximateApplicationCount"
	JreSortByApproximateManagedInstanceCount JreSortByEnum = "approximateManagedInstanceCount"
	JreSortByOsName                          JreSortByEnum = "osName"
	JreSortBySecurityStatus                  JreSortByEnum = "securityStatus"
)

var mappingJreSortByEnum = map[string]JreSortByEnum{
	"distribution":                    JreSortByDistribution,
	"timeFirstSeen":                   JreSortByTimeFirstSeen,
	"timeLastSeen":                    JreSortByTimeLastSeen,
	"vendor":                          JreSortByVendor,
	"version":                         JreSortByVersion,
	"approximateInstallationCount":    JreSortByApproximateInstallationCount,
	"approximateApplicationCount":     JreSortByApproximateApplicationCount,
	"approximateManagedInstanceCount": JreSortByApproximateManagedInstanceCount,
	"osName":                          JreSortByOsName,
	"securityStatus":                  JreSortBySecurityStatus,
}

// GetJreSortByEnumValues Enumerates the set of values for JreSortByEnum
func GetJreSortByEnumValues() []JreSortByEnum {
	values := make([]JreSortByEnum, 0)
	for _, v := range mappingJreSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJreSortByEnumStringValues Enumerates the set of values in String for JreSortByEnum
func GetJreSortByEnumStringValues() []string {
	return []string{
		"distribution",
		"timeFirstSeen",
		"timeLastSeen",
		"vendor",
		"version",
		"approximateInstallationCount",
		"approximateApplicationCount",
		"approximateManagedInstanceCount",
		"osName",
		"securityStatus",
	}
}

// GetMappingJreSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJreSortByEnum(val string) (JreSortByEnum, bool) {
	mappingJreSortByEnumIgnoreCase := make(map[string]JreSortByEnum)
	for k, v := range mappingJreSortByEnum {
		mappingJreSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingJreSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
