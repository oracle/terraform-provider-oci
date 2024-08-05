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

// PlatformArchitectureTypeEnum Enum with underlying type: string
type PlatformArchitectureTypeEnum string

// Set of constants representing the allowable values for PlatformArchitectureTypeEnum
const (
	PlatformArchitectureTypeX8664   PlatformArchitectureTypeEnum = "X86_64"
	PlatformArchitectureTypeX86     PlatformArchitectureTypeEnum = "X86"
	PlatformArchitectureTypeAarch64 PlatformArchitectureTypeEnum = "AARCH64"
)

var mappingPlatformArchitectureTypeEnum = map[string]PlatformArchitectureTypeEnum{
	"X86_64":  PlatformArchitectureTypeX8664,
	"X86":     PlatformArchitectureTypeX86,
	"AARCH64": PlatformArchitectureTypeAarch64,
}

var mappingPlatformArchitectureTypeEnumLowerCase = map[string]PlatformArchitectureTypeEnum{
	"x86_64":  PlatformArchitectureTypeX8664,
	"x86":     PlatformArchitectureTypeX86,
	"aarch64": PlatformArchitectureTypeAarch64,
}

// GetPlatformArchitectureTypeEnumValues Enumerates the set of values for PlatformArchitectureTypeEnum
func GetPlatformArchitectureTypeEnumValues() []PlatformArchitectureTypeEnum {
	values := make([]PlatformArchitectureTypeEnum, 0)
	for _, v := range mappingPlatformArchitectureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPlatformArchitectureTypeEnumStringValues Enumerates the set of values in String for PlatformArchitectureTypeEnum
func GetPlatformArchitectureTypeEnumStringValues() []string {
	return []string{
		"X86_64",
		"X86",
		"AARCH64",
	}
}

// GetMappingPlatformArchitectureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPlatformArchitectureTypeEnum(val string) (PlatformArchitectureTypeEnum, bool) {
	enum, ok := mappingPlatformArchitectureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
