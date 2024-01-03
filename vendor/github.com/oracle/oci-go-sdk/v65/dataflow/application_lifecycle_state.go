// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"strings"
)

// ApplicationLifecycleStateEnum Enum with underlying type: string
type ApplicationLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationLifecycleStateEnum
const (
	ApplicationLifecycleStateActive   ApplicationLifecycleStateEnum = "ACTIVE"
	ApplicationLifecycleStateDeleted  ApplicationLifecycleStateEnum = "DELETED"
	ApplicationLifecycleStateInactive ApplicationLifecycleStateEnum = "INACTIVE"
)

var mappingApplicationLifecycleStateEnum = map[string]ApplicationLifecycleStateEnum{
	"ACTIVE":   ApplicationLifecycleStateActive,
	"DELETED":  ApplicationLifecycleStateDeleted,
	"INACTIVE": ApplicationLifecycleStateInactive,
}

var mappingApplicationLifecycleStateEnumLowerCase = map[string]ApplicationLifecycleStateEnum{
	"active":   ApplicationLifecycleStateActive,
	"deleted":  ApplicationLifecycleStateDeleted,
	"inactive": ApplicationLifecycleStateInactive,
}

// GetApplicationLifecycleStateEnumValues Enumerates the set of values for ApplicationLifecycleStateEnum
func GetApplicationLifecycleStateEnumValues() []ApplicationLifecycleStateEnum {
	values := make([]ApplicationLifecycleStateEnum, 0)
	for _, v := range mappingApplicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationLifecycleStateEnumStringValues Enumerates the set of values in String for ApplicationLifecycleStateEnum
func GetApplicationLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"INACTIVE",
	}
}

// GetMappingApplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationLifecycleStateEnum(val string) (ApplicationLifecycleStateEnum, bool) {
	enum, ok := mappingApplicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
