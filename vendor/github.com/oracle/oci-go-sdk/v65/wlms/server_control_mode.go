// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"strings"
)

// ServerControlModeEnum Enum with underlying type: string
type ServerControlModeEnum string

// Set of constants representing the allowable values for ServerControlModeEnum
const (
	ServerControlModeUseNodeManager ServerControlModeEnum = "USE_NODE_MANAGER"
	ServerControlModeUseScripts     ServerControlModeEnum = "USE_SCRIPTS"
)

var mappingServerControlModeEnum = map[string]ServerControlModeEnum{
	"USE_NODE_MANAGER": ServerControlModeUseNodeManager,
	"USE_SCRIPTS":      ServerControlModeUseScripts,
}

var mappingServerControlModeEnumLowerCase = map[string]ServerControlModeEnum{
	"use_node_manager": ServerControlModeUseNodeManager,
	"use_scripts":      ServerControlModeUseScripts,
}

// GetServerControlModeEnumValues Enumerates the set of values for ServerControlModeEnum
func GetServerControlModeEnumValues() []ServerControlModeEnum {
	values := make([]ServerControlModeEnum, 0)
	for _, v := range mappingServerControlModeEnum {
		values = append(values, v)
	}
	return values
}

// GetServerControlModeEnumStringValues Enumerates the set of values in String for ServerControlModeEnum
func GetServerControlModeEnumStringValues() []string {
	return []string{
		"USE_NODE_MANAGER",
		"USE_SCRIPTS",
	}
}

// GetMappingServerControlModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServerControlModeEnum(val string) (ServerControlModeEnum, bool) {
	enum, ok := mappingServerControlModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
