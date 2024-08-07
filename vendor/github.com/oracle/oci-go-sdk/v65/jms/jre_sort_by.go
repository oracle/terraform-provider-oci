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

var mappingJreSortByEnumLowerCase = map[string]JreSortByEnum{
	"distribution":                    JreSortByDistribution,
	"timefirstseen":                   JreSortByTimeFirstSeen,
	"timelastseen":                    JreSortByTimeLastSeen,
	"vendor":                          JreSortByVendor,
	"version":                         JreSortByVersion,
	"approximateinstallationcount":    JreSortByApproximateInstallationCount,
	"approximateapplicationcount":     JreSortByApproximateApplicationCount,
	"approximatemanagedinstancecount": JreSortByApproximateManagedInstanceCount,
	"osname":                          JreSortByOsName,
	"securitystatus":                  JreSortBySecurityStatus,
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
	enum, ok := mappingJreSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
