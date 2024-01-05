// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// PluginStatusEnum Enum with underlying type: string
type PluginStatusEnum string

// Set of constants representing the allowable values for PluginStatusEnum
const (
	PluginStatusRunning PluginStatusEnum = "RUNNING"
	PluginStatusStopped PluginStatusEnum = "STOPPED"
	PluginStatusInvalid PluginStatusEnum = "INVALID"
	PluginStatusFailed  PluginStatusEnum = "FAILED"
)

var mappingPluginStatusEnum = map[string]PluginStatusEnum{
	"RUNNING": PluginStatusRunning,
	"STOPPED": PluginStatusStopped,
	"INVALID": PluginStatusInvalid,
	"FAILED":  PluginStatusFailed,
}

var mappingPluginStatusEnumLowerCase = map[string]PluginStatusEnum{
	"running": PluginStatusRunning,
	"stopped": PluginStatusStopped,
	"invalid": PluginStatusInvalid,
	"failed":  PluginStatusFailed,
}

// GetPluginStatusEnumValues Enumerates the set of values for PluginStatusEnum
func GetPluginStatusEnumValues() []PluginStatusEnum {
	values := make([]PluginStatusEnum, 0)
	for _, v := range mappingPluginStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPluginStatusEnumStringValues Enumerates the set of values in String for PluginStatusEnum
func GetPluginStatusEnumStringValues() []string {
	return []string{
		"RUNNING",
		"STOPPED",
		"INVALID",
		"FAILED",
	}
}

// GetMappingPluginStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluginStatusEnum(val string) (PluginStatusEnum, bool) {
	enum, ok := mappingPluginStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
