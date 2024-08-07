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

// AgentInstallerSortByEnum Enum with underlying type: string
type AgentInstallerSortByEnum string

// Set of constants representing the allowable values for AgentInstallerSortByEnum
const (
	AgentInstallerSortByAgentInstallerId     AgentInstallerSortByEnum = "agentInstallerId"
	AgentInstallerSortByOsFamily             AgentInstallerSortByEnum = "osFamily"
	AgentInstallerSortByPlatformArchitecture AgentInstallerSortByEnum = "platformArchitecture"
)

var mappingAgentInstallerSortByEnum = map[string]AgentInstallerSortByEnum{
	"agentInstallerId":     AgentInstallerSortByAgentInstallerId,
	"osFamily":             AgentInstallerSortByOsFamily,
	"platformArchitecture": AgentInstallerSortByPlatformArchitecture,
}

var mappingAgentInstallerSortByEnumLowerCase = map[string]AgentInstallerSortByEnum{
	"agentinstallerid":     AgentInstallerSortByAgentInstallerId,
	"osfamily":             AgentInstallerSortByOsFamily,
	"platformarchitecture": AgentInstallerSortByPlatformArchitecture,
}

// GetAgentInstallerSortByEnumValues Enumerates the set of values for AgentInstallerSortByEnum
func GetAgentInstallerSortByEnumValues() []AgentInstallerSortByEnum {
	values := make([]AgentInstallerSortByEnum, 0)
	for _, v := range mappingAgentInstallerSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentInstallerSortByEnumStringValues Enumerates the set of values in String for AgentInstallerSortByEnum
func GetAgentInstallerSortByEnumStringValues() []string {
	return []string{
		"agentInstallerId",
		"osFamily",
		"platformArchitecture",
	}
}

// GetMappingAgentInstallerSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentInstallerSortByEnum(val string) (AgentInstallerSortByEnum, bool) {
	enum, ok := mappingAgentInstallerSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
