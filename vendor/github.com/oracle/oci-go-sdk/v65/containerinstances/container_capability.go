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

// ContainerCapabilityEnum Enum with underlying type: string
type ContainerCapabilityEnum string

// Set of constants representing the allowable values for ContainerCapabilityEnum
const (
	ContainerCapabilityCapNetAdmin ContainerCapabilityEnum = "CAP_NET_ADMIN"
	ContainerCapabilityCapNetRaw   ContainerCapabilityEnum = "CAP_NET_RAW"
)

var mappingContainerCapabilityEnum = map[string]ContainerCapabilityEnum{
	"CAP_NET_ADMIN": ContainerCapabilityCapNetAdmin,
	"CAP_NET_RAW":   ContainerCapabilityCapNetRaw,
}

var mappingContainerCapabilityEnumLowerCase = map[string]ContainerCapabilityEnum{
	"cap_net_admin": ContainerCapabilityCapNetAdmin,
	"cap_net_raw":   ContainerCapabilityCapNetRaw,
}

// GetContainerCapabilityEnumValues Enumerates the set of values for ContainerCapabilityEnum
func GetContainerCapabilityEnumValues() []ContainerCapabilityEnum {
	values := make([]ContainerCapabilityEnum, 0)
	for _, v := range mappingContainerCapabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerCapabilityEnumStringValues Enumerates the set of values in String for ContainerCapabilityEnum
func GetContainerCapabilityEnumStringValues() []string {
	return []string{
		"CAP_NET_ADMIN",
		"CAP_NET_RAW",
	}
}

// GetMappingContainerCapabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerCapabilityEnum(val string) (ContainerCapabilityEnum, bool) {
	enum, ok := mappingContainerCapabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
