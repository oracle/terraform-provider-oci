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

// LibraryManagedInstanceSortByEnum Enum with underlying type: string
type LibraryManagedInstanceSortByEnum string

// Set of constants representing the allowable values for LibraryManagedInstanceSortByEnum
const (
	LibraryManagedInstanceSortByHostname                LibraryManagedInstanceSortByEnum = "hostname"
	LibraryManagedInstanceSortByApplicationCount        LibraryManagedInstanceSortByEnum = "applicationCount"
	LibraryManagedInstanceSortByLastDetectedDynamically LibraryManagedInstanceSortByEnum = "lastDetectedDynamically"
	LibraryManagedInstanceSortByFirstSeenInClasspath    LibraryManagedInstanceSortByEnum = "firstSeenInClasspath"
	LibraryManagedInstanceSortByLastSeenInClasspath     LibraryManagedInstanceSortByEnum = "lastSeenInClasspath"
)

var mappingLibraryManagedInstanceSortByEnum = map[string]LibraryManagedInstanceSortByEnum{
	"hostname":                LibraryManagedInstanceSortByHostname,
	"applicationCount":        LibraryManagedInstanceSortByApplicationCount,
	"lastDetectedDynamically": LibraryManagedInstanceSortByLastDetectedDynamically,
	"firstSeenInClasspath":    LibraryManagedInstanceSortByFirstSeenInClasspath,
	"lastSeenInClasspath":     LibraryManagedInstanceSortByLastSeenInClasspath,
}

var mappingLibraryManagedInstanceSortByEnumLowerCase = map[string]LibraryManagedInstanceSortByEnum{
	"hostname":                LibraryManagedInstanceSortByHostname,
	"applicationcount":        LibraryManagedInstanceSortByApplicationCount,
	"lastdetecteddynamically": LibraryManagedInstanceSortByLastDetectedDynamically,
	"firstseeninclasspath":    LibraryManagedInstanceSortByFirstSeenInClasspath,
	"lastseeninclasspath":     LibraryManagedInstanceSortByLastSeenInClasspath,
}

// GetLibraryManagedInstanceSortByEnumValues Enumerates the set of values for LibraryManagedInstanceSortByEnum
func GetLibraryManagedInstanceSortByEnumValues() []LibraryManagedInstanceSortByEnum {
	values := make([]LibraryManagedInstanceSortByEnum, 0)
	for _, v := range mappingLibraryManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetLibraryManagedInstanceSortByEnumStringValues Enumerates the set of values in String for LibraryManagedInstanceSortByEnum
func GetLibraryManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"hostname",
		"applicationCount",
		"lastDetectedDynamically",
		"firstSeenInClasspath",
		"lastSeenInClasspath",
	}
}

// GetMappingLibraryManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLibraryManagedInstanceSortByEnum(val string) (LibraryManagedInstanceSortByEnum, bool) {
	enum, ok := mappingLibraryManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
