// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"strings"
)

// TargetResourceEntityTypeEnum Enum with underlying type: string
type TargetResourceEntityTypeEnum string

// Set of constants representing the allowable values for TargetResourceEntityTypeEnum
const (
	TargetResourceEntityTypeInstance             TargetResourceEntityTypeEnum = "INSTANCE"
	TargetResourceEntityTypeGroup                TargetResourceEntityTypeEnum = "GROUP"
	TargetResourceEntityTypeCompartment          TargetResourceEntityTypeEnum = "COMPARTMENT"
	TargetResourceEntityTypeLifecycleEnvironment TargetResourceEntityTypeEnum = "LIFECYCLE_ENVIRONMENT"
	TargetResourceEntityTypeSoftwareSource       TargetResourceEntityTypeEnum = "SOFTWARE_SOURCE"
)

var mappingTargetResourceEntityTypeEnum = map[string]TargetResourceEntityTypeEnum{
	"INSTANCE":              TargetResourceEntityTypeInstance,
	"GROUP":                 TargetResourceEntityTypeGroup,
	"COMPARTMENT":           TargetResourceEntityTypeCompartment,
	"LIFECYCLE_ENVIRONMENT": TargetResourceEntityTypeLifecycleEnvironment,
	"SOFTWARE_SOURCE":       TargetResourceEntityTypeSoftwareSource,
}

var mappingTargetResourceEntityTypeEnumLowerCase = map[string]TargetResourceEntityTypeEnum{
	"instance":              TargetResourceEntityTypeInstance,
	"group":                 TargetResourceEntityTypeGroup,
	"compartment":           TargetResourceEntityTypeCompartment,
	"lifecycle_environment": TargetResourceEntityTypeLifecycleEnvironment,
	"software_source":       TargetResourceEntityTypeSoftwareSource,
}

// GetTargetResourceEntityTypeEnumValues Enumerates the set of values for TargetResourceEntityTypeEnum
func GetTargetResourceEntityTypeEnumValues() []TargetResourceEntityTypeEnum {
	values := make([]TargetResourceEntityTypeEnum, 0)
	for _, v := range mappingTargetResourceEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetResourceEntityTypeEnumStringValues Enumerates the set of values in String for TargetResourceEntityTypeEnum
func GetTargetResourceEntityTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
		"GROUP",
		"COMPARTMENT",
		"LIFECYCLE_ENVIRONMENT",
		"SOFTWARE_SOURCE",
	}
}

// GetMappingTargetResourceEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetResourceEntityTypeEnum(val string) (TargetResourceEntityTypeEnum, bool) {
	enum, ok := mappingTargetResourceEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
