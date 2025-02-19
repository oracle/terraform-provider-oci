// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PluginErrorReasonEnum Enum with underlying type: string
type PluginErrorReasonEnum string

// Set of constants representing the allowable values for PluginErrorReasonEnum
const (
	PluginErrorReasonRegistrationPolicy PluginErrorReasonEnum = "REGISTRATION_POLICY"
	PluginErrorReasonLogResourcePolicy  PluginErrorReasonEnum = "LOG_RESOURCE_POLICY"
	PluginErrorReasonNoFleet            PluginErrorReasonEnum = "NO_FLEET"
)

var mappingPluginErrorReasonEnum = map[string]PluginErrorReasonEnum{
	"REGISTRATION_POLICY": PluginErrorReasonRegistrationPolicy,
	"LOG_RESOURCE_POLICY": PluginErrorReasonLogResourcePolicy,
	"NO_FLEET":            PluginErrorReasonNoFleet,
}

var mappingPluginErrorReasonEnumLowerCase = map[string]PluginErrorReasonEnum{
	"registration_policy": PluginErrorReasonRegistrationPolicy,
	"log_resource_policy": PluginErrorReasonLogResourcePolicy,
	"no_fleet":            PluginErrorReasonNoFleet,
}

// GetPluginErrorReasonEnumValues Enumerates the set of values for PluginErrorReasonEnum
func GetPluginErrorReasonEnumValues() []PluginErrorReasonEnum {
	values := make([]PluginErrorReasonEnum, 0)
	for _, v := range mappingPluginErrorReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetPluginErrorReasonEnumStringValues Enumerates the set of values in String for PluginErrorReasonEnum
func GetPluginErrorReasonEnumStringValues() []string {
	return []string{
		"REGISTRATION_POLICY",
		"LOG_RESOURCE_POLICY",
		"NO_FLEET",
	}
}

// GetMappingPluginErrorReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluginErrorReasonEnum(val string) (PluginErrorReasonEnum, bool) {
	enum, ok := mappingPluginErrorReasonEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
