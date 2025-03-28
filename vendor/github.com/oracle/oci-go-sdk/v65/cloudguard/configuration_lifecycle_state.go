// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"strings"
)

// ConfigurationLifecycleStateEnum Enum with underlying type: string
type ConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ConfigurationLifecycleStateEnum
const (
	ConfigurationLifecycleStateEnabled  ConfigurationLifecycleStateEnum = "ENABLED"
	ConfigurationLifecycleStateDisabled ConfigurationLifecycleStateEnum = "DISABLED"
)

var mappingConfigurationLifecycleStateEnum = map[string]ConfigurationLifecycleStateEnum{
	"ENABLED":  ConfigurationLifecycleStateEnabled,
	"DISABLED": ConfigurationLifecycleStateDisabled,
}

var mappingConfigurationLifecycleStateEnumLowerCase = map[string]ConfigurationLifecycleStateEnum{
	"enabled":  ConfigurationLifecycleStateEnabled,
	"disabled": ConfigurationLifecycleStateDisabled,
}

// GetConfigurationLifecycleStateEnumValues Enumerates the set of values for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumValues() []ConfigurationLifecycleStateEnum {
	values := make([]ConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingConfigurationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationLifecycleStateEnumStringValues Enumerates the set of values in String for ConfigurationLifecycleStateEnum
func GetConfigurationLifecycleStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingConfigurationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationLifecycleStateEnum(val string) (ConfigurationLifecycleStateEnum, bool) {
	enum, ok := mappingConfigurationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
