// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// ConfigLifecycleStateEnum Enum with underlying type: string
type ConfigLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigLifecycleStateEnum
const (
	ConfigLifecycleStateActive  ConfigLifecycleStateEnum = "ACTIVE"
	ConfigLifecycleStateDeleted ConfigLifecycleStateEnum = "DELETED"
)

var mappingConfigLifecycleStateEnum = map[string]ConfigLifecycleStateEnum{
	"ACTIVE":  ConfigLifecycleStateActive,
	"DELETED": ConfigLifecycleStateDeleted,
}

var mappingConfigLifecycleStateEnumLowerCase = map[string]ConfigLifecycleStateEnum{
	"active":  ConfigLifecycleStateActive,
	"deleted": ConfigLifecycleStateDeleted,
}

// GetConfigLifecycleStateEnumValues Enumerates the set of values for ConfigLifecycleStateEnum
func GetConfigLifecycleStateEnumValues() []ConfigLifecycleStateEnum {
	values := make([]ConfigLifecycleStateEnum, 0)
	for _, v := range mappingConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigLifecycleStateEnumStringValues Enumerates the set of values in String for ConfigLifecycleStateEnum
func GetConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigLifecycleStateEnum(val string) (ConfigLifecycleStateEnum, bool) {
	enum, ok := mappingConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
