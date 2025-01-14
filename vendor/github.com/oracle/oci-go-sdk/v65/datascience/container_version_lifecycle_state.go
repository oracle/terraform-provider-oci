// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ContainerVersionLifecycleStateEnum Enum with underlying type: string
type ContainerVersionLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerVersionLifecycleStateEnum
const (
	ContainerVersionLifecycleStateActive   ContainerVersionLifecycleStateEnum = "ACTIVE"
	ContainerVersionLifecycleStateInactive ContainerVersionLifecycleStateEnum = "INACTIVE"
)

var mappingContainerVersionLifecycleStateEnum = map[string]ContainerVersionLifecycleStateEnum{
	"ACTIVE":   ContainerVersionLifecycleStateActive,
	"INACTIVE": ContainerVersionLifecycleStateInactive,
}

var mappingContainerVersionLifecycleStateEnumLowerCase = map[string]ContainerVersionLifecycleStateEnum{
	"active":   ContainerVersionLifecycleStateActive,
	"inactive": ContainerVersionLifecycleStateInactive,
}

// GetContainerVersionLifecycleStateEnumValues Enumerates the set of values for ContainerVersionLifecycleStateEnum
func GetContainerVersionLifecycleStateEnumValues() []ContainerVersionLifecycleStateEnum {
	values := make([]ContainerVersionLifecycleStateEnum, 0)
	for _, v := range mappingContainerVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerVersionLifecycleStateEnumStringValues Enumerates the set of values in String for ContainerVersionLifecycleStateEnum
func GetContainerVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingContainerVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerVersionLifecycleStateEnum(val string) (ContainerVersionLifecycleStateEnum, bool) {
	enum, ok := mappingContainerVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
