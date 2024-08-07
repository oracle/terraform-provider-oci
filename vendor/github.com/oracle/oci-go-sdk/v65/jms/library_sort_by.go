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

// LibrarySortByEnum Enum with underlying type: string
type LibrarySortByEnum string

// Set of constants representing the allowable values for LibrarySortByEnum
const (
	LibrarySortByApplicationCount         LibrarySortByEnum = "applicationCount"
	LibrarySortByJavaServerInstanceCount  LibrarySortByEnum = "javaServerInstanceCount"
	LibrarySortByCvssScore                LibrarySortByEnum = "cvssScore"
	LibrarySortByDeployedApplicationCount LibrarySortByEnum = "deployedApplicationCount"
	LibrarySortByLibraryName              LibrarySortByEnum = "libraryName"
	LibrarySortByLibraryVersion           LibrarySortByEnum = "libraryVersion"
	LibrarySortByManagedInstanceCount     LibrarySortByEnum = "managedInstanceCount"
	LibrarySortByTimeFirstSeen            LibrarySortByEnum = "timeFirstSeen"
	LibrarySortByTimeLastSeen             LibrarySortByEnum = "timeLastSeen"
)

var mappingLibrarySortByEnum = map[string]LibrarySortByEnum{
	"applicationCount":         LibrarySortByApplicationCount,
	"javaServerInstanceCount":  LibrarySortByJavaServerInstanceCount,
	"cvssScore":                LibrarySortByCvssScore,
	"deployedApplicationCount": LibrarySortByDeployedApplicationCount,
	"libraryName":              LibrarySortByLibraryName,
	"libraryVersion":           LibrarySortByLibraryVersion,
	"managedInstanceCount":     LibrarySortByManagedInstanceCount,
	"timeFirstSeen":            LibrarySortByTimeFirstSeen,
	"timeLastSeen":             LibrarySortByTimeLastSeen,
}

var mappingLibrarySortByEnumLowerCase = map[string]LibrarySortByEnum{
	"applicationcount":         LibrarySortByApplicationCount,
	"javaserverinstancecount":  LibrarySortByJavaServerInstanceCount,
	"cvssscore":                LibrarySortByCvssScore,
	"deployedapplicationcount": LibrarySortByDeployedApplicationCount,
	"libraryname":              LibrarySortByLibraryName,
	"libraryversion":           LibrarySortByLibraryVersion,
	"managedinstancecount":     LibrarySortByManagedInstanceCount,
	"timefirstseen":            LibrarySortByTimeFirstSeen,
	"timelastseen":             LibrarySortByTimeLastSeen,
}

// GetLibrarySortByEnumValues Enumerates the set of values for LibrarySortByEnum
func GetLibrarySortByEnumValues() []LibrarySortByEnum {
	values := make([]LibrarySortByEnum, 0)
	for _, v := range mappingLibrarySortByEnum {
		values = append(values, v)
	}
	return values
}

// GetLibrarySortByEnumStringValues Enumerates the set of values in String for LibrarySortByEnum
func GetLibrarySortByEnumStringValues() []string {
	return []string{
		"applicationCount",
		"javaServerInstanceCount",
		"cvssScore",
		"deployedApplicationCount",
		"libraryName",
		"libraryVersion",
		"managedInstanceCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingLibrarySortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibrarySortByEnum(val string) (LibrarySortByEnum, bool) {
	enum, ok := mappingLibrarySortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
