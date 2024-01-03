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

// JavaServerSortByEnum Enum with underlying type: string
type JavaServerSortByEnum string

// Set of constants representing the allowable values for JavaServerSortByEnum
const (
	JavaServerSortByServerName                          JavaServerSortByEnum = "serverName"
	JavaServerSortByServerVersion                       JavaServerSortByEnum = "serverVersion"
	JavaServerSortByServerInstanceCount                 JavaServerSortByEnum = "serverInstanceCount"
	JavaServerSortByApproximateDeployedApplicationCount JavaServerSortByEnum = "approximateDeployedApplicationCount"
	JavaServerSortByTimeFirstSeen                       JavaServerSortByEnum = "timeFirstSeen"
	JavaServerSortByTimeLastSeen                        JavaServerSortByEnum = "timeLastSeen"
)

var mappingJavaServerSortByEnum = map[string]JavaServerSortByEnum{
	"serverName":                          JavaServerSortByServerName,
	"serverVersion":                       JavaServerSortByServerVersion,
	"serverInstanceCount":                 JavaServerSortByServerInstanceCount,
	"approximateDeployedApplicationCount": JavaServerSortByApproximateDeployedApplicationCount,
	"timeFirstSeen":                       JavaServerSortByTimeFirstSeen,
	"timeLastSeen":                        JavaServerSortByTimeLastSeen,
}

var mappingJavaServerSortByEnumLowerCase = map[string]JavaServerSortByEnum{
	"servername":                          JavaServerSortByServerName,
	"serverversion":                       JavaServerSortByServerVersion,
	"serverinstancecount":                 JavaServerSortByServerInstanceCount,
	"approximatedeployedapplicationcount": JavaServerSortByApproximateDeployedApplicationCount,
	"timefirstseen":                       JavaServerSortByTimeFirstSeen,
	"timelastseen":                        JavaServerSortByTimeLastSeen,
}

// GetJavaServerSortByEnumValues Enumerates the set of values for JavaServerSortByEnum
func GetJavaServerSortByEnumValues() []JavaServerSortByEnum {
	values := make([]JavaServerSortByEnum, 0)
	for _, v := range mappingJavaServerSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetJavaServerSortByEnumStringValues Enumerates the set of values in String for JavaServerSortByEnum
func GetJavaServerSortByEnumStringValues() []string {
	return []string{
		"serverName",
		"serverVersion",
		"serverInstanceCount",
		"approximateDeployedApplicationCount",
		"timeFirstSeen",
		"timeLastSeen",
	}
}

// GetMappingJavaServerSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJavaServerSortByEnum(val string) (JavaServerSortByEnum, bool) {
	enum, ok := mappingJavaServerSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
