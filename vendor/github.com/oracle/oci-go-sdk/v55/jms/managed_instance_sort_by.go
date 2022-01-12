// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

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

var mappingManagedInstanceSortBy = map[string]ManagedInstanceSortByEnum{
	"timeFirstSeen":                ManagedInstanceSortByTimeFirstSeen,
	"timeLastSeen":                 ManagedInstanceSortByTimeLastSeen,
	"approximateJreCount":          ManagedInstanceSortByApproximateJreCount,
	"approximateInstallationCount": ManagedInstanceSortByApproximateInstallationCount,
	"approximateApplicationCount":  ManagedInstanceSortByApproximateApplicationCount,
	"osName":                       ManagedInstanceSortByOsName,
}

// GetManagedInstanceSortByEnumValues Enumerates the set of values for ManagedInstanceSortByEnum
func GetManagedInstanceSortByEnumValues() []ManagedInstanceSortByEnum {
	values := make([]ManagedInstanceSortByEnum, 0)
	for _, v := range mappingManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
