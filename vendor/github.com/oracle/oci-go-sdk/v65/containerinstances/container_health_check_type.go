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

// ContainerHealthCheckTypeEnum Enum with underlying type: string
type ContainerHealthCheckTypeEnum string

// Set of constants representing the allowable values for ContainerHealthCheckTypeEnum
const (
	ContainerHealthCheckTypeHttp    ContainerHealthCheckTypeEnum = "HTTP"
	ContainerHealthCheckTypeTcp     ContainerHealthCheckTypeEnum = "TCP"
	ContainerHealthCheckTypeCommand ContainerHealthCheckTypeEnum = "COMMAND"
)

var mappingContainerHealthCheckTypeEnum = map[string]ContainerHealthCheckTypeEnum{
	"HTTP":    ContainerHealthCheckTypeHttp,
	"TCP":     ContainerHealthCheckTypeTcp,
	"COMMAND": ContainerHealthCheckTypeCommand,
}

var mappingContainerHealthCheckTypeEnumLowerCase = map[string]ContainerHealthCheckTypeEnum{
	"http":    ContainerHealthCheckTypeHttp,
	"tcp":     ContainerHealthCheckTypeTcp,
	"command": ContainerHealthCheckTypeCommand,
}

// GetContainerHealthCheckTypeEnumValues Enumerates the set of values for ContainerHealthCheckTypeEnum
func GetContainerHealthCheckTypeEnumValues() []ContainerHealthCheckTypeEnum {
	values := make([]ContainerHealthCheckTypeEnum, 0)
	for _, v := range mappingContainerHealthCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerHealthCheckTypeEnumStringValues Enumerates the set of values in String for ContainerHealthCheckTypeEnum
func GetContainerHealthCheckTypeEnumStringValues() []string {
	return []string{
		"HTTP",
		"TCP",
		"COMMAND",
	}
}

// GetMappingContainerHealthCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerHealthCheckTypeEnum(val string) (ContainerHealthCheckTypeEnum, bool) {
	enum, ok := mappingContainerHealthCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
