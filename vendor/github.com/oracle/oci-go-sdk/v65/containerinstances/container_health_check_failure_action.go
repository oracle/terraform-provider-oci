// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"strings"
)

// ContainerHealthCheckFailureActionEnum Enum with underlying type: string
type ContainerHealthCheckFailureActionEnum string

// Set of constants representing the allowable values for ContainerHealthCheckFailureActionEnum
const (
	ContainerHealthCheckFailureActionKill ContainerHealthCheckFailureActionEnum = "KILL"
	ContainerHealthCheckFailureActionNone ContainerHealthCheckFailureActionEnum = "NONE"
)

var mappingContainerHealthCheckFailureActionEnum = map[string]ContainerHealthCheckFailureActionEnum{
	"KILL": ContainerHealthCheckFailureActionKill,
	"NONE": ContainerHealthCheckFailureActionNone,
}

var mappingContainerHealthCheckFailureActionEnumLowerCase = map[string]ContainerHealthCheckFailureActionEnum{
	"kill": ContainerHealthCheckFailureActionKill,
	"none": ContainerHealthCheckFailureActionNone,
}

// GetContainerHealthCheckFailureActionEnumValues Enumerates the set of values for ContainerHealthCheckFailureActionEnum
func GetContainerHealthCheckFailureActionEnumValues() []ContainerHealthCheckFailureActionEnum {
	values := make([]ContainerHealthCheckFailureActionEnum, 0)
	for _, v := range mappingContainerHealthCheckFailureActionEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerHealthCheckFailureActionEnumStringValues Enumerates the set of values in String for ContainerHealthCheckFailureActionEnum
func GetContainerHealthCheckFailureActionEnumStringValues() []string {
	return []string{
		"KILL",
		"NONE",
	}
}

// GetMappingContainerHealthCheckFailureActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerHealthCheckFailureActionEnum(val string) (ContainerHealthCheckFailureActionEnum, bool) {
	enum, ok := mappingContainerHealthCheckFailureActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
