// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
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
