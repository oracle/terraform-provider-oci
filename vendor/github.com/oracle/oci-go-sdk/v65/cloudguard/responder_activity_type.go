// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ResponderActivityTypeEnum Enum with underlying type: string
type ResponderActivityTypeEnum string

// Set of constants representing the allowable values for ResponderActivityTypeEnum
const (
	ResponderActivityTypeStarted   ResponderActivityTypeEnum = "STARTED"
	ResponderActivityTypeCompleted ResponderActivityTypeEnum = "COMPLETED"
)

var mappingResponderActivityTypeEnum = map[string]ResponderActivityTypeEnum{
	"STARTED":   ResponderActivityTypeStarted,
	"COMPLETED": ResponderActivityTypeCompleted,
}

var mappingResponderActivityTypeEnumLowerCase = map[string]ResponderActivityTypeEnum{
	"started":   ResponderActivityTypeStarted,
	"completed": ResponderActivityTypeCompleted,
}

// GetResponderActivityTypeEnumValues Enumerates the set of values for ResponderActivityTypeEnum
func GetResponderActivityTypeEnumValues() []ResponderActivityTypeEnum {
	values := make([]ResponderActivityTypeEnum, 0)
	for _, v := range mappingResponderActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResponderActivityTypeEnumStringValues Enumerates the set of values in String for ResponderActivityTypeEnum
func GetResponderActivityTypeEnumStringValues() []string {
	return []string{
		"STARTED",
		"COMPLETED",
	}
}

// GetMappingResponderActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResponderActivityTypeEnum(val string) (ResponderActivityTypeEnum, bool) {
	enum, ok := mappingResponderActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
