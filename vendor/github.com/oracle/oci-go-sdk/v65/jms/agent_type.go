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

// AgentTypeEnum Enum with underlying type: string
type AgentTypeEnum string

// Set of constants representing the allowable values for AgentTypeEnum
const (
	AgentTypeOma AgentTypeEnum = "OMA"
	AgentTypeOca AgentTypeEnum = "OCA"
)

var mappingAgentTypeEnum = map[string]AgentTypeEnum{
	"OMA": AgentTypeOma,
	"OCA": AgentTypeOca,
}

var mappingAgentTypeEnumLowerCase = map[string]AgentTypeEnum{
	"oma": AgentTypeOma,
	"oca": AgentTypeOca,
}

// GetAgentTypeEnumValues Enumerates the set of values for AgentTypeEnum
func GetAgentTypeEnumValues() []AgentTypeEnum {
	values := make([]AgentTypeEnum, 0)
	for _, v := range mappingAgentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentTypeEnumStringValues Enumerates the set of values in String for AgentTypeEnum
func GetAgentTypeEnumStringValues() []string {
	return []string{
		"OMA",
		"OCA",
	}
}

// GetMappingAgentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentTypeEnum(val string) (AgentTypeEnum, bool) {
	enum, ok := mappingAgentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
