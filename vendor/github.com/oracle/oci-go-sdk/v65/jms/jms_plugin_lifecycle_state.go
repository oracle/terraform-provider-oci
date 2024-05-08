// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// JmsPluginLifecycleStateEnum Enum with underlying type: string
type JmsPluginLifecycleStateEnum string

// Set of constants representing the allowable values for JmsPluginLifecycleStateEnum
const (
	JmsPluginLifecycleStateActive         JmsPluginLifecycleStateEnum = "ACTIVE"
	JmsPluginLifecycleStateInactive       JmsPluginLifecycleStateEnum = "INACTIVE"
	JmsPluginLifecycleStateNeedsAttention JmsPluginLifecycleStateEnum = "NEEDS_ATTENTION"
	JmsPluginLifecycleStateDeleted        JmsPluginLifecycleStateEnum = "DELETED"
)

var mappingJmsPluginLifecycleStateEnum = map[string]JmsPluginLifecycleStateEnum{
	"ACTIVE":          JmsPluginLifecycleStateActive,
	"INACTIVE":        JmsPluginLifecycleStateInactive,
	"NEEDS_ATTENTION": JmsPluginLifecycleStateNeedsAttention,
	"DELETED":         JmsPluginLifecycleStateDeleted,
}

var mappingJmsPluginLifecycleStateEnumLowerCase = map[string]JmsPluginLifecycleStateEnum{
	"active":          JmsPluginLifecycleStateActive,
	"inactive":        JmsPluginLifecycleStateInactive,
	"needs_attention": JmsPluginLifecycleStateNeedsAttention,
	"deleted":         JmsPluginLifecycleStateDeleted,
}

// GetJmsPluginLifecycleStateEnumValues Enumerates the set of values for JmsPluginLifecycleStateEnum
func GetJmsPluginLifecycleStateEnumValues() []JmsPluginLifecycleStateEnum {
	values := make([]JmsPluginLifecycleStateEnum, 0)
	for _, v := range mappingJmsPluginLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetJmsPluginLifecycleStateEnumStringValues Enumerates the set of values in String for JmsPluginLifecycleStateEnum
func GetJmsPluginLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingJmsPluginLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJmsPluginLifecycleStateEnum(val string) (JmsPluginLifecycleStateEnum, bool) {
	enum, ok := mappingJmsPluginLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
