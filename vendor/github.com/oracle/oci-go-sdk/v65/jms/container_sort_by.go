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

// ContainerSortByEnum Enum with underlying type: string
type ContainerSortByEnum string

// Set of constants representing the allowable values for ContainerSortByEnum
const (
	ContainerSortByDisplayName       ContainerSortByEnum = "displayName"
	ContainerSortByNamespace         ContainerSortByEnum = "namespace"
	ContainerSortByPodName           ContainerSortByEnum = "podName"
	ContainerSortByApplicationName   ContainerSortByEnum = "applicationName"
	ContainerSortByJreVersion        ContainerSortByEnum = "jreVersion"
	ContainerSortByJreSecurityStatus ContainerSortByEnum = "jreSecurityStatus"
	ContainerSortByTimeStarted       ContainerSortByEnum = "timeStarted"
)

var mappingContainerSortByEnum = map[string]ContainerSortByEnum{
	"displayName":       ContainerSortByDisplayName,
	"namespace":         ContainerSortByNamespace,
	"podName":           ContainerSortByPodName,
	"applicationName":   ContainerSortByApplicationName,
	"jreVersion":        ContainerSortByJreVersion,
	"jreSecurityStatus": ContainerSortByJreSecurityStatus,
	"timeStarted":       ContainerSortByTimeStarted,
}

var mappingContainerSortByEnumLowerCase = map[string]ContainerSortByEnum{
	"displayname":       ContainerSortByDisplayName,
	"namespace":         ContainerSortByNamespace,
	"podname":           ContainerSortByPodName,
	"applicationname":   ContainerSortByApplicationName,
	"jreversion":        ContainerSortByJreVersion,
	"jresecuritystatus": ContainerSortByJreSecurityStatus,
	"timestarted":       ContainerSortByTimeStarted,
}

// GetContainerSortByEnumValues Enumerates the set of values for ContainerSortByEnum
func GetContainerSortByEnumValues() []ContainerSortByEnum {
	values := make([]ContainerSortByEnum, 0)
	for _, v := range mappingContainerSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerSortByEnumStringValues Enumerates the set of values in String for ContainerSortByEnum
func GetContainerSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"namespace",
		"podName",
		"applicationName",
		"jreVersion",
		"jreSecurityStatus",
		"timeStarted",
	}
}

// GetMappingContainerSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerSortByEnum(val string) (ContainerSortByEnum, bool) {
	enum, ok := mappingContainerSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
