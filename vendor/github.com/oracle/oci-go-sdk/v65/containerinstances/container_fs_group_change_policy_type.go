// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerFsGroupChangePolicyTypeEnum Enum with underlying type: string
type ContainerFsGroupChangePolicyTypeEnum string

// Set of constants representing the allowable values for ContainerFsGroupChangePolicyTypeEnum
const (
	ContainerFsGroupChangePolicyTypeAlways         ContainerFsGroupChangePolicyTypeEnum = "ALWAYS"
	ContainerFsGroupChangePolicyTypeOnRootMismatch ContainerFsGroupChangePolicyTypeEnum = "ON_ROOT_MISMATCH"
)

var mappingContainerFsGroupChangePolicyTypeEnum = map[string]ContainerFsGroupChangePolicyTypeEnum{
	"ALWAYS":           ContainerFsGroupChangePolicyTypeAlways,
	"ON_ROOT_MISMATCH": ContainerFsGroupChangePolicyTypeOnRootMismatch,
}

var mappingContainerFsGroupChangePolicyTypeEnumLowerCase = map[string]ContainerFsGroupChangePolicyTypeEnum{
	"always":           ContainerFsGroupChangePolicyTypeAlways,
	"on_root_mismatch": ContainerFsGroupChangePolicyTypeOnRootMismatch,
}

// GetContainerFsGroupChangePolicyTypeEnumValues Enumerates the set of values for ContainerFsGroupChangePolicyTypeEnum
func GetContainerFsGroupChangePolicyTypeEnumValues() []ContainerFsGroupChangePolicyTypeEnum {
	values := make([]ContainerFsGroupChangePolicyTypeEnum, 0)
	for _, v := range mappingContainerFsGroupChangePolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerFsGroupChangePolicyTypeEnumStringValues Enumerates the set of values in String for ContainerFsGroupChangePolicyTypeEnum
func GetContainerFsGroupChangePolicyTypeEnumStringValues() []string {
	return []string{
		"ALWAYS",
		"ON_ROOT_MISMATCH",
	}
}

// GetMappingContainerFsGroupChangePolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerFsGroupChangePolicyTypeEnum(val string) (ContainerFsGroupChangePolicyTypeEnum, bool) {
	enum, ok := mappingContainerFsGroupChangePolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
