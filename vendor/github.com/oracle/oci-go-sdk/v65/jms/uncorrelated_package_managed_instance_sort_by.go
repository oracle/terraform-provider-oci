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

// UncorrelatedPackageManagedInstanceSortByEnum Enum with underlying type: string
type UncorrelatedPackageManagedInstanceSortByEnum string

// Set of constants representing the allowable values for UncorrelatedPackageManagedInstanceSortByEnum
const (
	UncorrelatedPackageManagedInstanceSortByHostname                UncorrelatedPackageManagedInstanceSortByEnum = "hostname"
	UncorrelatedPackageManagedInstanceSortByApplicationCount        UncorrelatedPackageManagedInstanceSortByEnum = "applicationCount"
	UncorrelatedPackageManagedInstanceSortByLastDetectedDynamically UncorrelatedPackageManagedInstanceSortByEnum = "lastDetectedDynamically"
)

var mappingUncorrelatedPackageManagedInstanceSortByEnum = map[string]UncorrelatedPackageManagedInstanceSortByEnum{
	"hostname":                UncorrelatedPackageManagedInstanceSortByHostname,
	"applicationCount":        UncorrelatedPackageManagedInstanceSortByApplicationCount,
	"lastDetectedDynamically": UncorrelatedPackageManagedInstanceSortByLastDetectedDynamically,
}

var mappingUncorrelatedPackageManagedInstanceSortByEnumLowerCase = map[string]UncorrelatedPackageManagedInstanceSortByEnum{
	"hostname":                UncorrelatedPackageManagedInstanceSortByHostname,
	"applicationcount":        UncorrelatedPackageManagedInstanceSortByApplicationCount,
	"lastdetecteddynamically": UncorrelatedPackageManagedInstanceSortByLastDetectedDynamically,
}

// GetUncorrelatedPackageManagedInstanceSortByEnumValues Enumerates the set of values for UncorrelatedPackageManagedInstanceSortByEnum
func GetUncorrelatedPackageManagedInstanceSortByEnumValues() []UncorrelatedPackageManagedInstanceSortByEnum {
	values := make([]UncorrelatedPackageManagedInstanceSortByEnum, 0)
	for _, v := range mappingUncorrelatedPackageManagedInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetUncorrelatedPackageManagedInstanceSortByEnumStringValues Enumerates the set of values in String for UncorrelatedPackageManagedInstanceSortByEnum
func GetUncorrelatedPackageManagedInstanceSortByEnumStringValues() []string {
	return []string{
		"hostname",
		"applicationCount",
		"lastDetectedDynamically",
	}
}

// GetMappingUncorrelatedPackageManagedInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUncorrelatedPackageManagedInstanceSortByEnum(val string) (UncorrelatedPackageManagedInstanceSortByEnum, bool) {
	enum, ok := mappingUncorrelatedPackageManagedInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
