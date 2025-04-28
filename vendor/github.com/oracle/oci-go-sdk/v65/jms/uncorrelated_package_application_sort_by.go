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

// UncorrelatedPackageApplicationSortByEnum Enum with underlying type: string
type UncorrelatedPackageApplicationSortByEnum string

// Set of constants representing the allowable values for UncorrelatedPackageApplicationSortByEnum
const (
	UncorrelatedPackageApplicationSortByApplicationName         UncorrelatedPackageApplicationSortByEnum = "applicationName"
	UncorrelatedPackageApplicationSortByManagedInstanceCount    UncorrelatedPackageApplicationSortByEnum = "managedInstanceCount"
	UncorrelatedPackageApplicationSortByLastDetectedDynamically UncorrelatedPackageApplicationSortByEnum = "lastDetectedDynamically"
)

var mappingUncorrelatedPackageApplicationSortByEnum = map[string]UncorrelatedPackageApplicationSortByEnum{
	"applicationName":         UncorrelatedPackageApplicationSortByApplicationName,
	"managedInstanceCount":    UncorrelatedPackageApplicationSortByManagedInstanceCount,
	"lastDetectedDynamically": UncorrelatedPackageApplicationSortByLastDetectedDynamically,
}

var mappingUncorrelatedPackageApplicationSortByEnumLowerCase = map[string]UncorrelatedPackageApplicationSortByEnum{
	"applicationname":         UncorrelatedPackageApplicationSortByApplicationName,
	"managedinstancecount":    UncorrelatedPackageApplicationSortByManagedInstanceCount,
	"lastdetecteddynamically": UncorrelatedPackageApplicationSortByLastDetectedDynamically,
}

// GetUncorrelatedPackageApplicationSortByEnumValues Enumerates the set of values for UncorrelatedPackageApplicationSortByEnum
func GetUncorrelatedPackageApplicationSortByEnumValues() []UncorrelatedPackageApplicationSortByEnum {
	values := make([]UncorrelatedPackageApplicationSortByEnum, 0)
	for _, v := range mappingUncorrelatedPackageApplicationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetUncorrelatedPackageApplicationSortByEnumStringValues Enumerates the set of values in String for UncorrelatedPackageApplicationSortByEnum
func GetUncorrelatedPackageApplicationSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"managedInstanceCount",
		"lastDetectedDynamically",
	}
}

// GetMappingUncorrelatedPackageApplicationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUncorrelatedPackageApplicationSortByEnum(val string) (UncorrelatedPackageApplicationSortByEnum, bool) {
	enum, ok := mappingUncorrelatedPackageApplicationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
