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

// LibraryApplicationSortByEnum Enum with underlying type: string
type LibraryApplicationSortByEnum string

// Set of constants representing the allowable values for LibraryApplicationSortByEnum
const (
	LibraryApplicationSortByApplicationName         LibraryApplicationSortByEnum = "applicationName"
	LibraryApplicationSortByManagedInstanceCount    LibraryApplicationSortByEnum = "managedInstanceCount"
	LibraryApplicationSortByLastDetectedDynamically LibraryApplicationSortByEnum = "lastDetectedDynamically"
	LibraryApplicationSortByFirstSeenInClasspath    LibraryApplicationSortByEnum = "firstSeenInClasspath"
	LibraryApplicationSortByLastSeenInClasspath     LibraryApplicationSortByEnum = "lastSeenInClasspath"
)

var mappingLibraryApplicationSortByEnum = map[string]LibraryApplicationSortByEnum{
	"applicationName":         LibraryApplicationSortByApplicationName,
	"managedInstanceCount":    LibraryApplicationSortByManagedInstanceCount,
	"lastDetectedDynamically": LibraryApplicationSortByLastDetectedDynamically,
	"firstSeenInClasspath":    LibraryApplicationSortByFirstSeenInClasspath,
	"lastSeenInClasspath":     LibraryApplicationSortByLastSeenInClasspath,
}

var mappingLibraryApplicationSortByEnumLowerCase = map[string]LibraryApplicationSortByEnum{
	"applicationname":         LibraryApplicationSortByApplicationName,
	"managedinstancecount":    LibraryApplicationSortByManagedInstanceCount,
	"lastdetecteddynamically": LibraryApplicationSortByLastDetectedDynamically,
	"firstseeninclasspath":    LibraryApplicationSortByFirstSeenInClasspath,
	"lastseeninclasspath":     LibraryApplicationSortByLastSeenInClasspath,
}

// GetLibraryApplicationSortByEnumValues Enumerates the set of values for LibraryApplicationSortByEnum
func GetLibraryApplicationSortByEnumValues() []LibraryApplicationSortByEnum {
	values := make([]LibraryApplicationSortByEnum, 0)
	for _, v := range mappingLibraryApplicationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryApplicationSortByEnumStringValues Enumerates the set of values in String for LibraryApplicationSortByEnum
func GetLibraryApplicationSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"managedInstanceCount",
		"lastDetectedDynamically",
		"firstSeenInClasspath",
		"lastSeenInClasspath",
	}
}

// GetMappingLibraryApplicationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryApplicationSortByEnum(val string) (LibraryApplicationSortByEnum, bool) {
	enum, ok := mappingLibraryApplicationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
