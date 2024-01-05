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

// JavaServerInstanceSortByEnum Enum with underlying type: string
type JavaServerInstanceSortByEnum string

// Set of constants representing the allowable values for JavaServerInstanceSortByEnum
const (
	JavaServerInstanceSortByServerInstanceName                  JavaServerInstanceSortByEnum = "serverInstanceName"
	JavaServerInstanceSortByManagedInstanceName                 JavaServerInstanceSortByEnum = "managedInstanceName"
	JavaServerInstanceSortByApproximateDeployedApplicationCount JavaServerInstanceSortByEnum = "approximateDeployedApplicationCount"
	JavaServerInstanceSortByTimeFirstSeen                       JavaServerInstanceSortByEnum = "timeFirstSeen"
	JavaServerInstanceSortByTimeLastSeen                        JavaServerInstanceSortByEnum = "timeLastSeen"
)

var mappingJavaServerInstanceSortByEnum = map[string]JavaServerInstanceSortByEnum{
	"serverInstanceName":                  JavaServerInstanceSortByServerInstanceName,
	"managedInstanceName":                 JavaServerInstanceSortByManagedInstanceName,
	"approximateDeployedApplicationCount": JavaServerInstanceSortByApproximateDeployedApplicationCount,
	"timeFirstSeen":                       JavaServerInstanceSortByTimeFirstSeen,
	"timeLastSeen":                        JavaServerInstanceSortByTimeLastSeen,
}

var mappingJavaServerInstanceSortByEnumLowerCase = map[string]JavaServerInstanceSortByEnum{
	"serverinstancename":                  JavaServerInstanceSortByServerInstanceName,
	"managedinstancename":                 JavaServerInstanceSortByManagedInstanceName,
	"approximatedeployedapplicationcount": JavaServerInstanceSortByApproximateDeployedApplicationCount,
	"timefirstseen":                       JavaServerInstanceSortByTimeFirstSeen,
	"timelastseen":                        JavaServerInstanceSortByTimeLastSeen,
}

// GetJavaServerInstanceSortByEnumValues Enumerates the set of values for JavaServerInstanceSortByEnum
func GetJavaServerInstanceSortByEnumValues() []JavaServerInstanceSortByEnum {
	values := make([]JavaServerInstanceSortByEnum, 0)
	for _, v := range mappingJavaServerInstanceSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaServerInstanceSortByEnumStringValues Enumerates the set of values in String for JavaServerInstanceSortByEnum
func GetJavaServerInstanceSortByEnumStringValues() []string {
	return []string{
		"serverInstanceName",
		"managedInstanceName",
		"approximateDeployedApplicationCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingJavaServerInstanceSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaServerInstanceSortByEnum(val string) (JavaServerInstanceSortByEnum, bool) {
	enum, ok := mappingJavaServerInstanceSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
