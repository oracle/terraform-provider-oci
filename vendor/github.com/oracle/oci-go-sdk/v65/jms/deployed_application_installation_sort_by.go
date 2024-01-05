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

// DeployedApplicationInstallationSortByEnum Enum with underlying type: string
type DeployedApplicationInstallationSortByEnum string

// Set of constants representing the allowable values for DeployedApplicationInstallationSortByEnum
const (
	DeployedApplicationInstallationSortByApplicationName         DeployedApplicationInstallationSortByEnum = "applicationName"
	DeployedApplicationInstallationSortByApplicationType         DeployedApplicationInstallationSortByEnum = "applicationType"
	DeployedApplicationInstallationSortByApplicationSourcePath   DeployedApplicationInstallationSortByEnum = "applicationSourcePath"
	DeployedApplicationInstallationSortByIsClustered             DeployedApplicationInstallationSortByEnum = "isClustered"
	DeployedApplicationInstallationSortByJavaServerInstanceCount DeployedApplicationInstallationSortByEnum = "javaServerInstanceCount"
	DeployedApplicationInstallationSortByTimeFirstSeen           DeployedApplicationInstallationSortByEnum = "timeFirstSeen"
	DeployedApplicationInstallationSortByTimeLastSeen            DeployedApplicationInstallationSortByEnum = "timeLastSeen"
)

var mappingDeployedApplicationInstallationSortByEnum = map[string]DeployedApplicationInstallationSortByEnum{
	"applicationName":         DeployedApplicationInstallationSortByApplicationName,
	"applicationType":         DeployedApplicationInstallationSortByApplicationType,
	"applicationSourcePath":   DeployedApplicationInstallationSortByApplicationSourcePath,
	"isClustered":             DeployedApplicationInstallationSortByIsClustered,
	"javaServerInstanceCount": DeployedApplicationInstallationSortByJavaServerInstanceCount,
	"timeFirstSeen":           DeployedApplicationInstallationSortByTimeFirstSeen,
	"timeLastSeen":            DeployedApplicationInstallationSortByTimeLastSeen,
}

var mappingDeployedApplicationInstallationSortByEnumLowerCase = map[string]DeployedApplicationInstallationSortByEnum{
	"applicationname":         DeployedApplicationInstallationSortByApplicationName,
	"applicationtype":         DeployedApplicationInstallationSortByApplicationType,
	"applicationsourcepath":   DeployedApplicationInstallationSortByApplicationSourcePath,
	"isclustered":             DeployedApplicationInstallationSortByIsClustered,
	"javaserverinstancecount": DeployedApplicationInstallationSortByJavaServerInstanceCount,
	"timefirstseen":           DeployedApplicationInstallationSortByTimeFirstSeen,
	"timelastseen":            DeployedApplicationInstallationSortByTimeLastSeen,
}

// GetDeployedApplicationInstallationSortByEnumValues Enumerates the set of values for DeployedApplicationInstallationSortByEnum
func GetDeployedApplicationInstallationSortByEnumValues() []DeployedApplicationInstallationSortByEnum {
	values := make([]DeployedApplicationInstallationSortByEnum, 0)
	for _, v := range mappingDeployedApplicationInstallationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetDeployedApplicationInstallationSortByEnumStringValues Enumerates the set of values in String for DeployedApplicationInstallationSortByEnum
func GetDeployedApplicationInstallationSortByEnumStringValues() []string {
	return []string{
		"applicationName",
		"applicationType",
		"applicationSourcePath",
		"isClustered",
		"javaServerInstanceCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingDeployedApplicationInstallationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeployedApplicationInstallationSortByEnum(val string) (DeployedApplicationInstallationSortByEnum, bool) {
	enum, ok := mappingDeployedApplicationInstallationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
