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

// ContainerVolumeTypeEnum Enum with underlying type: string
type ContainerVolumeTypeEnum string

// Set of constants representing the allowable values for ContainerVolumeTypeEnum
const (
	ContainerVolumeTypeEmptydir   ContainerVolumeTypeEnum = "EMPTYDIR"
	ContainerVolumeTypeConfigfile ContainerVolumeTypeEnum = "CONFIGFILE"
)

var mappingContainerVolumeTypeEnum = map[string]ContainerVolumeTypeEnum{
	"EMPTYDIR":   ContainerVolumeTypeEmptydir,
	"CONFIGFILE": ContainerVolumeTypeConfigfile,
}

var mappingContainerVolumeTypeEnumLowerCase = map[string]ContainerVolumeTypeEnum{
	"emptydir":   ContainerVolumeTypeEmptydir,
	"configfile": ContainerVolumeTypeConfigfile,
}

// GetContainerVolumeTypeEnumValues Enumerates the set of values for ContainerVolumeTypeEnum
func GetContainerVolumeTypeEnumValues() []ContainerVolumeTypeEnum {
	values := make([]ContainerVolumeTypeEnum, 0)
	for _, v := range mappingContainerVolumeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerVolumeTypeEnumStringValues Enumerates the set of values in String for ContainerVolumeTypeEnum
func GetContainerVolumeTypeEnumStringValues() []string {
	return []string{
		"EMPTYDIR",
		"CONFIGFILE",
	}
}

// GetMappingContainerVolumeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerVolumeTypeEnum(val string) (ContainerVolumeTypeEnum, bool) {
	enum, ok := mappingContainerVolumeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
