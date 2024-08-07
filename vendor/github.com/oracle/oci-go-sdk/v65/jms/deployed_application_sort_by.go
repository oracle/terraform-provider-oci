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

// DeployedApplicationSortByEnum Enum with underlying type: string
type DeployedApplicationSortByEnum string

// Set of constants representing the allowable values for DeployedApplicationSortByEnum
const (
	DeployedApplicationSortByApplicationName                    DeployedApplicationSortByEnum = "applicationName"
	DeployedApplicationSortByApplicationType                    DeployedApplicationSortByEnum = "applicationType"
	DeployedApplicationSortByIsClustered                        DeployedApplicationSortByEnum = "isClustered"
	DeployedApplicationSortByJavaServerInstanceCount            DeployedApplicationSortByEnum = "javaServerInstanceCount"
	DeployedApplicationSortByApproximateJavaServerInstanceCount DeployedApplicationSortByEnum = "approximateJavaServerInstanceCount"
	DeployedApplicationSortByApproximateLibraryCount            DeployedApplicationSortByEnum = "approximateLibraryCount"
	DeployedApplicationSortByTimeFirstSeen                      DeployedApplicationSortByEnum = "timeFirstSeen"
	DeployedApplicationSortByTimeLastSeen                       DeployedApplicationSortByEnum = "timeLastSeen"
)

var mappingDeployedApplicationSortByEnum = map[string]DeployedApplicationSortByEnum{
	"applicationName":                    DeployedApplicationSortByApplicationName,
	"applicationType":                    DeployedApplicationSortByApplicationType,
	"isClustered":                        DeployedApplicationSortByIsClustered,
	"javaServerInstanceCount":            DeployedApplicationSortByJavaServerInstanceCount,
	"approximateJavaServerInstanceCount": DeployedApplicationSortByApproximateJavaServerInstanceCount,
	"approximateLibraryCount":            DeployedApplicationSortByApproximateLibraryCount,
	"timeFirstSeen":                      DeployedApplicationSortByTimeFirstSeen,
	"timeLastSeen":                       DeployedApplicationSortByTimeLastSeen,
}

var mappingDeployedApplicationSortByEnumLowerCase = map[string]DeployedApplicationSortByEnum{
	"applicationname":                    DeployedApplicationSortByApplicationName,
	"applicationtype":                    DeployedApplicationSortByApplicationType,
	"isclustered":                        DeployedApplicationSortByIsClustered,
	"javaserverinstancecount":            DeployedApplicationSortByJavaServerInstanceCount,
	"approximatejavaserverinstancecount": DeployedApplicationSortByApproximateJavaServerInstanceCount,
	"approximatelibrarycount":            DeployedApplicationSortByApproximateLibraryCount,
	"timefirstseen":                      DeployedApplicationSortByTimeFirstSeen,
	"timelastseen":                       DeployedApplicationSortByTimeLastSeen,
}

// GetDeployedApplicationSortByEnumValues Enumerates the set of values for DeployedApplicationSortByEnum
func GetDeployedApplicationSortByEnumValues() []DeployedApplicationSortByEnum {
	values := make([]DeployedApplicationSortByEnum, 0)
	for _, v := range mappingDeployedApplicationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployedApplicationSortByEnumStringValues Enumerates the set of values in String for DeployedApplicationSortByEnum
func GetDeployedApplicationSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"applicationType",
		"isClustered",
		"javaServerInstanceCount",
		"approximateJavaServerInstanceCount",
		"approximateLibraryCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingDeployedApplicationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployedApplicationSortByEnum(val string) (DeployedApplicationSortByEnum, bool) {
	enum, ok := mappingDeployedApplicationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
