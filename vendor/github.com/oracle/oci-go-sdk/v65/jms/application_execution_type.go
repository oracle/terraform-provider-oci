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

// ApplicationExecutionTypeEnum Enum with underlying type: string
type ApplicationExecutionTypeEnum string

// Set of constants representing the allowable values for ApplicationExecutionTypeEnum
const (
	ApplicationExecutionTypeInstalled ApplicationExecutionTypeEnum = "INSTALLED"
	ApplicationExecutionTypeDeployed  ApplicationExecutionTypeEnum = "DEPLOYED"
)

var mappingApplicationExecutionTypeEnum = map[string]ApplicationExecutionTypeEnum{
	"INSTALLED": ApplicationExecutionTypeInstalled,
	"DEPLOYED":  ApplicationExecutionTypeDeployed,
}

var mappingApplicationExecutionTypeEnumLowerCase = map[string]ApplicationExecutionTypeEnum{
	"installed": ApplicationExecutionTypeInstalled,
	"deployed":  ApplicationExecutionTypeDeployed,
}

// GetApplicationExecutionTypeEnumValues Enumerates the set of values for ApplicationExecutionTypeEnum
func GetApplicationExecutionTypeEnumValues() []ApplicationExecutionTypeEnum {
	values := make([]ApplicationExecutionTypeEnum, 0)
	for _, v := range mappingApplicationExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationExecutionTypeEnumStringValues Enumerates the set of values in String for ApplicationExecutionTypeEnum
func GetApplicationExecutionTypeEnumStringValues() []string {
	return []string{
		"INSTALLED",
		"DEPLOYED",
	}
}

// GetMappingApplicationExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationExecutionTypeEnum(val string) (ApplicationExecutionTypeEnum, bool) {
	enum, ok := mappingApplicationExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
