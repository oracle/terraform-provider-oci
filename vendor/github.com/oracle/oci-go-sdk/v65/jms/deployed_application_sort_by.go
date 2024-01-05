// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeployedApplicationSortByEnum Enum with underlying type: string
type DeployedApplicationSortByEnum string

// Set of constants representing the allowable values for DeployedApplicationSortByEnum
const (
	DeployedApplicationSortByApplicationName         DeployedApplicationSortByEnum = "applicationName"
	DeployedApplicationSortByApplicationType         DeployedApplicationSortByEnum = "applicationType"
	DeployedApplicationSortByIsClustered             DeployedApplicationSortByEnum = "isClustered"
	DeployedApplicationSortByJavaServerInstanceCount DeployedApplicationSortByEnum = "javaServerInstanceCount"
	DeployedApplicationSortByTimeFirstSeen           DeployedApplicationSortByEnum = "timeFirstSeen"
	DeployedApplicationSortByTimeLastSeen            DeployedApplicationSortByEnum = "timeLastSeen"
)

var mappingDeployedApplicationSortByEnum = map[string]DeployedApplicationSortByEnum{
	"applicationName":         DeployedApplicationSortByApplicationName,
	"applicationType":         DeployedApplicationSortByApplicationType,
	"isClustered":             DeployedApplicationSortByIsClustered,
	"javaServerInstanceCount": DeployedApplicationSortByJavaServerInstanceCount,
	"timeFirstSeen":           DeployedApplicationSortByTimeFirstSeen,
	"timeLastSeen":            DeployedApplicationSortByTimeLastSeen,
}

var mappingDeployedApplicationSortByEnumLowerCase = map[string]DeployedApplicationSortByEnum{
	"applicationname":         DeployedApplicationSortByApplicationName,
	"applicationtype":         DeployedApplicationSortByApplicationType,
	"isclustered":             DeployedApplicationSortByIsClustered,
	"javaserverinstancecount": DeployedApplicationSortByJavaServerInstanceCount,
	"timefirstseen":           DeployedApplicationSortByTimeFirstSeen,
	"timelastseen":            DeployedApplicationSortByTimeLastSeen,
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
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingDeployedApplicationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployedApplicationSortByEnum(val string) (DeployedApplicationSortByEnum, bool) {
	enum, ok := mappingDeployedApplicationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
