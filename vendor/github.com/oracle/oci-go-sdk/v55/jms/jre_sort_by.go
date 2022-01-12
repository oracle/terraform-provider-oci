// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

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

var mappingJreSortBy = map[string]JreSortByEnum{
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
	for _, v := range mappingJreSortBy {
		values = append(values, v)
	}
	return values
}
